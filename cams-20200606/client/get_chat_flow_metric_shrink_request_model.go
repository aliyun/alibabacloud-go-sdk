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
	// example:
	//
	// 示例值示例值示例值
	BizCode         *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// 示例值示例值
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// example:
	//
	// 示例值示例值
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// example:
	//
	// 31
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	MetricName           *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricParamShrink    *string `json:"MetricParam,omitempty" xml:"MetricParam,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 81
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
