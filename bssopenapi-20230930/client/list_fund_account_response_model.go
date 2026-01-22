// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFundAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFundAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFundAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListFundAccountResponseBody) *ListFundAccountResponse
	GetBody() *ListFundAccountResponseBody
}

type ListFundAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFundAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFundAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountResponse) GoString() string {
	return s.String()
}

func (s *ListFundAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFundAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFundAccountResponse) GetBody() *ListFundAccountResponseBody {
	return s.Body
}

func (s *ListFundAccountResponse) SetHeaders(v map[string]*string) *ListFundAccountResponse {
	s.Headers = v
	return s
}

func (s *ListFundAccountResponse) SetStatusCode(v int32) *ListFundAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFundAccountResponse) SetBody(v *ListFundAccountResponseBody) *ListFundAccountResponse {
	s.Body = v
	return s
}

func (s *ListFundAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
