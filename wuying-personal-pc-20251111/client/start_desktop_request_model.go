// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDesktopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *StartDesktopRequest
	GetApiKey() *string
	SetDesktopId(v string) *StartDesktopRequest
	GetDesktopId() *string
	SetSessionId(v string) *StartDesktopRequest
	GetSessionId() *string
}

type StartDesktopRequest struct {
	// This parameter is required.
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StartDesktopRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDesktopRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *StartDesktopRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *StartDesktopRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartDesktopRequest) SetApiKey(v string) *StartDesktopRequest {
	s.ApiKey = &v
	return s
}

func (s *StartDesktopRequest) SetDesktopId(v string) *StartDesktopRequest {
	s.DesktopId = &v
	return s
}

func (s *StartDesktopRequest) SetSessionId(v string) *StartDesktopRequest {
	s.SessionId = &v
	return s
}

func (s *StartDesktopRequest) Validate() error {
	return dara.Validate(s)
}
