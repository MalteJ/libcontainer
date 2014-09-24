package network

// Network defines configuration for a container's networking stack
//
// The network configuration can be omited from a container causing the
// container to be setup with the host's networking stack
type Network struct {
	// Type sets the networks type, commonly veth and loopback
	Type string `json:"type,omitempty"`

	// Path to network namespace
	NsPath string `json:"ns_path,omitempty"`

	// The bridge to use.
	Bridge string `json:"bridge,omitempty"`

	// Prefix for the veth interfaces.
	VethPrefix string `json:"veth_prefix,omitempty"`

	// Addresses slice contains the IPv4 and/or IPv6 and mask to set on the network interface
	Addresses []string `json:"addresses,omitempty"`

	// Routes slice contains Route objects with destination, source, gateway strings
	Routes []Route `json:"routes,omitempty"`

	// Mtu sets the mtu value for the interface and will be mirrored on both the host and
	// container's interfaces if a pair is created, specifically in the case of type veth
	// Note: This does not apply to loopback interfaces.
	Mtu int `json:"mtu,omitempty"`
}

// Struct describing a network route. Maximum two parameters may be empty.
type Route struct {
	// Destination IPv4 or IPv6 and mask
	Destination string `json:"destination,omitempty"`

	// Source IPv4 or IPv6 address and mask
	Source string `json:"source,omitempty"`

	// Gateway IPv4 or IPv6 address for this route
	Gateway string `json:"gateway,omitempty"`
}

// Struct describing the network specific runtime state that will be maintained by libcontainer for all running containers
// Do not depend on it outside of libcontainer.
type NetworkState struct {
	// The name of the veth interface on the Host.
	VethHost string `json:"veth_host,omitempty"`
	// The name of the veth interface created inside the container for the child.
	VethChild string `json:"veth_child,omitempty"`
	// Net namespace path.
	NsPath string `json:"ns_path,omitempty"`
}
