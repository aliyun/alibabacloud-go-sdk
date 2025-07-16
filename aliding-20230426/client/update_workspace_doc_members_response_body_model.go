// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDocMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkspaceDocMembersResponseBody
	GetRequestId() *string
}

type UpdateWorkspaceDocMembersResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateWorkspaceDocMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceDocMembersResponseBody) SetRequestId(v string) *UpdateWorkspaceDocMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
