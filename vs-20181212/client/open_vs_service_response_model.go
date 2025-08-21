// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenVsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenVsServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenVsServiceResponseBody) *OpenVsServiceResponse
	GetBody() *OpenVsServiceResponseBody
}

type OpenVsServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenVsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenVsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenVsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenVsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenVsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenVsServiceResponse) GetBody() *OpenVsServiceResponseBody {
	return s.Body
}

func (s *OpenVsServiceResponse) SetHeaders(v map[string]*string) *OpenVsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenVsServiceResponse) SetStatusCode(v int32) *OpenVsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenVsServiceResponse) SetBody(v *OpenVsServiceResponseBody) *OpenVsServiceResponse {
	s.Body = v
	return s
}

func (s *OpenVsServiceResponse) Validate() error {
	return dara.Validate(s)
}
