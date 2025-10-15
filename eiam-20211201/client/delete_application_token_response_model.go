// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationTokenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationTokenResponseBody) *DeleteApplicationTokenResponse
	GetBody() *DeleteApplicationTokenResponseBody
}

type DeleteApplicationTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationTokenResponse) GetBody() *DeleteApplicationTokenResponseBody {
	return s.Body
}

func (s *DeleteApplicationTokenResponse) SetHeaders(v map[string]*string) *DeleteApplicationTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationTokenResponse) SetStatusCode(v int32) *DeleteApplicationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationTokenResponse) SetBody(v *DeleteApplicationTokenResponseBody) *DeleteApplicationTokenResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
