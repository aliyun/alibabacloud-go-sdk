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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
