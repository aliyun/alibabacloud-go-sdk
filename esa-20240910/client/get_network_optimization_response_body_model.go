// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkOptimizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetNetworkOptimizationResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetNetworkOptimizationResponseBody
	GetConfigType() *string
	SetGrpc(v string) *GetNetworkOptimizationResponseBody
	GetGrpc() *string
	SetHttp2Origin(v string) *GetNetworkOptimizationResponseBody
	GetHttp2Origin() *string
	SetRequestId(v string) *GetNetworkOptimizationResponseBody
	GetRequestId() *string
	SetRule(v string) *GetNetworkOptimizationResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetNetworkOptimizationResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetNetworkOptimizationResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetNetworkOptimizationResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetNetworkOptimizationResponseBody
	GetSiteVersion() *int32
	SetSmartRouting(v string) *GetNetworkOptimizationResponseBody
	GetSmartRouting() *string
	SetUploadMaxFilesize(v string) *GetNetworkOptimizationResponseBody
	GetUploadMaxFilesize() *string
	SetWebsocket(v string) *GetNetworkOptimizationResponseBody
	GetWebsocket() *string
}

type GetNetworkOptimizationResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Value range:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Whether to enable GRPC, default is disabled. Value range:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// example:
	//
	// on
	Grpc *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	// Whether to enable HTTP2 origin, default is disabled. Value range:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// example:
	//
	// on
	Http2Origin *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// Request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
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
	// 2
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site configuration version number. For sites with version management enabled, this parameter can specify the effective site version, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Whether to enable smart routing service, default is disabled. Value range:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// example:
	//
	// on
	SmartRouting *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	// Maximum upload file size in MB, with a range from 100 to 500.
	//
	// example:
	//
	// 500
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	// Whether to enable Websocket, default is enabled. Value range:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// example:
	//
	// on
	Websocket *string `json:"Websocket,omitempty" xml:"Websocket,omitempty"`
}

func (s GetNetworkOptimizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkOptimizationResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetNetworkOptimizationResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetNetworkOptimizationResponseBody) GetGrpc() *string {
	return s.Grpc
}

func (s *GetNetworkOptimizationResponseBody) GetHttp2Origin() *string {
	return s.Http2Origin
}

func (s *GetNetworkOptimizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkOptimizationResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetNetworkOptimizationResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetNetworkOptimizationResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetNetworkOptimizationResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetNetworkOptimizationResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetNetworkOptimizationResponseBody) GetSmartRouting() *string {
	return s.SmartRouting
}

func (s *GetNetworkOptimizationResponseBody) GetUploadMaxFilesize() *string {
	return s.UploadMaxFilesize
}

func (s *GetNetworkOptimizationResponseBody) GetWebsocket() *string {
	return s.Websocket
}

func (s *GetNetworkOptimizationResponseBody) SetConfigId(v int64) *GetNetworkOptimizationResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetConfigType(v string) *GetNetworkOptimizationResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetGrpc(v string) *GetNetworkOptimizationResponseBody {
	s.Grpc = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetHttp2Origin(v string) *GetNetworkOptimizationResponseBody {
	s.Http2Origin = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetRequestId(v string) *GetNetworkOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetRule(v string) *GetNetworkOptimizationResponseBody {
	s.Rule = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetRuleEnable(v string) *GetNetworkOptimizationResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetRuleName(v string) *GetNetworkOptimizationResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetSequence(v int32) *GetNetworkOptimizationResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetSiteVersion(v int32) *GetNetworkOptimizationResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetSmartRouting(v string) *GetNetworkOptimizationResponseBody {
	s.SmartRouting = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetUploadMaxFilesize(v string) *GetNetworkOptimizationResponseBody {
	s.UploadMaxFilesize = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) SetWebsocket(v string) *GetNetworkOptimizationResponseBody {
	s.Websocket = &v
	return s
}

func (s *GetNetworkOptimizationResponseBody) Validate() error {
	return dara.Validate(s)
}
