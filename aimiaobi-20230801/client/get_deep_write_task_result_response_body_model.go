// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeepWriteTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeepWriteTaskResultResponseBody
	GetCode() *string
	SetData(v *GetDeepWriteTaskResultResponseBodyData) *GetDeepWriteTaskResultResponseBody
	GetData() *GetDeepWriteTaskResultResponseBodyData
	SetHttpStatusCode(v int32) *GetDeepWriteTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDeepWriteTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeepWriteTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeepWriteTaskResultResponseBody
	GetSuccess() *bool
}

type GetDeepWriteTaskResultResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDeepWriteTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeepWriteTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeepWriteTaskResultResponseBody) GetData() *GetDeepWriteTaskResultResponseBodyData {
	return s.Data
}

func (s *GetDeepWriteTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDeepWriteTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeepWriteTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeepWriteTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeepWriteTaskResultResponseBody) SetCode(v string) *GetDeepWriteTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeepWriteTaskResultResponseBody) SetData(v *GetDeepWriteTaskResultResponseBodyData) *GetDeepWriteTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDeepWriteTaskResultResponseBody) SetHttpStatusCode(v int32) *GetDeepWriteTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDeepWriteTaskResultResponseBody) SetMessage(v string) *GetDeepWriteTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeepWriteTaskResultResponseBody) SetRequestId(v string) *GetDeepWriteTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeepWriteTaskResultResponseBody) SetSuccess(v bool) *GetDeepWriteTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeepWriteTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeepWriteTaskResultResponseBodyData struct {
	// example:
	//
	// https://aimiaobi-service-pre-hangzhou.oss-cn-hangzhou.aliyuncs.com/aimiaobi/deep-write-workspace/142***1/dbaaebd1-eb1b-41e8-9b99-******-result.zip?Expire=1111
	ArtifactUrl *string `json:"ArtifactUrl,omitempty" xml:"ArtifactUrl,omitempty"`
	// example:
	//
	// f8707efa-c30e-407f-a611-50871aa68952
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDeepWriteTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskResultResponseBodyData) GetArtifactUrl() *string {
	return s.ArtifactUrl
}

func (s *GetDeepWriteTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDeepWriteTaskResultResponseBodyData) SetArtifactUrl(v string) *GetDeepWriteTaskResultResponseBodyData {
	s.ArtifactUrl = &v
	return s
}

func (s *GetDeepWriteTaskResultResponseBodyData) SetTaskId(v string) *GetDeepWriteTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDeepWriteTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
