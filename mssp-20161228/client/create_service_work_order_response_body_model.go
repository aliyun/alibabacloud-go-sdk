// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceWorkOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceWorkOrderResponseBody
	GetCode() *string
	SetData(v *CreateServiceWorkOrderResponseBodyData) *CreateServiceWorkOrderResponseBody
	GetData() *CreateServiceWorkOrderResponseBodyData
	SetHttpStatusCode(v int32) *CreateServiceWorkOrderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateServiceWorkOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceWorkOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceWorkOrderResponseBody
	GetSuccess() *bool
}

type CreateServiceWorkOrderResponseBody struct {
	// Interface status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *CreateServiceWorkOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7DC44321-7AAE-51CD-8E5F-CEB968569042
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceWorkOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceWorkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceWorkOrderResponseBody) GetData() *CreateServiceWorkOrderResponseBodyData {
	return s.Data
}

func (s *CreateServiceWorkOrderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateServiceWorkOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceWorkOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceWorkOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceWorkOrderResponseBody) SetCode(v string) *CreateServiceWorkOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetData(v *CreateServiceWorkOrderResponseBodyData) *CreateServiceWorkOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetHttpStatusCode(v int32) *CreateServiceWorkOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetMessage(v string) *CreateServiceWorkOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetRequestId(v string) *CreateServiceWorkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetSuccess(v bool) *CreateServiceWorkOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceWorkOrderResponseBodyData struct {
	// Completion time.
	//
	// example:
	//
	// 1734788109000
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1734788109000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creator.
	//
	// example:
	//
	// 426556
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Customer ID.
	//
	// example:
	//
	// 1477832102462645
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// End time.
	//
	// example:
	//
	// 1734788109000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Primary key.
	//
	// example:
	//
	// 1978941
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Owner.
	//
	// example:
	//
	// 426556
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Start time.
	//
	// example:
	//
	// 1734788109000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Work order details.
	//
	// example:
	//
	// {"questionDetail":"测试工单","answerDetail":""}
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	// Work order name.
	//
	// example:
	//
	// Delivery task of safety monthly report
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// Work order source.
	//
	// example:
	//
	// Work order migration
	WorkOrderSource *string `json:"WorkOrderSource,omitempty" xml:"WorkOrderSource,omitempty"`
	// Work order status.
	//
	// example:
	//
	// UNPROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
	// Work order type.
	//
	// example:
	//
	// MONTH_REPORT
	WorkOrderType *string `json:"WorkOrderType,omitempty" xml:"WorkOrderType,omitempty"`
}

func (s CreateServiceWorkOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceWorkOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponseBodyData) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *CreateServiceWorkOrderResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateServiceWorkOrderResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *CreateServiceWorkOrderResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *CreateServiceWorkOrderResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateServiceWorkOrderResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateServiceWorkOrderResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateServiceWorkOrderResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateServiceWorkOrderResponseBodyData) GetWorkOrderDetail() *string {
	return s.WorkOrderDetail
}

func (s *CreateServiceWorkOrderResponseBodyData) GetWorkOrderName() *string {
	return s.WorkOrderName
}

func (s *CreateServiceWorkOrderResponseBodyData) GetWorkOrderSource() *string {
	return s.WorkOrderSource
}

func (s *CreateServiceWorkOrderResponseBodyData) GetWorkOrderStatus() *string {
	return s.WorkOrderStatus
}

func (s *CreateServiceWorkOrderResponseBodyData) GetWorkOrderType() *string {
	return s.WorkOrderType
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCompleteTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCreateTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCreator(v string) *CreateServiceWorkOrderResponseBodyData {
	s.Creator = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCustomerId(v string) *CreateServiceWorkOrderResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetEndTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetId(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetOwnerId(v string) *CreateServiceWorkOrderResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetStartTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderDetail(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderDetail = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderName(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderName = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderSource(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderSource = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderStatus(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderStatus = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderType(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderType = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
