// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPermissionVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPermissionVersionsRequest
	GetNextToken() *string
	SetPermissionName(v string) *ListPermissionVersionsRequest
	GetPermissionName() *string
}

type ListPermissionVersionsRequest struct {
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
	// The name of the permission.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunRSDefaultPermissionVSwitch
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
}

func (s ListPermissionVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPermissionVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionVersionsRequest) GetPermissionName() *string {
	return s.PermissionName
}

func (s *ListPermissionVersionsRequest) SetMaxResults(v int32) *ListPermissionVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionVersionsRequest) SetNextToken(v string) *ListPermissionVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPermissionVersionsRequest) SetPermissionName(v string) *ListPermissionVersionsRequest {
	s.PermissionName = &v
	return s
}

func (s *ListPermissionVersionsRequest) Validate() error {
	return dara.Validate(s)
}
