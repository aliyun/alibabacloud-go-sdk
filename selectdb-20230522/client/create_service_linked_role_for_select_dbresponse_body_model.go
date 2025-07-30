// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleForSelectDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceLinkedRoleForSelectDBResponseBody
	GetRequestId() *string
}

type CreateServiceLinkedRoleForSelectDBResponseBody struct {
	// example:
	//
	// F203FA74-3041-589F-BE66-E570793A0C91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleForSelectDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleForSelectDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForSelectDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceLinkedRoleForSelectDBResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleForSelectDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBResponseBody) Validate() error {
	return dara.Validate(s)
}
