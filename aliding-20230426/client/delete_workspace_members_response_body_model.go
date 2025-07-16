// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkspaceMembersResponseBody
	GetRequestId() *string
}

type DeleteWorkspaceMembersResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteWorkspaceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceMembersResponseBody) SetRequestId(v string) *DeleteWorkspaceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
