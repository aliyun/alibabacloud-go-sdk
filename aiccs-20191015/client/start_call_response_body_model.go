// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartCallResponseBody
	GetCode() *string
	SetMessage(v string) *StartCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartCallResponseBody
	GetSuccess() *bool
}

type StartCallResponseBody struct {
	// Status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartCallResponseBody) SetCode(v string) *StartCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallResponseBody) SetMessage(v string) *StartCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallResponseBody) SetRequestId(v string) *StartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallResponseBody) SetSuccess(v bool) *StartCallResponseBody {
	s.Success = &v
	return s
}

func (s *StartCallResponseBody) Validate() error {
	return dara.Validate(s)
}
