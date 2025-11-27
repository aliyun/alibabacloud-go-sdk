// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmsSignResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmsSignResponseBody) *DeleteSmsSignResponse
	GetBody() *DeleteSmsSignResponseBody
}

type DeleteSmsSignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsSignResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmsSignResponse) GetBody() *DeleteSmsSignResponseBody {
	return s.Body
}

func (s *DeleteSmsSignResponse) SetHeaders(v map[string]*string) *DeleteSmsSignResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmsSignResponse) SetStatusCode(v int32) *DeleteSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmsSignResponse) SetBody(v *DeleteSmsSignResponseBody) *DeleteSmsSignResponse {
	s.Body = v
	return s
}

func (s *DeleteSmsSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
