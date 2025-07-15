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
	SetHttpStatusCode(v int32) *DeleteSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSkillGroupResponseBody
	GetRequestId() *string
}

type DeleteSkillGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DeleteSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillGroupResponseBody) SetCode(v string) *DeleteSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetHttpStatusCode(v int32) *DeleteSkillGroupResponseBody {
	s.HttpStatusCode = &v
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

func (s *DeleteSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
