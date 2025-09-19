// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshRegistryTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshRegistryTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshRegistryTokenResponse
	GetStatusCode() *int32
	SetBody(v *RefreshRegistryTokenResponseBody) *RefreshRegistryTokenResponse
	GetBody() *RefreshRegistryTokenResponseBody
}

type RefreshRegistryTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshRegistryTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshRegistryTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshRegistryTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshRegistryTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshRegistryTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshRegistryTokenResponse) GetBody() *RefreshRegistryTokenResponseBody {
	return s.Body
}

func (s *RefreshRegistryTokenResponse) SetHeaders(v map[string]*string) *RefreshRegistryTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshRegistryTokenResponse) SetStatusCode(v int32) *RefreshRegistryTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshRegistryTokenResponse) SetBody(v *RefreshRegistryTokenResponseBody) *RefreshRegistryTokenResponse {
	s.Body = v
	return s
}

func (s *RefreshRegistryTokenResponse) Validate() error {
	return dara.Validate(s)
}
