// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitRecordingResponse
	GetStatusCode() *int32
	SetBody(v *SubmitRecordingResponseBody) *SubmitRecordingResponse
	GetBody() *SubmitRecordingResponseBody
}

type SubmitRecordingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitRecordingResponse) GoString() string {
	return s.String()
}

func (s *SubmitRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitRecordingResponse) GetBody() *SubmitRecordingResponseBody {
	return s.Body
}

func (s *SubmitRecordingResponse) SetHeaders(v map[string]*string) *SubmitRecordingResponse {
	s.Headers = v
	return s
}

func (s *SubmitRecordingResponse) SetStatusCode(v int32) *SubmitRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitRecordingResponse) SetBody(v *SubmitRecordingResponseBody) *SubmitRecordingResponse {
	s.Body = v
	return s
}

func (s *SubmitRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
