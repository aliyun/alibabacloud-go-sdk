// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceSharedAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceSharedAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceSharedAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceSharedAccountsResponseBody) *ListServiceSharedAccountsResponse
	GetBody() *ListServiceSharedAccountsResponseBody
}

type ListServiceSharedAccountsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceSharedAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceSharedAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceSharedAccountsResponse) GetBody() *ListServiceSharedAccountsResponseBody {
	return s.Body
}

func (s *ListServiceSharedAccountsResponse) SetHeaders(v map[string]*string) *ListServiceSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceSharedAccountsResponse) SetStatusCode(v int32) *ListServiceSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceSharedAccountsResponse) SetBody(v *ListServiceSharedAccountsResponseBody) *ListServiceSharedAccountsResponse {
	s.Body = v
	return s
}

func (s *ListServiceSharedAccountsResponse) Validate() error {
	return dara.Validate(s)
}
