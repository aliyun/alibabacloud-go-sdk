// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamInstanceRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListStreamInstanceRecordsResponseBody
	GetCode() *string
	SetData(v []*ListStreamInstanceRecordsResponseBodyData) *ListStreamInstanceRecordsResponseBody
	GetData() []*ListStreamInstanceRecordsResponseBodyData
	SetHttpStatusCode(v int32) *ListStreamInstanceRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListStreamInstanceRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListStreamInstanceRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListStreamInstanceRecordsResponseBody
	GetSuccess() *bool
}

type ListStreamInstanceRecordsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The historical running instance records of the real-time task. You can use the deploymentInstanceId to query the logs of a specific run.
	Data []*ListStreamInstanceRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListStreamInstanceRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamInstanceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamInstanceRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListStreamInstanceRecordsResponseBody) GetData() []*ListStreamInstanceRecordsResponseBodyData {
	return s.Data
}

func (s *ListStreamInstanceRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListStreamInstanceRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStreamInstanceRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamInstanceRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStreamInstanceRecordsResponseBody) SetCode(v string) *ListStreamInstanceRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBody) SetData(v []*ListStreamInstanceRecordsResponseBodyData) *ListStreamInstanceRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListStreamInstanceRecordsResponseBody) SetHttpStatusCode(v int32) *ListStreamInstanceRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBody) SetMessage(v string) *ListStreamInstanceRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBody) SetRequestId(v string) *ListStreamInstanceRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBody) SetSuccess(v bool) *ListStreamInstanceRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBody) Validate() error {
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

type ListStreamInstanceRecordsResponseBodyData struct {
	// The deployment ID of the task. A single deployment can be associated with multiple running instances.
	//
	// example:
	//
	// ac6e5a34bbd44d7f8b7851ef28c90d12
	DeploymentId *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	// The instance ID of the deployment running instance. This value corresponds to the DeploymentInstanceId request parameter in the log and TaskManager query operations.
	//
	// example:
	//
	// 7db678957dc449aab0086b931c9f8218
	DeploymentInstanceId *string `json:"DeploymentInstanceId,omitempty" xml:"DeploymentInstanceId,omitempty"`
	// The relative path for dirty data storage in this run, excluding the storage service domain name. A value of null indicates that no dirty data path exists.
	//
	// example:
	//
	// dirty-data/7333610201479808/7db678957dc449aab0086b931c9f8218/error.log
	ErrorDataPath *string `json:"ErrorDataPath,omitempty" xml:"ErrorDataPath,omitempty"`
	// The start time of this run. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-10-15 13:46:26
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The running instance status returned by the engine, such as RUNNING, FINISHED, and TRANSITIONING. This status uses a different enumeration from the TaskStatusEnum returned by the GetStreamTask operation.
	//
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListStreamInstanceRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStreamInstanceRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStreamInstanceRecordsResponseBodyData) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ListStreamInstanceRecordsResponseBodyData) GetDeploymentInstanceId() *string {
	return s.DeploymentInstanceId
}

func (s *ListStreamInstanceRecordsResponseBodyData) GetErrorDataPath() *string {
	return s.ErrorDataPath
}

func (s *ListStreamInstanceRecordsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListStreamInstanceRecordsResponseBodyData) GetState() *string {
	return s.State
}

func (s *ListStreamInstanceRecordsResponseBodyData) SetDeploymentId(v string) *ListStreamInstanceRecordsResponseBodyData {
	s.DeploymentId = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBodyData) SetDeploymentInstanceId(v string) *ListStreamInstanceRecordsResponseBodyData {
	s.DeploymentInstanceId = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBodyData) SetErrorDataPath(v string) *ListStreamInstanceRecordsResponseBodyData {
	s.ErrorDataPath = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBodyData) SetStartTime(v string) *ListStreamInstanceRecordsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBodyData) SetState(v string) *ListStreamInstanceRecordsResponseBodyData {
	s.State = &v
	return s
}

func (s *ListStreamInstanceRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
