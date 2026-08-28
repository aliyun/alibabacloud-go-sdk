// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSkillScopeResponseBody
	GetRequestId() *string
}

type UpdateSkillScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSkillScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillScopeResponseBody) SetRequestId(v string) *UpdateSkillScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
