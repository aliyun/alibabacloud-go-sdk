// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLiveStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLiveStreamingResponse
	GetStatusCode() *int32
	SetBody(v *StopLiveStreamingResponseBody) *StopLiveStreamingResponse
	GetBody() *StopLiveStreamingResponseBody
}

type StopLiveStreamingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLiveStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLiveStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLiveStreamingResponse) GoString() string {
	return s.String()
}

func (s *StopLiveStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLiveStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLiveStreamingResponse) GetBody() *StopLiveStreamingResponseBody {
	return s.Body
}

func (s *StopLiveStreamingResponse) SetHeaders(v map[string]*string) *StopLiveStreamingResponse {
	s.Headers = v
	return s
}

func (s *StopLiveStreamingResponse) SetStatusCode(v int32) *StopLiveStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLiveStreamingResponse) SetBody(v *StopLiveStreamingResponseBody) *StopLiveStreamingResponse {
	s.Body = v
	return s
}

func (s *StopLiveStreamingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
