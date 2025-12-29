// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushVoiceBoxCommandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushVoiceBoxCommandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushVoiceBoxCommandsResponse
	GetStatusCode() *int32
	SetBody(v *PushVoiceBoxCommandsResponseBody) *PushVoiceBoxCommandsResponse
	GetBody() *PushVoiceBoxCommandsResponseBody
}

type PushVoiceBoxCommandsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushVoiceBoxCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushVoiceBoxCommandsResponse) String() string {
	return dara.Prettify(s)
}

func (s PushVoiceBoxCommandsResponse) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushVoiceBoxCommandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushVoiceBoxCommandsResponse) GetBody() *PushVoiceBoxCommandsResponseBody {
	return s.Body
}

func (s *PushVoiceBoxCommandsResponse) SetHeaders(v map[string]*string) *PushVoiceBoxCommandsResponse {
	s.Headers = v
	return s
}

func (s *PushVoiceBoxCommandsResponse) SetStatusCode(v int32) *PushVoiceBoxCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *PushVoiceBoxCommandsResponse) SetBody(v *PushVoiceBoxCommandsResponseBody) *PushVoiceBoxCommandsResponse {
	s.Body = v
	return s
}

func (s *PushVoiceBoxCommandsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
