// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsByApiProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductId(v string) *DescribeAppsByApiProductRequest
	GetApiProductId() *string
	SetAppName(v string) *DescribeAppsByApiProductRequest
	GetAppName() *string
	SetPageNumber(v int32) *DescribeAppsByApiProductRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAppsByApiProductRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeAppsByApiProductRequest
	GetSecurityToken() *string
}

type DescribeAppsByApiProductRequest struct {
	// The ID of the API product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 117b7a64a8b3f064eaa4a47ac62aac5e
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The application name.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppsByApiProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsByApiProductRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiProductRequest) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *DescribeAppsByApiProductRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppsByApiProductRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAppsByApiProductRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppsByApiProductRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAppsByApiProductRequest) SetApiProductId(v string) *DescribeAppsByApiProductRequest {
	s.ApiProductId = &v
	return s
}

func (s *DescribeAppsByApiProductRequest) SetAppName(v string) *DescribeAppsByApiProductRequest {
	s.AppName = &v
	return s
}

func (s *DescribeAppsByApiProductRequest) SetPageNumber(v int32) *DescribeAppsByApiProductRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsByApiProductRequest) SetPageSize(v int32) *DescribeAppsByApiProductRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsByApiProductRequest) SetSecurityToken(v string) *DescribeAppsByApiProductRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppsByApiProductRequest) Validate() error {
	return dara.Validate(s)
}
