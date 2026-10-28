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
}

type UpdateSessionRequest struct {
	// Specifies whether to pin the session. If you do not specify this parameter, the pinned status remains unchanged.
	//
	// example:
	//
	// true
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// The new title of the session. If you do not need to update the title, do not specify this parameter.
	//
	// example:
	//
	// Updated inspection session
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
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

func (s *UpdateSessionRequest) SetPinned(v bool) *UpdateSessionRequest {
	s.Pinned = &v
	return s
}

func (s *UpdateSessionRequest) SetTitle(v string) *UpdateSessionRequest {
	s.Title = &v
	return s
}

func (s *UpdateSessionRequest) Validate() error {
	return dara.Validate(s)
}
