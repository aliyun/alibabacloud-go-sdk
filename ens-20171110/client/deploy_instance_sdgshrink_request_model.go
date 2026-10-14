// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployInstanceSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentType(v string) *DeployInstanceSDGShrinkRequest
	GetDeploymentType() *string
	SetDiskAccessProtocol(v string) *DeployInstanceSDGShrinkRequest
	GetDiskAccessProtocol() *string
	SetDiskType(v string) *DeployInstanceSDGShrinkRequest
	GetDiskType() *string
	SetInstanceIdsShrink(v string) *DeployInstanceSDGShrinkRequest
	GetInstanceIdsShrink() *string
	SetSDGId(v string) *DeployInstanceSDGShrinkRequest
	GetSDGId() *string
}

type DeployInstanceSDGShrinkRequest struct {
	// The SDG deployment type. Set the value to shared, which indicates read/write splitting shared deployment. In this mode, the SDG content is read-only, and data updates are written to the local storage of the instance.
	//
	// This parameter is required. Currently, only the value shared is supported. If this parameter is not specified, the service returns a Parameter.Invalid error.
	//
	// example:
	//
	// shared
	DeploymentType *string `json:"DeploymentType,omitempty" xml:"DeploymentType,omitempty"`
	// The access protocol type of the cloud disk. Valid values:
	//
	// - iscsi: the default protocol for standard cloud disks. Both RBD cloud disks and standard cloud disks support this protocol.
	//
	// - rbd: the default protocol for RBD cloud disks. Only RBD cloud disks support this mount protocol.
	//
	// example:
	//
	// iscsi
	DiskAccessProtocol *string `json:"DiskAccessProtocol,omitempty" xml:"DiskAccessProtocol,omitempty"`
	// The type of the shared cloud disk to be created. Valid values:
	//
	// - standard: standard cloud disk.
	//
	// - rbd: RBD cloud disk.
	//
	// >
	//
	// > - If this parameter is not specified, a standard cloud disk is created by default.
	//
	// > - An RBD shared cloud disk can be created only when the source disk type is RBD.
	//
	// example:
	//
	// standard
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The list of instance IDs. The value is a JSON array. A maximum of 100 IDs can be specified.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The ID of the SDG.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DeployInstanceSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeployInstanceSDGShrinkRequest) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DeployInstanceSDGShrinkRequest) GetDiskAccessProtocol() *string {
	return s.DiskAccessProtocol
}

func (s *DeployInstanceSDGShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *DeployInstanceSDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DeployInstanceSDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DeployInstanceSDGShrinkRequest) SetDeploymentType(v string) *DeployInstanceSDGShrinkRequest {
	s.DeploymentType = &v
	return s
}

func (s *DeployInstanceSDGShrinkRequest) SetDiskAccessProtocol(v string) *DeployInstanceSDGShrinkRequest {
	s.DiskAccessProtocol = &v
	return s
}

func (s *DeployInstanceSDGShrinkRequest) SetDiskType(v string) *DeployInstanceSDGShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *DeployInstanceSDGShrinkRequest) SetInstanceIdsShrink(v string) *DeployInstanceSDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DeployInstanceSDGShrinkRequest) SetSDGId(v string) *DeployInstanceSDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *DeployInstanceSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
