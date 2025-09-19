// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaliciousFileWhitelistConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMaliciousFileWhitelistConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMaliciousFileWhitelistConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListMaliciousFileWhitelistConfigsResponseBody) *ListMaliciousFileWhitelistConfigsResponse
	GetBody() *ListMaliciousFileWhitelistConfigsResponseBody
}

type ListMaliciousFileWhitelistConfigsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMaliciousFileWhitelistConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMaliciousFileWhitelistConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMaliciousFileWhitelistConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListMaliciousFileWhitelistConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMaliciousFileWhitelistConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMaliciousFileWhitelistConfigsResponse) GetBody() *ListMaliciousFileWhitelistConfigsResponseBody {
	return s.Body
}

func (s *ListMaliciousFileWhitelistConfigsResponse) SetHeaders(v map[string]*string) *ListMaliciousFileWhitelistConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponse) SetStatusCode(v int32) *ListMaliciousFileWhitelistConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponse) SetBody(v *ListMaliciousFileWhitelistConfigsResponseBody) *ListMaliciousFileWhitelistConfigsResponse {
	s.Body = v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponse) Validate() error {
	return dara.Validate(s)
}
