// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DeleteDesktopRequest
	GetApiKey() *string
	SetDesktopId(v string) *DeleteDesktopRequest
	GetDesktopId() *string
	SetSessionId(v string) *DeleteDesktopRequest
	GetSessionId() *string
}

type DeleteDesktopRequest struct {
	// This parameter is required.
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteDesktopRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DeleteDesktopRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DeleteDesktopRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteDesktopRequest) SetApiKey(v string) *DeleteDesktopRequest {
	s.ApiKey = &v
	return s
}

func (s *DeleteDesktopRequest) SetDesktopId(v string) *DeleteDesktopRequest {
	s.DesktopId = &v
	return s
}

func (s *DeleteDesktopRequest) SetSessionId(v string) *DeleteDesktopRequest {
	s.SessionId = &v
	return s
}

func (s *DeleteDesktopRequest) Validate() error {
	return dara.Validate(s)
}
