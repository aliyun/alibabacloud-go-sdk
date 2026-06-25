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
	// The HTTP status code or POP error code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A client-side error occurred.
	//
	// - **5xx**: A server-side error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the change order.
	Data *DescribeChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - If the request is successful, this parameter is not returned.
	//
	// - If the request fails, this parameter is returned. For more information, see the **error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned for the request.
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
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID of the request. This ID is used for troubleshooting.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeChangeOrderResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// bbbbb-3fd370b2-3646-4ba6-91f9-9423e19ab0cd-*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// app-test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Indicates whether gray tag routing is enabled.
	ApplicationEnableGreyTagRoute *bool `json:"ApplicationEnableGreyTagRoute,omitempty" xml:"ApplicationEnableGreyTagRoute,omitempty"`
	// The update strategy for the application.
	ApplicationUpdateStrategy *string `json:"ApplicationUpdateStrategy,omitempty" xml:"ApplicationUpdateStrategy,omitempty"`
	// The approval ID for the operation.
	//
	// example:
	//
	// 67de0b39-a9d4-4c09-a170-cf438208****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// Indicates whether to automatically release the change in batches. Valid values:
	//
	// - **true**: The change is automatically released.
	//
	// - **false**: The change is not automatically released.
	//
	// example:
	//
	// true
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// The number of batches.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// The release mode for batches. Valid values:
	//
	// - **auto**: Automatic release.
	//
	// - **manual**: Manual release.
	//
	// example:
	//
	// auto
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The wait time, in minutes, between batches in an automatic release.
	//
	// example:
	//
	// 0
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The change order ID.
	//
	// example:
	//
	// 765fa5c0-9ebb-4bb4-b383-1f885447**
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The change type. This parameter is a description of `CoTypeCode`.
	//
	// example:
	//
	// Batch Restart Instances
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The change type code. Valid values:
	//
	// - **CoBindSlb**: Binds an SLB instance.
	//
	// - **CoUnbindSlb**: Unbinds an SLB instance.
	//
	// - **CoCreateApp**: Creates an application.
	//
	// - **CoDeleteApp**: Deletes an application.
	//
	// - **CoDeploy**: Deploys an application.
	//
	// - **CoRestartApplication**: Restarts an application.
	//
	// - **CoRollback**: Rolls back an application.
	//
	// - **CoScaleIn**: Scales in an application.
	//
	// - **CoScaleOut**: Scales out an application.
	//
	// - **CoStart**: Starts an application.
	//
	// - **CoStop**: Stops an application.
	//
	// - **CoRescaleApplicationVertically**: Modifies instance specifications.
	//
	// - **CoDeployHistroy**: Rolls back to a historical version.
	//
	// - **CoBindNas**: Binds a NAS file system.
	//
	// - **CoUnbindNas**: Unbinds a NAS file system.
	//
	// - **CoBatchStartApplication**: Starts applications in batches.
	//
	// - **CoBatchStopApplication**: Stops applications in batches.
	//
	// - **CoRestartInstances**: Restarts instances.
	//
	// - **CoDeleteInstances**: Deletes instances.
	//
	// - **CoScaleInAppWithInstances**: Scales in an application by specifying instances.
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
	// - **0**: Preparing
	//
	// - **1**: In progress
	//
	// - **2**: Succeeded
	//
	// - **3**: Failed
	//
	// - **6**: Terminated
	//
	// - **8**: Awaiting manual confirmation
	//
	// - **9**: Awaiting automatic confirmation
	//
	// - **10**: Failed due to a system error
	//
	// - **11**: Pending approval
	//
	// - **12**: Approved and pending execution
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The substatus of the release order. This parameter indicates whether an exception occurred during the release. Valid values:
	//
	// - **0**: Normal.
	//
	// - **1**: Abnormal. For example, if a batch release fails, you must manually perform a rollback, leaving the release order in the In Progress state.
	//
	// example:
	//
	// 0
	SubStatus *int32 `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// Indicates whether rollback is supported. Valid values:
	//
	// - **true**: Rollback is supported.
	//
	// - **false**: Rollback is not supported.
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
	if s.Pipelines != nil {
		for _, item := range s.Pipelines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	// The batch ID.
	//
	// example:
	//
	// 0e4acf82-c9b1-4c1e-ac28-55776338****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The batch name.
	//
	// example:
	//
	// Batch 1 Change
	PipelineName *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	// The start time of the batch.
	//
	// example:
	//
	// 1562831689704
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the batch. Valid values:
	//
	// - **0**: Preparing
	//
	// - **1**: In progress
	//
	// - **2**: Succeeded
	//
	// - **3**: Failed
	//
	// - **6**: Terminated
	//
	// - **8**: Awaiting manual confirmation
	//
	// - **9**: Awaiting automatic confirmation
	//
	// - **10**: Failed due to a system error
	//
	// - **11**: Pending approval
	//
	// - **12**: Approved and pending execution
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the batch was last updated.
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
