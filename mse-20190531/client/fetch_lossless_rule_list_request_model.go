// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchLosslessRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *FetchLosslessRuleListRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *FetchLosslessRuleListRequest
	GetAppId() *string
	SetAppName(v string) *FetchLosslessRuleListRequest
	GetAppName() *string
	SetNamespace(v string) *FetchLosslessRuleListRequest
	GetNamespace() *string
	SetPageNumber(v int32) *FetchLosslessRuleListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *FetchLosslessRuleListRequest
	GetPageSize() *int32
	SetRegionId(v string) *FetchLosslessRuleListRequest
	GetRegionId() *string
}

type FetchLosslessRuleListRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// hyaziyb6sc@86827c61f5ed8fc
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// spring-boot-sample
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s FetchLosslessRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchLosslessRuleListRequest) GoString() string {
	return s.String()
}

func (s *FetchLosslessRuleListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *FetchLosslessRuleListRequest) GetAppId() *string {
	return s.AppId
}

func (s *FetchLosslessRuleListRequest) GetAppName() *string {
	return s.AppName
}

func (s *FetchLosslessRuleListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *FetchLosslessRuleListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *FetchLosslessRuleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FetchLosslessRuleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *FetchLosslessRuleListRequest) SetAcceptLanguage(v string) *FetchLosslessRuleListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *FetchLosslessRuleListRequest) SetAppId(v string) *FetchLosslessRuleListRequest {
	s.AppId = &v
	return s
}

func (s *FetchLosslessRuleListRequest) SetAppName(v string) *FetchLosslessRuleListRequest {
	s.AppName = &v
	return s
}

func (s *FetchLosslessRuleListRequest) SetNamespace(v string) *FetchLosslessRuleListRequest {
	s.Namespace = &v
	return s
}

func (s *FetchLosslessRuleListRequest) SetPageNumber(v int32) *FetchLosslessRuleListRequest {
	s.PageNumber = &v
	return s
}

func (s *FetchLosslessRuleListRequest) SetPageSize(v int32) *FetchLosslessRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *FetchLosslessRuleListRequest) SetRegionId(v string) *FetchLosslessRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *FetchLosslessRuleListRequest) Validate() error {
	return dara.Validate(s)
}
