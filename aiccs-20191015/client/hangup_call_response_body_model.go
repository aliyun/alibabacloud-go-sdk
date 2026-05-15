// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HangupCallResponseBody
	GetCode() *string
	SetMessage(v string) *HangupCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *HangupCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HangupCallResponseBody
	GetSuccess() *bool
}

type HangupCallResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxx
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

func (s HangupCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HangupCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *HangupCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HangupCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HangupCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HangupCallResponseBody) SetCode(v string) *HangupCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupCallResponseBody) SetMessage(v string) *HangupCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupCallResponseBody) SetRequestId(v string) *HangupCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupCallResponseBody) SetSuccess(v bool) *HangupCallResponseBody {
	s.Success = &v
	return s
}

func (s *HangupCallResponseBody) Validate() error {
	return dara.Validate(s)
}
