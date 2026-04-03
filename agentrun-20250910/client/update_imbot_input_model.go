// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIMBotInput interface {
	dara.Model
	String() string
	GoString() string
	SetBotBizId(v string) *UpdateIMBotInput
	GetBotBizId() *string
	SetBotBizSecret(v string) *UpdateIMBotInput
	GetBotBizSecret() *string
	SetBotBizType(v string) *UpdateIMBotInput
	GetBotBizType() *string
	SetBotMode(v string) *UpdateIMBotInput
	GetBotMode() *string
	SetBotName(v string) *UpdateIMBotInput
	GetBotName() *string
	SetDescription(v string) *UpdateIMBotInput
	GetDescription() *string
	SetMetadata(v *IMBotMetadata) *UpdateIMBotInput
	GetMetadata() *IMBotMetadata
	SetMinInstances(v int32) *UpdateIMBotInput
	GetMinInstances() *int32
	SetStatus(v string) *UpdateIMBotInput
	GetStatus() *string
}

type UpdateIMBotInput struct {
	BotBizId *string `json:"botBizId,omitempty" xml:"botBizId,omitempty"`
	// 若提供则不可为空字符串
	BotBizSecret *string `json:"botBizSecret,omitempty" xml:"botBizSecret,omitempty"`
	BotBizType   *string `json:"botBizType,omitempty" xml:"botBizType,omitempty"`
	// 不可与租户已固定的 deployment 模式冲突
	BotMode *string `json:"botMode,omitempty" xml:"botMode,omitempty"`
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty"`
	// Bot 描述信息
	//
	// if can be null:
	// true
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// if can be null:
	// true
	Metadata *IMBotMetadata `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// ≥ 1，更新账号级 FC 最小实例
	//
	// if can be null:
	// true
	MinInstances *int32  `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateIMBotInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateIMBotInput) GoString() string {
	return s.String()
}

func (s *UpdateIMBotInput) GetBotBizId() *string {
	return s.BotBizId
}

func (s *UpdateIMBotInput) GetBotBizSecret() *string {
	return s.BotBizSecret
}

func (s *UpdateIMBotInput) GetBotBizType() *string {
	return s.BotBizType
}

func (s *UpdateIMBotInput) GetBotMode() *string {
	return s.BotMode
}

func (s *UpdateIMBotInput) GetBotName() *string {
	return s.BotName
}

func (s *UpdateIMBotInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateIMBotInput) GetMetadata() *IMBotMetadata {
	return s.Metadata
}

func (s *UpdateIMBotInput) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *UpdateIMBotInput) GetStatus() *string {
	return s.Status
}

func (s *UpdateIMBotInput) SetBotBizId(v string) *UpdateIMBotInput {
	s.BotBizId = &v
	return s
}

func (s *UpdateIMBotInput) SetBotBizSecret(v string) *UpdateIMBotInput {
	s.BotBizSecret = &v
	return s
}

func (s *UpdateIMBotInput) SetBotBizType(v string) *UpdateIMBotInput {
	s.BotBizType = &v
	return s
}

func (s *UpdateIMBotInput) SetBotMode(v string) *UpdateIMBotInput {
	s.BotMode = &v
	return s
}

func (s *UpdateIMBotInput) SetBotName(v string) *UpdateIMBotInput {
	s.BotName = &v
	return s
}

func (s *UpdateIMBotInput) SetDescription(v string) *UpdateIMBotInput {
	s.Description = &v
	return s
}

func (s *UpdateIMBotInput) SetMetadata(v *IMBotMetadata) *UpdateIMBotInput {
	s.Metadata = v
	return s
}

func (s *UpdateIMBotInput) SetMinInstances(v int32) *UpdateIMBotInput {
	s.MinInstances = &v
	return s
}

func (s *UpdateIMBotInput) SetStatus(v string) *UpdateIMBotInput {
	s.Status = &v
	return s
}

func (s *UpdateIMBotInput) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
