// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeInstanceStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealtimeInstanceStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealtimeInstanceStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetRealtimeInstanceStatsResponseBody) *GetRealtimeInstanceStatsResponse
	GetBody() *GetRealtimeInstanceStatsResponseBody
}

type GetRealtimeInstanceStatsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeInstanceStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeInstanceStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatsResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealtimeInstanceStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealtimeInstanceStatsResponse) GetBody() *GetRealtimeInstanceStatsResponseBody {
	return s.Body
}

func (s *GetRealtimeInstanceStatsResponse) SetHeaders(v map[string]*string) *GetRealtimeInstanceStatsResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeInstanceStatsResponse) SetStatusCode(v int32) *GetRealtimeInstanceStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponse) SetBody(v *GetRealtimeInstanceStatsResponseBody) *GetRealtimeInstanceStatsResponse {
	s.Body = v
	return s
}

func (s *GetRealtimeInstanceStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
