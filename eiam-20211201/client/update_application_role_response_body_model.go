// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationRoleResponseBody
	GetRequestId() *string
}

type UpdateApplicationRoleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationRoleResponseBody) SetRequestId(v string) *UpdateApplicationRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
