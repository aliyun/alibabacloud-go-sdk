// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSkillVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishSkillVersionResponseBody
	GetRequestId() *string
}

type PublishSkillVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PublishSkillVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishSkillVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishSkillVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishSkillVersionResponseBody) SetRequestId(v string) *PublishSkillVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishSkillVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
