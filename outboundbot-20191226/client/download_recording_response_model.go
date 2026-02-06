// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadRecordingResponse
	GetStatusCode() *int32
	SetBody(v *DownloadRecordingResponseBody) *DownloadRecordingResponse
	GetBody() *DownloadRecordingResponseBody
}

type DownloadRecordingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadRecordingResponse) GoString() string {
	return s.String()
}

func (s *DownloadRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadRecordingResponse) GetBody() *DownloadRecordingResponseBody {
	return s.Body
}

func (s *DownloadRecordingResponse) SetHeaders(v map[string]*string) *DownloadRecordingResponse {
	s.Headers = v
	return s
}

func (s *DownloadRecordingResponse) SetStatusCode(v int32) *DownloadRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadRecordingResponse) SetBody(v *DownloadRecordingResponseBody) *DownloadRecordingResponse {
	s.Body = v
	return s
}

func (s *DownloadRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
