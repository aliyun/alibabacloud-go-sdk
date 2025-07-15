// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribePluginApisRequest
	GetApiId() *string
	SetApiName(v string) *DescribePluginApisRequest
	GetApiName() *string
	SetDescription(v string) *DescribePluginApisRequest
	GetDescription() *string
	SetGroupId(v string) *DescribePluginApisRequest
	GetGroupId() *string
	SetMethod(v string) *DescribePluginApisRequest
	GetMethod() *string
	SetPageNumber(v int32) *DescribePluginApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginApisRequest
	GetPageSize() *int32
	SetPath(v string) *DescribePluginApisRequest
	GetPath() *string
	SetPluginId(v string) *DescribePluginApisRequest
	GetPluginId() *string
	SetSecurityToken(v string) *DescribePluginApisRequest
	GetSecurityToken() *string
}

type DescribePluginApisRequest struct {
	// The ID of the API.
	//
	// example:
	//
	// c6b0dd188b0e4e408e12f926********
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// API
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// API description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 231a4bb81ee94da785733c29********
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The request HTTP method of the API.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// Default value:10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /sendVerifyCode
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the gateway plug-in.
	//
	// This parameter is required.
	//
	// example:
	//
	// bf6583efcef44c51adb00c4e********
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePluginApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginApisRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginApisRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribePluginApisRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribePluginApisRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePluginApisRequest) GetMethod() *string {
	return s.Method
}

func (s *DescribePluginApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginApisRequest) GetPath() *string {
	return s.Path
}

func (s *DescribePluginApisRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DescribePluginApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePluginApisRequest) SetApiId(v string) *DescribePluginApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribePluginApisRequest) SetApiName(v string) *DescribePluginApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePluginApisRequest) SetDescription(v string) *DescribePluginApisRequest {
	s.Description = &v
	return s
}

func (s *DescribePluginApisRequest) SetGroupId(v string) *DescribePluginApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePluginApisRequest) SetMethod(v string) *DescribePluginApisRequest {
	s.Method = &v
	return s
}

func (s *DescribePluginApisRequest) SetPageNumber(v int32) *DescribePluginApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginApisRequest) SetPageSize(v int32) *DescribePluginApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePluginApisRequest) SetPath(v string) *DescribePluginApisRequest {
	s.Path = &v
	return s
}

func (s *DescribePluginApisRequest) SetPluginId(v string) *DescribePluginApisRequest {
	s.PluginId = &v
	return s
}

func (s *DescribePluginApisRequest) SetSecurityToken(v string) *DescribePluginApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginApisRequest) Validate() error {
	return dara.Validate(s)
}
