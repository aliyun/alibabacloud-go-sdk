// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaasServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSaasServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSaasServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSaasServiceResponseBody) *DeleteSaasServiceResponse
	GetBody() *DeleteSaasServiceResponseBody
}

type DeleteSaasServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSaasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSaasServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaasServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSaasServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSaasServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSaasServiceResponse) GetBody() *DeleteSaasServiceResponseBody {
	return s.Body
}

func (s *DeleteSaasServiceResponse) SetHeaders(v map[string]*string) *DeleteSaasServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSaasServiceResponse) SetStatusCode(v int32) *DeleteSaasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSaasServiceResponse) SetBody(v *DeleteSaasServiceResponseBody) *DeleteSaasServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteSaasServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
