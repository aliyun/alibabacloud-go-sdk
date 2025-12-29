// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateMessageTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMessageTemplateResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateMessageTemplateResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *UpdateMessageTemplateResponseBody
	GetStatusCode() *int32
}

type UpdateMessageTemplateResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateMessageTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMessageTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMessageTemplateResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateMessageTemplateResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMessageTemplateResponseBody) SetMessage(v string) *UpdateMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMessageTemplateResponseBody) SetRequestId(v string) *UpdateMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMessageTemplateResponseBody) SetResult(v bool) *UpdateMessageTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateMessageTemplateResponseBody) SetStatusCode(v int32) *UpdateMessageTemplateResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
