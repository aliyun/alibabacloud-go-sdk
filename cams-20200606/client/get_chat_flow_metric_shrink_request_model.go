// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFlowMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *GetChatFlowMetricShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *GetChatFlowMetricShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *GetChatFlowMetricShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *GetChatFlowMetricShrinkRequest
	GetFlowVersion() *string
	SetFrom(v int64) *GetChatFlowMetricShrinkRequest
	GetFrom() *int64
	SetMetricName(v string) *GetChatFlowMetricShrinkRequest
	GetMetricName() *string
	SetMetricParamShrink(v string) *GetChatFlowMetricShrinkRequest
	GetMetricParamShrink() *string
	SetOwnerId(v int64) *GetChatFlowMetricShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetChatFlowMetricShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatFlowMetricShrinkRequest
	GetResourceOwnerId() *int64
	SetTo(v int64) *GetChatFlowMetricShrinkRequest
	GetTo() *int64
}

type GetChatFlowMetricShrinkRequest struct {
	// Business tenant code, default is “ALICOM_OPAAS”.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Business extension information, default is “{}”.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// Flow code.
	//
	// example:
	//
	// f4912c16943b4dfba44bd6fedacf****
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// Flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// Start timestamp in seconds.
	//
	// example:
	//
	// 1751299200
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Metric name.
	//
	// This parameter is required.
	//
	// example:
	//
	// nodeUsageStatistics
	MetricName           *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricParamShrink    *string `json:"MetricParam,omitempty" xml:"MetricParam,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// End timestamp in seconds.
	//
	// example:
	//
	// 1751385599
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetChatFlowMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetChatFlowMetricShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *GetChatFlowMetricShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *GetChatFlowMetricShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *GetChatFlowMetricShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *GetChatFlowMetricShrinkRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetChatFlowMetricShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetChatFlowMetricShrinkRequest) GetMetricParamShrink() *string {
	return s.MetricParamShrink
}

func (s *GetChatFlowMetricShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatFlowMetricShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatFlowMetricShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatFlowMetricShrinkRequest) GetTo() *int64 {
	return s.To
}

func (s *GetChatFlowMetricShrinkRequest) SetBizCode(v string) *GetChatFlowMetricShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetBizExtendShrink(v string) *GetChatFlowMetricShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetFlowCode(v string) *GetChatFlowMetricShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetFlowVersion(v string) *GetChatFlowMetricShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetFrom(v int64) *GetChatFlowMetricShrinkRequest {
	s.From = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetMetricName(v string) *GetChatFlowMetricShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetMetricParamShrink(v string) *GetChatFlowMetricShrinkRequest {
	s.MetricParamShrink = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetOwnerId(v int64) *GetChatFlowMetricShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetResourceOwnerAccount(v string) *GetChatFlowMetricShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetResourceOwnerId(v int64) *GetChatFlowMetricShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) SetTo(v int64) *GetChatFlowMetricShrinkRequest {
	s.To = &v
	return s
}

func (s *GetChatFlowMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
