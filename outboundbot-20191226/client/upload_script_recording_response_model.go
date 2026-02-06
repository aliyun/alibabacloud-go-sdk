// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadScriptRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadScriptRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadScriptRecordingResponse
	GetStatusCode() *int32
	SetBody(v *UploadScriptRecordingResponseBody) *UploadScriptRecordingResponse
	GetBody() *UploadScriptRecordingResponseBody
}

type UploadScriptRecordingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadScriptRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadScriptRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadScriptRecordingResponse) GoString() string {
	return s.String()
}

func (s *UploadScriptRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadScriptRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadScriptRecordingResponse) GetBody() *UploadScriptRecordingResponseBody {
	return s.Body
}

func (s *UploadScriptRecordingResponse) SetHeaders(v map[string]*string) *UploadScriptRecordingResponse {
	s.Headers = v
	return s
}

func (s *UploadScriptRecordingResponse) SetStatusCode(v int32) *UploadScriptRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadScriptRecordingResponse) SetBody(v *UploadScriptRecordingResponseBody) *UploadScriptRecordingResponse {
	s.Body = v
	return s
}

func (s *UploadScriptRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
