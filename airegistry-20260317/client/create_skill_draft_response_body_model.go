// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateSkillDraftResponseBody
	GetData() *string
	SetRequestId(v string) *CreateSkillDraftResponseBody
	GetRequestId() *string
}

type CreateSkillDraftResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSkillDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillDraftResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillDraftResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSkillDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillDraftResponseBody) SetData(v string) *CreateSkillDraftResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillDraftResponseBody) SetRequestId(v string) *CreateSkillDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillDraftResponseBody) Validate() error {
	return dara.Validate(s)
}
