// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDeploymentSetAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *ModifyRCDeploymentSetAttributeRequest
	GetDeploymentSetId() *string
	SetDeploymentSetName(v string) *ModifyRCDeploymentSetAttributeRequest
	GetDeploymentSetName() *string
	SetDescription(v string) *ModifyRCDeploymentSetAttributeRequest
	GetDescription() *string
	SetRegionId(v string) *ModifyRCDeploymentSetAttributeRequest
	GetRegionId() *string
}

type ModifyRCDeploymentSetAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// example:
	//
	// DeploymentSetTestName
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCDeploymentSetAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDeploymentSetAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCDeploymentSetAttributeRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ModifyRCDeploymentSetAttributeRequest) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *ModifyRCDeploymentSetAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRCDeploymentSetAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCDeploymentSetAttributeRequest) SetDeploymentSetId(v string) *ModifyRCDeploymentSetAttributeRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyRCDeploymentSetAttributeRequest) SetDeploymentSetName(v string) *ModifyRCDeploymentSetAttributeRequest {
	s.DeploymentSetName = &v
	return s
}

func (s *ModifyRCDeploymentSetAttributeRequest) SetDescription(v string) *ModifyRCDeploymentSetAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyRCDeploymentSetAttributeRequest) SetRegionId(v string) *ModifyRCDeploymentSetAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCDeploymentSetAttributeRequest) Validate() error {
	return dara.Validate(s)
}
