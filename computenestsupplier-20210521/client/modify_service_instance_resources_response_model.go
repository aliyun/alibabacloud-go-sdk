// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceInstanceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyServiceInstanceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyServiceInstanceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyServiceInstanceResourcesResponseBody) *ModifyServiceInstanceResourcesResponse
	GetBody() *ModifyServiceInstanceResourcesResponseBody
}

type ModifyServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyServiceInstanceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyServiceInstanceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyServiceInstanceResourcesResponse) GetBody() *ModifyServiceInstanceResourcesResponseBody {
	return s.Body
}

func (s *ModifyServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ModifyServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceInstanceResourcesResponse) SetStatusCode(v int32) *ModifyServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceInstanceResourcesResponse) SetBody(v *ModifyServiceInstanceResourcesResponseBody) *ModifyServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

func (s *ModifyServiceInstanceResourcesResponse) Validate() error {
	return dara.Validate(s)
}
