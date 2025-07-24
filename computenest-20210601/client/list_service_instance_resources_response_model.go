// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstanceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstanceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstanceResourcesResponseBody) *ListServiceInstanceResourcesResponse
	GetBody() *ListServiceInstanceResourcesResponseBody
}

type ListServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstanceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstanceResourcesResponse) GetBody() *ListServiceInstanceResourcesResponseBody {
	return s.Body
}

func (s *ListServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetStatusCode(v int32) *ListServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetBody(v *ListServiceInstanceResourcesResponseBody) *ListServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstanceResourcesResponse) Validate() error {
	return dara.Validate(s)
}
