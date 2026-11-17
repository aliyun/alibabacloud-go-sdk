// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAppInstanceCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunAppInstanceCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunAppInstanceCommandResponse
	GetStatusCode() *int32
	SetBody(v *RunAppInstanceCommandResponseBody) *RunAppInstanceCommandResponse
	GetBody() *RunAppInstanceCommandResponseBody
}

type RunAppInstanceCommandResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunAppInstanceCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunAppInstanceCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s RunAppInstanceCommandResponse) GoString() string {
	return s.String()
}

func (s *RunAppInstanceCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunAppInstanceCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunAppInstanceCommandResponse) GetBody() *RunAppInstanceCommandResponseBody {
	return s.Body
}

func (s *RunAppInstanceCommandResponse) SetHeaders(v map[string]*string) *RunAppInstanceCommandResponse {
	s.Headers = v
	return s
}

func (s *RunAppInstanceCommandResponse) SetStatusCode(v int32) *RunAppInstanceCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunAppInstanceCommandResponse) SetBody(v *RunAppInstanceCommandResponseBody) *RunAppInstanceCommandResponse {
	s.Body = v
	return s
}

func (s *RunAppInstanceCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
