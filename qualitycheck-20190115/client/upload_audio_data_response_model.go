// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAudioDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadAudioDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadAudioDataResponse
	GetStatusCode() *int32
	SetBody(v *UploadAudioDataResponseBody) *UploadAudioDataResponse
	GetBody() *UploadAudioDataResponseBody
}

type UploadAudioDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAudioDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAudioDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAudioDataResponse) GoString() string {
	return s.String()
}

func (s *UploadAudioDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAudioDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadAudioDataResponse) GetBody() *UploadAudioDataResponseBody {
	return s.Body
}

func (s *UploadAudioDataResponse) SetHeaders(v map[string]*string) *UploadAudioDataResponse {
	s.Headers = v
	return s
}

func (s *UploadAudioDataResponse) SetStatusCode(v int32) *UploadAudioDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAudioDataResponse) SetBody(v *UploadAudioDataResponseBody) *UploadAudioDataResponse {
	s.Body = v
	return s
}

func (s *UploadAudioDataResponse) Validate() error {
	return dara.Validate(s)
}
