package aanuevo

// TODO: json/toml tags

type AppArmorProfile struct {
	Capabilities  []Capability
	FileRules     []FileRule
	NetworkRules  []NetworkRule
	ChildProfiles []AppArmorProfile
}

type Capability string

type FileRule struct {
	Path       string
	Permission Permission
}

type Permission struct {
	Owner     bool
	Read      bool
	Write     bool
	Execute   bool // TODO: Various execute permission options
	MemoryMap bool
	Lock      bool // Validate, requires RW
	Link      bool
}

type NetworkRule struct {
}
