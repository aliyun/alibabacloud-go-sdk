// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesForExpressConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *TagResourcesForExpressConnectRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TagResourcesForExpressConnectRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TagResourcesForExpressConnectRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagResourcesForExpressConnectRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *TagResourcesForExpressConnectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TagResourcesForExpressConnectRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *TagResourcesForExpressConnectRequest
	GetResourceType() *string
	SetTag(v []*TagResourcesForExpressConnectRequestTag) *TagResourcesForExpressConnectRequest
	GetTag() []*TagResourcesForExpressConnectRequestTag
}

type TagResourcesForExpressConnectRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the resource is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 20 resource IDs.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **PHYSICALCONNECTION**: Express Connect circuit.
	//
	// 	- **VIRTUALBORDERROUTER**: virtual border router (VBR).
	//
	// 	- **ROUTERINTERFACE**: router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to add to the resource.
	//
	// This parameter is required.
	Tag []*TagResourcesForExpressConnectRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesForExpressConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesForExpressConnectRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesForExpressConnectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TagResourcesForExpressConnectRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagResourcesForExpressConnectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesForExpressConnectRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesForExpressConnectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TagResourcesForExpressConnectRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TagResourcesForExpressConnectRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesForExpressConnectRequest) GetTag() []*TagResourcesForExpressConnectRequestTag {
	return s.Tag
}

func (s *TagResourcesForExpressConnectRequest) SetOwnerAccount(v string) *TagResourcesForExpressConnectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesForExpressConnectRequest) SetOwnerId(v int64) *TagResourcesForExpressConnectRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesForExpressConnectRequest) SetRegionId(v string) *TagResourcesForExpressConnectRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesForExpressConnectRequest) SetResourceId(v []*string) *TagResourcesForExpressConnectRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesForExpressConnectRequest) SetResourceOwnerAccount(v string) *TagResourcesForExpressConnectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesForExpressConnectRequest) SetResourceOwnerId(v int64) *TagResourcesForExpressConnectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesForExpressConnectRequest) SetResourceType(v string) *TagResourcesForExpressConnectRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesForExpressConnectRequest) SetTag(v []*TagResourcesForExpressConnectRequestTag) *TagResourcesForExpressConnectRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesForExpressConnectRequest) Validate() error {
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

type TagResourcesForExpressConnectRequestTag struct {
	// The key of the tag to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesForExpressConnectRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesForExpressConnectRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesForExpressConnectRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagResourcesForExpressConnectRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagResourcesForExpressConnectRequestTag) SetKey(v string) *TagResourcesForExpressConnectRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesForExpressConnectRequestTag) SetValue(v string) *TagResourcesForExpressConnectRequestTag {
	s.Value = &v
	return s
}

func (s *TagResourcesForExpressConnectRequestTag) Validate() error {
	return dara.Validate(s)
}
