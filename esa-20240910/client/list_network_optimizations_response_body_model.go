// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkOptimizationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListNetworkOptimizationsResponseBodyConfigs) *ListNetworkOptimizationsResponseBody
	GetConfigs() []*ListNetworkOptimizationsResponseBodyConfigs
	SetPageNumber(v int32) *ListNetworkOptimizationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNetworkOptimizationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNetworkOptimizationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNetworkOptimizationsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListNetworkOptimizationsResponseBody
	GetTotalPage() *int32
}

type ListNetworkOptimizationsResponseBody struct {
	// Response body configurations.
	Configs []*ListNetworkOptimizationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The size of the page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListNetworkOptimizationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkOptimizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkOptimizationsResponseBody) GetConfigs() []*ListNetworkOptimizationsResponseBodyConfigs {
	return s.Configs
}

func (s *ListNetworkOptimizationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNetworkOptimizationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetworkOptimizationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkOptimizationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNetworkOptimizationsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListNetworkOptimizationsResponseBody) SetConfigs(v []*ListNetworkOptimizationsResponseBodyConfigs) *ListNetworkOptimizationsResponseBody {
	s.Configs = v
	return s
}

func (s *ListNetworkOptimizationsResponseBody) SetPageNumber(v int32) *ListNetworkOptimizationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBody) SetPageSize(v int32) *ListNetworkOptimizationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBody) SetRequestId(v string) *ListNetworkOptimizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBody) SetTotalCount(v int32) *ListNetworkOptimizationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBody) SetTotalPage(v int32) *ListNetworkOptimizationsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkOptimizationsResponseBodyConfigs struct {
	// Configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule-based configurations. The value range is as follows:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule-based configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Whether to enable GRPC, default is off. The value range is:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Grpc *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	// Whether to enable HTTP2 origin, defaulting to off. The value range is as follows:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Http2Origin *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. The value range is as follows:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site configuration version number. For sites with version management enabled, this parameter can specify the site version for which the configuration takes effect, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Whether to enable smart routing service, defaulting to off. The value range is as follows:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	SmartRouting *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	// Maximum file size for upload, in MB. The value range is 100 to 500.
	//
	// example:
	//
	// 500
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	// Whether to enable Websocket, enabled by default. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Websocket *string `json:"Websocket,omitempty" xml:"Websocket,omitempty"`
}

func (s ListNetworkOptimizationsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkOptimizationsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetGrpc() *string {
	return s.Grpc
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetHttp2Origin() *string {
	return s.Http2Origin
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetSmartRouting() *string {
	return s.SmartRouting
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetUploadMaxFilesize() *string {
	return s.UploadMaxFilesize
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) GetWebsocket() *string {
	return s.Websocket
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetConfigId(v int64) *ListNetworkOptimizationsResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetConfigType(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetGrpc(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.Grpc = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetHttp2Origin(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.Http2Origin = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetRule(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetRuleEnable(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetRuleName(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetSequence(v int32) *ListNetworkOptimizationsResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetSiteVersion(v int32) *ListNetworkOptimizationsResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetSmartRouting(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.SmartRouting = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetUploadMaxFilesize(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.UploadMaxFilesize = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) SetWebsocket(v string) *ListNetworkOptimizationsResponseBodyConfigs {
	s.Websocket = &v
	return s
}

func (s *ListNetworkOptimizationsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
