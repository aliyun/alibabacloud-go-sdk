// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackendResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackendResponseBody) *DeleteBackendResponse
	GetBody() *DeleteBackendResponseBody
}

type DeleteBackendResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackendResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackendResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackendResponse) GetBody() *DeleteBackendResponseBody {
	return s.Body
}

func (s *DeleteBackendResponse) SetHeaders(v map[string]*string) *DeleteBackendResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackendResponse) SetStatusCode(v int32) *DeleteBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackendResponse) SetBody(v *DeleteBackendResponseBody) *DeleteBackendResponse {
	s.Body = v
	return s
}

func (s *DeleteBackendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
