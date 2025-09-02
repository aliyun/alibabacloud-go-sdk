// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProjectMemberResponseBody
	GetRequestId() *string
}

type DeleteProjectMemberResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectMemberResponseBody) SetRequestId(v string) *DeleteProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
