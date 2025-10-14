// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkOptimizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkOptimizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkOptimizationResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkOptimizationResponseBody) *GetNetworkOptimizationResponse
	GetBody() *GetNetworkOptimizationResponseBody
}

type GetNetworkOptimizationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkOptimizationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkOptimizationResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkOptimizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkOptimizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkOptimizationResponse) GetBody() *GetNetworkOptimizationResponseBody {
	return s.Body
}

func (s *GetNetworkOptimizationResponse) SetHeaders(v map[string]*string) *GetNetworkOptimizationResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkOptimizationResponse) SetStatusCode(v int32) *GetNetworkOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkOptimizationResponse) SetBody(v *GetNetworkOptimizationResponseBody) *GetNetworkOptimizationResponse {
	s.Body = v
	return s
}

func (s *GetNetworkOptimizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
