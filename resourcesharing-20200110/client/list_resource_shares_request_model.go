// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceSharesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceSharesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceSharesRequest
	GetNextToken() *string
	SetPermissionName(v string) *ListResourceSharesRequest
	GetPermissionName() *string
	SetResourceGroupId(v string) *ListResourceSharesRequest
	GetResourceGroupId() *string
	SetResourceOwner(v string) *ListResourceSharesRequest
	GetResourceOwner() *string
	SetResourceShareIds(v []*string) *ListResourceSharesRequest
	GetResourceShareIds() []*string
	SetResourceShareName(v string) *ListResourceSharesRequest
	GetResourceShareName() *string
	SetResourceShareStatus(v string) *ListResourceSharesRequest
	GetResourceShareStatus() *string
	SetTag(v []*ListResourceSharesRequestTag) *ListResourceSharesRequest
	GetTag() []*ListResourceSharesRequestTag
}

type ListResourceSharesRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the permissions. For more information, see [Permission library](https://help.aliyun.com/document_detail/465474.html).
	//
	// example:
	//
	// AliyunRSDefaultPermissionVSwitch
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekz5nlvlak****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The owner of the resource shares. Valid values:
	//
	// 	- Self: the current account
	//
	// 	- OtherAccounts: an account other than the current account
	//
	// This parameter is required.
	//
	// example:
	//
	// Self
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The IDs of the resource shares.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five resource shares can be specified at a time.
	//
	// example:
	//
	// rs-PqysnzIj****
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The name of the resource share.
	//
	// example:
	//
	// test
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The status of the resource shares. Valid values:
	//
	// 	- Active
	//
	// 	- Pending
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// >  The system automatically deletes the records of resource shares in the Deleted state within 48 hours to 96 hours after you delete the resource shares.
	//
	// example:
	//
	// Active
	ResourceShareStatus *string `json:"ResourceShareStatus,omitempty" xml:"ResourceShareStatus,omitempty"`
	// The tags.
	Tag []*ListResourceSharesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceSharesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceSharesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceSharesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceSharesRequest) GetPermissionName() *string {
	return s.PermissionName
}

func (s *ListResourceSharesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceSharesRequest) GetResourceOwner() *string {
	return s.ResourceOwner
}

func (s *ListResourceSharesRequest) GetResourceShareIds() []*string {
	return s.ResourceShareIds
}

func (s *ListResourceSharesRequest) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *ListResourceSharesRequest) GetResourceShareStatus() *string {
	return s.ResourceShareStatus
}

func (s *ListResourceSharesRequest) GetTag() []*ListResourceSharesRequestTag {
	return s.Tag
}

func (s *ListResourceSharesRequest) SetMaxResults(v int32) *ListResourceSharesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceSharesRequest) SetNextToken(v string) *ListResourceSharesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharesRequest) SetPermissionName(v string) *ListResourceSharesRequest {
	s.PermissionName = &v
	return s
}

func (s *ListResourceSharesRequest) SetResourceGroupId(v string) *ListResourceSharesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceSharesRequest) SetResourceOwner(v string) *ListResourceSharesRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListResourceSharesRequest) SetResourceShareIds(v []*string) *ListResourceSharesRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceSharesRequest) SetResourceShareName(v string) *ListResourceSharesRequest {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceSharesRequest) SetResourceShareStatus(v string) *ListResourceSharesRequest {
	s.ResourceShareStatus = &v
	return s
}

func (s *ListResourceSharesRequest) SetTag(v []*ListResourceSharesRequestTag) *ListResourceSharesRequest {
	s.Tag = v
	return s
}

func (s *ListResourceSharesRequest) Validate() error {
	return dara.Validate(s)
}

type ListResourceSharesRequestTag struct {
	// The tag key.
	//
	// >  The tag key can be 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// >  The tag value can be 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceSharesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharesRequestTag) GoString() string {
	return s.String()
}

func (s *ListResourceSharesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListResourceSharesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListResourceSharesRequestTag) SetKey(v string) *ListResourceSharesRequestTag {
	s.Key = &v
	return s
}

func (s *ListResourceSharesRequestTag) SetValue(v string) *ListResourceSharesRequestTag {
	s.Value = &v
	return s
}

func (s *ListResourceSharesRequestTag) Validate() error {
	return dara.Validate(s)
}
