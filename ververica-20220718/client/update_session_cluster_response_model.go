// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSessionClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSessionClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSessionClusterResponseBody) *UpdateSessionClusterResponse
	GetBody() *UpdateSessionClusterResponseBody
}

type UpdateSessionClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSessionClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSessionClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSessionClusterResponse) GetBody() *UpdateSessionClusterResponseBody {
	return s.Body
}

func (s *UpdateSessionClusterResponse) SetHeaders(v map[string]*string) *UpdateSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateSessionClusterResponse) SetStatusCode(v int32) *UpdateSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSessionClusterResponse) SetBody(v *UpdateSessionClusterResponseBody) *UpdateSessionClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateSessionClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
