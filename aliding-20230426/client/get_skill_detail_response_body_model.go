// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetSkillDetailResponseBody
	GetData() interface{}
	SetRequestId(v string) *GetSkillDetailResponseBody
	GetRequestId() *string
}

type GetSkillDetailResponseBody struct {
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

func (s GetSkillDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetSkillDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillDetailResponseBody) SetData(v interface{}) *GetSkillDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillDetailResponseBody) SetRequestId(v string) *GetSkillDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
