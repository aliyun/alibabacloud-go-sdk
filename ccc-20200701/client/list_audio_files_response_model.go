// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAudioFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAudioFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAudioFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListAudioFilesResponseBody) *ListAudioFilesResponse
	GetBody() *ListAudioFilesResponseBody
}

type ListAudioFilesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAudioFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAudioFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAudioFilesResponse) GoString() string {
	return s.String()
}

func (s *ListAudioFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAudioFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAudioFilesResponse) GetBody() *ListAudioFilesResponseBody {
	return s.Body
}

func (s *ListAudioFilesResponse) SetHeaders(v map[string]*string) *ListAudioFilesResponse {
	s.Headers = v
	return s
}

func (s *ListAudioFilesResponse) SetStatusCode(v int32) *ListAudioFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAudioFilesResponse) SetBody(v *ListAudioFilesResponseBody) *ListAudioFilesResponse {
	s.Body = v
	return s
}

func (s *ListAudioFilesResponse) Validate() error {
	return dara.Validate(s)
}
