// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshWebofficeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshWebofficeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshWebofficeTokenResponse
	GetStatusCode() *int32
	SetBody(v *RefreshWebofficeTokenResponseBody) *RefreshWebofficeTokenResponse
	GetBody() *RefreshWebofficeTokenResponseBody
}

type RefreshWebofficeTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshWebofficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshWebofficeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshWebofficeTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshWebofficeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshWebofficeTokenResponse) GetBody() *RefreshWebofficeTokenResponseBody {
	return s.Body
}

func (s *RefreshWebofficeTokenResponse) SetHeaders(v map[string]*string) *RefreshWebofficeTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshWebofficeTokenResponse) SetStatusCode(v int32) *RefreshWebofficeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshWebofficeTokenResponse) SetBody(v *RefreshWebofficeTokenResponseBody) *RefreshWebofficeTokenResponse {
	s.Body = v
	return s
}

func (s *RefreshWebofficeTokenResponse) Validate() error {
	return dara.Validate(s)
}
