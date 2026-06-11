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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
