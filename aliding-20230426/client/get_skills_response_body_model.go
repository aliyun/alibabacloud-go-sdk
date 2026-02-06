// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetSkillsResponseBody
	GetData() interface{}
	SetRequestId(v string) *GetSkillsResponseBody
	GetRequestId() *string
}

type GetSkillsResponseBody struct {
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// RequestId
	//
	// example:
	//
	// A348BA5D-FFD4-57E4-9450-23A14D72F331
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillsResponseBody) SetData(v interface{}) *GetSkillsResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillsResponseBody) SetRequestId(v string) *GetSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillsResponseBody) Validate() error {
	return dara.Validate(s)
}
