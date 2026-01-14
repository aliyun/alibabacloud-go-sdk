// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAcceleratorServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenAcceleratorServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenAcceleratorServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenAcceleratorServiceResponseBody) *OpenAcceleratorServiceResponse
	GetBody() *OpenAcceleratorServiceResponseBody
}

type OpenAcceleratorServiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenAcceleratorServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenAcceleratorServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenAcceleratorServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenAcceleratorServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenAcceleratorServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenAcceleratorServiceResponse) GetBody() *OpenAcceleratorServiceResponseBody {
	return s.Body
}

func (s *OpenAcceleratorServiceResponse) SetHeaders(v map[string]*string) *OpenAcceleratorServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenAcceleratorServiceResponse) SetStatusCode(v int32) *OpenAcceleratorServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAcceleratorServiceResponse) SetBody(v *OpenAcceleratorServiceResponseBody) *OpenAcceleratorServiceResponse {
	s.Body = v
	return s
}

func (s *OpenAcceleratorServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
