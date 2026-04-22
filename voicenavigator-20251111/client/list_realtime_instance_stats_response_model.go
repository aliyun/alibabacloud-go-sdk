// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeInstanceStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRealtimeInstanceStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRealtimeInstanceStatsResponse
	GetStatusCode() *int32
	SetBody(v *ListRealtimeInstanceStatsResponseBody) *ListRealtimeInstanceStatsResponse
	GetBody() *ListRealtimeInstanceStatsResponseBody
}

type ListRealtimeInstanceStatsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRealtimeInstanceStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRealtimeInstanceStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeInstanceStatsResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeInstanceStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRealtimeInstanceStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRealtimeInstanceStatsResponse) GetBody() *ListRealtimeInstanceStatsResponseBody {
	return s.Body
}

func (s *ListRealtimeInstanceStatsResponse) SetHeaders(v map[string]*string) *ListRealtimeInstanceStatsResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeInstanceStatsResponse) SetStatusCode(v int32) *ListRealtimeInstanceStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRealtimeInstanceStatsResponse) SetBody(v *ListRealtimeInstanceStatsResponseBody) *ListRealtimeInstanceStatsResponse {
	s.Body = v
	return s
}

func (s *ListRealtimeInstanceStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
