// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRouteMapResponseBody) *UpdateRouteMapResponse
	GetBody() *UpdateRouteMapResponseBody
}

type UpdateRouteMapResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRouteMapResponse) GetBody() *UpdateRouteMapResponseBody {
	return s.Body
}

func (s *UpdateRouteMapResponse) SetHeaders(v map[string]*string) *UpdateRouteMapResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteMapResponse) SetStatusCode(v int32) *UpdateRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteMapResponse) SetBody(v *UpdateRouteMapResponseBody) *UpdateRouteMapResponse {
	s.Body = v
	return s
}

func (s *UpdateRouteMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
