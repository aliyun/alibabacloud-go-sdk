// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePortResponse
	GetStatusCode() *int32
	SetBody(v *DeletePortResponseBody) *DeletePortResponse
	GetBody() *DeletePortResponseBody
}

type DeletePortResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePortResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePortResponse) GoString() string {
	return s.String()
}

func (s *DeletePortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePortResponse) GetBody() *DeletePortResponseBody {
	return s.Body
}

func (s *DeletePortResponse) SetHeaders(v map[string]*string) *DeletePortResponse {
	s.Headers = v
	return s
}

func (s *DeletePortResponse) SetStatusCode(v int32) *DeletePortResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePortResponse) SetBody(v *DeletePortResponseBody) *DeletePortResponse {
	s.Body = v
	return s
}

func (s *DeletePortResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
