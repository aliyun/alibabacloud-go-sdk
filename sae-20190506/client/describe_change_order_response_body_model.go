// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChangeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeChangeOrderResponseBody
	GetCode() *string
	SetData(v *DescribeChangeOrderResponseBodyData) *DescribeChangeOrderResponseBody
	GetData() *DescribeChangeOrderResponseBodyData
	SetErrorCode(v string) *DescribeChangeOrderResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeChangeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeChangeOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeChangeOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeChangeOrderResponseBody
	GetTraceId() *string
}

type DescribeChangeOrderResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the change order.
	Data *DescribeChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the information of the change order was queried. Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The information failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeChangeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeChangeOrderResponseBody) GetData() *DescribeChangeOrderResponseBodyData {
	return s.Data
}

func (s *DescribeChangeOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeChangeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeChangeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChangeOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeChangeOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeChangeOrderResponseBody) SetCode(v string) *DescribeChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetData(v *DescribeChangeOrderResponseBodyData) *DescribeChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetErrorCode(v string) *DescribeChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetMessage(v string) *DescribeChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetRequestId(v string) *DescribeChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetSuccess(v bool) *DescribeChangeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetTraceId(v string) *DescribeChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChangeOrderResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// bbbbb-3fd370b2-3646-4ba6-91f9-9423e19ab0cd-*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// app-test
	AppName                       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ApplicationEnableGreyTagRoute *bool   `json:"ApplicationEnableGreyTagRoute,omitempty" xml:"ApplicationEnableGreyTagRoute,omitempty"`
	ApplicationUpdateStrategy     *string `json:"ApplicationUpdateStrategy,omitempty" xml:"ApplicationUpdateStrategy,omitempty"`
	// The approval ID of the change order.
	//
	// example:
	//
	// 67de0b39-a9d4-4c09-a170-cf438208****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// Indicates whether SAE automatically releases the batches. Valid values:
	//
	// 	- **true**: SAE automatically releases the batches.
	//
	// 	- **false**: SAE does not automatically release the batches.
	//
	// example:
	//
	// true
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// The number of release batches.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// The processing method for the batches. Valid values:
	//
	// 	- **auto**: SAE automatically releases the batches.
	//
	// 	- **Manual**: You must manually release the batches.
	//
	// example:
	//
	// auto
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The interval between batches in a phased release. SAE automatically releases batches at the specified interval. Unit: minutes.
	//
	// example:
	//
	// 0
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The ID of the change order.
	//
	// example:
	//
	// 765fa5c0-9ebb-4bb4-b383-1f885447**
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The change type, which corresponds to the **CoTypeCode*	- parameter.
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The code of the change type. Valid values:
	//
	// 	- **CoBindSlb**: associates a Sever Load Balancer (SLB) instance with the application.
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
	// 	- **CoRescaleApplicationVertically**: modifies the instance type.
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
	// 	- **CoScaleInAppWithInstances**: reduces the specified number of application instances.
	//
	// example:
	//
	// CoRestartInstances
	CoTypeCode *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	// The time when the change order was created.
	//
	// example:
	//
	// 2020-12-17 21:06:45
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the current batch.
	//
	// example:
	//
	// 0e4acf82-c9b1-4c1e-ac28-55776338****
	CurrentPipelineId *string `json:"CurrentPipelineId,omitempty" xml:"CurrentPipelineId,omitempty"`
	// The description of the change order.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The batch information.
	Pipelines []*DescribeChangeOrderResponseBodyDataPipelines `json:"Pipelines,omitempty" xml:"Pipelines,omitempty" type:"Repeated"`
	// The status of the change order. Valid values:
	//
	// 	- **0**: The change order is being prepared.
	//
	// 	- **1**: The change order is being executed.
	//
	// 	- **2**: The change order was executed.
	//
	// 	- **3**: The change order failed to be executed.
	//
	// 	- **6**: The change order was terminated.
	//
	// 	- **8**: The execution process is pending. You must manually release the batches.
	//
	// 	- **9**: The execution process is pending. SAE will automatically release the batches.
	//
	// 	- **10**: The execution failed due to a system exception.
	//
	// 	- **11**: The change order is pending approval.
	//
	// 	- **12**: The change order is approved and is pending execution.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The substatus of the change order. This parameter indicates whether an exception occurred while the change order was being executed. Valid values:
	//
	// 	- **0**: No exception occurred.
	//
	// 	- **1**: An exception occurred. For example, if an error occurs during a phased release, you must manually roll back the application. In this case, the change order cannot be completed, so the Status parameter is still displayed as "1", which indicates that the change order is being executed. You can check the value of this parameter to determine whether an exception occurs.
	//
	// example:
	//
	// 0
	SubStatus *int32 `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// Indicates whether the application can be rolled back. Valid values:
	//
	// 	- **true**: The application can be rolled back.
	//
	// 	- **false**: The application cannot be rolled back.
	//
	// example:
	//
	// false
	SupportRollback *bool `json:"SupportRollback,omitempty" xml:"SupportRollback,omitempty"`
}

func (s DescribeChangeOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChangeOrderResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeChangeOrderResponseBodyData) GetApplicationEnableGreyTagRoute() *bool {
	return s.ApplicationEnableGreyTagRoute
}

func (s *DescribeChangeOrderResponseBodyData) GetApplicationUpdateStrategy() *string {
	return s.ApplicationUpdateStrategy
}

func (s *DescribeChangeOrderResponseBodyData) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *DescribeChangeOrderResponseBodyData) GetAuto() *bool {
	return s.Auto
}

func (s *DescribeChangeOrderResponseBodyData) GetBatchCount() *int32 {
	return s.BatchCount
}

func (s *DescribeChangeOrderResponseBodyData) GetBatchType() *string {
	return s.BatchType
}

func (s *DescribeChangeOrderResponseBodyData) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *DescribeChangeOrderResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DescribeChangeOrderResponseBodyData) GetCoType() *string {
	return s.CoType
}

func (s *DescribeChangeOrderResponseBodyData) GetCoTypeCode() *string {
	return s.CoTypeCode
}

func (s *DescribeChangeOrderResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeChangeOrderResponseBodyData) GetCurrentPipelineId() *string {
	return s.CurrentPipelineId
}

func (s *DescribeChangeOrderResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeChangeOrderResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeChangeOrderResponseBodyData) GetPipelines() []*DescribeChangeOrderResponseBodyDataPipelines {
	return s.Pipelines
}

func (s *DescribeChangeOrderResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeChangeOrderResponseBodyData) GetSubStatus() *int32 {
	return s.SubStatus
}

func (s *DescribeChangeOrderResponseBodyData) GetSupportRollback() *bool {
	return s.SupportRollback
}

func (s *DescribeChangeOrderResponseBodyData) SetAppId(v string) *DescribeChangeOrderResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetAppName(v string) *DescribeChangeOrderResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetApplicationEnableGreyTagRoute(v bool) *DescribeChangeOrderResponseBodyData {
	s.ApplicationEnableGreyTagRoute = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetApplicationUpdateStrategy(v string) *DescribeChangeOrderResponseBodyData {
	s.ApplicationUpdateStrategy = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetApprovalId(v string) *DescribeChangeOrderResponseBodyData {
	s.ApprovalId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetAuto(v bool) *DescribeChangeOrderResponseBodyData {
	s.Auto = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetBatchCount(v int32) *DescribeChangeOrderResponseBodyData {
	s.BatchCount = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetBatchType(v string) *DescribeChangeOrderResponseBodyData {
	s.BatchType = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetBatchWaitTime(v int32) *DescribeChangeOrderResponseBodyData {
	s.BatchWaitTime = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetChangeOrderId(v string) *DescribeChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCoType(v string) *DescribeChangeOrderResponseBodyData {
	s.CoType = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCoTypeCode(v string) *DescribeChangeOrderResponseBodyData {
	s.CoTypeCode = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCreateTime(v string) *DescribeChangeOrderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCurrentPipelineId(v string) *DescribeChangeOrderResponseBodyData {
	s.CurrentPipelineId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetDescription(v string) *DescribeChangeOrderResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetErrorMessage(v string) *DescribeChangeOrderResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetPipelines(v []*DescribeChangeOrderResponseBodyDataPipelines) *DescribeChangeOrderResponseBodyData {
	s.Pipelines = v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetStatus(v int32) *DescribeChangeOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetSubStatus(v int32) *DescribeChangeOrderResponseBodyData {
	s.SubStatus = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetSupportRollback(v bool) *DescribeChangeOrderResponseBodyData {
	s.SupportRollback = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeChangeOrderResponseBodyDataPipelines struct {
	// The batch type.
	//
	// example:
	//
	// 0
	BatchType *int32 `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The number of parallel tasks in a batch.
	//
	// example:
	//
	// 0
	ParallelCount *int32 `json:"ParallelCount,omitempty" xml:"ParallelCount,omitempty"`
	// The ID of the batch.
	//
	// example:
	//
	// 0e4acf82-c9b1-4c1e-ac28-55776338****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The name of the batch.
	//
	// example:
	//
	// Batch 1 Change
	PipelineName *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	// The time when the batch processing started.
	//
	// example:
	//
	// 1562831689704
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the batch. Valid values:
	//
	// 	- **0**: The batch is being prepared.
	//
	// 	- **1**: The batch is being processed.
	//
	// 	- **2**: The batch was processed.
	//
	// 	- **3**: The batch failed to be processed.
	//
	// 	- **6**: The batch processing was terminated.
	//
	// 	- **8**: The execution process is pending. You must manually release the batch.
	//
	// 	- **9**: The execution process is pending. SAE will automatically release the batch.
	//
	// 	- **10**: The batch failed to be processed due to a system exception.
	//
	// 	- **11**: The batch is pending approval.
	//
	// 	- **12**: The batch is approved and is pending execution.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the batch information was last modified.
	//
	// example:
	//
	// 1562847178007
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeChangeOrderResponseBodyDataPipelines) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeOrderResponseBodyDataPipelines) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) GetBatchType() *int32 {
	return s.BatchType
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) GetParallelCount() *int32 {
	return s.ParallelCount
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) GetPipelineName() *string {
	return s.PipelineName
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetBatchType(v int32) *DescribeChangeOrderResponseBodyDataPipelines {
	s.BatchType = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetParallelCount(v int32) *DescribeChangeOrderResponseBodyDataPipelines {
	s.ParallelCount = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetPipelineId(v string) *DescribeChangeOrderResponseBodyDataPipelines {
	s.PipelineId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetPipelineName(v string) *DescribeChangeOrderResponseBodyDataPipelines {
	s.PipelineName = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetStartTime(v int64) *DescribeChangeOrderResponseBodyDataPipelines {
	s.StartTime = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetStatus(v int32) *DescribeChangeOrderResponseBodyDataPipelines {
	s.Status = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetUpdateTime(v int64) *DescribeChangeOrderResponseBodyDataPipelines {
	s.UpdateTime = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) Validate() error {
	return dara.Validate(s)
}
