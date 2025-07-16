// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListSkillResponseBody
	GetData() interface{}
	SetRequestId(v string) *ListSkillResponseBody
	GetRequestId() *string
}

type ListSkillResponseBody struct {
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 2715B4D3-A3FB-5FC7-AFA0-4471687B1EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillResponseBody) SetData(v interface{}) *ListSkillResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillResponseBody) SetRequestId(v string) *ListSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
