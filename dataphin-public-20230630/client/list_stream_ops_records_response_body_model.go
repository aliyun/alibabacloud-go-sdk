// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamOpsRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListStreamOpsRecordsResponseBody
	GetCode() *string
	SetData(v []*ListStreamOpsRecordsResponseBodyData) *ListStreamOpsRecordsResponseBody
	GetData() []*ListStreamOpsRecordsResponseBodyData
	SetHttpStatusCode(v int32) *ListStreamOpsRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListStreamOpsRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListStreamOpsRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListStreamOpsRecordsResponseBody
	GetSuccess() *bool
}

type ListStreamOpsRecordsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The start and stop operation records for real-time tasks. You can use the log storage identifiers in the list to query the log content by calling the GetStreamOpsRecord operation.
	Data []*ListStreamOpsRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListStreamOpsRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamOpsRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamOpsRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListStreamOpsRecordsResponseBody) GetData() []*ListStreamOpsRecordsResponseBodyData {
	return s.Data
}

func (s *ListStreamOpsRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListStreamOpsRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStreamOpsRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamOpsRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStreamOpsRecordsResponseBody) SetCode(v string) *ListStreamOpsRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBody) SetData(v []*ListStreamOpsRecordsResponseBodyData) *ListStreamOpsRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListStreamOpsRecordsResponseBody) SetHttpStatusCode(v int32) *ListStreamOpsRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBody) SetMessage(v string) *ListStreamOpsRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBody) SetRequestId(v string) *ListStreamOpsRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBody) SetSuccess(v bool) *ListStreamOpsRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStreamOpsRecordsResponseBodyData struct {
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
	// The end time of the operation. The time is in the yyyy-MM-dd HH:mm:ss format. If the operation has not ended, the value may be null.
	//
	// example:
	//
	// 2026-10-15 13:46:40
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The operation logs. In the ListStreamOpsRecords response, this field may contain a log storage identity. After you invoke the GetStreamOpsRecord operation, the log content or operation information is returned.
	//
	// example:
	//
	// 2026-10-15 13:46:26 INFO Starting Flink deployment
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// The type of the operation. Valid values: START, PAUSE, RESUME, and STOP. This field does not represent the running status of the task.
	//
	// example:
	//
	// START
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The ID of the operation record. This ID corresponds to the OpsRecordId request parameter of the GetStreamOpsRecord operation.
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

func (s ListStreamOpsRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStreamOpsRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStreamOpsRecordsResponseBodyData) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ListStreamOpsRecordsResponseBodyData) GetDeploymentInstanceId() *string {
	return s.DeploymentInstanceId
}

func (s *ListStreamOpsRecordsResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListStreamOpsRecordsResponseBodyData) GetLogContent() *string {
	return s.LogContent
}

func (s *ListStreamOpsRecordsResponseBodyData) GetOperate() *string {
	return s.Operate
}

func (s *ListStreamOpsRecordsResponseBodyData) GetOpsRecordId() *string {
	return s.OpsRecordId
}

func (s *ListStreamOpsRecordsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListStreamOpsRecordsResponseBodyData) SetDeploymentId(v string) *ListStreamOpsRecordsResponseBodyData {
	s.DeploymentId = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBodyData) SetDeploymentInstanceId(v string) *ListStreamOpsRecordsResponseBodyData {
	s.DeploymentInstanceId = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBodyData) SetEndTime(v string) *ListStreamOpsRecordsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBodyData) SetLogContent(v string) *ListStreamOpsRecordsResponseBodyData {
	s.LogContent = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBodyData) SetOperate(v string) *ListStreamOpsRecordsResponseBodyData {
	s.Operate = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBodyData) SetOpsRecordId(v string) *ListStreamOpsRecordsResponseBodyData {
	s.OpsRecordId = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBodyData) SetStartTime(v string) *ListStreamOpsRecordsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListStreamOpsRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
