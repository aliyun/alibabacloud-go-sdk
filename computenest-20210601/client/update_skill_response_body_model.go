// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSkillResponseBody
	GetRequestId() *string
}

type UpdateSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillResponseBody) SetRequestId(v string) *UpdateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
