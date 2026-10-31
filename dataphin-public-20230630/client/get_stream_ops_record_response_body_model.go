// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamOpsRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStreamOpsRecordResponseBody
	GetCode() *string
	SetData(v *GetStreamOpsRecordResponseBodyData) *GetStreamOpsRecordResponseBody
	GetData() *GetStreamOpsRecordResponseBodyData
	SetHttpStatusCode(v int32) *GetStreamOpsRecordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStreamOpsRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStreamOpsRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStreamOpsRecordResponseBody
	GetSuccess() *bool
}

type GetStreamOpsRecordResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The start or stop operation record of the real-time task. You can call the GetStreamOpsRecord operation to query the log content by using the log storage identifier in the list.
	Data *GetStreamOpsRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend response exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStreamOpsRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStreamOpsRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetStreamOpsRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStreamOpsRecordResponseBody) GetData() *GetStreamOpsRecordResponseBodyData {
	return s.Data
}

func (s *GetStreamOpsRecordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStreamOpsRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStreamOpsRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStreamOpsRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStreamOpsRecordResponseBody) SetCode(v string) *GetStreamOpsRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetStreamOpsRecordResponseBody) SetData(v *GetStreamOpsRecordResponseBodyData) *GetStreamOpsRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetStreamOpsRecordResponseBody) SetHttpStatusCode(v int32) *GetStreamOpsRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStreamOpsRecordResponseBody) SetMessage(v string) *GetStreamOpsRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetStreamOpsRecordResponseBody) SetRequestId(v string) *GetStreamOpsRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStreamOpsRecordResponseBody) SetSuccess(v bool) *GetStreamOpsRecordResponseBody {
	s.Success = &v
	return s
}

func (s *GetStreamOpsRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStreamOpsRecordResponseBodyData struct {
	// The deployment ID of the task. This ID is used to identify the deployment and is different from the specific running instance ID deploymentInstanceId.
	//
	// example:
	//
	// ac6e5a34bbd44d7f8b7851ef28c90d12
	DeploymentId *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	// The deployment instance ID associated with this operation. This ID corresponds to the DeploymentInstanceId request parameter in the log and TaskManager query operations.
	//
	// example:
	//
	// 7db678957dc449aab0086b931c9f8218
	DeploymentInstanceId *string `json:"DeploymentInstanceId,omitempty" xml:"DeploymentInstanceId,omitempty"`
	// The end time of the operation. The time is in the yyyy-MM-dd HH:mm:ss format. If the operation has not ended, this parameter may be null.
	//
	// example:
	//
	// 2026-10-15 13:46:40
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The operation logs. In the ListStreamOpsRecords operation, this parameter may be a log storage identity. After you invoke the GetStreamOpsRecord operation, the log content or operation information is returned.
	//
	// example:
	//
	// 2026-10-15 13:46:26 INFO Starting Flink deployment
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// The operation type. Valid values: START, PAUSE, RESUME, and STOP. This parameter does not indicate the running status of the task.
	//
	// example:
	//
	// START
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The ID of the operation record. This ID corresponds to the OpsRecordId request parameter in the GetStreamOpsRecord operation.
	//
	// example:
	//
	// 646
	OpsRecordId *string `json:"OpsRecordId,omitempty" xml:"OpsRecordId,omitempty"`
	// The start time of the operation. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-10-15 13:46:26
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStreamOpsRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStreamOpsRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStreamOpsRecordResponseBodyData) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *GetStreamOpsRecordResponseBodyData) GetDeploymentInstanceId() *string {
	return s.DeploymentInstanceId
}

func (s *GetStreamOpsRecordResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStreamOpsRecordResponseBodyData) GetLogContent() *string {
	return s.LogContent
}

func (s *GetStreamOpsRecordResponseBodyData) GetOperate() *string {
	return s.Operate
}

func (s *GetStreamOpsRecordResponseBodyData) GetOpsRecordId() *string {
	return s.OpsRecordId
}

func (s *GetStreamOpsRecordResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStreamOpsRecordResponseBodyData) SetDeploymentId(v string) *GetStreamOpsRecordResponseBodyData {
	s.DeploymentId = &v
	return s
}

func (s *GetStreamOpsRecordResponseBodyData) SetDeploymentInstanceId(v string) *GetStreamOpsRecordResponseBodyData {
	s.DeploymentInstanceId = &v
	return s
}

func (s *GetStreamOpsRecordResponseBodyData) SetEndTime(v string) *GetStreamOpsRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetStreamOpsRecordResponseBodyData) SetLogContent(v string) *GetStreamOpsRecordResponseBodyData {
	s.LogContent = &v
	return s
}

func (s *GetStreamOpsRecordResponseBodyData) SetOperate(v string) *GetStreamOpsRecordResponseBodyData {
	s.Operate = &v
	return s
}

func (s *GetStreamOpsRecordResponseBodyData) SetOpsRecordId(v string) *GetStreamOpsRecordResponseBodyData {
	s.OpsRecordId = &v
	return s
}

func (s *GetStreamOpsRecordResponseBodyData) SetStartTime(v string) *GetStreamOpsRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetStreamOpsRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
