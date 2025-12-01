// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceWorkOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *CreateServiceWorkOrderRequest
	GetCreator() *string
	SetCustomerId(v string) *CreateServiceWorkOrderRequest
	GetCustomerId() *string
	SetDurationDay(v string) *CreateServiceWorkOrderRequest
	GetDurationDay() *string
	SetIsAttachment(v string) *CreateServiceWorkOrderRequest
	GetIsAttachment() *string
	SetIsMilestone(v string) *CreateServiceWorkOrderRequest
	GetIsMilestone() *string
	SetIsWorkOrderNotify(v string) *CreateServiceWorkOrderRequest
	GetIsWorkOrderNotify() *string
	SetNotifyDay(v string) *CreateServiceWorkOrderRequest
	GetNotifyDay() *string
	SetNotifyId(v int64) *CreateServiceWorkOrderRequest
	GetNotifyId() *int64
	SetOperateRemark(v string) *CreateServiceWorkOrderRequest
	GetOperateRemark() *string
	SetOperateType(v string) *CreateServiceWorkOrderRequest
	GetOperateType() *string
	SetOperator(v string) *CreateServiceWorkOrderRequest
	GetOperator() *string
	SetOwnerId(v string) *CreateServiceWorkOrderRequest
	GetOwnerId() *string
	SetPriority(v int32) *CreateServiceWorkOrderRequest
	GetPriority() *int32
	SetStartTime(v int64) *CreateServiceWorkOrderRequest
	GetStartTime() *int64
	SetWorkOrderDetail(v string) *CreateServiceWorkOrderRequest
	GetWorkOrderDetail() *string
	SetWorkOrderName(v string) *CreateServiceWorkOrderRequest
	GetWorkOrderName() *string
	SetWorkOrderSource(v string) *CreateServiceWorkOrderRequest
	GetWorkOrderSource() *string
	SetWorkOrderStatus(v string) *CreateServiceWorkOrderRequest
	GetWorkOrderStatus() *string
	SetWorkOrderType(v string) *CreateServiceWorkOrderRequest
	GetWorkOrderType() *string
}

type CreateServiceWorkOrderRequest struct {
	// Creator.
	//
	// This parameter is required.
	//
	// example:
	//
	// 426556
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Customer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1477832102462645
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Duration in days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DurationDay *string `json:"DurationDay,omitempty" xml:"DurationDay,omitempty"`
	// Attachment requirement.
	//
	// This parameter is required.
	//
	// example:
	//
	// Y
	IsAttachment *string `json:"IsAttachment,omitempty" xml:"IsAttachment,omitempty"`
	// Is milestone.
	//
	// example:
	//
	// Y
	IsMilestone *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	// Whether a reminder is needed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Y
	IsWorkOrderNotify *string `json:"IsWorkOrderNotify,omitempty" xml:"IsWorkOrderNotify,omitempty"`
	// Number of days for advance notification.
	//
	// example:
	//
	// 5
	NotifyDay *string `json:"NotifyDay,omitempty" xml:"NotifyDay,omitempty"`
	// Notification ID.
	//
	// example:
	//
	// 10
	NotifyId *int64 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// Operation remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// newly built
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// Operation type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// Operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// 426556
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	OwnerId  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1734788109000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Work order details.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"questionDetail":"测试工单","answerDetail":""}
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	// Work order name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Delivery task of safety monthly report
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// Work order source.
	//
	// This parameter is required.
	//
	// example:
	//
	// Work order migration
	WorkOrderSource *string `json:"WorkOrderSource,omitempty" xml:"WorkOrderSource,omitempty"`
	// Work order status.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNPROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
	// Work order type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MONTH_REPORT
	WorkOrderType *string `json:"WorkOrderType,omitempty" xml:"WorkOrderType,omitempty"`
}

func (s CreateServiceWorkOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceWorkOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderRequest) GetCreator() *string {
	return s.Creator
}

func (s *CreateServiceWorkOrderRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *CreateServiceWorkOrderRequest) GetDurationDay() *string {
	return s.DurationDay
}

func (s *CreateServiceWorkOrderRequest) GetIsAttachment() *string {
	return s.IsAttachment
}

func (s *CreateServiceWorkOrderRequest) GetIsMilestone() *string {
	return s.IsMilestone
}

func (s *CreateServiceWorkOrderRequest) GetIsWorkOrderNotify() *string {
	return s.IsWorkOrderNotify
}

func (s *CreateServiceWorkOrderRequest) GetNotifyDay() *string {
	return s.NotifyDay
}

func (s *CreateServiceWorkOrderRequest) GetNotifyId() *int64 {
	return s.NotifyId
}

func (s *CreateServiceWorkOrderRequest) GetOperateRemark() *string {
	return s.OperateRemark
}

func (s *CreateServiceWorkOrderRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *CreateServiceWorkOrderRequest) GetOperator() *string {
	return s.Operator
}

func (s *CreateServiceWorkOrderRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateServiceWorkOrderRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateServiceWorkOrderRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateServiceWorkOrderRequest) GetWorkOrderDetail() *string {
	return s.WorkOrderDetail
}

func (s *CreateServiceWorkOrderRequest) GetWorkOrderName() *string {
	return s.WorkOrderName
}

func (s *CreateServiceWorkOrderRequest) GetWorkOrderSource() *string {
	return s.WorkOrderSource
}

func (s *CreateServiceWorkOrderRequest) GetWorkOrderStatus() *string {
	return s.WorkOrderStatus
}

func (s *CreateServiceWorkOrderRequest) GetWorkOrderType() *string {
	return s.WorkOrderType
}

func (s *CreateServiceWorkOrderRequest) SetCreator(v string) *CreateServiceWorkOrderRequest {
	s.Creator = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetCustomerId(v string) *CreateServiceWorkOrderRequest {
	s.CustomerId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetDurationDay(v string) *CreateServiceWorkOrderRequest {
	s.DurationDay = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsAttachment(v string) *CreateServiceWorkOrderRequest {
	s.IsAttachment = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsMilestone(v string) *CreateServiceWorkOrderRequest {
	s.IsMilestone = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsWorkOrderNotify(v string) *CreateServiceWorkOrderRequest {
	s.IsWorkOrderNotify = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetNotifyDay(v string) *CreateServiceWorkOrderRequest {
	s.NotifyDay = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetNotifyId(v int64) *CreateServiceWorkOrderRequest {
	s.NotifyId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperateRemark(v string) *CreateServiceWorkOrderRequest {
	s.OperateRemark = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperateType(v string) *CreateServiceWorkOrderRequest {
	s.OperateType = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperator(v string) *CreateServiceWorkOrderRequest {
	s.Operator = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOwnerId(v string) *CreateServiceWorkOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetPriority(v int32) *CreateServiceWorkOrderRequest {
	s.Priority = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetStartTime(v int64) *CreateServiceWorkOrderRequest {
	s.StartTime = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderDetail(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderDetail = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderName(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderName = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderSource(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderSource = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderStatus(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderStatus = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderType(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderType = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) Validate() error {
	return dara.Validate(s)
}
