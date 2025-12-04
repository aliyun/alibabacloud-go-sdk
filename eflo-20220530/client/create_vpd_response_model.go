// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpdResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpdResponseBody) *CreateVpdResponse
	GetBody() *CreateVpdResponseBody
}

type CreateVpdResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpdResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdResponse) GoString() string {
	return s.String()
}

func (s *CreateVpdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpdResponse) GetBody() *CreateVpdResponseBody {
	return s.Body
}

func (s *CreateVpdResponse) SetHeaders(v map[string]*string) *CreateVpdResponse {
	s.Headers = v
	return s
}

func (s *CreateVpdResponse) SetStatusCode(v int32) *CreateVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpdResponse) SetBody(v *CreateVpdResponseBody) *CreateVpdResponse {
	s.Body = v
	return s
}

func (s *CreateVpdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
