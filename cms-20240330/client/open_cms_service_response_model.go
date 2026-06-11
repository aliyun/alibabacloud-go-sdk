// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCmsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenCmsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenCmsServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenCmsServiceResponseBody) *OpenCmsServiceResponse
	GetBody() *OpenCmsServiceResponseBody
}

type OpenCmsServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenCmsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenCmsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCmsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenCmsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenCmsServiceResponse) GetBody() *OpenCmsServiceResponseBody {
	return s.Body
}

func (s *OpenCmsServiceResponse) SetHeaders(v map[string]*string) *OpenCmsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCmsServiceResponse) SetStatusCode(v int32) *OpenCmsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCmsServiceResponse) SetBody(v *OpenCmsServiceResponseBody) *OpenCmsServiceResponse {
	s.Body = v
	return s
}

func (s *OpenCmsServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
