// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkspaceMembersResponseBody
	GetRequestId() *string
}

type UpdateWorkspaceMembersResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateWorkspaceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceMembersResponseBody) SetRequestId(v string) *UpdateWorkspaceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
