// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTranscodeJobResponseBody) *SubmitTranscodeJobResponse
	GetBody() *SubmitTranscodeJobResponseBody
}

type SubmitTranscodeJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTranscodeJobResponse) GetBody() *SubmitTranscodeJobResponseBody {
	return s.Body
}

func (s *SubmitTranscodeJobResponse) SetHeaders(v map[string]*string) *SubmitTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitTranscodeJobResponse) SetStatusCode(v int32) *SubmitTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTranscodeJobResponse) SetBody(v *SubmitTranscodeJobResponseBody) *SubmitTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *SubmitTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
