// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiJupyterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaiJupyterRequest
	GetClientToken() *string
	SetEaisName(v string) *CreateEaiJupyterRequest
	GetEaisName() *string
	SetEaisType(v string) *CreateEaiJupyterRequest
	GetEaisType() *string
	SetEnvironmentVar(v []*CreateEaiJupyterRequestEnvironmentVar) *CreateEaiJupyterRequest
	GetEnvironmentVar() []*CreateEaiJupyterRequestEnvironmentVar
	SetRegionId(v string) *CreateEaiJupyterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaiJupyterRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaiJupyterRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaiJupyterRequestTag) *CreateEaiJupyterRequest
	GetTag() []*CreateEaiJupyterRequestTag
	SetVSwitchId(v string) *CreateEaiJupyterRequest
	GetVSwitchId() *string
}

type CreateEaiJupyterRequest struct {
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
	EaisType       *string                                  `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	EnvironmentVar []*CreateEaiJupyterRequestEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
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
	SecurityGroupId *string                       `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiJupyterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiJupyterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiJupyterRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaiJupyterRequest) GetEaisName() *string {
	return s.EaisName
}

func (s *CreateEaiJupyterRequest) GetEaisType() *string {
	return s.EaisType
}

func (s *CreateEaiJupyterRequest) GetEnvironmentVar() []*CreateEaiJupyterRequestEnvironmentVar {
	return s.EnvironmentVar
}

func (s *CreateEaiJupyterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaiJupyterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaiJupyterRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaiJupyterRequest) GetTag() []*CreateEaiJupyterRequestTag {
	return s.Tag
}

func (s *CreateEaiJupyterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaiJupyterRequest) SetClientToken(v string) *CreateEaiJupyterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetEaisName(v string) *CreateEaiJupyterRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetEaisType(v string) *CreateEaiJupyterRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetEnvironmentVar(v []*CreateEaiJupyterRequestEnvironmentVar) *CreateEaiJupyterRequest {
	s.EnvironmentVar = v
	return s
}

func (s *CreateEaiJupyterRequest) SetRegionId(v string) *CreateEaiJupyterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetResourceGroupId(v string) *CreateEaiJupyterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetSecurityGroupId(v string) *CreateEaiJupyterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiJupyterRequest) SetTag(v []*CreateEaiJupyterRequestTag) *CreateEaiJupyterRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiJupyterRequest) SetVSwitchId(v string) *CreateEaiJupyterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaiJupyterRequest) Validate() error {
	if s.EnvironmentVar != nil {
		for _, item := range s.EnvironmentVar {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateEaiJupyterRequestEnvironmentVar struct {
	// example:
	//
	// MY_USER_NAME
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiJupyterRequestEnvironmentVar) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiJupyterRequestEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterRequestEnvironmentVar) GetKey() *string {
	return s.Key
}

func (s *CreateEaiJupyterRequestEnvironmentVar) GetValue() *string {
	return s.Value
}

func (s *CreateEaiJupyterRequestEnvironmentVar) SetKey(v string) *CreateEaiJupyterRequestEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateEaiJupyterRequestEnvironmentVar) SetValue(v string) *CreateEaiJupyterRequestEnvironmentVar {
	s.Value = &v
	return s
}

func (s *CreateEaiJupyterRequestEnvironmentVar) Validate() error {
	return dara.Validate(s)
}

type CreateEaiJupyterRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiJupyterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiJupyterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaiJupyterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaiJupyterRequestTag) SetKey(v string) *CreateEaiJupyterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiJupyterRequestTag) SetValue(v string) *CreateEaiJupyterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaiJupyterRequestTag) Validate() error {
	return dara.Validate(s)
}
