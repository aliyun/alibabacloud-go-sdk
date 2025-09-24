// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeBillToOSSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnsubscribeBillToOSSResponseBody
	GetCode() *string
	SetMessage(v string) *UnsubscribeBillToOSSResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnsubscribeBillToOSSResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnsubscribeBillToOSSResponseBody
	GetSuccess() *bool
}

type UnsubscribeBillToOSSResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D23FE74C-742F-4624-A82B-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnsubscribeBillToOSSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeBillToOSSResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeBillToOSSResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnsubscribeBillToOSSResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnsubscribeBillToOSSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnsubscribeBillToOSSResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnsubscribeBillToOSSResponseBody) SetCode(v string) *UnsubscribeBillToOSSResponseBody {
	s.Code = &v
	return s
}

func (s *UnsubscribeBillToOSSResponseBody) SetMessage(v string) *UnsubscribeBillToOSSResponseBody {
	s.Message = &v
	return s
}

func (s *UnsubscribeBillToOSSResponseBody) SetRequestId(v string) *UnsubscribeBillToOSSResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeBillToOSSResponseBody) SetSuccess(v bool) *UnsubscribeBillToOSSResponseBody {
	s.Success = &v
	return s
}

func (s *UnsubscribeBillToOSSResponseBody) Validate() error {
	return dara.Validate(s)
}
