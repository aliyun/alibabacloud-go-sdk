// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentSkillResponseBody
	GetCode() *string
	SetMessage(v string) *CreateAgentSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentSkillResponseBody
	GetRequestId() *string
	SetSkillInfo(v []*CreateAgentSkillResponseBodySkillInfo) *CreateAgentSkillResponseBody
	GetSkillInfo() []*CreateAgentSkillResponseBodySkillInfo
}

type CreateAgentSkillResponseBody struct {
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
	// The skill information.
	SkillInfo []*CreateAgentSkillResponseBodySkillInfo `json:"SkillInfo,omitempty" xml:"SkillInfo,omitempty" type:"Repeated"`
}

func (s CreateAgentSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentSkillResponseBody) GetSkillInfo() []*CreateAgentSkillResponseBodySkillInfo {
	return s.SkillInfo
}

func (s *CreateAgentSkillResponseBody) SetCode(v string) *CreateAgentSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentSkillResponseBody) SetMessage(v string) *CreateAgentSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentSkillResponseBody) SetRequestId(v string) *CreateAgentSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentSkillResponseBody) SetSkillInfo(v []*CreateAgentSkillResponseBodySkillInfo) *CreateAgentSkillResponseBody {
	s.SkillInfo = v
	return s
}

func (s *CreateAgentSkillResponseBody) Validate() error {
	if s.SkillInfo != nil {
		for _, item := range s.SkillInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAgentSkillResponseBodySkillInfo struct {
	// The skill description.
	//
	// example:
	//
	// Current weather and forecasts with wttr.in via curl for locations, rain, temperature, travel planning.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The skill summary.
	//
	// example:
	//
	// Current weather and forecasts.
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The unique ID of the skill.
	//
	// example:
	//
	// sk-051j4pbwxzgol****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The skill name.
	//
	// example:
	//
	// dev-spec
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The skill type.
	//
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAgentSkillResponseBodySkillInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSkillResponseBodySkillInfo) GoString() string {
	return s.String()
}

func (s *CreateAgentSkillResponseBodySkillInfo) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentSkillResponseBodySkillInfo) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateAgentSkillResponseBodySkillInfo) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateAgentSkillResponseBodySkillInfo) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateAgentSkillResponseBodySkillInfo) GetStatus() *string {
	return s.Status
}

func (s *CreateAgentSkillResponseBodySkillInfo) GetType() *string {
	return s.Type
}

func (s *CreateAgentSkillResponseBodySkillInfo) SetDescription(v string) *CreateAgentSkillResponseBodySkillInfo {
	s.Description = &v
	return s
}

func (s *CreateAgentSkillResponseBodySkillInfo) SetInstruction(v string) *CreateAgentSkillResponseBodySkillInfo {
	s.Instruction = &v
	return s
}

func (s *CreateAgentSkillResponseBodySkillInfo) SetSkillId(v string) *CreateAgentSkillResponseBodySkillInfo {
	s.SkillId = &v
	return s
}

func (s *CreateAgentSkillResponseBodySkillInfo) SetSkillName(v string) *CreateAgentSkillResponseBodySkillInfo {
	s.SkillName = &v
	return s
}

func (s *CreateAgentSkillResponseBodySkillInfo) SetStatus(v string) *CreateAgentSkillResponseBodySkillInfo {
	s.Status = &v
	return s
}

func (s *CreateAgentSkillResponseBodySkillInfo) SetType(v string) *CreateAgentSkillResponseBodySkillInfo {
	s.Type = &v
	return s
}

func (s *CreateAgentSkillResponseBodySkillInfo) Validate() error {
	return dara.Validate(s)
}
