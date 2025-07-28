// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVpcAuthorizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *DescribeUserVpcAuthorizationsRequest
	GetAuthType() *string
	SetAuthorizedUserId(v int64) *DescribeUserVpcAuthorizationsRequest
	GetAuthorizedUserId() *int64
	SetPageNumber(v int32) *DescribeUserVpcAuthorizationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUserVpcAuthorizationsRequest
	GetPageSize() *int32
}

type DescribeUserVpcAuthorizationsRequest struct {
	// The authorization scope. Valid values:
	//
	// 	- NORMAL: general authorization
	//
	// 	- CLOUD_PRODUCT: cloud service-related authorization
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the Alibaba Cloud account to which the permissions on the resources are granted.
	//
	// example:
	//
	// 141339776561****
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserVpcAuthorizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeUserVpcAuthorizationsRequest) GetAuthorizedUserId() *int64 {
	return s.AuthorizedUserId
}

func (s *DescribeUserVpcAuthorizationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserVpcAuthorizationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserVpcAuthorizationsRequest) SetAuthType(v string) *DescribeUserVpcAuthorizationsRequest {
	s.AuthType = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsRequest) SetAuthorizedUserId(v int64) *DescribeUserVpcAuthorizationsRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsRequest) SetPageNumber(v int32) *DescribeUserVpcAuthorizationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsRequest) SetPageSize(v int32) *DescribeUserVpcAuthorizationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsRequest) Validate() error {
	return dara.Validate(s)
}
