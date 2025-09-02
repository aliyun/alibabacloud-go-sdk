// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateProjectMemberResponseBody
	GetRequestId() *string
}

type CreateProjectMemberResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectMemberResponseBody) SetRequestId(v string) *CreateProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
