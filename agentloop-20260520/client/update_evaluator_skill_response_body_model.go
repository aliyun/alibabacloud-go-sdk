// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluatorSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEvaluatorSkillResponseBody
	GetRequestId() *string
}

type UpdateEvaluatorSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateEvaluatorSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluatorSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEvaluatorSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEvaluatorSkillResponseBody) SetRequestId(v string) *UpdateEvaluatorSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEvaluatorSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
