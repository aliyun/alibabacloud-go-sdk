// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSkillVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SubmitSkillVersionResponseBody
	GetData() *string
	SetRequestId(v string) *SubmitSkillVersionResponseBody
	GetRequestId() *string
}

type SubmitSkillVersionResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSkillVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSkillVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSkillVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitSkillVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSkillVersionResponseBody) SetData(v string) *SubmitSkillVersionResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitSkillVersionResponseBody) SetRequestId(v string) *SubmitSkillVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSkillVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
