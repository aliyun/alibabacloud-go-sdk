// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBusinessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBusinessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBusinessResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBusinessResponseBody) *DeleteBusinessResponse
	GetBody() *DeleteBusinessResponseBody
}

type DeleteBusinessResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBusinessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBusinessResponse) GoString() string {
	return s.String()
}

func (s *DeleteBusinessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBusinessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBusinessResponse) GetBody() *DeleteBusinessResponseBody {
	return s.Body
}

func (s *DeleteBusinessResponse) SetHeaders(v map[string]*string) *DeleteBusinessResponse {
	s.Headers = v
	return s
}

func (s *DeleteBusinessResponse) SetStatusCode(v int32) *DeleteBusinessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBusinessResponse) SetBody(v *DeleteBusinessResponseBody) *DeleteBusinessResponse {
	s.Body = v
	return s
}

func (s *DeleteBusinessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
