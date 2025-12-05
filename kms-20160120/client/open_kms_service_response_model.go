// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenKmsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenKmsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenKmsServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenKmsServiceResponseBody) *OpenKmsServiceResponse
	GetBody() *OpenKmsServiceResponseBody
}

type OpenKmsServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenKmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenKmsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenKmsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenKmsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenKmsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenKmsServiceResponse) GetBody() *OpenKmsServiceResponseBody {
	return s.Body
}

func (s *OpenKmsServiceResponse) SetHeaders(v map[string]*string) *OpenKmsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenKmsServiceResponse) SetStatusCode(v int32) *OpenKmsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenKmsServiceResponse) SetBody(v *OpenKmsServiceResponseBody) *OpenKmsServiceResponse {
	s.Body = v
	return s
}

func (s *OpenKmsServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
