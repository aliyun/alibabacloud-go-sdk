// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptForArbitrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *InterruptForArbitrationRequest
	GetAppId() *string
	SetChatId(v string) *InterruptForArbitrationRequest
	GetChatId() *string
	SetHubRequestId(v string) *InterruptForArbitrationRequest
	GetHubRequestId() *string
	SetInterrupt(v *InterruptForArbitrationRequestInterrupt) *InterruptForArbitrationRequest
	GetInterrupt() *InterruptForArbitrationRequestInterrupt
	SetSessionId(v string) *InterruptForArbitrationRequest
	GetSessionId() *string
}

type InterruptForArbitrationRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// This parameter is required.
	HubRequestId *string `json:"HubRequestId,omitempty" xml:"HubRequestId,omitempty"`
	// This parameter is required.
	Interrupt *InterruptForArbitrationRequestInterrupt `json:"Interrupt,omitempty" xml:"Interrupt,omitempty" type:"Struct"`
	SessionId *string                                  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s InterruptForArbitrationRequest) String() string {
	return dara.Prettify(s)
}

func (s InterruptForArbitrationRequest) GoString() string {
	return s.String()
}

func (s *InterruptForArbitrationRequest) GetAppId() *string {
	return s.AppId
}

func (s *InterruptForArbitrationRequest) GetChatId() *string {
	return s.ChatId
}

func (s *InterruptForArbitrationRequest) GetHubRequestId() *string {
	return s.HubRequestId
}

func (s *InterruptForArbitrationRequest) GetInterrupt() *InterruptForArbitrationRequestInterrupt {
	return s.Interrupt
}

func (s *InterruptForArbitrationRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *InterruptForArbitrationRequest) SetAppId(v string) *InterruptForArbitrationRequest {
	s.AppId = &v
	return s
}

func (s *InterruptForArbitrationRequest) SetChatId(v string) *InterruptForArbitrationRequest {
	s.ChatId = &v
	return s
}

func (s *InterruptForArbitrationRequest) SetHubRequestId(v string) *InterruptForArbitrationRequest {
	s.HubRequestId = &v
	return s
}

func (s *InterruptForArbitrationRequest) SetInterrupt(v *InterruptForArbitrationRequestInterrupt) *InterruptForArbitrationRequest {
	s.Interrupt = v
	return s
}

func (s *InterruptForArbitrationRequest) SetSessionId(v string) *InterruptForArbitrationRequest {
	s.SessionId = &v
	return s
}

func (s *InterruptForArbitrationRequest) Validate() error {
	if s.Interrupt != nil {
		if err := s.Interrupt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InterruptForArbitrationRequestInterrupt struct {
	// This parameter is required.
	Submit *bool `json:"Submit,omitempty" xml:"Submit,omitempty"`
}

func (s InterruptForArbitrationRequestInterrupt) String() string {
	return dara.Prettify(s)
}

func (s InterruptForArbitrationRequestInterrupt) GoString() string {
	return s.String()
}

func (s *InterruptForArbitrationRequestInterrupt) GetSubmit() *bool {
	return s.Submit
}

func (s *InterruptForArbitrationRequestInterrupt) SetSubmit(v bool) *InterruptForArbitrationRequestInterrupt {
	s.Submit = &v
	return s
}

func (s *InterruptForArbitrationRequestInterrupt) Validate() error {
	return dara.Validate(s)
}
