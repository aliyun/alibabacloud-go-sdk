// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterAvailableResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterAvailableResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterAvailableResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterAvailableResourceResponseBody) *ListTransitRouterAvailableResourceResponse
	GetBody() *ListTransitRouterAvailableResourceResponseBody
}

type ListTransitRouterAvailableResourceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterAvailableResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterAvailableResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterAvailableResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterAvailableResourceResponse) GetBody() *ListTransitRouterAvailableResourceResponseBody {
	return s.Body
}

func (s *ListTransitRouterAvailableResourceResponse) SetHeaders(v map[string]*string) *ListTransitRouterAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterAvailableResourceResponse) SetStatusCode(v int32) *ListTransitRouterAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterAvailableResourceResponse) SetBody(v *ListTransitRouterAvailableResourceResponseBody) *ListTransitRouterAvailableResourceResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterAvailableResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
