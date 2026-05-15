// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProjectRoleResponseBody
	GetRequestId() *string
}

type DeleteProjectRoleResponseBody struct {
	// example:
	//
	// 5BCD2252-F184-55A8-9675-337C43BE0CD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectRoleResponseBody) SetRequestId(v string) *DeleteProjectRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
