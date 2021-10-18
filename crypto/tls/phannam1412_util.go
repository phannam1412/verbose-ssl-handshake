package tls

func TlsVersionString(vers uint16) string {
	if vers == VersionTLS10 {
		return "1.0"
	}
	if vers == VersionTLS11 {
		return "1.1"
	}
	if vers == VersionTLS12 {
		return "1.2"
	}
	if vers == VersionTLS10 {
		return "1.3"
	}
	return ""
}

func KeyUsageString(usage int) {}
