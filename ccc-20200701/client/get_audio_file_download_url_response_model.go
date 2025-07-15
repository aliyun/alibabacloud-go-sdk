// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAudioFileDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAudioFileDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAudioFileDownloadUrlResponseBody) *GetAudioFileDownloadUrlResponse
	GetBody() *GetAudioFileDownloadUrlResponseBody
}

type GetAudioFileDownloadUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAudioFileDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAudioFileDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAudioFileDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAudioFileDownloadUrlResponse) GetBody() *GetAudioFileDownloadUrlResponseBody {
	return s.Body
}

func (s *GetAudioFileDownloadUrlResponse) SetHeaders(v map[string]*string) *GetAudioFileDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAudioFileDownloadUrlResponse) SetStatusCode(v int32) *GetAudioFileDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAudioFileDownloadUrlResponse) SetBody(v *GetAudioFileDownloadUrlResponseBody) *GetAudioFileDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetAudioFileDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
