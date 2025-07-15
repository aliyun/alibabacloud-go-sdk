// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiProductsByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *DescribeApiProductsByAppRequest
	GetAppId() *int64
	SetPageNumber(v int32) *DescribeApiProductsByAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiProductsByAppRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApiProductsByAppRequest
	GetSecurityToken() *string
}

type DescribeApiProductsByAppRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110962435
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiProductsByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductsByAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiProductsByAppRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeApiProductsByAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiProductsByAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiProductsByAppRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiProductsByAppRequest) SetAppId(v int64) *DescribeApiProductsByAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApiProductsByAppRequest) SetPageNumber(v int32) *DescribeApiProductsByAppRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiProductsByAppRequest) SetPageSize(v int32) *DescribeApiProductsByAppRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiProductsByAppRequest) SetSecurityToken(v string) *DescribeApiProductsByAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiProductsByAppRequest) Validate() error {
	return dara.Validate(s)
}
