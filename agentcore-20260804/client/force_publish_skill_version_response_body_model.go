// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForcePublishSkillVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ForcePublishSkillVersionResponseBody
	GetRequestId() *string
}

type ForcePublishSkillVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ForcePublishSkillVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForcePublishSkillVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ForcePublishSkillVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForcePublishSkillVersionResponseBody) SetRequestId(v string) *ForcePublishSkillVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForcePublishSkillVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
