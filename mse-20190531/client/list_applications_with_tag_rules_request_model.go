// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsWithTagRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListApplicationsWithTagRulesRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ListApplicationsWithTagRulesRequest
	GetAppId() *string
	SetAppName(v string) *ListApplicationsWithTagRulesRequest
	GetAppName() *string
	SetNamespace(v string) *ListApplicationsWithTagRulesRequest
	GetNamespace() *string
	SetPageNumber(v int32) *ListApplicationsWithTagRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationsWithTagRulesRequest
	GetPageSize() *int32
	SetRegion(v string) *ListApplicationsWithTagRulesRequest
	GetRegion() *string
	SetSource(v string) *ListApplicationsWithTagRulesRequest
	GetSource() *string
}

type ListApplicationsWithTagRulesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Deprecated
	//
	// The ID of the application.
	//
	// example:
	//
	// xjpc0h9h4d@xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The MSE namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The source of the routing rule. Default value: edasmsc.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListApplicationsWithTagRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListApplicationsWithTagRulesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsWithTagRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationsWithTagRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListApplicationsWithTagRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationsWithTagRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsWithTagRulesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListApplicationsWithTagRulesRequest) GetSource() *string {
	return s.Source
}

func (s *ListApplicationsWithTagRulesRequest) SetAcceptLanguage(v string) *ListApplicationsWithTagRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) SetAppId(v string) *ListApplicationsWithTagRulesRequest {
	s.AppId = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) SetAppName(v string) *ListApplicationsWithTagRulesRequest {
	s.AppName = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) SetNamespace(v string) *ListApplicationsWithTagRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) SetPageNumber(v int32) *ListApplicationsWithTagRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) SetPageSize(v int32) *ListApplicationsWithTagRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) SetRegion(v string) *ListApplicationsWithTagRulesRequest {
	s.Region = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) SetSource(v string) *ListApplicationsWithTagRulesRequest {
	s.Source = &v
	return s
}

func (s *ListApplicationsWithTagRulesRequest) Validate() error {
	return dara.Validate(s)
}
