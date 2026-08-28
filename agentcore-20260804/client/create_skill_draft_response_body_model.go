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
	// The response data.
	//
	// example:
	//
	// skill-1234567890abcdef
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
