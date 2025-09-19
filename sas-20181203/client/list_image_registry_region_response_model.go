// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRegistryRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageRegistryRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageRegistryRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListImageRegistryRegionResponseBody) *ListImageRegistryRegionResponse
	GetBody() *ListImageRegistryRegionResponseBody
}

type ListImageRegistryRegionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageRegistryRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageRegistryRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageRegistryRegionResponse) GoString() string {
	return s.String()
}

func (s *ListImageRegistryRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageRegistryRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageRegistryRegionResponse) GetBody() *ListImageRegistryRegionResponseBody {
	return s.Body
}

func (s *ListImageRegistryRegionResponse) SetHeaders(v map[string]*string) *ListImageRegistryRegionResponse {
	s.Headers = v
	return s
}

func (s *ListImageRegistryRegionResponse) SetStatusCode(v int32) *ListImageRegistryRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageRegistryRegionResponse) SetBody(v *ListImageRegistryRegionResponseBody) *ListImageRegistryRegionResponse {
	s.Body = v
	return s
}

func (s *ListImageRegistryRegionResponse) Validate() error {
	return dara.Validate(s)
}
