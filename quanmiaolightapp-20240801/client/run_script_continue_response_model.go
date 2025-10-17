// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptContinueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunScriptContinueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunScriptContinueResponse
	GetStatusCode() *int32
	SetBody(v *RunScriptContinueResponseBody) *RunScriptContinueResponse
	GetBody() *RunScriptContinueResponseBody
}

type RunScriptContinueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunScriptContinueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptContinueResponse) String() string {
	return dara.Prettify(s)
}

func (s RunScriptContinueResponse) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunScriptContinueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunScriptContinueResponse) GetBody() *RunScriptContinueResponseBody {
	return s.Body
}

func (s *RunScriptContinueResponse) SetHeaders(v map[string]*string) *RunScriptContinueResponse {
	s.Headers = v
	return s
}

func (s *RunScriptContinueResponse) SetStatusCode(v int32) *RunScriptContinueResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptContinueResponse) SetBody(v *RunScriptContinueResponseBody) *RunScriptContinueResponse {
	s.Body = v
	return s
}

func (s *RunScriptContinueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
