// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptForArbitrationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *InterruptForArbitrationShrinkRequest
	GetAppId() *string
	SetChatId(v string) *InterruptForArbitrationShrinkRequest
	GetChatId() *string
	SetHubRequestId(v string) *InterruptForArbitrationShrinkRequest
	GetHubRequestId() *string
	SetInterruptShrink(v string) *InterruptForArbitrationShrinkRequest
	GetInterruptShrink() *string
	SetSessionId(v string) *InterruptForArbitrationShrinkRequest
	GetSessionId() *string
}

type InterruptForArbitrationShrinkRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// This parameter is required.
	HubRequestId *string `json:"HubRequestId,omitempty" xml:"HubRequestId,omitempty"`
	// This parameter is required.
	InterruptShrink *string `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	SessionId       *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s InterruptForArbitrationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InterruptForArbitrationShrinkRequest) GoString() string {
	return s.String()
}

func (s *InterruptForArbitrationShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *InterruptForArbitrationShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *InterruptForArbitrationShrinkRequest) GetHubRequestId() *string {
	return s.HubRequestId
}

func (s *InterruptForArbitrationShrinkRequest) GetInterruptShrink() *string {
	return s.InterruptShrink
}

func (s *InterruptForArbitrationShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *InterruptForArbitrationShrinkRequest) SetAppId(v string) *InterruptForArbitrationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *InterruptForArbitrationShrinkRequest) SetChatId(v string) *InterruptForArbitrationShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *InterruptForArbitrationShrinkRequest) SetHubRequestId(v string) *InterruptForArbitrationShrinkRequest {
	s.HubRequestId = &v
	return s
}

func (s *InterruptForArbitrationShrinkRequest) SetInterruptShrink(v string) *InterruptForArbitrationShrinkRequest {
	s.InterruptShrink = &v
	return s
}

func (s *InterruptForArbitrationShrinkRequest) SetSessionId(v string) *InterruptForArbitrationShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *InterruptForArbitrationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
