// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSkillGroupConfigResponseBody
	GetCode() *string
	SetData(v int64) *CreateSkillGroupConfigResponseBody
	GetData() *int64
	SetMessage(v string) *CreateSkillGroupConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSkillGroupConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSkillGroupConfigResponseBody
	GetSuccess() *bool
}

type CreateSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 223
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillGroupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSkillGroupConfigResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateSkillGroupConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSkillGroupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillGroupConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSkillGroupConfigResponseBody) SetCode(v string) *CreateSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetData(v int64) *CreateSkillGroupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetMessage(v string) *CreateSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetRequestId(v string) *CreateSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetSuccess(v bool) *CreateSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
