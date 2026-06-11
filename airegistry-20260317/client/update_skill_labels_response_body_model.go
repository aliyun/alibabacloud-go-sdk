// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSkillLabelsResponseBody
	GetRequestId() *string
}

type UpdateSkillLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSkillLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillLabelsResponseBody) SetRequestId(v string) *UpdateSkillLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
