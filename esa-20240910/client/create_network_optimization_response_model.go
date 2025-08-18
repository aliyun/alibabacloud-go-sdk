// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkOptimizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkOptimizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkOptimizationResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkOptimizationResponseBody) *CreateNetworkOptimizationResponse
	GetBody() *CreateNetworkOptimizationResponseBody
}

type CreateNetworkOptimizationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkOptimizationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkOptimizationResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkOptimizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkOptimizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkOptimizationResponse) GetBody() *CreateNetworkOptimizationResponseBody {
	return s.Body
}

func (s *CreateNetworkOptimizationResponse) SetHeaders(v map[string]*string) *CreateNetworkOptimizationResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkOptimizationResponse) SetStatusCode(v int32) *CreateNetworkOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkOptimizationResponse) SetBody(v *CreateNetworkOptimizationResponseBody) *CreateNetworkOptimizationResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkOptimizationResponse) Validate() error {
	return dara.Validate(s)
}
