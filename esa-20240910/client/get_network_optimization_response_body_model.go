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
	ConfigId          *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType        *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Grpc              *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	Http2Origin       *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rule              *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable        *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence          *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion       *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	SmartRouting      *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	Websocket         *string `json:"Websocket,omitempty" xml:"Websocket,omitempty"`
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
