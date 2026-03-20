// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateListenerAttributeResponseBody
	GetRequestId() *string
}

type UpdateListenerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7DBFC67C-A272-5952-8287-6C3EBE4E04D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateListenerAttributeResponseBody) SetRequestId(v string) *UpdateListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
