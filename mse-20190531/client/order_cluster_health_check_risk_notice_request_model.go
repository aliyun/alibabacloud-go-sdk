// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderClusterHealthCheckRiskNoticeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *OrderClusterHealthCheckRiskNoticeRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *OrderClusterHealthCheckRiskNoticeRequest
	GetInstanceId() *string
	SetMute(v bool) *OrderClusterHealthCheckRiskNoticeRequest
	GetMute() *bool
	SetNoticeType(v string) *OrderClusterHealthCheckRiskNoticeRequest
	GetNoticeType() *string
	SetRegionId(v string) *OrderClusterHealthCheckRiskNoticeRequest
	GetRegionId() *string
	SetRequestPars(v string) *OrderClusterHealthCheckRiskNoticeRequest
	GetRequestPars() *string
	SetRiskCode(v string) *OrderClusterHealthCheckRiskNoticeRequest
	GetRiskCode() *string
}

type OrderClusterHealthCheckRiskNoticeRequest struct {
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
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to disable the notification feature if the risk item occurs.
	//
	// 	- true: disabled
	//
	// 	- false: enabled
	//
	// example:
	//
	// false
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The region in which the cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The ID of the risk item.
	//
	// example:
	//
	// 30010010001
	RiskCode *string `json:"RiskCode,omitempty" xml:"RiskCode,omitempty"`
}

func (s OrderClusterHealthCheckRiskNoticeRequest) String() string {
	return dara.Prettify(s)
}

func (s OrderClusterHealthCheckRiskNoticeRequest) GoString() string {
	return s.String()
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) GetMute() *bool {
	return s.Mute
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) GetNoticeType() *string {
	return s.NoticeType
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) GetRiskCode() *string {
	return s.RiskCode
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) SetAcceptLanguage(v string) *OrderClusterHealthCheckRiskNoticeRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) SetInstanceId(v string) *OrderClusterHealthCheckRiskNoticeRequest {
	s.InstanceId = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) SetMute(v bool) *OrderClusterHealthCheckRiskNoticeRequest {
	s.Mute = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) SetNoticeType(v string) *OrderClusterHealthCheckRiskNoticeRequest {
	s.NoticeType = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) SetRegionId(v string) *OrderClusterHealthCheckRiskNoticeRequest {
	s.RegionId = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) SetRequestPars(v string) *OrderClusterHealthCheckRiskNoticeRequest {
	s.RequestPars = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) SetRiskCode(v string) *OrderClusterHealthCheckRiskNoticeRequest {
	s.RiskCode = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeRequest) Validate() error {
	return dara.Validate(s)
}
