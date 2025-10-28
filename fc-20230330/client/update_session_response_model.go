// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSessionResponse
	GetStatusCode() *int32
	SetBody(v *Session) *UpdateSessionResponse
	GetBody() *Session
}

type UpdateSessionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Session           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSessionResponse) GetBody() *Session {
	return s.Body
}

func (s *UpdateSessionResponse) SetHeaders(v map[string]*string) *UpdateSessionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSessionResponse) SetStatusCode(v int32) *UpdateSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSessionResponse) SetBody(v *Session) *UpdateSessionResponse {
	s.Body = v
	return s
}

func (s *UpdateSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
