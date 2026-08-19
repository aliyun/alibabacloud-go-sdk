// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBotSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotInstanceLevel(v string) *UpdateBotSpecRequest
	GetBotInstanceLevel() *string
	SetInstanceId(v string) *UpdateBotSpecRequest
	GetInstanceId() *string
}

type UpdateBotSpecRequest struct {
	// The bot instance specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// enterprise_bot
	BotInstanceLevel *string `json:"BotInstanceLevel,omitempty" xml:"BotInstanceLevel,omitempty"`
	// The bot instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-bot-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateBotSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBotSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateBotSpecRequest) GetBotInstanceLevel() *string {
	return s.BotInstanceLevel
}

func (s *UpdateBotSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateBotSpecRequest) SetBotInstanceLevel(v string) *UpdateBotSpecRequest {
	s.BotInstanceLevel = &v
	return s
}

func (s *UpdateBotSpecRequest) SetInstanceId(v string) *UpdateBotSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateBotSpecRequest) Validate() error {
	return dara.Validate(s)
}
