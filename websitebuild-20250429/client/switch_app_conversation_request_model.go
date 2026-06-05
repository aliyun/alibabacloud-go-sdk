// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchAppConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SwitchAppConversationRequest
	GetBizId() *string
	SetBotId(v string) *SwitchAppConversationRequest
	GetBotId() *string
	SetTaskType(v string) *SwitchAppConversationRequest
	GetTaskType() *string
}

type SwitchAppConversationRequest struct {
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// IMAGE_LOGO
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s SwitchAppConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchAppConversationRequest) GoString() string {
	return s.String()
}

func (s *SwitchAppConversationRequest) GetBizId() *string {
	return s.BizId
}

func (s *SwitchAppConversationRequest) GetBotId() *string {
	return s.BotId
}

func (s *SwitchAppConversationRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *SwitchAppConversationRequest) SetBizId(v string) *SwitchAppConversationRequest {
	s.BizId = &v
	return s
}

func (s *SwitchAppConversationRequest) SetBotId(v string) *SwitchAppConversationRequest {
	s.BotId = &v
	return s
}

func (s *SwitchAppConversationRequest) SetTaskType(v string) *SwitchAppConversationRequest {
	s.TaskType = &v
	return s
}

func (s *SwitchAppConversationRequest) Validate() error {
	return dara.Validate(s)
}
