// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeSolutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeSolutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeSolutionResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeSolutionResponseBody) *GetQueryOptimizeSolutionResponse
	GetBody() *GetQueryOptimizeSolutionResponseBody
}

type GetQueryOptimizeSolutionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeSolutionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeSolutionResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeSolutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeSolutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeSolutionResponse) GetBody() *GetQueryOptimizeSolutionResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeSolutionResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeSolutionResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeSolutionResponse) SetStatusCode(v int32) *GetQueryOptimizeSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeSolutionResponse) SetBody(v *GetQueryOptimizeSolutionResponseBody) *GetQueryOptimizeSolutionResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeSolutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
