// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPageAdAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPageAdAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPageAdAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListPageAdAccountResponseBody) *ListPageAdAccountResponse
	GetBody() *ListPageAdAccountResponseBody
}

type ListPageAdAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPageAdAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPageAdAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPageAdAccountResponse) GoString() string {
	return s.String()
}

func (s *ListPageAdAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPageAdAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPageAdAccountResponse) GetBody() *ListPageAdAccountResponseBody {
	return s.Body
}

func (s *ListPageAdAccountResponse) SetHeaders(v map[string]*string) *ListPageAdAccountResponse {
	s.Headers = v
	return s
}

func (s *ListPageAdAccountResponse) SetStatusCode(v int32) *ListPageAdAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPageAdAccountResponse) SetBody(v *ListPageAdAccountResponseBody) *ListPageAdAccountResponse {
	s.Body = v
	return s
}

func (s *ListPageAdAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
