// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileUploadParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAudioFileUploadParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAudioFileUploadParametersResponse
	GetStatusCode() *int32
	SetBody(v *GetAudioFileUploadParametersResponseBody) *GetAudioFileUploadParametersResponse
	GetBody() *GetAudioFileUploadParametersResponseBody
}

type GetAudioFileUploadParametersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAudioFileUploadParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAudioFileUploadParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileUploadParametersResponse) GoString() string {
	return s.String()
}

func (s *GetAudioFileUploadParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAudioFileUploadParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAudioFileUploadParametersResponse) GetBody() *GetAudioFileUploadParametersResponseBody {
	return s.Body
}

func (s *GetAudioFileUploadParametersResponse) SetHeaders(v map[string]*string) *GetAudioFileUploadParametersResponse {
	s.Headers = v
	return s
}

func (s *GetAudioFileUploadParametersResponse) SetStatusCode(v int32) *GetAudioFileUploadParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAudioFileUploadParametersResponse) SetBody(v *GetAudioFileUploadParametersResponseBody) *GetAudioFileUploadParametersResponse {
	s.Body = v
	return s
}

func (s *GetAudioFileUploadParametersResponse) Validate() error {
	return dara.Validate(s)
}
