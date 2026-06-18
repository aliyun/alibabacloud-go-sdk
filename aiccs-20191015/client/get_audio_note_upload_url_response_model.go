// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioNoteUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAudioNoteUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAudioNoteUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAudioNoteUploadUrlResponseBody) *GetAudioNoteUploadUrlResponse
	GetBody() *GetAudioNoteUploadUrlResponseBody
}

type GetAudioNoteUploadUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAudioNoteUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAudioNoteUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAudioNoteUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAudioNoteUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAudioNoteUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAudioNoteUploadUrlResponse) GetBody() *GetAudioNoteUploadUrlResponseBody {
	return s.Body
}

func (s *GetAudioNoteUploadUrlResponse) SetHeaders(v map[string]*string) *GetAudioNoteUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAudioNoteUploadUrlResponse) SetStatusCode(v int32) *GetAudioNoteUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponse) SetBody(v *GetAudioNoteUploadUrlResponseBody) *GetAudioNoteUploadUrlResponse {
	s.Body = v
	return s
}

func (s *GetAudioNoteUploadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
