// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteMessageTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMessageTemplateResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteMessageTemplateResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *DeleteMessageTemplateResponseBody
	GetStatusCode() *int32
}

type DeleteMessageTemplateResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
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

func (s DeleteMessageTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMessageTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMessageTemplateResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteMessageTemplateResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageTemplateResponseBody) SetMessage(v string) *DeleteMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMessageTemplateResponseBody) SetRequestId(v string) *DeleteMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessageTemplateResponseBody) SetResult(v bool) *DeleteMessageTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteMessageTemplateResponseBody) SetStatusCode(v int32) *DeleteMessageTemplateResponseBody {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
