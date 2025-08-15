// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceAccountResponseBody) *ListInstanceAccountResponse
	GetBody() *ListInstanceAccountResponseBody
}

type ListInstanceAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceAccountResponse) GetBody() *ListInstanceAccountResponseBody {
	return s.Body
}

func (s *ListInstanceAccountResponse) SetHeaders(v map[string]*string) *ListInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceAccountResponse) SetStatusCode(v int32) *ListInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceAccountResponse) SetBody(v *ListInstanceAccountResponseBody) *ListInstanceAccountResponse {
	s.Body = v
	return s
}

func (s *ListInstanceAccountResponse) Validate() error {
	return dara.Validate(s)
}
