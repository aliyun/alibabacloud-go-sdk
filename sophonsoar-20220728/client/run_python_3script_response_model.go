// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPython3ScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunPython3ScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunPython3ScriptResponse
	GetStatusCode() *int32
	SetBody(v *RunPython3ScriptResponseBody) *RunPython3ScriptResponse
	GetBody() *RunPython3ScriptResponseBody
}

type RunPython3ScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPython3ScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPython3ScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s RunPython3ScriptResponse) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunPython3ScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPython3ScriptResponse) GetBody() *RunPython3ScriptResponseBody {
	return s.Body
}

func (s *RunPython3ScriptResponse) SetHeaders(v map[string]*string) *RunPython3ScriptResponse {
	s.Headers = v
	return s
}

func (s *RunPython3ScriptResponse) SetStatusCode(v int32) *RunPython3ScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPython3ScriptResponse) SetBody(v *RunPython3ScriptResponseBody) *RunPython3ScriptResponse {
	s.Body = v
	return s
}

func (s *RunPython3ScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
