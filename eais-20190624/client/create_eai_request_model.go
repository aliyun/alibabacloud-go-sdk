// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaiRequest
	GetClientToken() *string
	SetImage(v string) *CreateEaiRequest
	GetImage() *string
	SetInstanceName(v string) *CreateEaiRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateEaiRequest
	GetInstanceType() *string
	SetRegionId(v string) *CreateEaiRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaiRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaiRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaiRequestTag) *CreateEaiRequest
	GetTag() []*CreateEaiRequestTag
	SetVSwitchId(v string) *CreateEaiRequest
	GetVSwitchId() *string
}

type CreateEaiRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Image       *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// eais-test01
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
	// cn-shenzhen
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string                `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaiRequest) GetImage() *string {
	return s.Image
}

func (s *CreateEaiRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateEaiRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateEaiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaiRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaiRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaiRequest) GetTag() []*CreateEaiRequestTag {
	return s.Tag
}

func (s *CreateEaiRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaiRequest) SetClientToken(v string) *CreateEaiRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiRequest) SetImage(v string) *CreateEaiRequest {
	s.Image = &v
	return s
}

func (s *CreateEaiRequest) SetInstanceName(v string) *CreateEaiRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateEaiRequest) SetInstanceType(v string) *CreateEaiRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateEaiRequest) SetRegionId(v string) *CreateEaiRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiRequest) SetResourceGroupId(v string) *CreateEaiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiRequest) SetSecurityGroupId(v string) *CreateEaiRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiRequest) SetTag(v []*CreateEaiRequestTag) *CreateEaiRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiRequest) SetVSwitchId(v string) *CreateEaiRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaiRequest) Validate() error {
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

type CreateEaiRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaiRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaiRequestTag) SetKey(v string) *CreateEaiRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiRequestTag) SetValue(v string) *CreateEaiRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaiRequestTag) Validate() error {
	return dara.Validate(s)
}
