// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *StartInstanceResponseBody
	GetCode() *int32
	SetMessage(v string) *StartInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartInstanceResponseBody
	GetSuccess() *bool
}

type StartInstanceResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *StartInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartInstanceResponseBody) SetCode(v int32) *StartInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartInstanceResponseBody) SetMessage(v string) *StartInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetSuccess(v bool) *StartInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
