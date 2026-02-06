// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadScriptRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadScriptRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadScriptRecordingResponse
	GetStatusCode() *int32
	SetBody(v *DownloadScriptRecordingResponseBody) *DownloadScriptRecordingResponse
	GetBody() *DownloadScriptRecordingResponseBody
}

type DownloadScriptRecordingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadScriptRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadScriptRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadScriptRecordingResponse) GoString() string {
	return s.String()
}

func (s *DownloadScriptRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadScriptRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadScriptRecordingResponse) GetBody() *DownloadScriptRecordingResponseBody {
	return s.Body
}

func (s *DownloadScriptRecordingResponse) SetHeaders(v map[string]*string) *DownloadScriptRecordingResponse {
	s.Headers = v
	return s
}

func (s *DownloadScriptRecordingResponse) SetStatusCode(v int32) *DownloadScriptRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadScriptRecordingResponse) SetBody(v *DownloadScriptRecordingResponseBody) *DownloadScriptRecordingResponse {
	s.Body = v
	return s
}

func (s *DownloadScriptRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
