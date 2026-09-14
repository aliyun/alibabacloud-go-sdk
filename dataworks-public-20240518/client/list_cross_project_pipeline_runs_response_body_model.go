// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectPipelineRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCrossProjectPipelineRunsResponseBodyData) *ListCrossProjectPipelineRunsResponseBody
	GetData() *ListCrossProjectPipelineRunsResponseBodyData
	SetRequestId(v string) *ListCrossProjectPipelineRunsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCrossProjectPipelineRunsResponseBody
	GetSuccess() *bool
}

type ListCrossProjectPipelineRunsResponseBody struct {
	// The business response data.
	//
	// example:
	//
	// {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PageNumber":1,"PageSize":10,"TotalCount":1,"PipelineRuns":[{"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PipelineRunId":"fcfd4160-e2ff-4603-9719-09128fe733df","DeploymentEnvironmentId":101,"ObjectId":"1","ObjectType":"ODPS_SQL","ObjectName":"object-1","ObjectVersion":"7","ChangeType":"ADD","Status":"Ready","Description":"Publish objects that have been published in the source project to the target project","Creator":"creator","CreateTime":1788739200000}]}
	Data *ListCrossProjectPipelineRunsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListCrossProjectPipelineRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunsResponseBody) GetData() *ListCrossProjectPipelineRunsResponseBodyData {
	return s.Data
}

func (s *ListCrossProjectPipelineRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectPipelineRunsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCrossProjectPipelineRunsResponseBody) SetData(v *ListCrossProjectPipelineRunsResponseBodyData) *ListCrossProjectPipelineRunsResponseBody {
	s.Data = v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBody) SetRequestId(v string) *ListCrossProjectPipelineRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBody) SetSuccess(v bool) *ListCrossProjectPipelineRunsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrossProjectPipelineRunsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of cross-workspace publish pipelines that match the query conditions.
	//
	// example:
	//
	// [{"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PipelineRunId":"fcfd4160-e2ff-4603-9719-09128fe733df","DeploymentEnvironmentId":101,"ObjectId":"1","ObjectType":"ODPS_SQL","ObjectName":"object-1","ObjectVersion":"7","ChangeType":"ADD","Status":"Ready","Description":"Publish objects that have been published in the source project to the target project","Creator":"creator","CreateTime":1788739200000}]
	PipelineRuns []*ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns `json:"PipelineRuns,omitempty" xml:"PipelineRuns,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrossProjectPipelineRunsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) GetPipelineRuns() []*ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	return s.PipelineRuns
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) SetPageNumber(v int32) *ListCrossProjectPipelineRunsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) SetPageSize(v int32) *ListCrossProjectPipelineRunsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) SetPipelineRuns(v []*ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) *ListCrossProjectPipelineRunsResponseBodyData {
	s.PipelineRuns = v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) SetRequestId(v string) *ListCrossProjectPipelineRunsResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) SetTotalCount(v int32) *ListCrossProjectPipelineRunsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyData) Validate() error {
	if s.PipelineRuns != nil {
		for _, item := range s.PipelineRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns struct {
	// The termination time. This value is a UNIX timestamp in milliseconds. This parameter is returned only after the pipeline is terminated.
	//
	// example:
	//
	// 1788739260000
	AbolishTime *int64 `json:"AbolishTime,omitempty" xml:"AbolishTime,omitempty"`
	// The user who terminated the pipeline.
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
	// The cross-workspace publish environment ID.
	//
	// example:
	//
	// 101
	DeploymentEnvironmentId *int64 `json:"DeploymentEnvironmentId,omitempty" xml:"DeploymentEnvironmentId,omitempty"`
	// The description of the publish operation.
	//
	// example:
	//
	// Publish objects that have been published in the source project to the target project
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
	// The publish object ID.
	//
	// example:
	//
	// 1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the published object.
	//
	// example:
	//
	// object-1
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The publish object type.
	//
	// example:
	//
	// ODPS_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The version of the published object.
	//
	// example:
	//
	// 7
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
	// The ID of the cross-workspace publish pipeline.
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
	// The publish flow status. Valid values:
	//
	// - Building: Building.
	//
	// - Ready: Ready and waiting for execution.
	//
	// - Running: Running.
	//
	// - Termination: Terminated.
	//
	// - Success: Succeeded.
	//
	// - Fail: Failed.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetAbolishTime() *int64 {
	return s.AbolishTime
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetAbolisher() *string {
	return s.Abolisher
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetChangeType() *string {
	return s.ChangeType
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetCreator() *string {
	return s.Creator
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetDeploymentEnvironmentId() *int64 {
	return s.DeploymentEnvironmentId
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetDescription() *string {
	return s.Description
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetExecutor() *string {
	return s.Executor
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) GetStatus() *string {
	return s.Status
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetAbolishTime(v int64) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.AbolishTime = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetAbolisher(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.Abolisher = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetChangeType(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ChangeType = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetCreateTime(v int64) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.CreateTime = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetCreator(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.Creator = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetDeploymentEnvironmentId(v int64) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.DeploymentEnvironmentId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetDescription(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.Description = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetErrorCode(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ErrorCode = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetErrorMessage(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ErrorMessage = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetExecuteTime(v int64) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ExecuteTime = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetExecutor(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.Executor = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetFinishTime(v int64) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.FinishTime = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetObjectId(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ObjectId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetObjectName(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ObjectName = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetObjectType(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ObjectType = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetObjectVersion(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.ObjectVersion = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetPipelineRunId(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.PipelineRunId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetRequestId(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) SetStatus(v string) *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns {
	s.Status = &v
	return s
}

func (s *ListCrossProjectPipelineRunsResponseBodyDataPipelineRuns) Validate() error {
	return dara.Validate(s)
}
