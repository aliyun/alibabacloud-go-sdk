// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcCloudRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRtcCloudRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRtcCloudRecordingResponse
	GetStatusCode() *int32
	SetBody(v *StopRtcCloudRecordingResponseBody) *StopRtcCloudRecordingResponse
	GetBody() *StopRtcCloudRecordingResponseBody
}

type StopRtcCloudRecordingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRtcCloudRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRtcCloudRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRtcCloudRecordingResponse) GoString() string {
	return s.String()
}

func (s *StopRtcCloudRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRtcCloudRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRtcCloudRecordingResponse) GetBody() *StopRtcCloudRecordingResponseBody {
	return s.Body
}

func (s *StopRtcCloudRecordingResponse) SetHeaders(v map[string]*string) *StopRtcCloudRecordingResponse {
	s.Headers = v
	return s
}

func (s *StopRtcCloudRecordingResponse) SetStatusCode(v int32) *StopRtcCloudRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRtcCloudRecordingResponse) SetBody(v *StopRtcCloudRecordingResponseBody) *StopRtcCloudRecordingResponse {
	s.Body = v
	return s
}

func (s *StopRtcCloudRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
