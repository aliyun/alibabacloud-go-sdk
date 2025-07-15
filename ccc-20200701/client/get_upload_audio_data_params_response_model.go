// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadAudioDataParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadAudioDataParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadAudioDataParamsResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadAudioDataParamsResponseBody) *GetUploadAudioDataParamsResponse
	GetBody() *GetUploadAudioDataParamsResponseBody
}

type GetUploadAudioDataParamsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadAudioDataParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadAudioDataParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadAudioDataParamsResponse) GoString() string {
	return s.String()
}

func (s *GetUploadAudioDataParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadAudioDataParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadAudioDataParamsResponse) GetBody() *GetUploadAudioDataParamsResponseBody {
	return s.Body
}

func (s *GetUploadAudioDataParamsResponse) SetHeaders(v map[string]*string) *GetUploadAudioDataParamsResponse {
	s.Headers = v
	return s
}

func (s *GetUploadAudioDataParamsResponse) SetStatusCode(v int32) *GetUploadAudioDataParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadAudioDataParamsResponse) SetBody(v *GetUploadAudioDataParamsResponseBody) *GetUploadAudioDataParamsResponse {
	s.Body = v
	return s
}

func (s *GetUploadAudioDataParamsResponse) Validate() error {
	return dara.Validate(s)
}
