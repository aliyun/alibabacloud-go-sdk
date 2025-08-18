// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkOptimizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkOptimizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkOptimizationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkOptimizationResponseBody) *DeleteNetworkOptimizationResponse
	GetBody() *DeleteNetworkOptimizationResponseBody
}

type DeleteNetworkOptimizationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkOptimizationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkOptimizationResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkOptimizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkOptimizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkOptimizationResponse) GetBody() *DeleteNetworkOptimizationResponseBody {
	return s.Body
}

func (s *DeleteNetworkOptimizationResponse) SetHeaders(v map[string]*string) *DeleteNetworkOptimizationResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkOptimizationResponse) SetStatusCode(v int32) *DeleteNetworkOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkOptimizationResponse) SetBody(v *DeleteNetworkOptimizationResponseBody) *DeleteNetworkOptimizationResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkOptimizationResponse) Validate() error {
	return dara.Validate(s)
}
