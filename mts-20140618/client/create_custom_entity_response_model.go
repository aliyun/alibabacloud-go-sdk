// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomEntityResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomEntityResponseBody) *CreateCustomEntityResponse
	GetBody() *CreateCustomEntityResponseBody
}

type CreateCustomEntityResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomEntityResponse) GetBody() *CreateCustomEntityResponseBody {
	return s.Body
}

func (s *CreateCustomEntityResponse) SetHeaders(v map[string]*string) *CreateCustomEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomEntityResponse) SetStatusCode(v int32) *CreateCustomEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomEntityResponse) SetBody(v *CreateCustomEntityResponseBody) *CreateCustomEntityResponse {
	s.Body = v
	return s
}

func (s *CreateCustomEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
