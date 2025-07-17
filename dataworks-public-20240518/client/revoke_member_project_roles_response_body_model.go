// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemberProjectRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeMemberProjectRolesResponseBody
	GetRequestId() *string
}

type RevokeMemberProjectRolesResponseBody struct {
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeMemberProjectRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemberProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeMemberProjectRolesResponseBody) SetRequestId(v string) *RevokeMemberProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeMemberProjectRolesResponseBody) Validate() error {
	return dara.Validate(s)
}
