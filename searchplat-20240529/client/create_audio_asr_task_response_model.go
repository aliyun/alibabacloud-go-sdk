// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAudioAsrTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAudioAsrTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAudioAsrTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAudioAsrTaskResponseBody) *CreateAudioAsrTaskResponse
	GetBody() *CreateAudioAsrTaskResponseBody
}

type CreateAudioAsrTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAudioAsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAudioAsrTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioAsrTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAudioAsrTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAudioAsrTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAudioAsrTaskResponse) GetBody() *CreateAudioAsrTaskResponseBody {
	return s.Body
}

func (s *CreateAudioAsrTaskResponse) SetHeaders(v map[string]*string) *CreateAudioAsrTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAudioAsrTaskResponse) SetStatusCode(v int32) *CreateAudioAsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAudioAsrTaskResponse) SetBody(v *CreateAudioAsrTaskResponseBody) *CreateAudioAsrTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAudioAsrTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
