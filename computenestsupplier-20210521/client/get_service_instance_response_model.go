// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse
	GetBody() *GetServiceInstanceResponseBody
}

type GetServiceInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceInstanceResponse) GetBody() *GetServiceInstanceResponseBody {
	return s.Body
}

func (s *GetServiceInstanceResponse) SetHeaders(v map[string]*string) *GetServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceResponse) SetStatusCode(v int32) *GetServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *GetServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
