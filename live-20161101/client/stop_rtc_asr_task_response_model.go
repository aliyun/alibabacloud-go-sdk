// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcAsrTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRtcAsrTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRtcAsrTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopRtcAsrTaskResponseBody) *StopRtcAsrTaskResponse
	GetBody() *StopRtcAsrTaskResponseBody
}

type StopRtcAsrTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRtcAsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRtcAsrTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRtcAsrTaskResponse) GoString() string {
	return s.String()
}

func (s *StopRtcAsrTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRtcAsrTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRtcAsrTaskResponse) GetBody() *StopRtcAsrTaskResponseBody {
	return s.Body
}

func (s *StopRtcAsrTaskResponse) SetHeaders(v map[string]*string) *StopRtcAsrTaskResponse {
	s.Headers = v
	return s
}

func (s *StopRtcAsrTaskResponse) SetStatusCode(v int32) *StopRtcAsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRtcAsrTaskResponse) SetBody(v *StopRtcAsrTaskResponseBody) *StopRtcAsrTaskResponse {
	s.Body = v
	return s
}

func (s *StopRtcAsrTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
