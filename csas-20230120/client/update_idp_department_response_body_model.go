// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdpDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIdpDepartmentResponseBody
	GetRequestId() *string
}

type UpdateIdpDepartmentResponseBody struct {
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIdpDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdpDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIdpDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIdpDepartmentResponseBody) SetRequestId(v string) *UpdateIdpDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIdpDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
