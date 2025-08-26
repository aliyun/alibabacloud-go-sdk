// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodcastTaskResultQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PodcastTaskResultQueryResponseBody
	GetCode() *string
	SetData(v *PodcastTaskResultQueryResponseBodyData) *PodcastTaskResultQueryResponseBody
	GetData() *PodcastTaskResultQueryResponseBodyData
	SetHttpStatusCode(v string) *PodcastTaskResultQueryResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PodcastTaskResultQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *PodcastTaskResultQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PodcastTaskResultQueryResponseBody
	GetSuccess() *bool
}

type PodcastTaskResultQueryResponseBody struct {
	// example:
	//
	// "success"
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *PodcastTaskResultQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// "success"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C38F034D-7F36-531C-95AC-0C752F80E840
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PodcastTaskResultQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskResultQueryResponseBody) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *PodcastTaskResultQueryResponseBody) GetData() *PodcastTaskResultQueryResponseBodyData {
	return s.Data
}

func (s *PodcastTaskResultQueryResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PodcastTaskResultQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PodcastTaskResultQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PodcastTaskResultQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PodcastTaskResultQueryResponseBody) SetCode(v string) *PodcastTaskResultQueryResponseBody {
	s.Code = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetData(v *PodcastTaskResultQueryResponseBodyData) *PodcastTaskResultQueryResponseBody {
	s.Data = v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetHttpStatusCode(v string) *PodcastTaskResultQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetMessage(v string) *PodcastTaskResultQueryResponseBody {
	s.Message = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetRequestId(v string) *PodcastTaskResultQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) SetSuccess(v bool) *PodcastTaskResultQueryResponseBody {
	s.Success = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type PodcastTaskResultQueryResponseBodyData struct {
	ExtraResult interface{} `json:"extraResult,omitempty" xml:"extraResult,omitempty"`
	// example:
	//
	// {
	//
	//             "audio": "http://note-ai-file.oss-cn-beijing.aliyuncs.com/202503241702148295/audio.mp3?OSSAccessKeyId=LTAI5tPLWJfJHNkZbfnQv245&Expires=1742810788&Signature=b5p83nh443Gr7foqdvgrI4%2FKZVM%3D",
	//
	//             "script": "http://note-ai-file.oss-cn-beijing.aliyuncs.com/202503241702148295/script.txt?OSSAccessKeyId=LTAI5tPLWJfJHNkZbfnQv245&Expires=1742810622&Signature=TBBdikHzOWW3YqDw3sNMTXiMo6A%3D"
	//
	// }
	ResultUrl interface{} `json:"resultUrl,omitempty" xml:"resultUrl,omitempty"`
	Script    *string     `json:"script,omitempty" xml:"script,omitempty"`
	// example:
	//
	// 63c4e0eaab3b4c0db208ecafa990e8d1
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// SUCCEEDED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s PodcastTaskResultQueryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskResultQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryResponseBodyData) GetExtraResult() interface{} {
	return s.ExtraResult
}

func (s *PodcastTaskResultQueryResponseBodyData) GetResultUrl() interface{} {
	return s.ResultUrl
}

func (s *PodcastTaskResultQueryResponseBodyData) GetScript() *string {
	return s.Script
}

func (s *PodcastTaskResultQueryResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PodcastTaskResultQueryResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *PodcastTaskResultQueryResponseBodyData) SetExtraResult(v interface{}) *PodcastTaskResultQueryResponseBodyData {
	s.ExtraResult = v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) SetResultUrl(v interface{}) *PodcastTaskResultQueryResponseBodyData {
	s.ResultUrl = v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) SetScript(v string) *PodcastTaskResultQueryResponseBodyData {
	s.Script = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) SetTaskId(v string) *PodcastTaskResultQueryResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) SetTaskStatus(v string) *PodcastTaskResultQueryResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *PodcastTaskResultQueryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
