// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillBizTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSkillBizTagsResponseBody
	GetRequestId() *string
}

type UpdateSkillBizTagsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSkillBizTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillBizTagsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillBizTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillBizTagsResponseBody) SetRequestId(v string) *UpdateSkillBizTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillBizTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
