// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenOnsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenOnsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenOnsServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenOnsServiceResponseBody) *OpenOnsServiceResponse
	GetBody() *OpenOnsServiceResponseBody
}

type OpenOnsServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenOnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenOnsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenOnsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenOnsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenOnsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenOnsServiceResponse) GetBody() *OpenOnsServiceResponseBody {
	return s.Body
}

func (s *OpenOnsServiceResponse) SetHeaders(v map[string]*string) *OpenOnsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenOnsServiceResponse) SetStatusCode(v int32) *OpenOnsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenOnsServiceResponse) SetBody(v *OpenOnsServiceResponseBody) *OpenOnsServiceResponse {
	s.Body = v
	return s
}

func (s *OpenOnsServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
