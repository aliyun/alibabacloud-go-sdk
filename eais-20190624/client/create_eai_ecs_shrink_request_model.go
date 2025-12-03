// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEcsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaiEcsShrinkRequest
	GetClientToken() *string
	SetEaisName(v string) *CreateEaiEcsShrinkRequest
	GetEaisName() *string
	SetEaisType(v string) *CreateEaiEcsShrinkRequest
	GetEaisType() *string
	SetEcsShrink(v string) *CreateEaiEcsShrinkRequest
	GetEcsShrink() *string
	SetRegionId(v string) *CreateEaiEcsShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaiEcsShrinkRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaiEcsShrinkRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaiEcsShrinkRequestTag) *CreateEaiEcsShrinkRequest
	GetTag() []*CreateEaiEcsShrinkRequestTag
	SetVSwitchId(v string) *CreateEaiEcsShrinkRequest
	GetVSwitchId() *string
}

type CreateEaiEcsShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// eais-test01
	EaisName *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType  *string `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EcsShrink *string `json:"Ecs,omitempty" xml:"Ecs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmvpuy4a5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiEcsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEcsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEcsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaiEcsShrinkRequest) GetEaisName() *string {
	return s.EaisName
}

func (s *CreateEaiEcsShrinkRequest) GetEaisType() *string {
	return s.EaisType
}

func (s *CreateEaiEcsShrinkRequest) GetEcsShrink() *string {
	return s.EcsShrink
}

func (s *CreateEaiEcsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaiEcsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaiEcsShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaiEcsShrinkRequest) GetTag() []*CreateEaiEcsShrinkRequestTag {
	return s.Tag
}

func (s *CreateEaiEcsShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaiEcsShrinkRequest) SetClientToken(v string) *CreateEaiEcsShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetEaisName(v string) *CreateEaiEcsShrinkRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetEaisType(v string) *CreateEaiEcsShrinkRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetEcsShrink(v string) *CreateEaiEcsShrinkRequest {
	s.EcsShrink = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetRegionId(v string) *CreateEaiEcsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetResourceGroupId(v string) *CreateEaiEcsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetSecurityGroupId(v string) *CreateEaiEcsShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetTag(v []*CreateEaiEcsShrinkRequestTag) *CreateEaiEcsShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEcsShrinkRequest) SetVSwitchId(v string) *CreateEaiEcsShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaiEcsShrinkRequest) Validate() error {
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

type CreateEaiEcsShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEcsShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEcsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaiEcsShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaiEcsShrinkRequestTag) SetKey(v string) *CreateEaiEcsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEcsShrinkRequestTag) SetValue(v string) *CreateEaiEcsShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaiEcsShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
