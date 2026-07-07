// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAuthTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshAuthTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshAuthTokensResponse
	GetStatusCode() *int32
	SetBody(v *RefreshAuthTokensResponseBody) *RefreshAuthTokensResponse
	GetBody() *RefreshAuthTokensResponseBody
}

type RefreshAuthTokensResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAuthTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAuthTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshAuthTokensResponse) GoString() string {
	return s.String()
}

func (s *RefreshAuthTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshAuthTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshAuthTokensResponse) GetBody() *RefreshAuthTokensResponseBody {
	return s.Body
}

func (s *RefreshAuthTokensResponse) SetHeaders(v map[string]*string) *RefreshAuthTokensResponse {
	s.Headers = v
	return s
}

func (s *RefreshAuthTokensResponse) SetStatusCode(v int32) *RefreshAuthTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAuthTokensResponse) SetBody(v *RefreshAuthTokensResponseBody) *RefreshAuthTokensResponse {
	s.Body = v
	return s
}

func (s *RefreshAuthTokensResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
