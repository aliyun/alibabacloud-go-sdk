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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: success
	//
	// - **3xx**: redirection
	//
	// - **4xx**: client error
	//
	// - **5xx**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *ListNamespaceChangeOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - Request succeeded: The **ErrorCode*	- parameter is not returned.
	//
	// - Request failed: The **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// Resouce.no.permission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the API call result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc3915638507554994370d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the change order list was successfully retrieved. Valid values:
	//
	// - **true**: The list was successfully retrieved.
	//
	// - **false**: The retrieval failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query the exact call information.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNamespaceChangeOrdersResponseBodyData struct {
	// The list of change orders.
	ChangeOrderList []*ListNamespaceChangeOrdersResponseBodyDataChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
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
	if s.ChangeOrderList != nil {
		for _, item := range s.ChangeOrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespaceChangeOrdersResponseBodyDataChangeOrderList struct {
	// The number of batches.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// The batch type.
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The change order ID.
	//
	// example:
	//
	// 7fa5c0-9ebb-4bb4-b383-1f885447****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The type of the change order, which is the description of **CoTypeCode**.
	//
	// example:
	//
	// msg.docker.app.actions.CoBatchStartApplication
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The code of the change order type. Valid values:
	//
	// - **CoBatchStartApplication**: Starts applications in batches.
	//
	// - **CoBatchStopApplication**: Stops applications in batches.
	//
	// example:
	//
	// CoBatchStartApplication
	CoTypeCode *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	// The creation time.
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
	//
	// example:
	//
	// Batch Start Applications
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2019-07-11 20:12:58
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The group ID.
	//
	// example:
	//
	// c9ecd2-cf6c-46c3-9f20-525de202****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-shanghai:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The change pipelines.
	//
	// example:
	//
	// xxxx
	Pipelines *string `json:"Pipelines,omitempty" xml:"Pipelines,omitempty"`
	// The source from which the change order was initiated.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the change order. Valid values:
	//
	// - **0**: preparing
	//
	// - **1**: executing
	//
	// - **2**: succeeded
	//
	// - **3**: failed
	//
	// - **6**: terminated
	//
	// - **10**: failed due to a system exception
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
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
