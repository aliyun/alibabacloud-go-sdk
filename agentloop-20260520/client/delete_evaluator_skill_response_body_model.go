// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluatorSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEvaluatorSkillResponseBody
	GetRequestId() *string
}

type DeleteEvaluatorSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEvaluatorSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluatorSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEvaluatorSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEvaluatorSkillResponseBody) SetRequestId(v string) *DeleteEvaluatorSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEvaluatorSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
