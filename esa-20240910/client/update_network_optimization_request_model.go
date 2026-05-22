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
	// This parameter is required.
	ConfigId    *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Grpc        *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	Http2Origin *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	Rule        *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable  *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence    *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId            *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SmartRouting      *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	Websocket         *string `json:"Websocket,omitempty" xml:"Websocket,omitempty"`
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
