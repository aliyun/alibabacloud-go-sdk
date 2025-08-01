// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetApplicationListRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetApplicationListRequest
	GetAppId() *string
	SetAppName(v string) *GetApplicationListRequest
	GetAppName() *string
	SetLanguage(v string) *GetApplicationListRequest
	GetLanguage() *string
	SetNamespace(v string) *GetApplicationListRequest
	GetNamespace() *string
	SetPageNumber(v int32) *GetApplicationListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetApplicationListRequest
	GetPageSize() *int32
	SetRegion(v string) *GetApplicationListRequest
	GetRegion() *string
	SetSentinelEnable(v bool) *GetApplicationListRequest
	GetSentinelEnable() *bool
	SetSource(v string) *GetApplicationListRequest
	GetSource() *string
	SetSwitchEnable(v bool) *GetApplicationListRequest
	GetSwitchEnable() *bool
	SetTags(v []*GetApplicationListRequestTags) *GetApplicationListRequest
	GetTags() []*GetApplicationListRequestTags
}

type GetApplicationListRequest struct {
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
	SwitchEnable *bool                            `json:"SwitchEnable,omitempty" xml:"SwitchEnable,omitempty"`
	Tags         []*GetApplicationListRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetApplicationListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationListRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetApplicationListRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationListRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationListRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetApplicationListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetApplicationListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetApplicationListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetApplicationListRequest) GetRegion() *string {
	return s.Region
}

func (s *GetApplicationListRequest) GetSentinelEnable() *bool {
	return s.SentinelEnable
}

func (s *GetApplicationListRequest) GetSource() *string {
	return s.Source
}

func (s *GetApplicationListRequest) GetSwitchEnable() *bool {
	return s.SwitchEnable
}

func (s *GetApplicationListRequest) GetTags() []*GetApplicationListRequestTags {
	return s.Tags
}

func (s *GetApplicationListRequest) SetAcceptLanguage(v string) *GetApplicationListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetApplicationListRequest) SetAppId(v string) *GetApplicationListRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationListRequest) SetAppName(v string) *GetApplicationListRequest {
	s.AppName = &v
	return s
}

func (s *GetApplicationListRequest) SetLanguage(v string) *GetApplicationListRequest {
	s.Language = &v
	return s
}

func (s *GetApplicationListRequest) SetNamespace(v string) *GetApplicationListRequest {
	s.Namespace = &v
	return s
}

func (s *GetApplicationListRequest) SetPageNumber(v int32) *GetApplicationListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetApplicationListRequest) SetPageSize(v int32) *GetApplicationListRequest {
	s.PageSize = &v
	return s
}

func (s *GetApplicationListRequest) SetRegion(v string) *GetApplicationListRequest {
	s.Region = &v
	return s
}

func (s *GetApplicationListRequest) SetSentinelEnable(v bool) *GetApplicationListRequest {
	s.SentinelEnable = &v
	return s
}

func (s *GetApplicationListRequest) SetSource(v string) *GetApplicationListRequest {
	s.Source = &v
	return s
}

func (s *GetApplicationListRequest) SetSwitchEnable(v bool) *GetApplicationListRequest {
	s.SwitchEnable = &v
	return s
}

func (s *GetApplicationListRequest) SetTags(v []*GetApplicationListRequestTags) *GetApplicationListRequest {
	s.Tags = v
	return s
}

func (s *GetApplicationListRequest) Validate() error {
	return dara.Validate(s)
}

type GetApplicationListRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetApplicationListRequestTags) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationListRequestTags) GoString() string {
	return s.String()
}

func (s *GetApplicationListRequestTags) GetKey() *string {
	return s.Key
}

func (s *GetApplicationListRequestTags) GetValue() *string {
	return s.Value
}

func (s *GetApplicationListRequestTags) SetKey(v string) *GetApplicationListRequestTags {
	s.Key = &v
	return s
}

func (s *GetApplicationListRequestTags) SetValue(v string) *GetApplicationListRequestTags {
	s.Value = &v
	return s
}

func (s *GetApplicationListRequestTags) Validate() error {
	return dara.Validate(s)
}
