// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcCloudTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRtcCloudTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRtcCloudTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *StopRtcCloudTranscodeResponseBody) *StopRtcCloudTranscodeResponse
	GetBody() *StopRtcCloudTranscodeResponseBody
}

type StopRtcCloudTranscodeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRtcCloudTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRtcCloudTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRtcCloudTranscodeResponse) GoString() string {
	return s.String()
}

func (s *StopRtcCloudTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRtcCloudTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRtcCloudTranscodeResponse) GetBody() *StopRtcCloudTranscodeResponseBody {
	return s.Body
}

func (s *StopRtcCloudTranscodeResponse) SetHeaders(v map[string]*string) *StopRtcCloudTranscodeResponse {
	s.Headers = v
	return s
}

func (s *StopRtcCloudTranscodeResponse) SetStatusCode(v int32) *StopRtcCloudTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRtcCloudTranscodeResponse) SetBody(v *StopRtcCloudTranscodeResponseBody) *StopRtcCloudTranscodeResponse {
	s.Body = v
	return s
}

func (s *StopRtcCloudTranscodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
