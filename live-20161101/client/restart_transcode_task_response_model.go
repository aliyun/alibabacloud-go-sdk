// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartTranscodeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartTranscodeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartTranscodeTaskResponse
	GetStatusCode() *int32
	SetBody(v *RestartTranscodeTaskResponseBody) *RestartTranscodeTaskResponse
	GetBody() *RestartTranscodeTaskResponseBody
}

type RestartTranscodeTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartTranscodeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartTranscodeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartTranscodeTaskResponse) GoString() string {
	return s.String()
}

func (s *RestartTranscodeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartTranscodeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartTranscodeTaskResponse) GetBody() *RestartTranscodeTaskResponseBody {
	return s.Body
}

func (s *RestartTranscodeTaskResponse) SetHeaders(v map[string]*string) *RestartTranscodeTaskResponse {
	s.Headers = v
	return s
}

func (s *RestartTranscodeTaskResponse) SetStatusCode(v int32) *RestartTranscodeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartTranscodeTaskResponse) SetBody(v *RestartTranscodeTaskResponseBody) *RestartTranscodeTaskResponse {
	s.Body = v
	return s
}

func (s *RestartTranscodeTaskResponse) Validate() error {
	return dara.Validate(s)
}
