// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ManualCallbackResponseBody
	GetCode() *int32
	SetMessage(v string) *ManualCallbackResponseBody
	GetMessage() *string
	SetRequestId(v string) *ManualCallbackResponseBody
	GetRequestId() *string
}

type ManualCallbackResponseBody struct {
	// Error code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message information
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManualCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ManualCallbackResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ManualCallbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ManualCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManualCallbackResponseBody) SetCode(v int32) *ManualCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *ManualCallbackResponseBody) SetMessage(v string) *ManualCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *ManualCallbackResponseBody) SetRequestId(v string) *ManualCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManualCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
