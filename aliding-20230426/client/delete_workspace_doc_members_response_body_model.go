// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceDocMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkspaceDocMembersResponseBody
	GetRequestId() *string
}

type DeleteWorkspaceDocMembersResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteWorkspaceDocMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceDocMembersResponseBody) SetRequestId(v string) *DeleteWorkspaceDocMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
