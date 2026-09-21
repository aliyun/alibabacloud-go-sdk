// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentSkillResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAgentSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentSkillResponseBody
	GetRequestId() *string
}

type DeleteAgentSkillResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgentSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentSkillResponseBody) SetCode(v string) *DeleteAgentSkillResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentSkillResponseBody) SetMessage(v string) *DeleteAgentSkillResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentSkillResponseBody) SetRequestId(v string) *DeleteAgentSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
