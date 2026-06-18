// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupThirdCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HangupThirdCallResponseBody
	GetCode() *string
	SetMessage(v string) *HangupThirdCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *HangupThirdCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HangupThirdCallResponseBody
	GetSuccess() *bool
}

type HangupThirdCallResponseBody struct {
	// Fault encoding
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Fault description
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, used to trail the cause of a fault
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangupThirdCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HangupThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *HangupThirdCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HangupThirdCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HangupThirdCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HangupThirdCallResponseBody) SetCode(v string) *HangupThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetMessage(v string) *HangupThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetRequestId(v string) *HangupThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetSuccess(v bool) *HangupThirdCallResponseBody {
	s.Success = &v
	return s
}

func (s *HangupThirdCallResponseBody) Validate() error {
	return dara.Validate(s)
}
