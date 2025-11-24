// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceAccountsResponseBody) *ListServiceAccountsResponse
	GetBody() *ListServiceAccountsResponseBody
}

type ListServiceAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceAccountsResponse) GetBody() *ListServiceAccountsResponseBody {
	return s.Body
}

func (s *ListServiceAccountsResponse) SetHeaders(v map[string]*string) *ListServiceAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceAccountsResponse) SetStatusCode(v int32) *ListServiceAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceAccountsResponse) SetBody(v *ListServiceAccountsResponseBody) *ListServiceAccountsResponse {
	s.Body = v
	return s
}

func (s *ListServiceAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
