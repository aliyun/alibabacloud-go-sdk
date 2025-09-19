// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshAssetsResponse
	GetStatusCode() *int32
	SetBody(v *RefreshAssetsResponseBody) *RefreshAssetsResponse
	GetBody() *RefreshAssetsResponseBody
}

type RefreshAssetsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshAssetsResponse) GoString() string {
	return s.String()
}

func (s *RefreshAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshAssetsResponse) GetBody() *RefreshAssetsResponseBody {
	return s.Body
}

func (s *RefreshAssetsResponse) SetHeaders(v map[string]*string) *RefreshAssetsResponse {
	s.Headers = v
	return s
}

func (s *RefreshAssetsResponse) SetStatusCode(v int32) *RefreshAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAssetsResponse) SetBody(v *RefreshAssetsResponseBody) *RefreshAssetsResponse {
	s.Body = v
	return s
}

func (s *RefreshAssetsResponse) Validate() error {
	return dara.Validate(s)
}
