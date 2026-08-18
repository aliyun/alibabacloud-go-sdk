// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextDBResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextDBResponseBody) *CreateContextDBResponse
	GetBody() *CreateContextDBResponseBody
}

type CreateContextDBResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextDBResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDBResponse) GoString() string {
	return s.String()
}

func (s *CreateContextDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextDBResponse) GetBody() *CreateContextDBResponseBody {
	return s.Body
}

func (s *CreateContextDBResponse) SetHeaders(v map[string]*string) *CreateContextDBResponse {
	s.Headers = v
	return s
}

func (s *CreateContextDBResponse) SetStatusCode(v int32) *CreateContextDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextDBResponse) SetBody(v *CreateContextDBResponseBody) *CreateContextDBResponse {
	s.Body = v
	return s
}

func (s *CreateContextDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
