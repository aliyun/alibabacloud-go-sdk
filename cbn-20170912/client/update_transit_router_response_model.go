// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterResponseBody) *UpdateTransitRouterResponse
	GetBody() *UpdateTransitRouterResponseBody
}

type UpdateTransitRouterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterResponse) GetBody() *UpdateTransitRouterResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterResponse) SetStatusCode(v int32) *UpdateTransitRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterResponse) SetBody(v *UpdateTransitRouterResponseBody) *UpdateTransitRouterResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterResponse) Validate() error {
	return dara.Validate(s)
}
