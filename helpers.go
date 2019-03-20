package urlfilter

import (
	"strings"
)

// splitWithEscapeCharacter splits string by the specified separator if it is not escaped
func splitWithEscapeCharacter(str string, sep byte, escapeCharacter byte, preserveAllTokens bool) []string {
	parts := make([]string, 0)

	if str == "" {
		return parts
	}

	var sb strings.Builder
	escaped := false
	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == escapeCharacter {
			escaped = true
		} else if c == sep {
			if escaped {
				sb.WriteByte(c)
				escaped = false
			} else {
				if preserveAllTokens || sb.Len() > 0 {
					parts = append(parts, sb.String())
					sb.Reset()
				}
			}
		} else {
			if escaped {
				escaped = false
				sb.WriteByte(escapeCharacter)
			}
			sb.WriteByte(c)
		}
	}

	if preserveAllTokens || sb.Len() > 0 {
		parts = append(parts, sb.String())
	}

	return parts
}

// getSubdomains splits the specified hostname and returns all subdomains (including the hostname itself)
func getSubdomains(hostname string) []string {
	parts := strings.Split(hostname, ".")
	var subdomains []string
	var domain = ""
	for i := len(parts) - 1; i >= 0; i-- {
		if domain == "" {
			domain = parts[i]
		} else {
			domain = parts[i] + "." + domain
		}
		subdomains = append(subdomains, domain)
	}
	return subdomains
}
