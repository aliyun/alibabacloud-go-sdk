// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemberRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddMemberRoleResponseBody
	GetRequestId() *string
}

type AddMemberRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMemberRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMemberRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMemberRoleResponseBody) SetRequestId(v string) *AddMemberRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMemberRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
