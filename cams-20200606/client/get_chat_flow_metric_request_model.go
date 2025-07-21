// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFlowMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *GetChatFlowMetricRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *GetChatFlowMetricRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *GetChatFlowMetricRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *GetChatFlowMetricRequest
	GetFlowVersion() *string
	SetFrom(v int64) *GetChatFlowMetricRequest
	GetFrom() *int64
	SetMetricName(v string) *GetChatFlowMetricRequest
	GetMetricName() *string
	SetMetricParam(v map[string]interface{}) *GetChatFlowMetricRequest
	GetMetricParam() map[string]interface{}
	SetOwnerId(v int64) *GetChatFlowMetricRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetChatFlowMetricRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatFlowMetricRequest
	GetResourceOwnerId() *int64
	SetTo(v int64) *GetChatFlowMetricRequest
	GetTo() *int64
}

type GetChatFlowMetricRequest struct {
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
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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
	MetricName           *string                `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricParam          map[string]interface{} `json:"MetricParam,omitempty" xml:"MetricParam,omitempty"`
	OwnerId              *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// End timestamp in seconds.
	//
	// example:
	//
	// 1751385599
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetChatFlowMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowMetricRequest) GoString() string {
	return s.String()
}

func (s *GetChatFlowMetricRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *GetChatFlowMetricRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *GetChatFlowMetricRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *GetChatFlowMetricRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *GetChatFlowMetricRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetChatFlowMetricRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetChatFlowMetricRequest) GetMetricParam() map[string]interface{} {
	return s.MetricParam
}

func (s *GetChatFlowMetricRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatFlowMetricRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatFlowMetricRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatFlowMetricRequest) GetTo() *int64 {
	return s.To
}

func (s *GetChatFlowMetricRequest) SetBizCode(v string) *GetChatFlowMetricRequest {
	s.BizCode = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetBizExtend(v map[string]interface{}) *GetChatFlowMetricRequest {
	s.BizExtend = v
	return s
}

func (s *GetChatFlowMetricRequest) SetFlowCode(v string) *GetChatFlowMetricRequest {
	s.FlowCode = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetFlowVersion(v string) *GetChatFlowMetricRequest {
	s.FlowVersion = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetFrom(v int64) *GetChatFlowMetricRequest {
	s.From = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetMetricName(v string) *GetChatFlowMetricRequest {
	s.MetricName = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetMetricParam(v map[string]interface{}) *GetChatFlowMetricRequest {
	s.MetricParam = v
	return s
}

func (s *GetChatFlowMetricRequest) SetOwnerId(v int64) *GetChatFlowMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetResourceOwnerAccount(v string) *GetChatFlowMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetResourceOwnerId(v int64) *GetChatFlowMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatFlowMetricRequest) SetTo(v int64) *GetChatFlowMetricRequest {
	s.To = &v
	return s
}

func (s *GetChatFlowMetricRequest) Validate() error {
	return dara.Validate(s)
}
