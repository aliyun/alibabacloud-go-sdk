// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateVirtualBorderRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminateVirtualBorderRouterResponseBody
	GetRequestId() *string
}

type TerminateVirtualBorderRouterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateVirtualBorderRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateVirtualBorderRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateVirtualBorderRouterResponseBody) SetRequestId(v string) *TerminateVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateVirtualBorderRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
