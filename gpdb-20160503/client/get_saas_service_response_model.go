// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSaasServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSaasServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetSaasServiceResponseBody) *GetSaasServiceResponse
	GetBody() *GetSaasServiceResponseBody
}

type GetSaasServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSaasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSaasServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSaasServiceResponse) GoString() string {
	return s.String()
}

func (s *GetSaasServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSaasServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSaasServiceResponse) GetBody() *GetSaasServiceResponseBody {
	return s.Body
}

func (s *GetSaasServiceResponse) SetHeaders(v map[string]*string) *GetSaasServiceResponse {
	s.Headers = v
	return s
}

func (s *GetSaasServiceResponse) SetStatusCode(v int32) *GetSaasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSaasServiceResponse) SetBody(v *GetSaasServiceResponseBody) *GetSaasServiceResponse {
	s.Body = v
	return s
}

func (s *GetSaasServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
