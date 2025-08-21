// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentType(v string) *DeploySDGShrinkRequest
	GetDeploymentType() *string
	SetInstanceIdsShrink(v string) *DeploySDGShrinkRequest
	GetInstanceIdsShrink() *string
	SetSDGId(v string) *DeploySDGShrinkRequest
	GetSDGId() *string
}

type DeploySDGShrinkRequest struct {
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
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The SDG ID. This parameter is used to create a disk, which is attached to an instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DeploySDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeploySDGShrinkRequest) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DeploySDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DeploySDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DeploySDGShrinkRequest) SetDeploymentType(v string) *DeploySDGShrinkRequest {
	s.DeploymentType = &v
	return s
}

func (s *DeploySDGShrinkRequest) SetInstanceIdsShrink(v string) *DeploySDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DeploySDGShrinkRequest) SetSDGId(v string) *DeploySDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *DeploySDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
