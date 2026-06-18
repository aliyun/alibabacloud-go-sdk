// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinThirdCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *JoinThirdCallResponseBody
	GetCode() *string
	SetMessage(v string) *JoinThirdCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *JoinThirdCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *JoinThirdCallResponseBody
	GetSuccess() *bool
}

type JoinThirdCallResponseBody struct {
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
	// Whether the API call succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinThirdCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *JoinThirdCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *JoinThirdCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinThirdCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *JoinThirdCallResponseBody) SetCode(v string) *JoinThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetMessage(v string) *JoinThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetRequestId(v string) *JoinThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetSuccess(v bool) *JoinThirdCallResponseBody {
	s.Success = &v
	return s
}

func (s *JoinThirdCallResponseBody) Validate() error {
	return dara.Validate(s)
}
