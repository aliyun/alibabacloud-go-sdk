// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationsForDynamicRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateAccessApplicationsForDynamicRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateAccessApplicationsForDynamicRouteResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateAccessApplicationsForDynamicRouteResponseBody) *ListPrivateAccessApplicationsForDynamicRouteResponse
	GetBody() *ListPrivateAccessApplicationsForDynamicRouteResponseBody
}

type ListPrivateAccessApplicationsForDynamicRouteResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessApplicationsForDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) GetBody() *ListPrivateAccessApplicationsForDynamicRouteResponseBody {
	return s.Body
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) SetHeaders(v map[string]*string) *ListPrivateAccessApplicationsForDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) SetStatusCode(v int32) *ListPrivateAccessApplicationsForDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) SetBody(v *ListPrivateAccessApplicationsForDynamicRouteResponseBody) *ListPrivateAccessApplicationsForDynamicRouteResponse {
	s.Body = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
