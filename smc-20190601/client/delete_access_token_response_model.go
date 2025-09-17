// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessTokenResponseBody) *DeleteAccessTokenResponse
	GetBody() *DeleteAccessTokenResponseBody
}

type DeleteAccessTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessTokenResponse) GetBody() *DeleteAccessTokenResponseBody {
	return s.Body
}

func (s *DeleteAccessTokenResponse) SetHeaders(v map[string]*string) *DeleteAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessTokenResponse) SetStatusCode(v int32) *DeleteAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessTokenResponse) SetBody(v *DeleteAccessTokenResponseBody) *DeleteAccessTokenResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessTokenResponse) Validate() error {
	return dara.Validate(s)
}
