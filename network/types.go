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

	// MacAddress
	MacAddress string `json:"mac_address,omitempty"`

	// Addresses slice contains the IPv4 and/or IPv6 and mask to set on the network interface
	Addresses []string `json:"addresses,omitempty"`

	// Routes slice contains Route objects with destination, source, gateway strings
	Routes []Route `json:"routes,omitempty"`

	// Mtu sets the mtu value for the interface and will be mirrored on both the host and
	// container's interfaces if a pair is created, specifically in the case of type veth
	// Note: This does not apply to loopback interfaces.
	Mtu int `json:"mtu,omitempty"`
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


// Routes can be specified to create entries in the route table as the container is started
//
// All of destination, source, and gateway should be either IPv4 or IPv6.
// One of the three options must be present, and ommitted entries will use their
// IP family default for the route table.  For IPv4 for example, setting the
// gateway to 1.2.3.4 and the interface to eth0 will set up a standard
// destination of 0.0.0.0(or *) when viewed in the route table.
type Route struct {
	// Sets the destination and mask, should be a CIDR.  Accepts IPv4 and IPv6
	Destination string `json:"destination,omitempty"`

	// Sets the source and mask, should be a CIDR.  Accepts IPv4 and IPv6
	Source string `json:"source,omitempty"`

	// Sets the gateway.  Accepts IPv4 and IPv6
	Gateway string `json:"gateway,omitempty"`

	// The device to set this route up for, for example: eth0
	InterfaceName string `json:"interface_name,omitempty"`
}