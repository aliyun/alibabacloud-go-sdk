// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaasServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSaasServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSaasServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSaasServiceResponseBody) *CreateSaasServiceResponse
	GetBody() *CreateSaasServiceResponseBody
}

type CreateSaasServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSaasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSaasServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSaasServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateSaasServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSaasServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSaasServiceResponse) GetBody() *CreateSaasServiceResponseBody {
	return s.Body
}

func (s *CreateSaasServiceResponse) SetHeaders(v map[string]*string) *CreateSaasServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateSaasServiceResponse) SetStatusCode(v int32) *CreateSaasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSaasServiceResponse) SetBody(v *CreateSaasServiceResponseBody) *CreateSaasServiceResponse {
	s.Body = v
	return s
}

func (s *CreateSaasServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
