// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubeOptimizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCubeOptimizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCubeOptimizationResponse
	GetStatusCode() *int32
	SetBody(v *QueryCubeOptimizationResponseBody) *QueryCubeOptimizationResponse
	GetBody() *QueryCubeOptimizationResponseBody
}

type QueryCubeOptimizationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCubeOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCubeOptimizationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCubeOptimizationResponse) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCubeOptimizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCubeOptimizationResponse) GetBody() *QueryCubeOptimizationResponseBody {
	return s.Body
}

func (s *QueryCubeOptimizationResponse) SetHeaders(v map[string]*string) *QueryCubeOptimizationResponse {
	s.Headers = v
	return s
}

func (s *QueryCubeOptimizationResponse) SetStatusCode(v int32) *QueryCubeOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCubeOptimizationResponse) SetBody(v *QueryCubeOptimizationResponseBody) *QueryCubeOptimizationResponse {
	s.Body = v
	return s
}

func (s *QueryCubeOptimizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
