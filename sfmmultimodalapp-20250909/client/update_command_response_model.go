// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCommandResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCommandResponseBody) *UpdateCommandResponse
	GetBody() *UpdateCommandResponseBody
}

type UpdateCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCommandResponse) GetBody() *UpdateCommandResponseBody {
	return s.Body
}

func (s *UpdateCommandResponse) SetHeaders(v map[string]*string) *UpdateCommandResponse {
	s.Headers = v
	return s
}

func (s *UpdateCommandResponse) SetStatusCode(v int32) *UpdateCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCommandResponse) SetBody(v *UpdateCommandResponseBody) *UpdateCommandResponse {
	s.Body = v
	return s
}

func (s *UpdateCommandResponse) Validate() error {
	return dara.Validate(s)
}
