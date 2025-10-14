// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserCommandResponse
	GetStatusCode() *int32
	SetBody(v *GetUserCommandResponseBody) *GetUserCommandResponse
	GetBody() *GetUserCommandResponseBody
}

type GetUserCommandResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserCommandResponse) GoString() string {
	return s.String()
}

func (s *GetUserCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserCommandResponse) GetBody() *GetUserCommandResponseBody {
	return s.Body
}

func (s *GetUserCommandResponse) SetHeaders(v map[string]*string) *GetUserCommandResponse {
	s.Headers = v
	return s
}

func (s *GetUserCommandResponse) SetStatusCode(v int32) *GetUserCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserCommandResponse) SetBody(v *GetUserCommandResponseBody) *GetUserCommandResponse {
	s.Body = v
	return s
}

func (s *GetUserCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
