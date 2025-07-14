// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGreyTagRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGreyTagRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGreyTagRouteResponse
	GetStatusCode() *int32
	SetBody(v *ListGreyTagRouteResponseBody) *ListGreyTagRouteResponse
	GetBody() *ListGreyTagRouteResponseBody
}

type ListGreyTagRouteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGreyTagRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGreyTagRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGreyTagRouteResponse) GetBody() *ListGreyTagRouteResponseBody {
	return s.Body
}

func (s *ListGreyTagRouteResponse) SetHeaders(v map[string]*string) *ListGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *ListGreyTagRouteResponse) SetStatusCode(v int32) *ListGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGreyTagRouteResponse) SetBody(v *ListGreyTagRouteResponseBody) *ListGreyTagRouteResponse {
	s.Body = v
	return s
}

func (s *ListGreyTagRouteResponse) Validate() error {
	return dara.Validate(s)
}
