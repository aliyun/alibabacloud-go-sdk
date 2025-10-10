// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagValuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetResourceId(v string) *ListTagValuesRequest
	GetResourceId() *string
	SetResourceType(v string) *ListTagValuesRequest
	GetResourceType() *string
	SetTagKey(v string) *ListTagValuesRequest
	GetTagKey() *string
}

type ListTagValuesRequest struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// eip-resource-test
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **loadbalancer**: an Application Load Balancer (ALB) instance
	//
	// 	- **acl**: an access control list (ACL)
	//
	// 	- **securitypolicy**: a security policy
	//
	// 	- **servergroup**: a server group
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagValuesRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceId(v string) *ListTagValuesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
