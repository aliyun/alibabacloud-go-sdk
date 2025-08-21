// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployInstanceSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentType(v string) *DeployInstanceSDGRequest
	GetDeploymentType() *string
	SetDiskAccessProtocol(v string) *DeployInstanceSDGRequest
	GetDiskAccessProtocol() *string
	SetDiskType(v string) *DeployInstanceSDGRequest
	GetDiskType() *string
	SetInstanceIds(v []*string) *DeployInstanceSDGRequest
	GetInstanceIds() []*string
	SetSDGId(v string) *DeployInstanceSDGRequest
	GetSDGId() *string
}

type DeployInstanceSDGRequest struct {
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
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the SDG.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DeployInstanceSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceSDGRequest) GoString() string {
	return s.String()
}

func (s *DeployInstanceSDGRequest) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DeployInstanceSDGRequest) GetDiskAccessProtocol() *string {
	return s.DiskAccessProtocol
}

func (s *DeployInstanceSDGRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *DeployInstanceSDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeployInstanceSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DeployInstanceSDGRequest) SetDeploymentType(v string) *DeployInstanceSDGRequest {
	s.DeploymentType = &v
	return s
}

func (s *DeployInstanceSDGRequest) SetDiskAccessProtocol(v string) *DeployInstanceSDGRequest {
	s.DiskAccessProtocol = &v
	return s
}

func (s *DeployInstanceSDGRequest) SetDiskType(v string) *DeployInstanceSDGRequest {
	s.DiskType = &v
	return s
}

func (s *DeployInstanceSDGRequest) SetInstanceIds(v []*string) *DeployInstanceSDGRequest {
	s.InstanceIds = v
	return s
}

func (s *DeployInstanceSDGRequest) SetSDGId(v string) *DeployInstanceSDGRequest {
	s.SDGId = &v
	return s
}

func (s *DeployInstanceSDGRequest) Validate() error {
	return dara.Validate(s)
}
