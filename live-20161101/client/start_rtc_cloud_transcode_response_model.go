// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRtcCloudTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRtcCloudTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *StartRtcCloudTranscodeResponseBody) *StartRtcCloudTranscodeResponse
	GetBody() *StartRtcCloudTranscodeResponseBody
}

type StartRtcCloudTranscodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRtcCloudTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRtcCloudTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudTranscodeResponse) GoString() string {
	return s.String()
}

func (s *StartRtcCloudTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRtcCloudTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRtcCloudTranscodeResponse) GetBody() *StartRtcCloudTranscodeResponseBody {
	return s.Body
}

func (s *StartRtcCloudTranscodeResponse) SetHeaders(v map[string]*string) *StartRtcCloudTranscodeResponse {
	s.Headers = v
	return s
}

func (s *StartRtcCloudTranscodeResponse) SetStatusCode(v int32) *StartRtcCloudTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRtcCloudTranscodeResponse) SetBody(v *StartRtcCloudTranscodeResponseBody) *StartRtcCloudTranscodeResponse {
	s.Body = v
	return s
}

func (s *StartRtcCloudTranscodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
