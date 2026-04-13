// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSessionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSessionResponseBody) *DeleteSessionResponse
	GetBody() *DeleteSessionResponseBody
}

type DeleteSessionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSessionResponse) GetBody() *DeleteSessionResponseBody {
	return s.Body
}

func (s *DeleteSessionResponse) SetHeaders(v map[string]*string) *DeleteSessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSessionResponse) SetStatusCode(v int32) *DeleteSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSessionResponse) SetBody(v *DeleteSessionResponseBody) *DeleteSessionResponse {
	s.Body = v
	return s
}

func (s *DeleteSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
