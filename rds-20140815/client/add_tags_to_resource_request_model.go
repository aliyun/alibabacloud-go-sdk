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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// >  You can enter up to 30 instance IDs in a single request. If you enter more than one instance ID, you must separate the instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hagnzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that you want to add. Each tag consists of a tag key and a tag value. You can specify a maximum of five tags in the following format for each request: {"key1":"value1","key2":"value2"...}.
	//
	// >  The tag key is required and the tag value is optional.
	//
	// example:
	//
	// {“key1”:”value1”,“key2”:””}
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
	// The tag key of the first tag that you want to add. Each tag consists of a tag key and a tag value. You can specify up to five tags in a single request. You cannot specify an empty string as the tag key. You can specify an empty string as the tag value.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value of the first tag that you want to add. Each tag consists of a tag key and a tag value. You can specify up to five tags in a single request. You cannot specify an empty string as the tag key. You can specify an empty string as the tag value.
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
