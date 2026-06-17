// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSaasServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSaasServiceResponse
	GetStatusCode() *int32
	SetBody(v *ListSaasServiceResponseBody) *ListSaasServiceResponse
	GetBody() *ListSaasServiceResponseBody
}

type ListSaasServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSaasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSaasServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSaasServiceResponse) GoString() string {
	return s.String()
}

func (s *ListSaasServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSaasServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSaasServiceResponse) GetBody() *ListSaasServiceResponseBody {
	return s.Body
}

func (s *ListSaasServiceResponse) SetHeaders(v map[string]*string) *ListSaasServiceResponse {
	s.Headers = v
	return s
}

func (s *ListSaasServiceResponse) SetStatusCode(v int32) *ListSaasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSaasServiceResponse) SetBody(v *ListSaasServiceResponseBody) *ListSaasServiceResponse {
	s.Body = v
	return s
}

func (s *ListSaasServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
