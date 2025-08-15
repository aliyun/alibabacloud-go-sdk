// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceAccountResponseBody) *CreateInstanceAccountResponse
	GetBody() *CreateInstanceAccountResponseBody
}

type CreateInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceAccountResponse) GetBody() *CreateInstanceAccountResponseBody {
	return s.Body
}

func (s *CreateInstanceAccountResponse) SetHeaders(v map[string]*string) *CreateInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceAccountResponse) SetStatusCode(v int32) *CreateInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceAccountResponse) SetBody(v *CreateInstanceAccountResponseBody) *CreateInstanceAccountResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceAccountResponse) Validate() error {
	return dara.Validate(s)
}
