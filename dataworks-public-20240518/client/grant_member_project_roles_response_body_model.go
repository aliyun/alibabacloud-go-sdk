// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantMemberProjectRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantMemberProjectRolesResponseBody
	GetRequestId() *string
}

type GrantMemberProjectRolesResponseBody struct {
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 2d9ced66-38ef-4923-baf6-391dd3a7e656
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantMemberProjectRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantMemberProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantMemberProjectRolesResponseBody) SetRequestId(v string) *GrantMemberProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantMemberProjectRolesResponseBody) Validate() error {
	return dara.Validate(s)
}
