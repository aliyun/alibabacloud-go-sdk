// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkOptimizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateNetworkOptimizationRequest
	GetConfigId() *int64
	SetGrpc(v string) *UpdateNetworkOptimizationRequest
	GetGrpc() *string
	SetHttp2Origin(v string) *UpdateNetworkOptimizationRequest
	GetHttp2Origin() *string
	SetRule(v string) *UpdateNetworkOptimizationRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateNetworkOptimizationRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateNetworkOptimizationRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateNetworkOptimizationRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateNetworkOptimizationRequest
	GetSiteId() *int64
	SetSmartRouting(v string) *UpdateNetworkOptimizationRequest
	GetSmartRouting() *string
	SetUploadMaxFilesize(v string) *UpdateNetworkOptimizationRequest
	GetUploadMaxFilesize() *string
	SetWebsocket(v string) *UpdateNetworkOptimizationRequest
	GetWebsocket() *string
}

type UpdateNetworkOptimizationRequest struct {
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Controls whether gRPC is enabled. This feature is disabled by default. Valid values:
	//
	// - on: gRPC is enabled.
	//
	// - off: gRPC is disabled.
	//
	// example:
	//
	// on
	Grpc *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	// Controls whether HTTP/2 to origin is enabled. This feature is disabled by default. Valid values:
	//
	// - on: HTTP/2 to origin is enabled.
	//
	// - off: HTTP/2 to origin is disabled.
	//
	// example:
	//
	// on
	Http2Origin *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// The conditional expression used to match requests. This parameter is optional for global configurations.
	//
	// - To match all incoming requests, set the value to true.
	//
	// - To match specific requests, set the value to a custom expression, for example, (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Controls whether the rule is enabled. This parameter is optional for global configurations. Valid values:
	//
	// - on: The rule is enabled.
	//
	// - off: The rule is disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is optional for global configurations.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. Smaller values have higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Controls whether smart routing is enabled. This feature is disabled by default. Valid values:
	//
	// - on: Smart routing is enabled.
	//
	// - off: Smart routing is disabled.
	//
	// example:
	//
	// on
	SmartRouting *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	// The maximum upload file size, in MB. The value must be an integer from 100 to 500.
	//
	// example:
	//
	// 100
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	// Controls whether WebSocket is enabled. This feature is enabled by default. Valid values:
	//
	// - on: WebSocket is enabled.
	//
	// - off: WebSocket is disabled.
	//
	// example:
	//
	// on
	Websocket *string `json:"Websocket,omitempty" xml:"Websocket,omitempty"`
}

func (s UpdateNetworkOptimizationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkOptimizationRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkOptimizationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateNetworkOptimizationRequest) GetGrpc() *string {
	return s.Grpc
}

func (s *UpdateNetworkOptimizationRequest) GetHttp2Origin() *string {
	return s.Http2Origin
}

func (s *UpdateNetworkOptimizationRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateNetworkOptimizationRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateNetworkOptimizationRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateNetworkOptimizationRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateNetworkOptimizationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateNetworkOptimizationRequest) GetSmartRouting() *string {
	return s.SmartRouting
}

func (s *UpdateNetworkOptimizationRequest) GetUploadMaxFilesize() *string {
	return s.UploadMaxFilesize
}

func (s *UpdateNetworkOptimizationRequest) GetWebsocket() *string {
	return s.Websocket
}

func (s *UpdateNetworkOptimizationRequest) SetConfigId(v int64) *UpdateNetworkOptimizationRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetGrpc(v string) *UpdateNetworkOptimizationRequest {
	s.Grpc = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetHttp2Origin(v string) *UpdateNetworkOptimizationRequest {
	s.Http2Origin = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetRule(v string) *UpdateNetworkOptimizationRequest {
	s.Rule = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetRuleEnable(v string) *UpdateNetworkOptimizationRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetRuleName(v string) *UpdateNetworkOptimizationRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetSequence(v int32) *UpdateNetworkOptimizationRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetSiteId(v int64) *UpdateNetworkOptimizationRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetSmartRouting(v string) *UpdateNetworkOptimizationRequest {
	s.SmartRouting = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetUploadMaxFilesize(v string) *UpdateNetworkOptimizationRequest {
	s.UploadMaxFilesize = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) SetWebsocket(v string) *UpdateNetworkOptimizationRequest {
	s.Websocket = &v
	return s
}

func (s *UpdateNetworkOptimizationRequest) Validate() error {
	return dara.Validate(s)
}
