// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneRunningStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPtsSceneRunningStatusResponseBody
	GetCode() *string
	SetCreateTime(v string) *GetPtsSceneRunningStatusResponseBody
	GetCreateTime() *string
	SetHttpStatusCode(v int32) *GetPtsSceneRunningStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPtsSceneRunningStatusResponseBody
	GetMessage() *string
	SetModifiedTime(v string) *GetPtsSceneRunningStatusResponseBody
	GetModifiedTime() *string
	SetRequestId(v string) *GetPtsSceneRunningStatusResponseBody
	GetRequestId() *string
	SetSceneName(v string) *GetPtsSceneRunningStatusResponseBody
	GetSceneName() *string
	SetStatus(v string) *GetPtsSceneRunningStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetPtsSceneRunningStatusResponseBody
	GetSuccess() *bool
}

type GetPtsSceneRunningStatusResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned
	//
	// example:
	//
	// 4001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the scenario was created.
	//
	// example:
	//
	// 2021-03-01 16:05:56
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The request status code. If the operation is successful, this parameter is not returned
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The last modification time of the scenario.
	//
	// example:
	//
	// 2021-03-26 16:03:56
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DC4E3177-6745-4925-B423-4E89VV34221A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the scenario.
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the scenario. Valid values:
	//
	// 	- CREATED
	//
	// 	- SYNCING
	//
	// 	- SYNC_DONE
	//
	// 	- UPLOADING
	//
	// 	- UPLOADED
	//
	// 	- PREPARING
	//
	// 	- READY
	//
	// 	- RUNNING
	//
	// 	- STOPPING
	//
	// 	- STOPPED
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsSceneRunningStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPtsSceneRunningStatusResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPtsSceneRunningStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPtsSceneRunningStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPtsSceneRunningStatusResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetPtsSceneRunningStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPtsSceneRunningStatusResponseBody) GetSceneName() *string {
	return s.SceneName
}

func (s *GetPtsSceneRunningStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetPtsSceneRunningStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPtsSceneRunningStatusResponseBody) SetCode(v string) *GetPtsSceneRunningStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetCreateTime(v string) *GetPtsSceneRunningStatusResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetHttpStatusCode(v int32) *GetPtsSceneRunningStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetMessage(v string) *GetPtsSceneRunningStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetModifiedTime(v string) *GetPtsSceneRunningStatusResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetRequestId(v string) *GetPtsSceneRunningStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetSceneName(v string) *GetPtsSceneRunningStatusResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetStatus(v string) *GetPtsSceneRunningStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetSuccess(v bool) *GetPtsSceneRunningStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
