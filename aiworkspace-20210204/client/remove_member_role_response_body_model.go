// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMemberRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveMemberRoleResponseBody
	GetRequestId() *string
}

type RemoveMemberRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveMemberRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveMemberRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMemberRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveMemberRoleResponseBody) SetRequestId(v string) *RemoveMemberRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveMemberRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
