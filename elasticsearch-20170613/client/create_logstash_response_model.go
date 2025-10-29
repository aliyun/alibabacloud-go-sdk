// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogstashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogstashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogstashResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogstashResponseBody) *CreateLogstashResponse
	GetBody() *CreateLogstashResponseBody
}

type CreateLogstashResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogstashResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogstashResponse) GoString() string {
	return s.String()
}

func (s *CreateLogstashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogstashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogstashResponse) GetBody() *CreateLogstashResponseBody {
	return s.Body
}

func (s *CreateLogstashResponse) SetHeaders(v map[string]*string) *CreateLogstashResponse {
	s.Headers = v
	return s
}

func (s *CreateLogstashResponse) SetStatusCode(v int32) *CreateLogstashResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogstashResponse) SetBody(v *CreateLogstashResponseBody) *CreateLogstashResponse {
	s.Body = v
	return s
}

func (s *CreateLogstashResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
