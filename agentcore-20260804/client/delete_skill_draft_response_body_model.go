// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSkillDraftResponseBody
	GetRequestId() *string
}

type DeleteSkillDraftResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSkillDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillDraftResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillDraftResponseBody) SetRequestId(v string) *DeleteSkillDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillDraftResponseBody) Validate() error {
	return dara.Validate(s)
}
