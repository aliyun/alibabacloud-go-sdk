// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSkillResponseBody
	GetRequestId() *string
}

type DeleteSkillResponseBody struct {
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillResponseBody) SetRequestId(v string) *DeleteSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
