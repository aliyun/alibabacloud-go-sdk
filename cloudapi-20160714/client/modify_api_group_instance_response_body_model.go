// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApiGroupInstanceResponseBody
	GetRequestId() *string
}

type ModifyApiGroupInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E07AEFF0-9FB0-599E-8F12-B418D8AE1F3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiGroupInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApiGroupInstanceResponseBody) SetRequestId(v string) *ModifyApiGroupInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiGroupInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
