// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChangeOrdersResponseBody
	GetCode() *string
	SetData(v *ListChangeOrdersResponseBodyData) *ListChangeOrdersResponseBody
	GetData() *ListChangeOrdersResponseBodyData
	SetErrorCode(v string) *ListChangeOrdersResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListChangeOrdersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChangeOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChangeOrdersResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListChangeOrdersResponseBody
	GetTraceId() *string
}

type ListChangeOrdersResponseBody struct {
	// The HTTP status code or the POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirect.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the change orders.
	Data *ListChangeOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is not returned on successful requests.
	//
	// - Returned if the request fails. For more information, see the **error code*	- list in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information about the response.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65E1F-43BA-4D0C-8E61-E4D1337F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the list of change orders was retrieved. Valid values:
	//
	// - **true**: The list was retrieved.
	//
	// - **false**: The list could not be retrieved.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query request details.
	//
	// example:
	//
	// 0bb6f815638568884597879d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListChangeOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChangeOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChangeOrdersResponseBody) GetData() *ListChangeOrdersResponseBodyData {
	return s.Data
}

func (s *ListChangeOrdersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListChangeOrdersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChangeOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChangeOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChangeOrdersResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListChangeOrdersResponseBody) SetCode(v string) *ListChangeOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetData(v *ListChangeOrdersResponseBodyData) *ListChangeOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListChangeOrdersResponseBody) SetErrorCode(v string) *ListChangeOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetMessage(v string) *ListChangeOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetRequestId(v string) *ListChangeOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetSuccess(v bool) *ListChangeOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetTraceId(v string) *ListChangeOrdersResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListChangeOrdersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChangeOrdersResponseBodyData struct {
	// The list of change orders.
	ChangeOrderList []*ListChangeOrdersResponseBodyDataChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of change orders.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListChangeOrdersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChangeOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponseBodyData) GetChangeOrderList() []*ListChangeOrdersResponseBodyDataChangeOrderList {
	return s.ChangeOrderList
}

func (s *ListChangeOrdersResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListChangeOrdersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChangeOrdersResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListChangeOrdersResponseBodyData) SetChangeOrderList(v []*ListChangeOrdersResponseBodyDataChangeOrderList) *ListChangeOrdersResponseBodyData {
	s.ChangeOrderList = v
	return s
}

func (s *ListChangeOrdersResponseBodyData) SetCurrentPage(v int32) *ListChangeOrdersResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListChangeOrdersResponseBodyData) SetPageSize(v int32) *ListChangeOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListChangeOrdersResponseBodyData) SetTotalSize(v int32) *ListChangeOrdersResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListChangeOrdersResponseBodyData) Validate() error {
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

type ListChangeOrdersResponseBodyDataChangeOrderList struct {
	// The application ID.
	//
	// example:
	//
	// 164341c-9708-4967-b3ec-24933767****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of batches.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// The batch type. Valid values:
	//
	// - **auto**: Automatic.
	//
	// - **manual**: Manual.
	//
	// example:
	//
	// auto
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The change order ID.
	//
	// example:
	//
	// 7fa5c0-9ebb-4bb4-b383-1f885447****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The description of the change type code (**CoTypeCode**).
	//
	// example:
	//
	// Create application
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The code of the change type. Valid values:
	//
	// - **CoBindSlb**: Bind an SLB instance.
	//
	// - **CoUnbindSlb**: Unbind an SLB instance.
	//
	// - **CoCreateApp**: Create an application.
	//
	// - **CoDeleteApp**: Delete an application.
	//
	// - **CoDeploy**: Deploy an application.
	//
	// - **CoRestartApplication**: Restart an application.
	//
	// - **CoRollback**: Roll back an application.
	//
	// - **CoScaleIn**: Scale in an application.
	//
	// - **CoScaleOut**: Scale out an application.
	//
	// - **CoStartApplication**: Start an application.
	//
	// - **CoStopApplication**: Stop an application.
	//
	// - **CoRescaleApplicationVertically**: Change the instance type.
	//
	// - **CoDeployHistroy**: Roll back to a previous version.
	//
	// - **CoBindNas**: Bind a NAS file system.
	//
	// - **CoUnbindNas**: Unbind a NAS file system.
	//
	// - **CoBatchStartApplication**: Start multiple applications.
	//
	// - **CoBatchStopApplication**: Stop multiple applications.
	//
	// - **CoRestartInstances**: Restart instances.
	//
	// - **CoDeleteInstances**: Delete instances.
	//
	// - **CoScaleInAppWithInstances**: Scale in an application by specifying instances.
	//
	// example:
	//
	// CoCreateApp
	CoTypeCode *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	// The time the change order was created.
	//
	// example:
	//
	// 2019-07-11 15:54:49
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the change order.
	//
	// example:
	//
	// sae-beta-test
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The description.
	//
	// example:
	//
	// Version: 1.0 | image name: nginx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time the change order was completed.
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
	// The source of the change order.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the change order. Valid values:
	//
	// - **0**: Preparing.
	//
	// - **1**: In progress.
	//
	// - **2**: Succeeded.
	//
	// - **3**: Failed.
	//
	// - **6**: Aborted.
	//
	// - **8**: Paused for manual confirmation.
	//
	// - **9**: Paused for automatic confirmation.
	//
	// - **10**: Failed due to a system exception.
	//
	// - **11**: Pending approval.
	//
	// - **12**: Approved and pending execution.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
	//
	// example:
	//
	// sae-beta-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListChangeOrdersResponseBodyDataChangeOrderList) String() string {
	return dara.Prettify(s)
}

func (s ListChangeOrdersResponseBodyDataChangeOrderList) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetAppId() *string {
	return s.AppId
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetBatchCount() *int32 {
	return s.BatchCount
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetBatchType() *string {
	return s.BatchType
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetCoType() *string {
	return s.CoType
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetCoTypeCode() *string {
	return s.CoTypeCode
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetDescription() *string {
	return s.Description
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetSource() *string {
	return s.Source
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetStatus() *int32 {
	return s.Status
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) GetUserId() *string {
	return s.UserId
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetAppId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.AppId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetBatchCount(v int32) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchCount = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetBatchType(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchType = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetChangeOrderId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.ChangeOrderId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCoType(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CoType = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCoTypeCode(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CoTypeCode = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCreateTime(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateTime = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCreateUserId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateUserId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetDescription(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.Description = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetFinishTime(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.FinishTime = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetGroupId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.GroupId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetSource(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.Source = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetStatus(v int32) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.Status = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetUserId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.UserId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) Validate() error {
	return dara.Validate(s)
}
