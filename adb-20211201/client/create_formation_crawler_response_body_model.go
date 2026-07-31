// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFormationCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFormationCrawlerResponseBody
	GetCode() *string
	SetData(v bool) *CreateFormationCrawlerResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CreateFormationCrawlerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateFormationCrawlerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFormationCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFormationCrawlerResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateFormationCrawlerResponseBody
	GetTaskId() *string
}

type CreateFormationCrawlerResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is processed.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message. A value of OK indicates success.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019F3BE7-E8FA-3DC5-8EE7-501A90B5A54D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: The call is successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The operation result. A value of true indicates that the task is created.
	//
	// example:
	//
	// 241
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFormationCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFormationCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFormationCrawlerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFormationCrawlerResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateFormationCrawlerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateFormationCrawlerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFormationCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFormationCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFormationCrawlerResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateFormationCrawlerResponseBody) SetCode(v string) *CreateFormationCrawlerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFormationCrawlerResponseBody) SetData(v bool) *CreateFormationCrawlerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFormationCrawlerResponseBody) SetHttpStatusCode(v int32) *CreateFormationCrawlerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFormationCrawlerResponseBody) SetMessage(v string) *CreateFormationCrawlerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFormationCrawlerResponseBody) SetRequestId(v string) *CreateFormationCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFormationCrawlerResponseBody) SetSuccess(v bool) *CreateFormationCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFormationCrawlerResponseBody) SetTaskId(v string) *CreateFormationCrawlerResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateFormationCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
