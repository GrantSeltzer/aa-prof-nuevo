package aanuevo

// TODO: json/toml tags

type AppArmorProfile struct {
	ProfileName   string
	Capabilities  []Capability
	FileRules     []FileRule
	NetworkRules  []NetworkRule
	ChildProfiles []AppArmorProfile
}

type Capability string // TODO: Decide on accepted syntax (CAP_CHOWN, chown, CHOWN, etc...)

type FileRule struct {
	Path       string
	Permission FilePermission
}

type FilePermission struct {
	Owner     bool
	Read      bool
	Write     bool
	Execute   bool // TODO: Various execute permission options
	MemoryMap bool
	Lock      bool // Validate, requires RW
	Link      bool
}

type NetworkRule struct {
	// The rule qualifiers modify network rules in the same manner that they are applied to file rules, ie. denied rules take precedence over allow rules and will both subtract permissions and quiet auditing.
	Deny       bool
	Permission NetworkPermission
}

//TODO: Take rest of comments from: http://wiki.apparmor.net/index.php/AppArmor_Core_Policy_Reference
// NetworkPermission specifies what operations are allowed on the network socket
// Should this just be a string with NetworkRules having slices of them? Validate in translation layer?
// '-> Check other members of runtime-spec
type NetworkPermission struct {
	Create      bool
	Accept      bool
	Bind        bool
	Connect     bool
	Listen      bool // listen[=X] - permission to listen on a socket of the domain, type, protocl specified. If the optional backlog parameter is specified it limits the maximum size that can be pasted to listen.
	Read        bool
	Write       bool
	Send        bool
	Receive     bool
	GetSockName bool
	GetPeerName bool
	GetSockOpt  bool
	SetSockOpt  bool // set socket options
	Fcntl       bool
	Ioctl       bool
	Shutdown    bool
	GetPeerSec  bool
}

type NetworkType struct {
	//The actually set of domain names is growing and ??? should be consulted for a complete updated list
	//	stream' | 'dgram' | 'seqpacket' | 'raw' | 'rdm' | 'packet' | 'dccp'
}
