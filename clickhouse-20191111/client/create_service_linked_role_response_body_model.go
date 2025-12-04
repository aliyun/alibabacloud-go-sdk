// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceLinkedRoleResponseBody
	GetRequestId() *string
}

type CreateServiceLinkedRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 66428721-FFEC-5023-B4E5-3BD1B67837E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
