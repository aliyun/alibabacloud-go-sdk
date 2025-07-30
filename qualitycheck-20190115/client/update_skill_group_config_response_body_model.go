// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillGroupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSkillGroupConfigResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSkillGroupConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSkillGroupConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSkillGroupConfigResponseBody
	GetSuccess() *bool
}

type UpdateSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 38E7E948-0876-4FEE-B0AA-6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSkillGroupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSkillGroupConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSkillGroupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillGroupConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSkillGroupConfigResponseBody) SetCode(v string) *UpdateSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetMessage(v string) *UpdateSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetRequestId(v string) *UpdateSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetSuccess(v bool) *UpdateSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
