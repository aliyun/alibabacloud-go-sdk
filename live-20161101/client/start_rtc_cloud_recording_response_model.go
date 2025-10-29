// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRtcCloudRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRtcCloudRecordingResponse
	GetStatusCode() *int32
	SetBody(v *StartRtcCloudRecordingResponseBody) *StartRtcCloudRecordingResponse
	GetBody() *StartRtcCloudRecordingResponseBody
}

type StartRtcCloudRecordingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRtcCloudRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRtcCloudRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingResponse) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRtcCloudRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRtcCloudRecordingResponse) GetBody() *StartRtcCloudRecordingResponseBody {
	return s.Body
}

func (s *StartRtcCloudRecordingResponse) SetHeaders(v map[string]*string) *StartRtcCloudRecordingResponse {
	s.Headers = v
	return s
}

func (s *StartRtcCloudRecordingResponse) SetStatusCode(v int32) *StartRtcCloudRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRtcCloudRecordingResponse) SetBody(v *StartRtcCloudRecordingResponseBody) *StartRtcCloudRecordingResponse {
	s.Body = v
	return s
}

func (s *StartRtcCloudRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
