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
	// The business tenant code. Default value: ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The business extension information. Default value: an empty collection.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The flow code. You can view the flow code on the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. On the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page, click the flow name to open the canvas and view the flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The start time. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1751299200
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// The metric name. Valid values:
	//
	// - nodeUsageStatistics: node usage statistics.
	//
	// - nodeErrorDetails: node error details.
	//
	// This parameter is required.
	//
	// example:
	//
	// nodeUsageStatistics
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The extended parameter for the metric query. When MetricName is set to nodeErrorDetails, pass in a JSON string. Valid values for the JSON fields:
	//
	// - pageNo: the current page number.
	//
	// - pageSize: the number of entries per page.
	//
	// - nodeId: the node ID. On the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page, click the flow name to open the canvas and copy the node ID.
	//
	// example:
	//
	// {
	//
	//   "pageNo": 1,
	//
	//   "pageSize": 20,
	//
	//   "nodeId": "SendWhatsAppMessage#H7fKq5rM"
	//
	// }
	MetricParamShrink    *string `json:"MetricParam,omitempty" xml:"MetricParam,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The end time. This value is a UNIX timestamp. Unit: seconds.
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
