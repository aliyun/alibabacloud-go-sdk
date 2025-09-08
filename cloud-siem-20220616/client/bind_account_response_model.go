// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAccountResponse
	GetStatusCode() *int32
	SetBody(v *BindAccountResponseBody) *BindAccountResponse
	GetBody() *BindAccountResponseBody
}

type BindAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAccountResponse) GoString() string {
	return s.String()
}

func (s *BindAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAccountResponse) GetBody() *BindAccountResponseBody {
	return s.Body
}

func (s *BindAccountResponse) SetHeaders(v map[string]*string) *BindAccountResponse {
	s.Headers = v
	return s
}

func (s *BindAccountResponse) SetStatusCode(v int32) *BindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAccountResponse) SetBody(v *BindAccountResponseBody) *BindAccountResponse {
	s.Body = v
	return s
}

func (s *BindAccountResponse) Validate() error {
	return dara.Validate(s)
}
