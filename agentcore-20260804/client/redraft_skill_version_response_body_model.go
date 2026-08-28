// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedraftSkillVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RedraftSkillVersionResponseBody
	GetRequestId() *string
}

type RedraftSkillVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RedraftSkillVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedraftSkillVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RedraftSkillVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedraftSkillVersionResponseBody) SetRequestId(v string) *RedraftSkillVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedraftSkillVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
