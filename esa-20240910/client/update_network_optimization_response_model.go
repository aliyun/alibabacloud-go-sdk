// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkOptimizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkOptimizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkOptimizationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkOptimizationResponseBody) *UpdateNetworkOptimizationResponse
	GetBody() *UpdateNetworkOptimizationResponseBody
}

type UpdateNetworkOptimizationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkOptimizationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkOptimizationResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkOptimizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkOptimizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkOptimizationResponse) GetBody() *UpdateNetworkOptimizationResponseBody {
	return s.Body
}

func (s *UpdateNetworkOptimizationResponse) SetHeaders(v map[string]*string) *UpdateNetworkOptimizationResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkOptimizationResponse) SetStatusCode(v int32) *UpdateNetworkOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkOptimizationResponse) SetBody(v *UpdateNetworkOptimizationResponseBody) *UpdateNetworkOptimizationResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkOptimizationResponse) Validate() error {
	return dara.Validate(s)
}
