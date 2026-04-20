// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGrantInstanceToTransitRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGrantInstanceToTransitRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGrantInstanceToTransitRouterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGrantInstanceToTransitRouterResponseBody) *ModifyGrantInstanceToTransitRouterResponse
	GetBody() *ModifyGrantInstanceToTransitRouterResponseBody
}

type ModifyGrantInstanceToTransitRouterResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGrantInstanceToTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGrantInstanceToTransitRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGrantInstanceToTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *ModifyGrantInstanceToTransitRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGrantInstanceToTransitRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGrantInstanceToTransitRouterResponse) GetBody() *ModifyGrantInstanceToTransitRouterResponseBody {
	return s.Body
}

func (s *ModifyGrantInstanceToTransitRouterResponse) SetHeaders(v map[string]*string) *ModifyGrantInstanceToTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterResponse) SetStatusCode(v int32) *ModifyGrantInstanceToTransitRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterResponse) SetBody(v *ModifyGrantInstanceToTransitRouterResponseBody) *ModifyGrantInstanceToTransitRouterResponse {
	s.Body = v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
