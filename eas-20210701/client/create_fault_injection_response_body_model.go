// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaultInjectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateFaultInjectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFaultInjectionResponseBody
	GetRequestId() *string
}

type CreateFaultInjectionResponseBody struct {
	// example:
	//
	// Addon prometheus_discovery is successfully reinstalled
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFaultInjectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFaultInjectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaultInjectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFaultInjectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFaultInjectionResponseBody) SetMessage(v string) *CreateFaultInjectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFaultInjectionResponseBody) SetRequestId(v string) *CreateFaultInjectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFaultInjectionResponseBody) Validate() error {
	return dara.Validate(s)
}
