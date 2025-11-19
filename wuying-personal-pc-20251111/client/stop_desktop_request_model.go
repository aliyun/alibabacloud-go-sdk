// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *StopDesktopRequest
	GetApiKey() *string
	SetDesktopId(v string) *StopDesktopRequest
	GetDesktopId() *string
	SetSessionId(v string) *StopDesktopRequest
	GetSessionId() *string
}

type StopDesktopRequest struct {
	// This parameter is required.
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StopDesktopRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDesktopRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *StopDesktopRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *StopDesktopRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopDesktopRequest) SetApiKey(v string) *StopDesktopRequest {
	s.ApiKey = &v
	return s
}

func (s *StopDesktopRequest) SetDesktopId(v string) *StopDesktopRequest {
	s.DesktopId = &v
	return s
}

func (s *StopDesktopRequest) SetSessionId(v string) *StopDesktopRequest {
	s.SessionId = &v
	return s
}

func (s *StopDesktopRequest) Validate() error {
	return dara.Validate(s)
}
