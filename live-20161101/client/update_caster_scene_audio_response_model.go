// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterSceneAudioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCasterSceneAudioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCasterSceneAudioResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCasterSceneAudioResponseBody) *UpdateCasterSceneAudioResponse
	GetBody() *UpdateCasterSceneAudioResponseBody
}

type UpdateCasterSceneAudioResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCasterSceneAudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCasterSceneAudioResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterSceneAudioResponse) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneAudioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCasterSceneAudioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCasterSceneAudioResponse) GetBody() *UpdateCasterSceneAudioResponseBody {
	return s.Body
}

func (s *UpdateCasterSceneAudioResponse) SetHeaders(v map[string]*string) *UpdateCasterSceneAudioResponse {
	s.Headers = v
	return s
}

func (s *UpdateCasterSceneAudioResponse) SetStatusCode(v int32) *UpdateCasterSceneAudioResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCasterSceneAudioResponse) SetBody(v *UpdateCasterSceneAudioResponseBody) *UpdateCasterSceneAudioResponse {
	s.Body = v
	return s
}

func (s *UpdateCasterSceneAudioResponse) Validate() error {
	return dara.Validate(s)
}
