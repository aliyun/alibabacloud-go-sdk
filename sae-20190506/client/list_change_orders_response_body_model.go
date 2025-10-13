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
	// Indicates whether the list of change orders was obtained. Valid values:
	//
	// 	- **true**: indicates that the list was obtained.
	//
	// 	- **false**: indicates that the list could not be obtained.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about change orders.
	Data *ListChangeOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the trace. It is used to query the details of a request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned message.
	//
	// example:
	//
	// 65E1F-43BA-4D0C-8E61-E4D1337F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about change orders.
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
	// The change orders.
	ChangeOrderList []*ListChangeOrdersResponseBodyDataChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" type:"Repeated"`
	// The total number of change orders.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of change orders.
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
	// The number of entries returned on each page.
	//
	// example:
	//
	// 164341c-9708-4967-b3ec-24933767****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the user who created the change order.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// auto
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The mode in which the release batches are determined. Valid values:
	//
	// 	- **auto**: SAE automatically determines the release batches.
	//
	// 	- **manual**: You must manually determine the release batches.
	//
	// example:
	//
	// 7fa5c0-9ebb-4bb4-b383-1f885447****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The ID of the application.
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The code of the change order. Valid values:
	//
	// 	- **CoBindSlb**: associates the Server Load Balancer (SLB) instance with the application.
	//
	// 	- **CoUnbindSlb**: disassociates an SLB instance from the application.
	//
	// 	- **CoCreateApp**: creates the application.
	//
	// 	- **CoDeleteApp**: deletes the application.
	//
	// 	- **CoDeploy**: deploys the application.
	//
	// 	- **CoRestartApplication**: restarts the application.
	//
	// 	- **CoRollback**: rolls back the application.
	//
	// 	- **CoScaleIn**: scales in the application.
	//
	// 	- **CoScaleOut**: scales out the application.
	//
	// 	- **CoStartApplication**: starts the application.
	//
	// 	- **CoStopApplication**: stops the application.
	//
	// 	- **CoRescaleApplicationVertically**: modifies the instance type.
	//
	// 	- **CoDeployHistroy**: rolls back the application to an earlier version.
	//
	// 	- **CoBindNas**: associates a network-attached storage (NAS) file system with the application.
	//
	// 	- **CoUnbindNas**: disassociates a NAS file system from the application.
	//
	// 	- **CoBatchStartApplication**: starts multiple applications concurrently.
	//
	// 	- **CoBatchStopApplication**: stops multiple applications concurrently.
	//
	// 	- **CoRestartInstances**: restarts the instance.
	//
	// 	- **CoDeleteInstances**: deletes the instance.
	//
	// 	- **CoScaleInAppWithInstances**: reduces the specified number of application instances.
	//
	// example:
	//
	// CoCreateApp
	CoTypeCode *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 2019-07-11 15:54:49
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The code of the change type. Valid values:
	//
	// 	- **CoBindSlb**: associates an SLB instance with the application.
	//
	// 	- **CoUnbindSlb**: disassociates the SLB instance from the application.
	//
	// 	- **CoCreateApp**: creates the application.
	//
	// 	- **CoDeleteApp**: deletes the application.
	//
	// 	- **CoDeploy**: deploys the application.
	//
	// 	- **CoRestartApplication**: restarts the application.
	//
	// 	- **CoRollback**: rolls back the application.
	//
	// 	- **CoScaleIn**: scales in the application.
	//
	// 	- **CoScaleOut**: scales out the application.
	//
	// 	- **CoStart**: starts the application.
	//
	// 	- **CoStop**: stops the application.
	//
	// 	- **CoRescaleApplicationVertically**: modifies the instance specifications.
	//
	// 	- **CoDeployHistroy**: rolls back the application to a historical version.
	//
	// 	- **CoBindNas**: associates a NAS file system with the application.
	//
	// 	- **CoUnbindNas**: disassociates the NAS file system from the application.
	//
	// 	- **CoBatchStartApplication**: starts multiple applications concurrently.
	//
	// 	- **CoBatchStopApplication**: stops multiple applications concurrently.
	//
	// 	- **CoRestartInstances**: restarts the instances.
	//
	// 	- **CoDeleteInstances**: deletes the instances.
	//
	// 	- **CoScaleInAppWithInstances**: reduces the number of the specified application instances.
	//
	// example:
	//
	// sae-beta-test
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The change type, which corresponds to the **CoTypeCode*	- parameter.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the change order was created.
	//
	// example:
	//
	// 2019-07-11 20:12:58
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The description about the application.
	//
	// example:
	//
	// c9ecd2-cf6c-46c3-9f20-525de202****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of release batches.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The time when the change order was completed.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The source of the change order.
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
