// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConnectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConnectionResponseBody) *UpdateConnectionResponse
	GetBody() *UpdateConnectionResponseBody
}

type UpdateConnectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConnectionResponse) GetBody() *UpdateConnectionResponseBody {
	return s.Body
}

func (s *UpdateConnectionResponse) SetHeaders(v map[string]*string) *UpdateConnectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnectionResponse) SetStatusCode(v int32) *UpdateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnectionResponse) SetBody(v *UpdateConnectionResponseBody) *UpdateConnectionResponse {
	s.Body = v
	return s
}

func (s *UpdateConnectionResponse) Validate() error {
	return dara.Validate(s)
}
