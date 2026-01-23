// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardResponseBody) *DeleteStandardResponse
	GetBody() *DeleteStandardResponseBody
}

type DeleteStandardResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardResponse) GetBody() *DeleteStandardResponseBody {
	return s.Body
}

func (s *DeleteStandardResponse) SetHeaders(v map[string]*string) *DeleteStandardResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardResponse) SetStatusCode(v int32) *DeleteStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardResponse) SetBody(v *DeleteStandardResponseBody) *DeleteStandardResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
