// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *StopServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopServiceResponseBody
	GetRequestId() *string
}

type StopServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to auto scale service [foo]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopServiceResponseBody) SetMessage(v string) *StopServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StopServiceResponseBody) SetRequestId(v string) *StopServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
