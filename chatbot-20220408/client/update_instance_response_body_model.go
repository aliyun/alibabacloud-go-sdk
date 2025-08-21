// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
}

type UpdateInstanceResponseBody struct {
	// example:
	//
	// C45BFEE4-F657-1332-8B47-2C757B94C972
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
