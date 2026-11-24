// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTitle(v string) *CreateSessionRequest
	GetTitle() *string
	SetUiEnabled(v bool) *CreateSessionRequest
	GetUiEnabled() *bool
}

type CreateSessionRequest struct {
	// The session title.
	//
	// example:
	//
	// Cloud resource inspection example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Specifies whether to enable session UI enhancement. Only common sessions support this feature.
	UiEnabled *bool `json:"uiEnabled,omitempty" xml:"uiEnabled,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateSessionRequest) GetUiEnabled() *bool {
	return s.UiEnabled
}

func (s *CreateSessionRequest) SetTitle(v string) *CreateSessionRequest {
	s.Title = &v
	return s
}

func (s *CreateSessionRequest) SetUiEnabled(v bool) *CreateSessionRequest {
	s.UiEnabled = &v
	return s
}

func (s *CreateSessionRequest) Validate() error {
	return dara.Validate(s)
}
