// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenHbrServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenHbrServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenHbrServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenHbrServiceResponseBody) *OpenHbrServiceResponse
	GetBody() *OpenHbrServiceResponseBody
}

type OpenHbrServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenHbrServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenHbrServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenHbrServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenHbrServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenHbrServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenHbrServiceResponse) GetBody() *OpenHbrServiceResponseBody {
	return s.Body
}

func (s *OpenHbrServiceResponse) SetHeaders(v map[string]*string) *OpenHbrServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenHbrServiceResponse) SetStatusCode(v int32) *OpenHbrServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenHbrServiceResponse) SetBody(v *OpenHbrServiceResponseBody) *OpenHbrServiceResponse {
	s.Body = v
	return s
}

func (s *OpenHbrServiceResponse) Validate() error {
	return dara.Validate(s)
}
