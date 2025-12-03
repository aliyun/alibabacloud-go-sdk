// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEaiEcsRequest
	GetClientToken() *string
	SetEaisName(v string) *CreateEaiEcsRequest
	GetEaisName() *string
	SetEaisType(v string) *CreateEaiEcsRequest
	GetEaisType() *string
	SetEcs(v *CreateEaiEcsRequestEcs) *CreateEaiEcsRequest
	GetEcs() *CreateEaiEcsRequestEcs
	SetRegionId(v string) *CreateEaiEcsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEaiEcsRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateEaiEcsRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateEaiEcsRequestTag) *CreateEaiEcsRequest
	GetTag() []*CreateEaiEcsRequestTag
	SetVSwitchId(v string) *CreateEaiEcsRequest
	GetVSwitchId() *string
}

type CreateEaiEcsRequest struct {
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
	EaisType *string                 `json:"EaisType,omitempty" xml:"EaisType,omitempty"`
	Ecs      *CreateEaiEcsRequestEcs `json:"Ecs,omitempty" xml:"Ecs,omitempty" type:"Struct"`
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
	SecurityGroupId *string                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag             []*CreateEaiEcsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6h3rbwbm90urjwa****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateEaiEcsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEcsRequest) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEaiEcsRequest) GetEaisName() *string {
	return s.EaisName
}

func (s *CreateEaiEcsRequest) GetEaisType() *string {
	return s.EaisType
}

func (s *CreateEaiEcsRequest) GetEcs() *CreateEaiEcsRequestEcs {
	return s.Ecs
}

func (s *CreateEaiEcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEaiEcsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEaiEcsRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEaiEcsRequest) GetTag() []*CreateEaiEcsRequestTag {
	return s.Tag
}

func (s *CreateEaiEcsRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEaiEcsRequest) SetClientToken(v string) *CreateEaiEcsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEaiEcsRequest) SetEaisName(v string) *CreateEaiEcsRequest {
	s.EaisName = &v
	return s
}

func (s *CreateEaiEcsRequest) SetEaisType(v string) *CreateEaiEcsRequest {
	s.EaisType = &v
	return s
}

func (s *CreateEaiEcsRequest) SetEcs(v *CreateEaiEcsRequestEcs) *CreateEaiEcsRequest {
	s.Ecs = v
	return s
}

func (s *CreateEaiEcsRequest) SetRegionId(v string) *CreateEaiEcsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEaiEcsRequest) SetResourceGroupId(v string) *CreateEaiEcsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEaiEcsRequest) SetSecurityGroupId(v string) *CreateEaiEcsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEaiEcsRequest) SetTag(v []*CreateEaiEcsRequestTag) *CreateEaiEcsRequest {
	s.Tag = v
	return s
}

func (s *CreateEaiEcsRequest) SetVSwitchId(v string) *CreateEaiEcsRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEaiEcsRequest) Validate() error {
	if s.Ecs != nil {
		if err := s.Ecs.Validate(); err != nil {
			return err
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

type CreateEaiEcsRequestEcs struct {
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200324.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 10
	InternetMaxBandwidthIn *string `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// example:
	//
	// 10
	InternetMaxBandwidthOut *string `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// example:
	//
	// 40
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// example:
	//
	// ecs.g7.4xlarge
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateEaiEcsRequestEcs) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEcsRequestEcs) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsRequestEcs) GetImageId() *string {
	return s.ImageId
}

func (s *CreateEaiEcsRequestEcs) GetInternetMaxBandwidthIn() *string {
	return s.InternetMaxBandwidthIn
}

func (s *CreateEaiEcsRequestEcs) GetInternetMaxBandwidthOut() *string {
	return s.InternetMaxBandwidthOut
}

func (s *CreateEaiEcsRequestEcs) GetName() *string {
	return s.Name
}

func (s *CreateEaiEcsRequestEcs) GetPassword() *string {
	return s.Password
}

func (s *CreateEaiEcsRequestEcs) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateEaiEcsRequestEcs) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *CreateEaiEcsRequestEcs) GetType() *string {
	return s.Type
}

func (s *CreateEaiEcsRequestEcs) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateEaiEcsRequestEcs) SetImageId(v string) *CreateEaiEcsRequestEcs {
	s.ImageId = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetInternetMaxBandwidthIn(v string) *CreateEaiEcsRequestEcs {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetInternetMaxBandwidthOut(v string) *CreateEaiEcsRequestEcs {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetName(v string) *CreateEaiEcsRequestEcs {
	s.Name = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetPassword(v string) *CreateEaiEcsRequestEcs {
	s.Password = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetSystemDiskCategory(v string) *CreateEaiEcsRequestEcs {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetSystemDiskSize(v int64) *CreateEaiEcsRequestEcs {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetType(v string) *CreateEaiEcsRequestEcs {
	s.Type = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) SetZoneId(v string) *CreateEaiEcsRequestEcs {
	s.ZoneId = &v
	return s
}

func (s *CreateEaiEcsRequestEcs) Validate() error {
	return dara.Validate(s)
}

type CreateEaiEcsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEaiEcsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEcsRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEaiEcsRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEaiEcsRequestTag) SetKey(v string) *CreateEaiEcsRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEaiEcsRequestTag) SetValue(v string) *CreateEaiEcsRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEaiEcsRequestTag) Validate() error {
	return dara.Validate(s)
}
