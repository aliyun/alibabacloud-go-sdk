// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentType(v string) *DeploySDGRequest
	GetDeploymentType() *string
	SetInstanceIds(v []*string) *DeploySDGRequest
	GetInstanceIds() []*string
	SetSDGId(v string) *DeploySDGRequest
	GetSDGId() *string
}

type DeploySDGRequest struct {
	// The SDG deployment type. Valid values:
	//
	// 	- common (default): read/write deployment. Data updates are written to disks.
	//
	// 	- overlay: read/write splitting deployment. Content in SDGs is read-only. Data updates are written to the local storage of the instance.
	//
	// example:
	//
	// common
	DeploymentType *string `json:"DeploymentType,omitempty" xml:"DeploymentType,omitempty"`
	// The IDs of instances on which you want to deploy SDGs. You can deploy SDGs on a maximum of 100 instances at a time.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The SDG ID. This parameter is used to create a disk, which is attached to an instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DeploySDGRequest) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGRequest) GoString() string {
	return s.String()
}

func (s *DeploySDGRequest) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DeploySDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeploySDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DeploySDGRequest) SetDeploymentType(v string) *DeploySDGRequest {
	s.DeploymentType = &v
	return s
}

func (s *DeploySDGRequest) SetInstanceIds(v []*string) *DeploySDGRequest {
	s.InstanceIds = v
	return s
}

func (s *DeploySDGRequest) SetSDGId(v string) *DeploySDGRequest {
	s.SDGId = &v
	return s
}

func (s *DeploySDGRequest) Validate() error {
	return dara.Validate(s)
}
