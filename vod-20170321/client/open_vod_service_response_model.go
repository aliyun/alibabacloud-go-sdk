// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVodServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenVodServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenVodServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenVodServiceResponseBody) *OpenVodServiceResponse
	GetBody() *OpenVodServiceResponseBody
}

type OpenVodServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenVodServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenVodServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenVodServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenVodServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenVodServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenVodServiceResponse) GetBody() *OpenVodServiceResponseBody {
	return s.Body
}

func (s *OpenVodServiceResponse) SetHeaders(v map[string]*string) *OpenVodServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenVodServiceResponse) SetStatusCode(v int32) *OpenVodServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenVodServiceResponse) SetBody(v *OpenVodServiceResponseBody) *OpenVodServiceResponse {
	s.Body = v
	return s
}

func (s *OpenVodServiceResponse) Validate() error {
	return dara.Validate(s)
}
