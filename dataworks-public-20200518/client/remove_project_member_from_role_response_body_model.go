// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectMemberFromRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveProjectMemberFromRoleResponseBody
	GetRequestId() *string
}

type RemoveProjectMemberFromRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveProjectMemberFromRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberFromRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberFromRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveProjectMemberFromRoleResponseBody) SetRequestId(v string) *RemoveProjectMemberFromRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveProjectMemberFromRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
