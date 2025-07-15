// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAudioFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAudioFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAudioFileResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAudioFileResponseBody) *ModifyAudioFileResponse
	GetBody() *ModifyAudioFileResponseBody
}

type ModifyAudioFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAudioFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAudioFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAudioFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyAudioFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAudioFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAudioFileResponse) GetBody() *ModifyAudioFileResponseBody {
	return s.Body
}

func (s *ModifyAudioFileResponse) SetHeaders(v map[string]*string) *ModifyAudioFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyAudioFileResponse) SetStatusCode(v int32) *ModifyAudioFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAudioFileResponse) SetBody(v *ModifyAudioFileResponseBody) *ModifyAudioFileResponse {
	s.Body = v
	return s
}

func (s *ModifyAudioFileResponse) Validate() error {
	return dara.Validate(s)
}
