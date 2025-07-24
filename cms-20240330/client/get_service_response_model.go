// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceResponseBody) *GetServiceResponse
	GetBody() *GetServiceResponseBody
}

type GetServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceResponse) GetBody() *GetServiceResponseBody {
	return s.Body
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

func (s *GetServiceResponse) Validate() error {
	return dara.Validate(s)
}
