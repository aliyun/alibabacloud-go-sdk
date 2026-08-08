// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrossAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrossAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListCrossAccountsResponseBody) *ListCrossAccountsResponse
	GetBody() *ListCrossAccountsResponseBody
}

type ListCrossAccountsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrossAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrossAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrossAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListCrossAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrossAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrossAccountsResponse) GetBody() *ListCrossAccountsResponseBody {
	return s.Body
}

func (s *ListCrossAccountsResponse) SetHeaders(v map[string]*string) *ListCrossAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListCrossAccountsResponse) SetStatusCode(v int32) *ListCrossAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrossAccountsResponse) SetBody(v *ListCrossAccountsResponseBody) *ListCrossAccountsResponse {
	s.Body = v
	return s
}

func (s *ListCrossAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
