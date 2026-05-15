// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSkillGroupResponseBody
	GetCode() *string
	SetData(v int64) *CreateSkillGroupResponseBody
	GetData() *int64
	SetMessage(v string) *CreateSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSkillGroupResponseBody
	GetSuccess() *bool
}

type CreateSkillGroupResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123456
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSkillGroupResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSkillGroupResponseBody) SetCode(v string) *CreateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetData(v int64) *CreateSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetMessage(v string) *CreateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetRequestId(v string) *CreateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetSuccess(v bool) *CreateSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
