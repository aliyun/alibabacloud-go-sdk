// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunRCCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunRCCommandResponse
	GetStatusCode() *int32
	SetBody(v *RunRCCommandResponseBody) *RunRCCommandResponse
	GetBody() *RunRCCommandResponseBody
}

type RunRCCommandResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunRCCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunRCCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s RunRCCommandResponse) GoString() string {
	return s.String()
}

func (s *RunRCCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunRCCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunRCCommandResponse) GetBody() *RunRCCommandResponseBody {
	return s.Body
}

func (s *RunRCCommandResponse) SetHeaders(v map[string]*string) *RunRCCommandResponse {
	s.Headers = v
	return s
}

func (s *RunRCCommandResponse) SetStatusCode(v int32) *RunRCCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunRCCommandResponse) SetBody(v *RunRCCommandResponseBody) *RunRCCommandResponse {
	s.Body = v
	return s
}

func (s *RunRCCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
