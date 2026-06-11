// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSkillDraftResponseBody
	GetRequestId() *string
}

type UpdateSkillDraftResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSkillDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillDraftResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillDraftResponseBody) SetRequestId(v string) *UpdateSkillDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillDraftResponseBody) Validate() error {
	return dara.Validate(s)
}
