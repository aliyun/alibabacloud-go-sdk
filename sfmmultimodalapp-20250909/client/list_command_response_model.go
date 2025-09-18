// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommandResponse
	GetStatusCode() *int32
	SetBody(v *ListCommandResponseBody) *ListCommandResponse
	GetBody() *ListCommandResponseBody
}

type ListCommandResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommandResponse) GoString() string {
	return s.String()
}

func (s *ListCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommandResponse) GetBody() *ListCommandResponseBody {
	return s.Body
}

func (s *ListCommandResponse) SetHeaders(v map[string]*string) *ListCommandResponse {
	s.Headers = v
	return s
}

func (s *ListCommandResponse) SetStatusCode(v int32) *ListCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommandResponse) SetBody(v *ListCommandResponseBody) *ListCommandResponse {
	s.Body = v
	return s
}

func (s *ListCommandResponse) Validate() error {
	return dara.Validate(s)
}
