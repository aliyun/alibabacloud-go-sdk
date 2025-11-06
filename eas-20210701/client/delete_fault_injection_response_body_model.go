// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaultInjectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteFaultInjectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFaultInjectionResponseBody
	GetRequestId() *string
}

type DeleteFaultInjectionResponseBody struct {
	// example:
	//
	// Successfully delete acl policy for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaultInjectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaultInjectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaultInjectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFaultInjectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaultInjectionResponseBody) SetMessage(v string) *DeleteFaultInjectionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFaultInjectionResponseBody) SetRequestId(v string) *DeleteFaultInjectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaultInjectionResponseBody) Validate() error {
	return dara.Validate(s)
}
