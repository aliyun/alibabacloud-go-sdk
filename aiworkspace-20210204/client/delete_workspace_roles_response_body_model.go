// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkspaceRolesResponseBody
	GetRequestId() *string
}

type DeleteWorkspaceRolesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A519F77D-28A0-52F5-AB82-5********8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkspaceRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceRolesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceRolesResponseBody) SetRequestId(v string) *DeleteWorkspaceRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceRolesResponseBody) Validate() error {
	return dara.Validate(s)
}
