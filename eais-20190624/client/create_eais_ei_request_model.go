// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaisEiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaisEiRequest
	GetClientToken() *string
	SetInstanceName(v string) *CreateEaisEiRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateEaisEiRequest
	GetInstanceType() *string
	SetRegionId(v string) *CreateEaisEiRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaisEiRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaisEiRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaisEiRequestTag) *CreateEaisEiRequest
	GetTag() []*CreateEaisEiRequestTag
	SetVSwitchId(v string) *CreateEaisEiRequest
	GetVSwitchId() *string
}

type CreateEaisEiRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// test_ei
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-bp16jgp51ttnkbdr****
	SecurityGroupId *string                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaisEiRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp17wmd1wb6fwlimk****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaisEiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaisEiRequest) GoString() string {
	return s.String()
}

func (s *CreateEaisEiRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaisEiRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateEaisEiRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateEaisEiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaisEiRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaisEiRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaisEiRequest) GetTag() []*CreateEaisEiRequestTag {
	return s.Tag
}

func (s *CreateEaisEiRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaisEiRequest) SetClientToken(v string) *CreateEaisEiRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaisEiRequest) SetInstanceName(v string) *CreateEaisEiRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateEaisEiRequest) SetInstanceType(v string) *CreateEaisEiRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateEaisEiRequest) SetRegionId(v string) *CreateEaisEiRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaisEiRequest) SetResourceGroupId(v string) *CreateEaisEiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaisEiRequest) SetSecurityGroupId(v string) *CreateEaisEiRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaisEiRequest) SetTag(v []*CreateEaisEiRequestTag) *CreateEaisEiRequest {
	s.Tag = v
	return s
}

func (s *CreateEaisEiRequest) SetVSwitchId(v string) *CreateEaisEiRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaisEiRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEaisEiRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaisEiRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaisEiRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaisEiRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaisEiRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaisEiRequestTag) SetKey(v string) *CreateEaisEiRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaisEiRequestTag) SetValue(v string) *CreateEaisEiRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaisEiRequestTag) Validate() error {
	return dara.Validate(s)
}
