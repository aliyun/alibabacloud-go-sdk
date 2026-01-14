// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageResponse
	GetStatusCode() *int32
	SetBody(v *PageResponseBody) *PageResponse
	GetBody() *PageResponseBody
}

type PageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageResponse) String() string {
	return dara.Prettify(s)
}

func (s PageResponse) GoString() string {
	return s.String()
}

func (s *PageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageResponse) GetBody() *PageResponseBody {
	return s.Body
}

func (s *PageResponse) SetHeaders(v map[string]*string) *PageResponse {
	s.Headers = v
	return s
}

func (s *PageResponse) SetStatusCode(v int32) *PageResponse {
	s.StatusCode = &v
	return s
}

func (s *PageResponse) SetBody(v *PageResponseBody) *PageResponse {
	s.Body = v
	return s
}

func (s *PageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
