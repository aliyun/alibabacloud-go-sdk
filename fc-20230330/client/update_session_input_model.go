// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSessionIdReuse(v bool) *UpdateSessionInput
	GetDisableSessionIdReuse() *bool
	SetSessionIdleTimeoutInSeconds(v int64) *UpdateSessionInput
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *UpdateSessionInput
	GetSessionTTLInSeconds() *int64
}

type UpdateSessionInput struct {
	DisableSessionIdReuse *bool `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// example:
	//
	// 21600
	SessionTTLInSeconds *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
}

func (s UpdateSessionInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionInput) GoString() string {
	return s.String()
}

func (s *UpdateSessionInput) GetDisableSessionIdReuse() *bool {
	return s.DisableSessionIdReuse
}

func (s *UpdateSessionInput) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *UpdateSessionInput) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *UpdateSessionInput) SetDisableSessionIdReuse(v bool) *UpdateSessionInput {
	s.DisableSessionIdReuse = &v
	return s
}

func (s *UpdateSessionInput) SetSessionIdleTimeoutInSeconds(v int64) *UpdateSessionInput {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *UpdateSessionInput) SetSessionTTLInSeconds(v int64) *UpdateSessionInput {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *UpdateSessionInput) Validate() error {
	return dara.Validate(s)
}
