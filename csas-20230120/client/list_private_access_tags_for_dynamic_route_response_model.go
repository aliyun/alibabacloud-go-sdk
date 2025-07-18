// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessTagsForDynamicRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateAccessTagsForDynamicRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateAccessTagsForDynamicRouteResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateAccessTagsForDynamicRouteResponseBody) *ListPrivateAccessTagsForDynamicRouteResponse
	GetBody() *ListPrivateAccessTagsForDynamicRouteResponseBody
}

type ListPrivateAccessTagsForDynamicRouteResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessTagsForDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) GetBody() *ListPrivateAccessTagsForDynamicRouteResponseBody {
	return s.Body
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) SetHeaders(v map[string]*string) *ListPrivateAccessTagsForDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) SetStatusCode(v int32) *ListPrivateAccessTagsForDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) SetBody(v *ListPrivateAccessTagsForDynamicRouteResponseBody) *ListPrivateAccessTagsForDynamicRouteResponse {
	s.Body = v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) Validate() error {
	return dara.Validate(s)
}
