// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *DescribeAppAttributesRequest
	GetAppCode() *string
	SetAppId(v int64) *DescribeAppAttributesRequest
	GetAppId() *int64
	SetAppKey(v string) *DescribeAppAttributesRequest
	GetAppKey() *string
	SetAppName(v string) *DescribeAppAttributesRequest
	GetAppName() *string
	SetEnableTagAuth(v bool) *DescribeAppAttributesRequest
	GetEnableTagAuth() *bool
	SetExtend(v string) *DescribeAppAttributesRequest
	GetExtend() *string
	SetPageNumber(v int32) *DescribeAppAttributesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAppAttributesRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeAppAttributesRequest
	GetSecurityToken() *string
	SetSort(v string) *DescribeAppAttributesRequest
	GetSort() *string
	SetTag(v []*DescribeAppAttributesRequestTag) *DescribeAppAttributesRequest
	GetTag() []*DescribeAppAttributesRequestTag
}

type DescribeAppAttributesRequest struct {
	// The AppCode of the app.
	//
	// example:
	//
	// 23552160
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// The ID of the app.
	//
	// example:
	//
	// 20112314518278
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The app key that is used to make an API call.
	//
	// example:
	//
	// 203708622
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the app.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to enable tag verification.
	//
	// example:
	//
	// true
	EnableTagAuth *bool `json:"EnableTagAuth,omitempty" xml:"EnableTagAuth,omitempty"`
	// 扩展信息
	//
	// example:
	//
	// 110243810311
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
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
	// The order. Valid values: asc and desc. Default value: desc.
	//
	// 	- asc: The apps are displayed in ascending order of modification time.
	//
	// 	- desc: The apps are displayed in descending order of modification time.
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*DescribeAppAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAppAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DescribeAppAttributesRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppAttributesRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeAppAttributesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppAttributesRequest) GetEnableTagAuth() *bool {
	return s.EnableTagAuth
}

func (s *DescribeAppAttributesRequest) GetExtend() *string {
	return s.Extend
}

func (s *DescribeAppAttributesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAppAttributesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppAttributesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAppAttributesRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeAppAttributesRequest) GetTag() []*DescribeAppAttributesRequestTag {
	return s.Tag
}

func (s *DescribeAppAttributesRequest) SetAppCode(v string) *DescribeAppAttributesRequest {
	s.AppCode = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetAppId(v int64) *DescribeAppAttributesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetAppKey(v string) *DescribeAppAttributesRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetAppName(v string) *DescribeAppAttributesRequest {
	s.AppName = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetEnableTagAuth(v bool) *DescribeAppAttributesRequest {
	s.EnableTagAuth = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetExtend(v string) *DescribeAppAttributesRequest {
	s.Extend = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetPageNumber(v int32) *DescribeAppAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetPageSize(v int32) *DescribeAppAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetSecurityToken(v string) *DescribeAppAttributesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetSort(v string) *DescribeAppAttributesRequest {
	s.Sort = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetTag(v []*DescribeAppAttributesRequestTag) *DescribeAppAttributesRequest {
	s.Tag = v
	return s
}

func (s *DescribeAppAttributesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppAttributesRequestTag struct {
	// The key of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// This parameter is required.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The key of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// example:
	//
	// \\" \\"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAppAttributesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAppAttributesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAppAttributesRequestTag) SetKey(v string) *DescribeAppAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAppAttributesRequestTag) SetValue(v string) *DescribeAppAttributesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAppAttributesRequestTag) Validate() error {
	return dara.Validate(s)
}
