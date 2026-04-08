// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBusinessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBusinessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBusinessResponse
	GetStatusCode() *int32
	SetBody(v *CreateBusinessResponseBody) *CreateBusinessResponse
	GetBody() *CreateBusinessResponseBody
}

type CreateBusinessResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBusinessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBusinessResponse) GoString() string {
	return s.String()
}

func (s *CreateBusinessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBusinessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBusinessResponse) GetBody() *CreateBusinessResponseBody {
	return s.Body
}

func (s *CreateBusinessResponse) SetHeaders(v map[string]*string) *CreateBusinessResponse {
	s.Headers = v
	return s
}

func (s *CreateBusinessResponse) SetStatusCode(v int32) *CreateBusinessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBusinessResponse) SetBody(v *CreateBusinessResponseBody) *CreateBusinessResponse {
	s.Body = v
	return s
}

func (s *CreateBusinessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
