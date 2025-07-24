// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewServiceInstanceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewServiceInstanceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewServiceInstanceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *RenewServiceInstanceResourcesResponseBody) *RenewServiceInstanceResourcesResponse
	GetBody() *RenewServiceInstanceResourcesResponseBody
}

type RenewServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewServiceInstanceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewServiceInstanceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewServiceInstanceResourcesResponse) GetBody() *RenewServiceInstanceResourcesResponseBody {
	return s.Body
}

func (s *RenewServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *RenewServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *RenewServiceInstanceResourcesResponse) SetStatusCode(v int32) *RenewServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponse) SetBody(v *RenewServiceInstanceResourcesResponseBody) *RenewServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

func (s *RenewServiceInstanceResourcesResponse) Validate() error {
	return dara.Validate(s)
}
