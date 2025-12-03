// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiJupyterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaiJupyterShrinkRequest
	GetClientToken() *string
	SetEaisName(v string) *CreateEaiJupyterShrinkRequest
	GetEaisName() *string
	SetEaisType(v string) *CreateEaiJupyterShrinkRequest
	GetEaisType() *string
	SetEnvironmentVarShrink(v string) *CreateEaiJupyterShrinkRequest
	GetEnvironmentVarShrink() *string
	SetRegionId(v string) *CreateEaiJupyterShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaiJupyterShrinkRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaiJupyterShrinkRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaiJupyterShrinkRequestTag) *CreateEaiJupyterShrinkRequest
	GetTag() []*CreateEaiJupyterShrinkRequestTag
	SetVSwitchId(v string) *CreateEaiJupyterShrinkRequest
	GetVSwitchId() *string
}

type CreateEaiJupyterShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EaisName    *string `json:"EaisName,omitempty" xml:"EaisName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais.ei-a6.2xlarge
	EaisType             *string `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EnvironmentVarShrink *string `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                             `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiJupyterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiJupyterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiJupyterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaiJupyterShrinkRequest) GetEaisName() *string {
	return s.EaisName
}

func (s *CreateEaiJupyterShrinkRequest) GetEaisType() *string {
	return s.EaisType
}

func (s *CreateEaiJupyterShrinkRequest) GetEnvironmentVarShrink() *string {
	return s.EnvironmentVarShrink
}

func (s *CreateEaiJupyterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaiJupyterShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaiJupyterShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaiJupyterShrinkRequest) GetTag() []*CreateEaiJupyterShrinkRequestTag {
	return s.Tag
}

func (s *CreateEaiJupyterShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaiJupyterShrinkRequest) SetClientToken(v string) *CreateEaiJupyterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetEaisName(v string) *CreateEaiJupyterShrinkRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetEaisType(v string) *CreateEaiJupyterShrinkRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetEnvironmentVarShrink(v string) *CreateEaiJupyterShrinkRequest {
	s.EnvironmentVarShrink = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetRegionId(v string) *CreateEaiJupyterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetResourceGroupId(v string) *CreateEaiJupyterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetSecurityGroupId(v string) *CreateEaiJupyterShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetTag(v []*CreateEaiJupyterShrinkRequestTag) *CreateEaiJupyterShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) SetVSwitchId(v string) *CreateEaiJupyterShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequest) Validate() error {
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

type CreateEaiJupyterShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiJupyterShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiJupyterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaiJupyterShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaiJupyterShrinkRequestTag) SetKey(v string) *CreateEaiJupyterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequestTag) SetValue(v string) *CreateEaiJupyterShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaiJupyterShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
