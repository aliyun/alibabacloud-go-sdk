// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UnbindSlbResponseBody
	GetCode() *int32
	SetData(v string) *UnbindSlbResponseBody
	GetData() *string
	SetMessage(v string) *UnbindSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindSlbResponseBody
	GetRequestId() *string
}

type UnbindSlbResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is left empty. It has no meaning.
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Unbind slb success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnbindSlbResponseBody) GetData() *string {
	return s.Data
}

func (s *UnbindSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindSlbResponseBody) SetCode(v int32) *UnbindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSlbResponseBody) SetData(v string) *UnbindSlbResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindSlbResponseBody) SetMessage(v string) *UnbindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSlbResponseBody) SetRequestId(v string) *UnbindSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
