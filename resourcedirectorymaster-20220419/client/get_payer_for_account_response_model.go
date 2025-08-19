// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPayerForAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPayerForAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPayerForAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetPayerForAccountResponseBody) *GetPayerForAccountResponse
	GetBody() *GetPayerForAccountResponseBody
}

type GetPayerForAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPayerForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPayerForAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPayerForAccountResponse) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPayerForAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPayerForAccountResponse) GetBody() *GetPayerForAccountResponseBody {
	return s.Body
}

func (s *GetPayerForAccountResponse) SetHeaders(v map[string]*string) *GetPayerForAccountResponse {
	s.Headers = v
	return s
}

func (s *GetPayerForAccountResponse) SetStatusCode(v int32) *GetPayerForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPayerForAccountResponse) SetBody(v *GetPayerForAccountResponseBody) *GetPayerForAccountResponse {
	s.Body = v
	return s
}

func (s *GetPayerForAccountResponse) Validate() error {
	return dara.Validate(s)
}
