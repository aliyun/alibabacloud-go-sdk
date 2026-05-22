// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkOptimizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrpc(v string) *CreateNetworkOptimizationRequest
	GetGrpc() *string
	SetHttp2Origin(v string) *CreateNetworkOptimizationRequest
	GetHttp2Origin() *string
	SetRule(v string) *CreateNetworkOptimizationRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateNetworkOptimizationRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateNetworkOptimizationRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateNetworkOptimizationRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateNetworkOptimizationRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateNetworkOptimizationRequest
	GetSiteVersion() *int32
	SetSmartRouting(v string) *CreateNetworkOptimizationRequest
	GetSmartRouting() *string
	SetUploadMaxFilesize(v string) *CreateNetworkOptimizationRequest
	GetUploadMaxFilesize() *string
	SetWebsocket(v string) *CreateNetworkOptimizationRequest
	GetWebsocket() *string
}

type CreateNetworkOptimizationRequest struct {
	// Whether to enable GRPC, disabled by default. Possible values:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// example:
	//
	// on
	Grpc *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	// Whether to enable HTTP2 origin, disabled by default. Possible values:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// example:
	//
	// on
	Http2Origin *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
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
	// Rule switch. This parameter is not required when adding a global configuration. Possible values:
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
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 340035003106221
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, this parameter can specify the effective version of the configuration, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Whether to enable smart routing service, disabled by default. Possible values:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// example:
	//
	// on
	SmartRouting *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	// Maximum upload file size in MB, range: 100～500.
	//
	// example:
	//
	// 100
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	// Whether to enable Websocket, enabled by default. Possible values:
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

func (s CreateNetworkOptimizationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkOptimizationRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkOptimizationRequest) GetGrpc() *string {
	return s.Grpc
}

func (s *CreateNetworkOptimizationRequest) GetHttp2Origin() *string {
	return s.Http2Origin
}

func (s *CreateNetworkOptimizationRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateNetworkOptimizationRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateNetworkOptimizationRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateNetworkOptimizationRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateNetworkOptimizationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateNetworkOptimizationRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateNetworkOptimizationRequest) GetSmartRouting() *string {
	return s.SmartRouting
}

func (s *CreateNetworkOptimizationRequest) GetUploadMaxFilesize() *string {
	return s.UploadMaxFilesize
}

func (s *CreateNetworkOptimizationRequest) GetWebsocket() *string {
	return s.Websocket
}

func (s *CreateNetworkOptimizationRequest) SetGrpc(v string) *CreateNetworkOptimizationRequest {
	s.Grpc = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetHttp2Origin(v string) *CreateNetworkOptimizationRequest {
	s.Http2Origin = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetRule(v string) *CreateNetworkOptimizationRequest {
	s.Rule = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetRuleEnable(v string) *CreateNetworkOptimizationRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetRuleName(v string) *CreateNetworkOptimizationRequest {
	s.RuleName = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetSequence(v int32) *CreateNetworkOptimizationRequest {
	s.Sequence = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetSiteId(v int64) *CreateNetworkOptimizationRequest {
	s.SiteId = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetSiteVersion(v int32) *CreateNetworkOptimizationRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetSmartRouting(v string) *CreateNetworkOptimizationRequest {
	s.SmartRouting = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetUploadMaxFilesize(v string) *CreateNetworkOptimizationRequest {
	s.UploadMaxFilesize = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) SetWebsocket(v string) *CreateNetworkOptimizationRequest {
	s.Websocket = &v
	return s
}

func (s *CreateNetworkOptimizationRequest) Validate() error {
	return dara.Validate(s)
}
