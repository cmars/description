// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package description

import (
	"time"

	"github.com/juju/version"
	"gopkg.in/juju/names.v2"
)

// AgentTools represent the version and related binary file
// that the machine and unit agents are using.
type AgentTools interface {
	Version() version.Binary
	URL() string
	SHA256() string
	Size() int64
}

// Unit represents an instance of an application in a model.
type Unit interface {
	HasAnnotations
	HasConstraints

	Tag() names.UnitTag
	Name() string
	Machine() names.MachineTag

	PasswordHash() string

	Principal() names.UnitTag
	Subordinates() []names.UnitTag

	MeterStatusCode() string
	MeterStatusInfo() string

	Tools() AgentTools
	SetTools(AgentToolsArgs)

	WorkloadStatus() Status
	SetWorkloadStatus(StatusArgs)

	WorkloadStatusHistory() []Status
	SetWorkloadStatusHistory([]StatusArgs)

	WorkloadVersion() string

	WorkloadVersionHistory() []Status
	SetWorkloadVersionHistory([]StatusArgs)

	AgentStatus() Status
	SetAgentStatus(StatusArgs)

	AgentStatusHistory() []Status
	SetAgentStatusHistory([]StatusArgs)

	AddResource(UnitResourceArgs) UnitResource
	Resources() []UnitResource

	AddPayload(PayloadArgs) Payload
	Payloads() []Payload

	Validate() error
}

// Space represents a network space, which is a named collection of subnets.
type Space interface {
	Name() string
	Public() bool
	ProviderID() string
}

// LinkLayerDevice represents a link layer device.
type LinkLayerDevice interface {
	Name() string
	MTU() uint
	ProviderID() string
	MachineID() string
	Type() string
	MACAddress() string
	IsAutoStart() bool
	IsUp() bool
	ParentName() string
}

// Subnet represents a network subnet.
type Subnet interface {
	ProviderId() string
	ProviderNetworkId() string
	CIDR() string
	VLANTag() int
	AvailabilityZone() string
	SpaceName() string
	AllocatableIPHigh() string
	AllocatableIPLow() string
}

// IPAddress represents an IP address.
type IPAddress interface {
	ProviderID() string
	DeviceName() string
	MachineID() string
	SubnetCIDR() string
	ConfigMethod() string
	Value() string
	DNSServers() []string
	DNSSearchDomains() []string
	GatewayAddress() string
}

// SSHHostKey represents an ssh host key.
type SSHHostKey interface {
	MachineID() string
	Keys() []string
}

// CloudImageMetadata represents an IP cloudimagemetadata.
type CloudImageMetadata interface {
	Stream() string
	Region() string
	Version() string
	Series() string
	Arch() string
	VirtType() string
	RootStorageType() string
	RootStorageSize() (uint64, bool)
	DateCreated() int64
	Source() string
	Priority() int
	ImageId() string
}

// Action represents an IP action.
type Action interface {
	Id() string
	Receiver() string
	Name() string
	Parameters() map[string]interface{}
	Enqueued() time.Time
	Started() time.Time
	Completed() time.Time
	Results() map[string]interface{}
	Status() string
	Message() string
}

// Volume represents a volume (disk, logical volume, etc.) in the model.
type Volume interface {
	HasStatus
	HasStatusHistory

	Tag() names.VolumeTag
	Storage() names.StorageTag

	Provisioned() bool

	Size() uint64
	Pool() string

	HardwareID() string
	VolumeID() string
	Persistent() bool

	Attachments() []VolumeAttachment
	AddAttachment(VolumeAttachmentArgs) VolumeAttachment
}

// VolumeAttachment represents a volume attached to a machine.
type VolumeAttachment interface {
	Machine() names.MachineTag
	Provisioned() bool
	ReadOnly() bool
	DeviceName() string
	DeviceLink() string
	BusAddress() string
}

// Filesystem represents a filesystem in the model.
type Filesystem interface {
	HasStatus
	HasStatusHistory

	Tag() names.FilesystemTag
	Volume() names.VolumeTag
	Storage() names.StorageTag

	Provisioned() bool

	Size() uint64
	Pool() string

	FilesystemID() string

	Attachments() []FilesystemAttachment
	AddAttachment(FilesystemAttachmentArgs) FilesystemAttachment
}

// FilesystemAttachment represents a filesystem attached to a machine.
type FilesystemAttachment interface {
	Machine() names.MachineTag
	Provisioned() bool
	MountPoint() string
	ReadOnly() bool
}

// Storage represents the state of a unit or application-wide storage instance
// in the model.
type Storage interface {
	Tag() names.StorageTag
	Kind() string
	// Owner returns the tag of the application or unit that owns this storage
	// instance.
	Owner() (names.Tag, error)
	Name() string

	Attachments() []names.UnitTag

	Validate() error
}

// StoragePool represents a named storage pool and its settings.
type StoragePool interface {
	Name() string
	Provider() string
	Attributes() map[string]interface{}
}

// StorageConstraint repressents the user-specified constraints for
// provisioning storage instances for an application unit.
type StorageConstraint interface {
	// Pool is the name of the storage pool from which to provision the
	// storage instances.
	Pool() string
	// Size is the required size of the storage instances, in MiB.
	Size() uint64
	// Count is the required number of storage instances.
	Count() uint64
}
