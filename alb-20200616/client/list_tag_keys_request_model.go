// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListTagKeysRequest
	GetCategory() *string
	SetKeyword(v string) *ListTagKeysRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListTagKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysRequest
	GetNextToken() *string
	SetResourceType(v string) *ListTagKeysRequest
	GetResourceType() *string
}

type ListTagKeysRequest struct {
	// The type of the tag.
	//
	// Valid values: **Custom**, **System**, and **All**.
	//
	// Default value: **All**.
	//
	// example:
	//
	// System
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
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
	// The type of the resource. Valid values:
	//
	// 	- **acl**: an access control list (ACL)
	//
	// 	- **loadbalancer**: an Application Load Balancer (ALB) instance
	//
	// 	- **securitypolicy**: a security policy
	//
	// 	- **servergroup**: a server group
	//
	// This parameter is required.
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) GetCategory() *string {
	return s.Category
}

func (s *ListTagKeysRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTagKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagKeysRequest) SetCategory(v string) *ListTagKeysRequest {
	s.Category = &v
	return s
}

func (s *ListTagKeysRequest) SetKeyword(v string) *ListTagKeysRequest {
	s.Keyword = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
