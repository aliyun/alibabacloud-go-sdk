// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *DescribeAppsRequest
	GetAppId() *int64
	SetAppOwner(v int64) *DescribeAppsRequest
	GetAppOwner() *int64
	SetPageNumber(v int32) *DescribeAppsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAppsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeAppsRequest
	GetSecurityToken() *string
}

type DescribeAppsRequest struct {
	// The ID of the app.
	//
	// example:
	//
	// 20112314518278
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The Alibaba Cloud account of the app owner. For more information, see [Account Management](https://account.console.aliyun.com/?spm=a2c4g.11186623.2.15.3a8c196eVWxvQB#/secure).
	//
	// example:
	//
	// 1546564
	AppOwner *int64 `json:"AppOwner,omitempty" xml:"AppOwner,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppsRequest) GetAppOwner() *int64 {
	return s.AppOwner
}

func (s *DescribeAppsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAppsRequest) SetAppId(v int64) *DescribeAppsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppsRequest) SetAppOwner(v int64) *DescribeAppsRequest {
	s.AppOwner = &v
	return s
}

func (s *DescribeAppsRequest) SetPageNumber(v int32) *DescribeAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsRequest) SetPageSize(v int32) *DescribeAppsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsRequest) SetSecurityToken(v string) *DescribeAppsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppsRequest) Validate() error {
	return dara.Validate(s)
}
