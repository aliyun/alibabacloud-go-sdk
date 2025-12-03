// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEciShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaiEciShrinkRequest
	GetClientToken() *string
	SetEaisName(v string) *CreateEaiEciShrinkRequest
	GetEaisName() *string
	SetEaisType(v string) *CreateEaiEciShrinkRequest
	GetEaisType() *string
	SetEciShrink(v string) *CreateEaiEciShrinkRequest
	GetEciShrink() *string
	SetRegionId(v string) *CreateEaiEciShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaiEciShrinkRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaiEciShrinkRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaiEciShrinkRequestTag) *CreateEaiEciShrinkRequest
	GetTag() []*CreateEaiEciShrinkRequestTag
	SetVSwitchId(v string) *CreateEaiEciShrinkRequest
	GetVSwitchId() *string
}

type CreateEaiEciShrinkRequest struct {
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
	EciShrink *string `json:"Eci,omitempty" xml:"Eci,omitempty"`
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
	Tag             []*CreateEaiEciShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEciShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEciShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaiEciShrinkRequest) GetEaisName() *string {
	return s.EaisName
}

func (s *CreateEaiEciShrinkRequest) GetEaisType() *string {
	return s.EaisType
}

func (s *CreateEaiEciShrinkRequest) GetEciShrink() *string {
	return s.EciShrink
}

func (s *CreateEaiEciShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaiEciShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaiEciShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaiEciShrinkRequest) GetTag() []*CreateEaiEciShrinkRequestTag {
	return s.Tag
}

func (s *CreateEaiEciShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaiEciShrinkRequest) SetClientToken(v string) *CreateEaiEciShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetEaisName(v string) *CreateEaiEciShrinkRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetEaisType(v string) *CreateEaiEciShrinkRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetEciShrink(v string) *CreateEaiEciShrinkRequest {
	s.EciShrink = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetRegionId(v string) *CreateEaiEciShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetResourceGroupId(v string) *CreateEaiEciShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetSecurityGroupId(v string) *CreateEaiEciShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetTag(v []*CreateEaiEciShrinkRequestTag) *CreateEaiEciShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEciShrinkRequest) SetVSwitchId(v string) *CreateEaiEciShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaiEciShrinkRequest) Validate() error {
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

type CreateEaiEciShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEciShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEciShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaiEciShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaiEciShrinkRequestTag) SetKey(v string) *CreateEaiEciShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEciShrinkRequestTag) SetValue(v string) *CreateEaiEciShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaiEciShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
