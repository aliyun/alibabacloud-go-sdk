// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIMBotInput interface {
	dara.Model
	String() string
	GoString() string
	SetBotBizId(v string) *CreateIMBotInput
	GetBotBizId() *string
	SetBotBizSecret(v string) *CreateIMBotInput
	GetBotBizSecret() *string
	SetBotBizType(v string) *CreateIMBotInput
	GetBotBizType() *string
	SetBotMode(v string) *CreateIMBotInput
	GetBotMode() *string
	SetBotName(v string) *CreateIMBotInput
	GetBotName() *string
	SetDescription(v string) *CreateIMBotInput
	GetDescription() *string
	SetMetadata(v *IMBotMetadata) *CreateIMBotInput
	GetMetadata() *IMBotMetadata
	SetMinInstances(v int32) *CreateIMBotInput
	GetMinInstances() *int32
}

type CreateIMBotInput struct {
	// This parameter is required.
	BotBizId *string `json:"botBizId,omitempty" xml:"botBizId,omitempty"`
	// This parameter is required.
	BotBizSecret *string `json:"botBizSecret,omitempty" xml:"botBizSecret,omitempty"`
	// This parameter is required.
	BotBizType *string `json:"botBizType,omitempty" xml:"botBizType,omitempty"`
	// standard、custom；不传时可能继承租户已有部署模式
	//
	// if can be null:
	// true
	BotMode *string `json:"botMode,omitempty" xml:"botMode,omitempty"`
	// This parameter is required.
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty"`
	// Bot 描述信息
	//
	// if can be null:
	// true
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// botMode 为 standard 时条件必填；custom 模式不进行标准运行时绑定校验
	//
	// if can be null:
	// true
	Metadata *IMBotMetadata `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// FC 最小实例数，≥ 1；与租户账号级配置相关
	//
	// if can be null:
	// true
	MinInstances *int32 `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
}

func (s CreateIMBotInput) String() string {
	return dara.Prettify(s)
}

func (s CreateIMBotInput) GoString() string {
	return s.String()
}

func (s *CreateIMBotInput) GetBotBizId() *string {
	return s.BotBizId
}

func (s *CreateIMBotInput) GetBotBizSecret() *string {
	return s.BotBizSecret
}

func (s *CreateIMBotInput) GetBotBizType() *string {
	return s.BotBizType
}

func (s *CreateIMBotInput) GetBotMode() *string {
	return s.BotMode
}

func (s *CreateIMBotInput) GetBotName() *string {
	return s.BotName
}

func (s *CreateIMBotInput) GetDescription() *string {
	return s.Description
}

func (s *CreateIMBotInput) GetMetadata() *IMBotMetadata {
	return s.Metadata
}

func (s *CreateIMBotInput) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *CreateIMBotInput) SetBotBizId(v string) *CreateIMBotInput {
	s.BotBizId = &v
	return s
}

func (s *CreateIMBotInput) SetBotBizSecret(v string) *CreateIMBotInput {
	s.BotBizSecret = &v
	return s
}

func (s *CreateIMBotInput) SetBotBizType(v string) *CreateIMBotInput {
	s.BotBizType = &v
	return s
}

func (s *CreateIMBotInput) SetBotMode(v string) *CreateIMBotInput {
	s.BotMode = &v
	return s
}

func (s *CreateIMBotInput) SetBotName(v string) *CreateIMBotInput {
	s.BotName = &v
	return s
}

func (s *CreateIMBotInput) SetDescription(v string) *CreateIMBotInput {
	s.Description = &v
	return s
}

func (s *CreateIMBotInput) SetMetadata(v *IMBotMetadata) *CreateIMBotInput {
	s.Metadata = v
	return s
}

func (s *CreateIMBotInput) SetMinInstances(v int32) *CreateIMBotInput {
	s.MinInstances = &v
	return s
}

func (s *CreateIMBotInput) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
