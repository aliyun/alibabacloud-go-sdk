// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCdnServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenCdnServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenCdnServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenCdnServiceResponseBody) *OpenCdnServiceResponse
	GetBody() *OpenCdnServiceResponseBody
}

type OpenCdnServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenCdnServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenCdnServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenCdnServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenCdnServiceResponse) GetBody() *OpenCdnServiceResponseBody {
	return s.Body
}

func (s *OpenCdnServiceResponse) SetHeaders(v map[string]*string) *OpenCdnServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCdnServiceResponse) SetStatusCode(v int32) *OpenCdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCdnServiceResponse) SetBody(v *OpenCdnServiceResponseBody) *OpenCdnServiceResponse {
	s.Body = v
	return s
}

func (s *OpenCdnServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
