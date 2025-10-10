// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v []*string) *ListAclsRequest
	GetAclIds() []*string
	SetAclNames(v []*string) *ListAclsRequest
	GetAclNames() []*string
	SetMaxResults(v int32) *ListAclsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAclsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListAclsRequest
	GetResourceGroupId() *string
	SetTag(v []*ListAclsRequestTag) *ListAclsRequest
	GetTag() []*ListAclsRequestTag
}

type ListAclsRequest struct {
	// Filter access control lists (ACLs) by ACL ID. You can specify at most 20 ACL IDs in each call.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The ACL names. You can specify up to 10 ACL names in each call.
	AclNames []*string `json:"AclNames,omitempty" xml:"AclNames,omitempty" type:"Repeated"`
	// The maximum number of entries to return. This parameter is optional. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query or no next query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group. You can filter the query results based on the specified ID.
	//
	// example:
	//
	// rg-atstuj3rtopty****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListAclsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAclsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAclsRequest) GoString() string {
	return s.String()
}

func (s *ListAclsRequest) GetAclIds() []*string {
	return s.AclIds
}

func (s *ListAclsRequest) GetAclNames() []*string {
	return s.AclNames
}

func (s *ListAclsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAclsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAclsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAclsRequest) GetTag() []*ListAclsRequestTag {
	return s.Tag
}

func (s *ListAclsRequest) SetAclIds(v []*string) *ListAclsRequest {
	s.AclIds = v
	return s
}

func (s *ListAclsRequest) SetAclNames(v []*string) *ListAclsRequest {
	s.AclNames = v
	return s
}

func (s *ListAclsRequest) SetMaxResults(v int32) *ListAclsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAclsRequest) SetNextToken(v string) *ListAclsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclsRequest) SetResourceGroupId(v string) *ListAclsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAclsRequest) SetTag(v []*ListAclsRequestTag) *ListAclsRequest {
	s.Tag = v
	return s
}

func (s *ListAclsRequest) Validate() error {
	return dara.Validate(s)
}

type ListAclsRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAclsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAclsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAclsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAclsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAclsRequestTag) SetKey(v string) *ListAclsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAclsRequestTag) SetValue(v string) *ListAclsRequestTag {
	s.Value = &v
	return s
}

func (s *ListAclsRequestTag) Validate() error {
	return dara.Validate(s)
}
