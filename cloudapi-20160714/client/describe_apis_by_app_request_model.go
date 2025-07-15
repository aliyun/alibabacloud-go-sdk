// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DescribeApisByAppRequest
	GetApiName() *string
	SetApiUid(v string) *DescribeApisByAppRequest
	GetApiUid() *string
	SetAppId(v int64) *DescribeApisByAppRequest
	GetAppId() *int64
	SetDescription(v string) *DescribeApisByAppRequest
	GetDescription() *string
	SetMethod(v string) *DescribeApisByAppRequest
	GetMethod() *string
	SetPageNumber(v int32) *DescribeApisByAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByAppRequest
	GetPageSize() *int32
	SetPath(v string) *DescribeApisByAppRequest
	GetPath() *string
	SetSecurityToken(v string) *DescribeApisByAppRequest
	GetSecurityToken() *string
}

type DescribeApisByAppRequest struct {
	// The name of the API. The name is used for fuzzy match.
	//
	// example:
	//
	// getPersonInfo
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The API ID.
	//
	// example:
	//
	// b19240592b1b4e74961fb8438ed7550c
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 333486644
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request HTTP method of the API.
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The number of the current page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /tt
	Path          *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApisByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisByAppRequest) GetApiUid() *string {
	return s.ApiUid
}

func (s *DescribeApisByAppRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeApisByAppRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisByAppRequest) GetMethod() *string {
	return s.Method
}

func (s *DescribeApisByAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByAppRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeApisByAppRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisByAppRequest) SetApiName(v string) *DescribeApisByAppRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByAppRequest) SetApiUid(v string) *DescribeApisByAppRequest {
	s.ApiUid = &v
	return s
}

func (s *DescribeApisByAppRequest) SetAppId(v int64) *DescribeApisByAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApisByAppRequest) SetDescription(v string) *DescribeApisByAppRequest {
	s.Description = &v
	return s
}

func (s *DescribeApisByAppRequest) SetMethod(v string) *DescribeApisByAppRequest {
	s.Method = &v
	return s
}

func (s *DescribeApisByAppRequest) SetPageNumber(v int32) *DescribeApisByAppRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByAppRequest) SetPageSize(v int32) *DescribeApisByAppRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByAppRequest) SetPath(v string) *DescribeApisByAppRequest {
	s.Path = &v
	return s
}

func (s *DescribeApisByAppRequest) SetSecurityToken(v string) *DescribeApisByAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisByAppRequest) Validate() error {
	return dara.Validate(s)
}
