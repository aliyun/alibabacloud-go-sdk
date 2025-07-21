// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSuppressionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserSuppressionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserSuppressionResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserSuppressionResponseBody) *CreateUserSuppressionResponse
	GetBody() *CreateUserSuppressionResponseBody
}

type CreateUserSuppressionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserSuppressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserSuppressionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSuppressionResponse) GoString() string {
	return s.String()
}

func (s *CreateUserSuppressionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserSuppressionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserSuppressionResponse) GetBody() *CreateUserSuppressionResponseBody {
	return s.Body
}

func (s *CreateUserSuppressionResponse) SetHeaders(v map[string]*string) *CreateUserSuppressionResponse {
	s.Headers = v
	return s
}

func (s *CreateUserSuppressionResponse) SetStatusCode(v int32) *CreateUserSuppressionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserSuppressionResponse) SetBody(v *CreateUserSuppressionResponseBody) *CreateUserSuppressionResponse {
	s.Body = v
	return s
}

func (s *CreateUserSuppressionResponse) Validate() error {
	return dara.Validate(s)
}
