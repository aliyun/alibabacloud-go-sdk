// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyBastionAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyBastionAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyBastionAccountResponse
	GetStatusCode() *int32
	SetBody(v *ApplyBastionAccountResponseBody) *ApplyBastionAccountResponse
	GetBody() *ApplyBastionAccountResponseBody
}

type ApplyBastionAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyBastionAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyBastionAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyBastionAccountResponse) GoString() string {
	return s.String()
}

func (s *ApplyBastionAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyBastionAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyBastionAccountResponse) GetBody() *ApplyBastionAccountResponseBody {
	return s.Body
}

func (s *ApplyBastionAccountResponse) SetHeaders(v map[string]*string) *ApplyBastionAccountResponse {
	s.Headers = v
	return s
}

func (s *ApplyBastionAccountResponse) SetStatusCode(v int32) *ApplyBastionAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyBastionAccountResponse) SetBody(v *ApplyBastionAccountResponseBody) *ApplyBastionAccountResponse {
	s.Body = v
	return s
}

func (s *ApplyBastionAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
