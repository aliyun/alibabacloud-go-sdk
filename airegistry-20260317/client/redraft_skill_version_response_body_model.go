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
	// example:
	//
	// 3920CB4B-90A2-5C97-80DA-578517F2C066
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
