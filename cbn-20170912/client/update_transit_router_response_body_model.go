// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E9963DD7-998B-4F92-892D-8293CB49EE81
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterResponseBody) SetRequestId(v string) *UpdateTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
