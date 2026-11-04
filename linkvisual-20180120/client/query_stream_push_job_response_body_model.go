// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStreamPushJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryStreamPushJobResponseBody
	GetCode() *string
	SetData(v *QueryStreamPushJobResponseBodyData) *QueryStreamPushJobResponseBody
	GetData() *QueryStreamPushJobResponseBodyData
	SetErrorMessage(v string) *QueryStreamPushJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryStreamPushJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryStreamPushJobResponseBody
	GetSuccess() *bool
}

type QueryStreamPushJobResponseBody struct {
	// The response code.
	//
	// - **200**: The call is successful.
	//
	// - Other values: An error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 400
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result returned if the call is successful.
	Data *QueryStreamPushJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Request parameter error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// E55***B7-4***-4***-8***-D3******F565
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStreamPushJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryStreamPushJobResponseBody) GetData() *QueryStreamPushJobResponseBodyData {
	return s.Data
}

func (s *QueryStreamPushJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryStreamPushJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryStreamPushJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryStreamPushJobResponseBody) SetCode(v string) *QueryStreamPushJobResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetData(v *QueryStreamPushJobResponseBodyData) *QueryStreamPushJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetErrorMessage(v string) *QueryStreamPushJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetRequestId(v string) *QueryStreamPushJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStreamPushJobResponseBody) SetSuccess(v bool) *QueryStreamPushJobResponseBody {
	s.Success = &v
	return s
}

func (s *QueryStreamPushJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryStreamPushJobResponseBodyData struct {
	// The time when the job was created. Unit: seconds.
	//
	// example:
	//
	// 1581595942
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the stream push job. Valid values:
	//
	// - **1**: Continuous job.
	//
	// - **2**: Triggered job, which ends after an interruption.
	//
	// example:
	//
	// 1
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The current stream push status. Valid values:
	//
	// - **1**: Pushing.
	//
	// - **2**: Waiting to push.
	//
	// - **3**: Push ended.
	//
	// example:
	//
	// 1
	PushStatus *int32 `json:"PushStatus,omitempty" xml:"PushStatus,omitempty"`
	// The stream push URL.
	//
	// example:
	//
	// rtmp://****
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// The stream type. Valid values:
	//
	// - **0**: Main stream.
	//
	// - **1**: Sub stream.
	//
	// example:
	//
	// 1
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryStreamPushJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobResponseBodyData) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *QueryStreamPushJobResponseBodyData) GetJobType() *int32 {
	return s.JobType
}

func (s *QueryStreamPushJobResponseBodyData) GetPushStatus() *int32 {
	return s.PushStatus
}

func (s *QueryStreamPushJobResponseBodyData) GetPushUrl() *string {
	return s.PushUrl
}

func (s *QueryStreamPushJobResponseBodyData) GetStreamType() *int32 {
	return s.StreamType
}

func (s *QueryStreamPushJobResponseBodyData) SetCreateTime(v int32) *QueryStreamPushJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetJobType(v int32) *QueryStreamPushJobResponseBodyData {
	s.JobType = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetPushStatus(v int32) *QueryStreamPushJobResponseBodyData {
	s.PushStatus = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetPushUrl(v string) *QueryStreamPushJobResponseBodyData {
	s.PushUrl = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) SetStreamType(v int32) *QueryStreamPushJobResponseBodyData {
	s.StreamType = &v
	return s
}

func (s *QueryStreamPushJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
