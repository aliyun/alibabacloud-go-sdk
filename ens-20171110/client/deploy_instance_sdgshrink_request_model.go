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
	// The deployment type of the SDG. shared: shared read/write splitting deployment. The content of the SDG is read-only, and data updates are written to the local storage of the instance.
	//
	// example:
	//
	// shared
	DeploymentType     *string `json:"DeploymentType,omitempty" xml:"DeploymentType,omitempty"`
	DiskAccessProtocol *string `json:"DiskAccessProtocol,omitempty" xml:"DiskAccessProtocol,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The IDs of the instances. The value is a JSON array that consists of up to 100 IDs.
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
