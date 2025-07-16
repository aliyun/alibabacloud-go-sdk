// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceDocMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWorkspaceDocMembersResponseBody
	GetRequestId() *string
}

type AddWorkspaceDocMembersResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddWorkspaceDocMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorkspaceDocMembersResponseBody) SetRequestId(v string) *AddWorkspaceDocMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkspaceDocMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
