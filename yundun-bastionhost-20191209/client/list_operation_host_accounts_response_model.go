// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHostAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationHostAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationHostAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationHostAccountsResponseBody) *ListOperationHostAccountsResponse
	GetBody() *ListOperationHostAccountsResponseBody
}

type ListOperationHostAccountsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationHostAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationHostAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationHostAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationHostAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationHostAccountsResponse) GetBody() *ListOperationHostAccountsResponseBody {
	return s.Body
}

func (s *ListOperationHostAccountsResponse) SetHeaders(v map[string]*string) *ListOperationHostAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationHostAccountsResponse) SetStatusCode(v int32) *ListOperationHostAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationHostAccountsResponse) SetBody(v *ListOperationHostAccountsResponseBody) *ListOperationHostAccountsResponse {
	s.Body = v
	return s
}

func (s *ListOperationHostAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
