// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormationCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFormationCrawlerResponseBody
	GetCode() *string
	SetData(v bool) *UpdateFormationCrawlerResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateFormationCrawlerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateFormationCrawlerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFormationCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFormationCrawlerResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *UpdateFormationCrawlerResponseBody
	GetTaskId() *int64
}

type UpdateFormationCrawlerResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is processed successfully.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information about the call result. Valid values:
	//
	// - If the request is successful, OK is returned.
	//
	// - If the request fails, a specific error code is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 15
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateFormationCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormationCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFormationCrawlerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFormationCrawlerResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateFormationCrawlerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateFormationCrawlerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFormationCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFormationCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFormationCrawlerResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateFormationCrawlerResponseBody) SetCode(v string) *UpdateFormationCrawlerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFormationCrawlerResponseBody) SetData(v bool) *UpdateFormationCrawlerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateFormationCrawlerResponseBody) SetHttpStatusCode(v int32) *UpdateFormationCrawlerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFormationCrawlerResponseBody) SetMessage(v string) *UpdateFormationCrawlerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFormationCrawlerResponseBody) SetRequestId(v string) *UpdateFormationCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFormationCrawlerResponseBody) SetSuccess(v bool) *UpdateFormationCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFormationCrawlerResponseBody) SetTaskId(v int64) *UpdateFormationCrawlerResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateFormationCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
