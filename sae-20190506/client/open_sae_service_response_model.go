// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSaeServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenSaeServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenSaeServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenSaeServiceResponseBody) *OpenSaeServiceResponse
	GetBody() *OpenSaeServiceResponseBody
}

type OpenSaeServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenSaeServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenSaeServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenSaeServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenSaeServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenSaeServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenSaeServiceResponse) GetBody() *OpenSaeServiceResponseBody {
	return s.Body
}

func (s *OpenSaeServiceResponse) SetHeaders(v map[string]*string) *OpenSaeServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenSaeServiceResponse) SetStatusCode(v int32) *OpenSaeServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenSaeServiceResponse) SetBody(v *OpenSaeServiceResponseBody) *OpenSaeServiceResponse {
	s.Body = v
	return s
}

func (s *OpenSaeServiceResponse) Validate() error {
	return dara.Validate(s)
}
