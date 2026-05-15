// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSkillGroupResponseBody
	GetCode() *string
	SetData(v bool) *DeleteSkillGroupResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSkillGroupResponseBody
	GetSuccess() *bool
}

type DeleteSkillGroupResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxxx
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

func (s DeleteSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSkillGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSkillGroupResponseBody) SetCode(v string) *DeleteSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetData(v bool) *DeleteSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetMessage(v string) *DeleteSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetRequestId(v string) *DeleteSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetSuccess(v bool) *DeleteSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
