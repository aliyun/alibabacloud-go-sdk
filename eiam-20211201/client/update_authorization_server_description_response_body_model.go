// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationServerDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAuthorizationServerDescriptionResponseBody
	GetRequestId() *string
}

type UpdateAuthorizationServerDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthorizationServerDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationServerDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationServerDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorizationServerDescriptionResponseBody) SetRequestId(v string) *UpdateAuthorizationServerDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorizationServerDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
