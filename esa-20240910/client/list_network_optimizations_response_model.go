// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkOptimizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkOptimizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkOptimizationsResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkOptimizationsResponseBody) *ListNetworkOptimizationsResponse
	GetBody() *ListNetworkOptimizationsResponseBody
}

type ListNetworkOptimizationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkOptimizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkOptimizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkOptimizationsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkOptimizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkOptimizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkOptimizationsResponse) GetBody() *ListNetworkOptimizationsResponseBody {
	return s.Body
}

func (s *ListNetworkOptimizationsResponse) SetHeaders(v map[string]*string) *ListNetworkOptimizationsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkOptimizationsResponse) SetStatusCode(v int32) *ListNetworkOptimizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkOptimizationsResponse) SetBody(v *ListNetworkOptimizationsResponseBody) *ListNetworkOptimizationsResponse {
	s.Body = v
	return s
}

func (s *ListNetworkOptimizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
