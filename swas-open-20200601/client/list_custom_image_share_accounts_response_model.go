// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomImageShareAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomImageShareAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomImageShareAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomImageShareAccountsResponseBody) *ListCustomImageShareAccountsResponse
	GetBody() *ListCustomImageShareAccountsResponseBody
}

type ListCustomImageShareAccountsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomImageShareAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomImageShareAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImageShareAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomImageShareAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomImageShareAccountsResponse) GetBody() *ListCustomImageShareAccountsResponseBody {
	return s.Body
}

func (s *ListCustomImageShareAccountsResponse) SetHeaders(v map[string]*string) *ListCustomImageShareAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImageShareAccountsResponse) SetStatusCode(v int32) *ListCustomImageShareAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomImageShareAccountsResponse) SetBody(v *ListCustomImageShareAccountsResponseBody) *ListCustomImageShareAccountsResponse {
	s.Body = v
	return s
}

func (s *ListCustomImageShareAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
