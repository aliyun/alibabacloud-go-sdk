// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceIpWhitelistResponseBody) *ListInstanceIpWhitelistResponse
	GetBody() *ListInstanceIpWhitelistResponseBody
}

type ListInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceIpWhitelistResponse) GetBody() *ListInstanceIpWhitelistResponseBody {
	return s.Body
}

func (s *ListInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *ListInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceIpWhitelistResponse) SetStatusCode(v int32) *ListInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceIpWhitelistResponse) SetBody(v *ListInstanceIpWhitelistResponseBody) *ListInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *ListInstanceIpWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
