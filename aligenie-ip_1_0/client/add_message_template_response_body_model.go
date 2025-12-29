// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMessageTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddMessageTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *AddMessageTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMessageTemplateResponseBody
	GetRequestId() *string
	SetResult(v int64) *AddMessageTemplateResponseBody
	GetResult() *int64
}

type AddMessageTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 11
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddMessageTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMessageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddMessageTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMessageTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMessageTemplateResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *AddMessageTemplateResponseBody) SetCode(v int32) *AddMessageTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *AddMessageTemplateResponseBody) SetMessage(v string) *AddMessageTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *AddMessageTemplateResponseBody) SetRequestId(v string) *AddMessageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMessageTemplateResponseBody) SetResult(v int64) *AddMessageTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *AddMessageTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
