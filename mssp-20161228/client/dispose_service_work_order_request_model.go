// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposeServiceWorkOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentName(v string) *DisposeServiceWorkOrderRequest
	GetAttachmentName() *string
	SetEndTime(v int64) *DisposeServiceWorkOrderRequest
	GetEndTime() *int64
	SetForwardOwnerId(v string) *DisposeServiceWorkOrderRequest
	GetForwardOwnerId() *string
	SetId(v int64) *DisposeServiceWorkOrderRequest
	GetId() *int64
	SetIsAttachment(v string) *DisposeServiceWorkOrderRequest
	GetIsAttachment() *string
	SetIsWorkOrderNotify(v string) *DisposeServiceWorkOrderRequest
	GetIsWorkOrderNotify() *string
	SetNotifyId(v int64) *DisposeServiceWorkOrderRequest
	GetNotifyId() *int64
	SetOperateRemark(v string) *DisposeServiceWorkOrderRequest
	GetOperateRemark() *string
	SetOperateType(v string) *DisposeServiceWorkOrderRequest
	GetOperateType() *string
	SetOperator(v string) *DisposeServiceWorkOrderRequest
	GetOperator() *string
	SetStartTime(v int64) *DisposeServiceWorkOrderRequest
	GetStartTime() *int64
	SetUpgradeOwnerId(v string) *DisposeServiceWorkOrderRequest
	GetUpgradeOwnerId() *string
	SetWorkOrderDetail(v string) *DisposeServiceWorkOrderRequest
	GetWorkOrderDetail() *string
	SetWorkOrderName(v string) *DisposeServiceWorkOrderRequest
	GetWorkOrderName() *string
	SetWorkOrderStatus(v string) *DisposeServiceWorkOrderRequest
	GetWorkOrderStatus() *string
}

type DisposeServiceWorkOrderRequest struct {
	// Attachment name.
	//
	// example:
	//
	// bbaa133c-0ac2-489f-9fc8-39f91c2e770c_20230301-20240403-服务工单列表.xlsx
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// End time.
	//
	// example:
	//
	// 2024-04-14 00:00:00
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Forward to owner.
	//
	// example:
	//
	// 405639
	ForwardOwnerId *string `json:"ForwardOwnerId,omitempty" xml:"ForwardOwnerId,omitempty"`
	// Work order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23172
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Attachment requirement.
	//
	// example:
	//
	// Y
	IsAttachment *string `json:"IsAttachment,omitempty" xml:"IsAttachment,omitempty"`
	// Work order notification.
	//
	// example:
	//
	// Y
	IsWorkOrderNotify *string `json:"IsWorkOrderNotify,omitempty" xml:"IsWorkOrderNotify,omitempty"`
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
	// 处理完成
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// Processing type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROCESSED
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// Operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// 396120
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Start time.
	//
	// example:
	//
	// 2024-04-02 00:00:00
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Upgrade owner.
	//
	// example:
	//
	// 336333
	UpgradeOwnerId *string `json:"UpgradeOwnerId,omitempty" xml:"UpgradeOwnerId,omitempty"`
	// Work order details.
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
	// 安全产品配置问题与超量提醒
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// Work order status.
	//
	// example:
	//
	// PROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
}

func (s DisposeServiceWorkOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DisposeServiceWorkOrderRequest) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderRequest) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *DisposeServiceWorkOrderRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DisposeServiceWorkOrderRequest) GetForwardOwnerId() *string {
	return s.ForwardOwnerId
}

func (s *DisposeServiceWorkOrderRequest) GetId() *int64 {
	return s.Id
}

func (s *DisposeServiceWorkOrderRequest) GetIsAttachment() *string {
	return s.IsAttachment
}

func (s *DisposeServiceWorkOrderRequest) GetIsWorkOrderNotify() *string {
	return s.IsWorkOrderNotify
}

func (s *DisposeServiceWorkOrderRequest) GetNotifyId() *int64 {
	return s.NotifyId
}

func (s *DisposeServiceWorkOrderRequest) GetOperateRemark() *string {
	return s.OperateRemark
}

func (s *DisposeServiceWorkOrderRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *DisposeServiceWorkOrderRequest) GetOperator() *string {
	return s.Operator
}

func (s *DisposeServiceWorkOrderRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DisposeServiceWorkOrderRequest) GetUpgradeOwnerId() *string {
	return s.UpgradeOwnerId
}

func (s *DisposeServiceWorkOrderRequest) GetWorkOrderDetail() *string {
	return s.WorkOrderDetail
}

func (s *DisposeServiceWorkOrderRequest) GetWorkOrderName() *string {
	return s.WorkOrderName
}

func (s *DisposeServiceWorkOrderRequest) GetWorkOrderStatus() *string {
	return s.WorkOrderStatus
}

func (s *DisposeServiceWorkOrderRequest) SetAttachmentName(v string) *DisposeServiceWorkOrderRequest {
	s.AttachmentName = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetEndTime(v int64) *DisposeServiceWorkOrderRequest {
	s.EndTime = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetForwardOwnerId(v string) *DisposeServiceWorkOrderRequest {
	s.ForwardOwnerId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetId(v int64) *DisposeServiceWorkOrderRequest {
	s.Id = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetIsAttachment(v string) *DisposeServiceWorkOrderRequest {
	s.IsAttachment = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetIsWorkOrderNotify(v string) *DisposeServiceWorkOrderRequest {
	s.IsWorkOrderNotify = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetNotifyId(v int64) *DisposeServiceWorkOrderRequest {
	s.NotifyId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperateRemark(v string) *DisposeServiceWorkOrderRequest {
	s.OperateRemark = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperateType(v string) *DisposeServiceWorkOrderRequest {
	s.OperateType = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperator(v string) *DisposeServiceWorkOrderRequest {
	s.Operator = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetStartTime(v int64) *DisposeServiceWorkOrderRequest {
	s.StartTime = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetUpgradeOwnerId(v string) *DisposeServiceWorkOrderRequest {
	s.UpgradeOwnerId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderDetail(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderDetail = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderName(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderName = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderStatus(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderStatus = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) Validate() error {
	return dara.Validate(s)
}
