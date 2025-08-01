// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetApplicationListShrinkRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetApplicationListShrinkRequest
	GetAppId() *string
	SetAppName(v string) *GetApplicationListShrinkRequest
	GetAppName() *string
	SetLanguage(v string) *GetApplicationListShrinkRequest
	GetLanguage() *string
	SetNamespace(v string) *GetApplicationListShrinkRequest
	GetNamespace() *string
	SetPageNumber(v int32) *GetApplicationListShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetApplicationListShrinkRequest
	GetPageSize() *int32
	SetRegion(v string) *GetApplicationListShrinkRequest
	GetRegion() *string
	SetSentinelEnable(v bool) *GetApplicationListShrinkRequest
	GetSentinelEnable() *bool
	SetSource(v string) *GetApplicationListShrinkRequest
	GetSource() *string
	SetSwitchEnable(v bool) *GetApplicationListShrinkRequest
	GetSwitchEnable() *bool
	SetTagsShrink(v string) *GetApplicationListShrinkRequest
	GetTagsShrink() *string
}

type GetApplicationListShrinkRequest struct {
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
	// The ID of an application.
	//
	// example:
	//
	// hkhon1po62@c3df23522b*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of an application.
	//
	// example:
	//
	// rest-container
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The programming language of the application, such as Java and Go.
	//
	// example:
	//
	// Java
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The microservice namespace to which the application belongs.
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Specifies whether to enable the Sentinel-compatible mode.
	//
	// example:
	//
	// true
	SentinelEnable *bool `json:"SentinelEnable,omitempty" xml:"SentinelEnable,omitempty"`
	// The source of the application. The value is fixed as edasmsc.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Specifies whether to enable switching.
	//
	// example:
	//
	// true
	SwitchEnable *bool   `json:"SwitchEnable,omitempty" xml:"SwitchEnable,omitempty"`
	TagsShrink   *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetApplicationListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationListShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetApplicationListShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationListShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationListShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetApplicationListShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetApplicationListShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetApplicationListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetApplicationListShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *GetApplicationListShrinkRequest) GetSentinelEnable() *bool {
	return s.SentinelEnable
}

func (s *GetApplicationListShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *GetApplicationListShrinkRequest) GetSwitchEnable() *bool {
	return s.SwitchEnable
}

func (s *GetApplicationListShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *GetApplicationListShrinkRequest) SetAcceptLanguage(v string) *GetApplicationListShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetAppId(v string) *GetApplicationListShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetAppName(v string) *GetApplicationListShrinkRequest {
	s.AppName = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetLanguage(v string) *GetApplicationListShrinkRequest {
	s.Language = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetNamespace(v string) *GetApplicationListShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetPageNumber(v int32) *GetApplicationListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetPageSize(v int32) *GetApplicationListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetRegion(v string) *GetApplicationListShrinkRequest {
	s.Region = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetSentinelEnable(v bool) *GetApplicationListShrinkRequest {
	s.SentinelEnable = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetSource(v string) *GetApplicationListShrinkRequest {
	s.Source = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetSwitchEnable(v bool) *GetApplicationListShrinkRequest {
	s.SwitchEnable = &v
	return s
}

func (s *GetApplicationListShrinkRequest) SetTagsShrink(v string) *GetApplicationListShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *GetApplicationListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
