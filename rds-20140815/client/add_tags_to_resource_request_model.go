// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsToResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*AddTagsToResourceRequestTag) *AddTagsToResourceRequest
	GetTag() []*AddTagsToResourceRequestTag
	SetClientToken(v string) *AddTagsToResourceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *AddTagsToResourceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *AddTagsToResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddTagsToResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddTagsToResourceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AddTagsToResourceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AddTagsToResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTagsToResourceRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *AddTagsToResourceRequest
	GetTags() *string
	SetProxyId(v string) *AddTagsToResourceRequest
	GetProxyId() *string
}

type AddTagsToResourceRequest struct {
	Tag []*AddTagsToResourceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// > You can specify up to 30 instance IDs for a batch operation. Separate multiple instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags to bind, including TagKey and TagValue. You can specify up to 5 pairs at a time. Format: {"key1":"value1","key2":"value2"...}.
	//
	// > TagKey cannot be empty, but TagValue can be empty.
	//
	// example:
	//
	// {"key1":"value1","key2":""}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the proxy mode.
	//
	// example:
	//
	// API
	ProxyId *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s AddTagsToResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagsToResourceRequest) GoString() string {
	return s.String()
}

func (s *AddTagsToResourceRequest) GetTag() []*AddTagsToResourceRequestTag {
	return s.Tag
}

func (s *AddTagsToResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddTagsToResourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AddTagsToResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddTagsToResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTagsToResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTagsToResourceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddTagsToResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTagsToResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTagsToResourceRequest) GetTags() *string {
	return s.Tags
}

func (s *AddTagsToResourceRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *AddTagsToResourceRequest) SetTag(v []*AddTagsToResourceRequestTag) *AddTagsToResourceRequest {
	s.Tag = v
	return s
}

func (s *AddTagsToResourceRequest) SetClientToken(v string) *AddTagsToResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddTagsToResourceRequest) SetDBInstanceId(v string) *AddTagsToResourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetOwnerAccount(v string) *AddTagsToResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTagsToResourceRequest) SetOwnerId(v int64) *AddTagsToResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetRegionId(v string) *AddTagsToResourceRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetResourceGroupId(v string) *AddTagsToResourceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetResourceOwnerAccount(v string) *AddTagsToResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTagsToResourceRequest) SetResourceOwnerId(v int64) *AddTagsToResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetTags(v string) *AddTagsToResourceRequest {
	s.Tags = &v
	return s
}

func (s *AddTagsToResourceRequest) SetProxyId(v string) *AddTagsToResourceRequest {
	s.ProxyId = &v
	return s
}

func (s *AddTagsToResourceRequest) Validate() error {
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

type AddTagsToResourceRequestTag struct {
	// The key of the first tag. Each tag consists of a TagKey and a TagValue. You can specify up to 5 pairs at a time. TagKey cannot be empty, but TagValue can be empty.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the first tag. Each tag consists of a TagKey and a TagValue. You can specify up to 5 pairs at a time. TagKey cannot be empty, but TagValue can be empty.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddTagsToResourceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddTagsToResourceRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsToResourceRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddTagsToResourceRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddTagsToResourceRequestTag) SetKey(v string) *AddTagsToResourceRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsToResourceRequestTag) SetValue(v string) *AddTagsToResourceRequestTag {
	s.Value = &v
	return s
}

func (s *AddTagsToResourceRequestTag) Validate() error {
	return dara.Validate(s)
}
