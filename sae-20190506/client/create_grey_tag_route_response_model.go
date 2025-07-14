// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGreyTagRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGreyTagRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGreyTagRouteResponse
	GetStatusCode() *int32
	SetBody(v *CreateGreyTagRouteResponseBody) *CreateGreyTagRouteResponse
	GetBody() *CreateGreyTagRouteResponseBody
}

type CreateGreyTagRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGreyTagRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGreyTagRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGreyTagRouteResponse) GetBody() *CreateGreyTagRouteResponseBody {
	return s.Body
}

func (s *CreateGreyTagRouteResponse) SetHeaders(v map[string]*string) *CreateGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateGreyTagRouteResponse) SetStatusCode(v int32) *CreateGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGreyTagRouteResponse) SetBody(v *CreateGreyTagRouteResponseBody) *CreateGreyTagRouteResponse {
	s.Body = v
	return s
}

func (s *CreateGreyTagRouteResponse) Validate() error {
	return dara.Validate(s)
}
