// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceResponseBody) *CreateServiceResponse
	GetBody() *CreateServiceResponseBody
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceResponse) GetBody() *CreateServiceResponseBody {
	return s.Body
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

func (s *CreateServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
