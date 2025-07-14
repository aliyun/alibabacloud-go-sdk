// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespaceChangeOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNamespaceChangeOrdersResponseBody
	GetCode() *string
	SetData(v *ListNamespaceChangeOrdersResponseBodyData) *ListNamespaceChangeOrdersResponseBody
	GetData() *ListNamespaceChangeOrdersResponseBodyData
	SetErrorCode(v string) *ListNamespaceChangeOrdersResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListNamespaceChangeOrdersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNamespaceChangeOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNamespaceChangeOrdersResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListNamespaceChangeOrdersResponseBody
	GetTraceId() *string
}

type ListNamespaceChangeOrdersResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListNamespaceChangeOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3915638507554994370d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the list of change orders was obtained. Valid values:
	//
	// 	- **true**: indicates that the list was obtained.
	//
	// 	- **false**: indicates that the list could not be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. It is used to query the details of a request.
	//
	// example:
	//
	// 0bc3915638507554994370d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListNamespaceChangeOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNamespaceChangeOrdersResponseBody) GetData() *ListNamespaceChangeOrdersResponseBodyData {
	return s.Data
}

func (s *ListNamespaceChangeOrdersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNamespaceChangeOrdersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNamespaceChangeOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespaceChangeOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNamespaceChangeOrdersResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListNamespaceChangeOrdersResponseBody) SetCode(v string) *ListNamespaceChangeOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetData(v *ListNamespaceChangeOrdersResponseBodyData) *ListNamespaceChangeOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetErrorCode(v string) *ListNamespaceChangeOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetMessage(v string) *ListNamespaceChangeOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetRequestId(v string) *ListNamespaceChangeOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetSuccess(v bool) *ListNamespaceChangeOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetTraceId(v string) *ListNamespaceChangeOrdersResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNamespaceChangeOrdersResponseBodyData struct {
	// The list of change orders.
	ChangeOrderList []*ListNamespaceChangeOrdersResponseBodyDataChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" type:"Repeated"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of change orders.
	//
	// example:
	//
	// 32
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListNamespaceChangeOrdersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponseBodyData) GetChangeOrderList() []*ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	return s.ChangeOrderList
}

func (s *ListNamespaceChangeOrdersResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListNamespaceChangeOrdersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespaceChangeOrdersResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetChangeOrderList(v []*ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) *ListNamespaceChangeOrdersResponseBodyData {
	s.ChangeOrderList = v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetCurrentPage(v int32) *ListNamespaceChangeOrdersResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetPageSize(v int32) *ListNamespaceChangeOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetTotalSize(v int32) *ListNamespaceChangeOrdersResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListNamespaceChangeOrdersResponseBodyDataChangeOrderList struct {
	// The number of release batches.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// The mode in which the release batches are determined. Valid values:
	//
	// 	- **auto**: SAE automatically determines the release batches.
	//
	// 	- **manual**: You must manually determine the release batches.
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The ID of the change order.
	//
	// example:
	//
	// 7fa5c0-9ebb-4bb4-b383-1f885447****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The type of the change order, which corresponds the **CoTypeCode*	- parameter.
	//
	// example:
	//
	// msg.docker.app.actions.CoBatchStartApplication
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The code of the change order type. Valid values:
	//
	// 	- **CoBatchStartApplication**: starts multiple applications concurrently.
	//
	// 	- **CoBatchStopApplication**: stops multiple applications concurrently.
	//
	// example:
	//
	// CoBatchStartApplication
	CoTypeCode *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	// The time when the change order was created.
	//
	// example:
	//
	// 2019-07-11 15:54:49
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the change order.
	//
	// example:
	//
	// test@aliyun.com
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The description of the change order.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the change order was completed.
	//
	// example:
	//
	// 2019-07-11 20:12:58
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// c9ecd2-cf6c-46c3-9f20-525de202****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-shanghai:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The information about release batches.
	//
	// example:
	//
	// xxxx
	Pipelines *string `json:"Pipelines,omitempty" xml:"Pipelines,omitempty"`
	// The source of the change order.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the change order. Valid values:
	//
	// 	- **0**: The change order is being prepared.
	//
	// 	- **1**: The change order is being executed.
	//
	// 	- **2**: The change order was executed.
	//
	// 	- **3**: The change order could not be executed.
	//
	// 	- **6**: The change order was terminated.
	//
	// 	- **10**: The change order could not be executed due to a system exception.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// test_sae
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetBatchCount() *int32 {
	return s.BatchCount
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetBatchType() *string {
	return s.BatchType
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetCoType() *string {
	return s.CoType
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetCoTypeCode() *string {
	return s.CoTypeCode
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetDescription() *string {
	return s.Description
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetPipelines() *string {
	return s.Pipelines
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetSource() *string {
	return s.Source
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetStatus() *int32 {
	return s.Status
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GetUserId() *string {
	return s.UserId
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetBatchCount(v int32) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchCount = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetBatchType(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchType = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetChangeOrderId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.ChangeOrderId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCoType(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CoType = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCoTypeCode(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CoTypeCode = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCreateTime(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateTime = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCreateUserId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateUserId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetDescription(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Description = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetFinishTime(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.FinishTime = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetGroupId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.GroupId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetNamespaceId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetPipelines(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Pipelines = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetSource(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Source = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetStatus(v int32) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Status = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetUserId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.UserId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) Validate() error {
	return dara.Validate(s)
}
