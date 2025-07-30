// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillGroupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSkillGroupConfigResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSkillGroupConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSkillGroupConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSkillGroupConfigResponseBody
	GetSuccess() *bool
}

type DeleteSkillGroupConfigResponseBody struct {
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
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSkillGroupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSkillGroupConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSkillGroupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillGroupConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSkillGroupConfigResponseBody) SetCode(v string) *DeleteSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetMessage(v string) *DeleteSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetRequestId(v string) *DeleteSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetSuccess(v bool) *DeleteSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
