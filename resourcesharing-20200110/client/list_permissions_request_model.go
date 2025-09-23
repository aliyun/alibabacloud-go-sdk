// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPermissionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPermissionsRequest
	GetNextToken() *string
	SetResourceType(v string) *ListPermissionsRequest
	GetResourceType() *string
}

type ListPermissionsRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPermissionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPermissionsRequest) SetMaxResults(v int32) *ListPermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionsRequest) SetNextToken(v string) *ListPermissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPermissionsRequest) SetResourceType(v string) *ListPermissionsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
