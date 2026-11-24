// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPinned(v bool) *UpdateSessionRequest
	GetPinned() *bool
	SetTitle(v string) *UpdateSessionRequest
	GetTitle() *string
	SetUiEnabled(v bool) *UpdateSessionRequest
	GetUiEnabled() *bool
}

type UpdateSessionRequest struct {
	// Specifies whether to pin the session. If this parameter is not specified, the pinned status remains unchanged.
	//
	// example:
	//
	// true
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// The updated session title. You can leave this parameter empty if no update is required.
	//
	// example:
	//
	// Updated inspection session
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Specifies whether to enable session UI enhancement. This feature is supported only for common sessions.
	UiEnabled *bool `json:"uiEnabled,omitempty" xml:"uiEnabled,omitempty"`
}

func (s UpdateSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSessionRequest) GetPinned() *bool {
	return s.Pinned
}

func (s *UpdateSessionRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateSessionRequest) GetUiEnabled() *bool {
	return s.UiEnabled
}

func (s *UpdateSessionRequest) SetPinned(v bool) *UpdateSessionRequest {
	s.Pinned = &v
	return s
}

func (s *UpdateSessionRequest) SetTitle(v string) *UpdateSessionRequest {
	s.Title = &v
	return s
}

func (s *UpdateSessionRequest) SetUiEnabled(v bool) *UpdateSessionRequest {
	s.UiEnabled = &v
	return s
}

func (s *UpdateSessionRequest) Validate() error {
	return dara.Validate(s)
}
