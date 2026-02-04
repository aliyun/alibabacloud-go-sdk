// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToTransitRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantInstanceToTransitRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantInstanceToTransitRouterResponse
	GetStatusCode() *int32
	SetBody(v *GrantInstanceToTransitRouterResponseBody) *GrantInstanceToTransitRouterResponse
	GetBody() *GrantInstanceToTransitRouterResponseBody
}

type GrantInstanceToTransitRouterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantInstanceToTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantInstanceToTransitRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToTransitRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantInstanceToTransitRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantInstanceToTransitRouterResponse) GetBody() *GrantInstanceToTransitRouterResponseBody {
	return s.Body
}

func (s *GrantInstanceToTransitRouterResponse) SetHeaders(v map[string]*string) *GrantInstanceToTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToTransitRouterResponse) SetStatusCode(v int32) *GrantInstanceToTransitRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantInstanceToTransitRouterResponse) SetBody(v *GrantInstanceToTransitRouterResponseBody) *GrantInstanceToTransitRouterResponse {
	s.Body = v
	return s
}

func (s *GrantInstanceToTransitRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
