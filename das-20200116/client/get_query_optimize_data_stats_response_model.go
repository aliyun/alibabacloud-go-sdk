// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeDataStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeDataStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeDataStatsResponseBody) *GetQueryOptimizeDataStatsResponse
	GetBody() *GetQueryOptimizeDataStatsResponseBody
}

type GetQueryOptimizeDataStatsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeDataStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeDataStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeDataStatsResponse) GetBody() *GetQueryOptimizeDataStatsResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeDataStatsResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeDataStatsResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponse) SetStatusCode(v int32) *GetQueryOptimizeDataStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponse) SetBody(v *GetQueryOptimizeDataStatsResponseBody) *GetQueryOptimizeDataStatsResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponse) Validate() error {
	return dara.Validate(s)
}
