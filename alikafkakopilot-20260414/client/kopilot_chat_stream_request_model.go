// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotChatStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *KopilotChatStreamRequest
	GetMessage() *string
	SetRegionId(v string) *KopilotChatStreamRequest
	GetRegionId() *string
	SetSessionId(v string) *KopilotChatStreamRequest
	GetSessionId() *string
}

type KopilotChatStreamRequest struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s KopilotChatStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotChatStreamRequest) GoString() string {
	return s.String()
}

func (s *KopilotChatStreamRequest) GetMessage() *string {
	return s.Message
}

func (s *KopilotChatStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotChatStreamRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *KopilotChatStreamRequest) SetMessage(v string) *KopilotChatStreamRequest {
	s.Message = &v
	return s
}

func (s *KopilotChatStreamRequest) SetRegionId(v string) *KopilotChatStreamRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotChatStreamRequest) SetSessionId(v string) *KopilotChatStreamRequest {
	s.SessionId = &v
	return s
}

func (s *KopilotChatStreamRequest) Validate() error {
	return dara.Validate(s)
}
