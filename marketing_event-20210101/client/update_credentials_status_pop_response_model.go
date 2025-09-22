// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialsStatusPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCredentialsStatusPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCredentialsStatusPopResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCredentialsStatusPopResponseBody) *UpdateCredentialsStatusPopResponse
	GetBody() *UpdateCredentialsStatusPopResponseBody
}

type UpdateCredentialsStatusPopResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCredentialsStatusPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialsStatusPopResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialsStatusPopResponse) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsStatusPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCredentialsStatusPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCredentialsStatusPopResponse) GetBody() *UpdateCredentialsStatusPopResponseBody {
	return s.Body
}

func (s *UpdateCredentialsStatusPopResponse) SetHeaders(v map[string]*string) *UpdateCredentialsStatusPopResponse {
	s.Headers = v
	return s
}

func (s *UpdateCredentialsStatusPopResponse) SetStatusCode(v int32) *UpdateCredentialsStatusPopResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCredentialsStatusPopResponse) SetBody(v *UpdateCredentialsStatusPopResponseBody) *UpdateCredentialsStatusPopResponse {
	s.Body = v
	return s
}

func (s *UpdateCredentialsStatusPopResponse) Validate() error {
	return dara.Validate(s)
}
