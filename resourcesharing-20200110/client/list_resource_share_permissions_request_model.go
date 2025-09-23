// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceSharePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceSharePermissionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceSharePermissionsRequest
	GetNextToken() *string
	SetResourceOwner(v string) *ListResourceSharePermissionsRequest
	GetResourceOwner() *string
	SetResourceShareId(v string) *ListResourceSharePermissionsRequest
	GetResourceShareId() *string
}

type ListResourceSharePermissionsRequest struct {
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
	// The owner of the resource share. Valid values:
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
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s ListResourceSharePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharePermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceSharePermissionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceSharePermissionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceSharePermissionsRequest) GetResourceOwner() *string {
	return s.ResourceOwner
}

func (s *ListResourceSharePermissionsRequest) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *ListResourceSharePermissionsRequest) SetMaxResults(v int32) *ListResourceSharePermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceSharePermissionsRequest) SetNextToken(v string) *ListResourceSharePermissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceSharePermissionsRequest) SetResourceOwner(v string) *ListResourceSharePermissionsRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListResourceSharePermissionsRequest) SetResourceShareId(v string) *ListResourceSharePermissionsRequest {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceSharePermissionsRequest) Validate() error {
	return dara.Validate(s)
}
