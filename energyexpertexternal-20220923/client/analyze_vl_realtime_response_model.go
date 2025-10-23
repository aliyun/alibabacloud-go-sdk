// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeVlRealtimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeVlRealtimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeVlRealtimeResponse
	GetStatusCode() *int32
	SetBody(v *AnalyzeVlRealtimeResponseBody) *AnalyzeVlRealtimeResponse
	GetBody() *AnalyzeVlRealtimeResponseBody
}

type AnalyzeVlRealtimeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeVlRealtimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeVlRealtimeResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeVlRealtimeResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeVlRealtimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeVlRealtimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeVlRealtimeResponse) GetBody() *AnalyzeVlRealtimeResponseBody {
	return s.Body
}

func (s *AnalyzeVlRealtimeResponse) SetHeaders(v map[string]*string) *AnalyzeVlRealtimeResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeVlRealtimeResponse) SetStatusCode(v int32) *AnalyzeVlRealtimeResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeVlRealtimeResponse) SetBody(v *AnalyzeVlRealtimeResponseBody) *AnalyzeVlRealtimeResponse {
	s.Body = v
	return s
}

func (s *AnalyzeVlRealtimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
