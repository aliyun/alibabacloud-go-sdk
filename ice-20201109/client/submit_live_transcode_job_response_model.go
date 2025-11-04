// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitLiveTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitLiveTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitLiveTranscodeJobResponseBody) *SubmitLiveTranscodeJobResponse
	GetBody() *SubmitLiveTranscodeJobResponseBody
}

type SubmitLiveTranscodeJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitLiveTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitLiveTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitLiveTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitLiveTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitLiveTranscodeJobResponse) GetBody() *SubmitLiveTranscodeJobResponseBody {
	return s.Body
}

func (s *SubmitLiveTranscodeJobResponse) SetHeaders(v map[string]*string) *SubmitLiveTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitLiveTranscodeJobResponse) SetStatusCode(v int32) *SubmitLiveTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitLiveTranscodeJobResponse) SetBody(v *SubmitLiveTranscodeJobResponseBody) *SubmitLiveTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *SubmitLiveTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
