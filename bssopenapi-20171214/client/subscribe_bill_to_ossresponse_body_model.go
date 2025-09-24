// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeBillToOSSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubscribeBillToOSSResponseBody
	GetCode() *string
	SetMessage(v string) *SubscribeBillToOSSResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubscribeBillToOSSResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubscribeBillToOSSResponseBody
	GetSuccess() *bool
}

type SubscribeBillToOSSResponseBody struct {
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
	// F61FCE4B-9B56-4FD9-A17E-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubscribeBillToOSSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubscribeBillToOSSResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeBillToOSSResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubscribeBillToOSSResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubscribeBillToOSSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubscribeBillToOSSResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubscribeBillToOSSResponseBody) SetCode(v string) *SubscribeBillToOSSResponseBody {
	s.Code = &v
	return s
}

func (s *SubscribeBillToOSSResponseBody) SetMessage(v string) *SubscribeBillToOSSResponseBody {
	s.Message = &v
	return s
}

func (s *SubscribeBillToOSSResponseBody) SetRequestId(v string) *SubscribeBillToOSSResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeBillToOSSResponseBody) SetSuccess(v bool) *SubscribeBillToOSSResponseBody {
	s.Success = &v
	return s
}

func (s *SubscribeBillToOSSResponseBody) Validate() error {
	return dara.Validate(s)
}
