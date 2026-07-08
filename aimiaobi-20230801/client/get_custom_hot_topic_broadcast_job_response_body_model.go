// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomHotTopicBroadcastJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomHotTopicBroadcastJobResponseBody
	GetCode() *string
	SetData(v *GetCustomHotTopicBroadcastJobResponseBodyData) *GetCustomHotTopicBroadcastJobResponseBody
	GetData() *GetCustomHotTopicBroadcastJobResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomHotTopicBroadcastJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomHotTopicBroadcastJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomHotTopicBroadcastJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomHotTopicBroadcastJobResponseBody
	GetSuccess() *bool
}

type GetCustomHotTopicBroadcastJobResponseBody struct {
	// The status code.
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetCustomHotTopicBroadcastJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error details.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomHotTopicBroadcastJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHotTopicBroadcastJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) GetData() *GetCustomHotTopicBroadcastJobResponseBodyData {
	return s.Data
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) SetCode(v string) *GetCustomHotTopicBroadcastJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) SetData(v *GetCustomHotTopicBroadcastJobResponseBodyData) *GetCustomHotTopicBroadcastJobResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) SetHttpStatusCode(v int32) *GetCustomHotTopicBroadcastJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) SetMessage(v string) *GetCustomHotTopicBroadcastJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) SetRequestId(v string) *GetCustomHotTopicBroadcastJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) SetSuccess(v bool) *GetCustomHotTopicBroadcastJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomHotTopicBroadcastJobResponseBodyData struct {
	// The error message.
	//
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The version number of the hot topic.
	//
	// example:
	//
	// 热点话题版本号标识
	HotTopicVersion *string `json:"HotTopicVersion,omitempty" xml:"HotTopicVersion,omitempty"`
	// The status of the task.
	//
	// - PENDING: The task is pending.
	//
	// - RUNNING: The task is running.
	//
	// - SUCCESSED: The task is successful.
	//
	// - SUSPENDED: The task is paused.
	//
	// - FAILED: The task failed.
	//
	// - CANCELED: The task is canceled.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCustomHotTopicBroadcastJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHotTopicBroadcastJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomHotTopicBroadcastJobResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCustomHotTopicBroadcastJobResponseBodyData) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *GetCustomHotTopicBroadcastJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCustomHotTopicBroadcastJobResponseBodyData) SetErrorMessage(v string) *GetCustomHotTopicBroadcastJobResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBodyData) SetHotTopicVersion(v string) *GetCustomHotTopicBroadcastJobResponseBodyData {
	s.HotTopicVersion = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBodyData) SetStatus(v string) *GetCustomHotTopicBroadcastJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
