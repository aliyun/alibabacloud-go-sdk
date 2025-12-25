// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteThreadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteThreadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteThreadResponse
	GetStatusCode() *int32
	SetBody(v *DeleteThreadResponseBody) *DeleteThreadResponse
	GetBody() *DeleteThreadResponseBody
}

type DeleteThreadResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteThreadResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteThreadResponse) GoString() string {
	return s.String()
}

func (s *DeleteThreadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteThreadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteThreadResponse) GetBody() *DeleteThreadResponseBody {
	return s.Body
}

func (s *DeleteThreadResponse) SetHeaders(v map[string]*string) *DeleteThreadResponse {
	s.Headers = v
	return s
}

func (s *DeleteThreadResponse) SetStatusCode(v int32) *DeleteThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteThreadResponse) SetBody(v *DeleteThreadResponseBody) *DeleteThreadResponse {
	s.Body = v
	return s
}

func (s *DeleteThreadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
