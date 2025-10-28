// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeExecErrorStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeExecErrorStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeExecErrorStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeExecErrorStatsResponseBody) *GetQueryOptimizeExecErrorStatsResponse
	GetBody() *GetQueryOptimizeExecErrorStatsResponseBody
}

type GetQueryOptimizeExecErrorStatsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeExecErrorStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeExecErrorStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeExecErrorStatsResponse) GetBody() *GetQueryOptimizeExecErrorStatsResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeExecErrorStatsResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeExecErrorStatsResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponse) SetStatusCode(v int32) *GetQueryOptimizeExecErrorStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponse) SetBody(v *GetQueryOptimizeExecErrorStatsResponseBody) *GetQueryOptimizeExecErrorStatsResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
