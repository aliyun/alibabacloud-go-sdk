// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacedConfigMapsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNamespacedConfigMapsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNamespacedConfigMapsResponse
	GetStatusCode() *int32
	SetBody(v *ListNamespacedConfigMapsResponseBody) *ListNamespacedConfigMapsResponse
	GetBody() *ListNamespacedConfigMapsResponseBody
}

type ListNamespacedConfigMapsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNamespacedConfigMapsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNamespacedConfigMapsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacedConfigMapsResponse) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNamespacedConfigMapsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNamespacedConfigMapsResponse) GetBody() *ListNamespacedConfigMapsResponseBody {
	return s.Body
}

func (s *ListNamespacedConfigMapsResponse) SetHeaders(v map[string]*string) *ListNamespacedConfigMapsResponse {
	s.Headers = v
	return s
}

func (s *ListNamespacedConfigMapsResponse) SetStatusCode(v int32) *ListNamespacedConfigMapsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespacedConfigMapsResponse) SetBody(v *ListNamespacedConfigMapsResponseBody) *ListNamespacedConfigMapsResponse {
	s.Body = v
	return s
}

func (s *ListNamespacedConfigMapsResponse) Validate() error {
	return dara.Validate(s)
}
