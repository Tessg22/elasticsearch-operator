package constants

const (
	ProxyName                  = "cluster"
	TrustedCABundleKey         = "ca-bundle.crt"
	TrustedCABundleMountDir    = "/etc/pki/ca-trust/extracted/pem/"
	TrustedCABundleMountFile   = "tls-ca-bundle.pem"
	InjectTrustedCABundleLabel = "config.openshift.io/inject-trusted-cabundle"
	TrustedCABundleHashName    = "logging.openshift.io/hash"
	KibanaTrustedCAName        = "kibana-trusted-ca-bundle"
)

var ReconcileForGlobalProxyList = []string{KibanaTrustedCAName}