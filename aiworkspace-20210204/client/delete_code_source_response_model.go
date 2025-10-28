// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCodeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCodeSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCodeSourceResponseBody) *DeleteCodeSourceResponse
	GetBody() *DeleteCodeSourceResponseBody
}

type DeleteCodeSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCodeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCodeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCodeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCodeSourceResponse) GetBody() *DeleteCodeSourceResponseBody {
	return s.Body
}

func (s *DeleteCodeSourceResponse) SetHeaders(v map[string]*string) *DeleteCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCodeSourceResponse) SetStatusCode(v int32) *DeleteCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCodeSourceResponse) SetBody(v *DeleteCodeSourceResponseBody) *DeleteCodeSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteCodeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
