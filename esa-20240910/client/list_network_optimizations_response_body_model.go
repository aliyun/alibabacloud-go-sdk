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
	Configs    []*ListNetworkOptimizationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                         `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	ConfigId          *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType        *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Grpc              *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	Http2Origin       *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	Rule              *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable        *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence          *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion       *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	SmartRouting      *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	Websocket         *string `json:"Websocket,omitempty" xml:"Websocket,omitempty"`
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
