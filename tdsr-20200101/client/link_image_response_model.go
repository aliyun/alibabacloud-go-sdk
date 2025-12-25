// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LinkImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LinkImageResponse
	GetStatusCode() *int32
	SetBody(v *LinkImageResponseBody) *LinkImageResponse
	GetBody() *LinkImageResponseBody
}

type LinkImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LinkImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LinkImageResponse) String() string {
	return dara.Prettify(s)
}

func (s LinkImageResponse) GoString() string {
	return s.String()
}

func (s *LinkImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LinkImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LinkImageResponse) GetBody() *LinkImageResponseBody {
	return s.Body
}

func (s *LinkImageResponse) SetHeaders(v map[string]*string) *LinkImageResponse {
	s.Headers = v
	return s
}

func (s *LinkImageResponse) SetStatusCode(v int32) *LinkImageResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkImageResponse) SetBody(v *LinkImageResponseBody) *LinkImageResponse {
	s.Body = v
	return s
}

func (s *LinkImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
