// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrossProjectPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCrossProjectPipelineRunResponseBodyData) *GetCrossProjectPipelineRunResponseBody
	GetData() *GetCrossProjectPipelineRunResponseBodyData
	SetRequestId(v string) *GetCrossProjectPipelineRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCrossProjectPipelineRunResponseBody
	GetSuccess() *bool
}

type GetCrossProjectPipelineRunResponseBody struct {
	// The business response.
	//
	// example:
	//
	// {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PipelineRunId":"fcfd4160-e2ff-4603-9719-09128fe733df","DeploymentEnvironmentId":101,"ObjectId":"1","ObjectType":"ODPS_SQL","ObjectName":"object-1","ObjectVersion":"7","ChangeType":"ADD","Status":"Ready","Description":"Publish objects that are published in the source project to the target project","Creator":"creator","CreateTime":1788739200000}
	Data *GetCrossProjectPipelineRunResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot this API call.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCrossProjectPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCrossProjectPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrossProjectPipelineRunResponseBody) GetData() *GetCrossProjectPipelineRunResponseBodyData {
	return s.Data
}

func (s *GetCrossProjectPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCrossProjectPipelineRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCrossProjectPipelineRunResponseBody) SetData(v *GetCrossProjectPipelineRunResponseBodyData) *GetCrossProjectPipelineRunResponseBody {
	s.Data = v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBody) SetRequestId(v string) *GetCrossProjectPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBody) SetSuccess(v bool) *GetCrossProjectPipelineRunResponseBody {
	s.Success = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCrossProjectPipelineRunResponseBodyData struct {
	// The termination time. This value is a UNIX timestamp in milliseconds. This parameter is returned only after the flow is terminated.
	//
	// example:
	//
	// 1788739260000
	AbolishTime *int64 `json:"AbolishTime,omitempty" xml:"AbolishTime,omitempty"`
	// The user who terminated the flow.
	//
	// example:
	//
	// operator
	Abolisher *string `json:"Abolisher,omitempty" xml:"Abolisher,omitempty"`
	// The change type.
	//
	// example:
	//
	// ADD
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The creation time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788739200000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// creator
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The cross-workspace deployment environment ID.
	//
	// example:
	//
	// 101
	DeploymentEnvironmentId *int64 `json:"DeploymentEnvironmentId,omitempty" xml:"DeploymentEnvironmentId,omitempty"`
	// The deployment description.
	//
	// example:
	//
	// Publish objects that are published in the source project to the target project
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error code.
	//
	// example:
	//
	// DeploymentFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Deployment failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The execution time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788739260000
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The executor.
	//
	// example:
	//
	// executor
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The completion time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788739320000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the deployment object.
	//
	// example:
	//
	// 1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the deployment object.
	//
	// example:
	//
	// object-1
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The object type of the publish object.
	//
	// example:
	//
	// ODPS_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The version of the deployment object.
	//
	// example:
	//
	// 7
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
	// The cross-workspace deployment flow ID.
	//
	// example:
	//
	// fcfd4160-e2ff-4603-9719-09128fe733df
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the deployment flow. Valid values:
	//
	// - Building: Building.
	//
	// - Ready: Ready and waiting for execution.
	//
	// - Running: Running.
	//
	// - Termination: Terminated.
	//
	// - Success: Execution succeeded.
	//
	// - Fail: Execution failed.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCrossProjectPipelineRunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCrossProjectPipelineRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetAbolishTime() *int64 {
	return s.AbolishTime
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetAbolisher() *string {
	return s.Abolisher
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetChangeType() *string {
	return s.ChangeType
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetDeploymentEnvironmentId() *int64 {
	return s.DeploymentEnvironmentId
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetExecutor() *string {
	return s.Executor
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCrossProjectPipelineRunResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetAbolishTime(v int64) *GetCrossProjectPipelineRunResponseBodyData {
	s.AbolishTime = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetAbolisher(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.Abolisher = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetChangeType(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.ChangeType = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetCreateTime(v int64) *GetCrossProjectPipelineRunResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetCreator(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetDeploymentEnvironmentId(v int64) *GetCrossProjectPipelineRunResponseBodyData {
	s.DeploymentEnvironmentId = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetDescription(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetErrorCode(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetErrorMessage(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetExecuteTime(v int64) *GetCrossProjectPipelineRunResponseBodyData {
	s.ExecuteTime = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetExecutor(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.Executor = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetFinishTime(v int64) *GetCrossProjectPipelineRunResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetObjectId(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.ObjectId = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetObjectName(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.ObjectName = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetObjectType(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.ObjectType = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetObjectVersion(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.ObjectVersion = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetPipelineRunId(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.PipelineRunId = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetRequestId(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) SetStatus(v string) *GetCrossProjectPipelineRunResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCrossProjectPipelineRunResponseBodyData) Validate() error {
	return dara.Validate(s)
}
