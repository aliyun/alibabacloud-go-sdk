// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeScriptStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealtimeScriptStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealtimeScriptStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetRealtimeScriptStatsResponseBody) *GetRealtimeScriptStatsResponse
	GetBody() *GetRealtimeScriptStatsResponseBody
}

type GetRealtimeScriptStatsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeScriptStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeScriptStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeScriptStatsResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeScriptStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealtimeScriptStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealtimeScriptStatsResponse) GetBody() *GetRealtimeScriptStatsResponseBody {
	return s.Body
}

func (s *GetRealtimeScriptStatsResponse) SetHeaders(v map[string]*string) *GetRealtimeScriptStatsResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeScriptStatsResponse) SetStatusCode(v int32) *GetRealtimeScriptStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeScriptStatsResponse) SetBody(v *GetRealtimeScriptStatsResponseBody) *GetRealtimeScriptStatsResponse {
	s.Body = v
	return s
}

func (s *GetRealtimeScriptStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
