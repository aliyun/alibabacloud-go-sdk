// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisWithStageNameIntegratedByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DescribeApisWithStageNameIntegratedByAppRequest
	GetApiName() *string
	SetApiUid(v string) *DescribeApisWithStageNameIntegratedByAppRequest
	GetApiUid() *string
	SetAppId(v int64) *DescribeApisWithStageNameIntegratedByAppRequest
	GetAppId() *int64
	SetDescription(v string) *DescribeApisWithStageNameIntegratedByAppRequest
	GetDescription() *string
	SetMethod(v string) *DescribeApisWithStageNameIntegratedByAppRequest
	GetMethod() *string
	SetPageNumber(v int32) *DescribeApisWithStageNameIntegratedByAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisWithStageNameIntegratedByAppRequest
	GetPageSize() *int32
	SetPath(v string) *DescribeApisWithStageNameIntegratedByAppRequest
	GetPath() *string
	SetSecurityToken(v string) *DescribeApisWithStageNameIntegratedByAppRequest
	GetSecurityToken() *string
}

type DescribeApisWithStageNameIntegratedByAppRequest struct {
	// The API name.
	//
	// example:
	//
	// ApiName
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
	// 2386789
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The API description. The description can be up to 200 characters in length.
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
	// The page number of the page to return.
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

func (s DescribeApisWithStageNameIntegratedByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisWithStageNameIntegratedByAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetApiUid() *string {
	return s.ApiUid
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetMethod() *string {
	return s.Method
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetApiName(v string) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetApiUid(v string) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.ApiUid = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetAppId(v int64) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetDescription(v string) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.Description = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetMethod(v string) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.Method = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetPageNumber(v int32) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetPageSize(v int32) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetPath(v string) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.Path = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) SetSecurityToken(v string) *DescribeApisWithStageNameIntegratedByAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppRequest) Validate() error {
	return dara.Validate(s)
}
