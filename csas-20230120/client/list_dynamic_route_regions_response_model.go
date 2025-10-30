// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicRouteRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDynamicRouteRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDynamicRouteRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDynamicRouteRegionsResponseBody) *ListDynamicRouteRegionsResponse
	GetBody() *ListDynamicRouteRegionsResponseBody
}

type ListDynamicRouteRegionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicRouteRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicRouteRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicRouteRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicRouteRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDynamicRouteRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDynamicRouteRegionsResponse) GetBody() *ListDynamicRouteRegionsResponseBody {
	return s.Body
}

func (s *ListDynamicRouteRegionsResponse) SetHeaders(v map[string]*string) *ListDynamicRouteRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicRouteRegionsResponse) SetStatusCode(v int32) *ListDynamicRouteRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicRouteRegionsResponse) SetBody(v *ListDynamicRouteRegionsResponseBody) *ListDynamicRouteRegionsResponse {
	s.Body = v
	return s
}

func (s *ListDynamicRouteRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
