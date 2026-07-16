// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVersionResponseBody) *DeleteVersionResponse
	GetBody() *DeleteVersionResponseBody
}

type DeleteVersionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVersionResponse) GetBody() *DeleteVersionResponseBody {
	return s.Body
}

func (s *DeleteVersionResponse) SetHeaders(v map[string]*string) *DeleteVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteVersionResponse) SetStatusCode(v int32) *DeleteVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVersionResponse) SetBody(v *DeleteVersionResponseBody) *DeleteVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
