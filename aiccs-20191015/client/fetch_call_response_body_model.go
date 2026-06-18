// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FetchCallResponseBody
	GetCode() *string
	SetMessage(v string) *FetchCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchCallResponseBody
	GetSuccess() *bool
}

type FetchCallResponseBody struct {
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

func (s FetchCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchCallResponseBody) GoString() string {
	return s.String()
}

func (s *FetchCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *FetchCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchCallResponseBody) SetCode(v string) *FetchCallResponseBody {
	s.Code = &v
	return s
}

func (s *FetchCallResponseBody) SetMessage(v string) *FetchCallResponseBody {
	s.Message = &v
	return s
}

func (s *FetchCallResponseBody) SetRequestId(v string) *FetchCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchCallResponseBody) SetSuccess(v bool) *FetchCallResponseBody {
	s.Success = &v
	return s
}

func (s *FetchCallResponseBody) Validate() error {
	return dara.Validate(s)
}
