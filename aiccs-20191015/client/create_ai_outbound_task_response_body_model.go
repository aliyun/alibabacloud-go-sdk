// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAiOutboundTaskResponseBody
	GetCode() *string
	SetData(v int64) *CreateAiOutboundTaskResponseBody
	GetData() *int64
	SetMessage(v string) *CreateAiOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAiOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAiOutboundTaskResponseBody
	GetSuccess() *bool
}

type CreateAiOutboundTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123456
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAiOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAiOutboundTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateAiOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAiOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAiOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAiOutboundTaskResponseBody) SetCode(v string) *CreateAiOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAiOutboundTaskResponseBody) SetData(v int64) *CreateAiOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAiOutboundTaskResponseBody) SetMessage(v string) *CreateAiOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAiOutboundTaskResponseBody) SetRequestId(v string) *CreateAiOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiOutboundTaskResponseBody) SetSuccess(v bool) *CreateAiOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAiOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
