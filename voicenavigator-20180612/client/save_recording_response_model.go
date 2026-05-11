// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveRecordingResponse
	GetStatusCode() *int32
	SetBody(v *SaveRecordingResponseBody) *SaveRecordingResponse
	GetBody() *SaveRecordingResponseBody
}

type SaveRecordingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveRecordingResponse) GoString() string {
	return s.String()
}

func (s *SaveRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveRecordingResponse) GetBody() *SaveRecordingResponseBody {
	return s.Body
}

func (s *SaveRecordingResponse) SetHeaders(v map[string]*string) *SaveRecordingResponse {
	s.Headers = v
	return s
}

func (s *SaveRecordingResponse) SetStatusCode(v int32) *SaveRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveRecordingResponse) SetBody(v *SaveRecordingResponseBody) *SaveRecordingResponse {
	s.Body = v
	return s
}

func (s *SaveRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
