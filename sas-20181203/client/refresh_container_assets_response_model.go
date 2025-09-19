// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshContainerAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshContainerAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshContainerAssetsResponse
	GetStatusCode() *int32
	SetBody(v *RefreshContainerAssetsResponseBody) *RefreshContainerAssetsResponse
	GetBody() *RefreshContainerAssetsResponseBody
}

type RefreshContainerAssetsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshContainerAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshContainerAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshContainerAssetsResponse) GoString() string {
	return s.String()
}

func (s *RefreshContainerAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshContainerAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshContainerAssetsResponse) GetBody() *RefreshContainerAssetsResponseBody {
	return s.Body
}

func (s *RefreshContainerAssetsResponse) SetHeaders(v map[string]*string) *RefreshContainerAssetsResponse {
	s.Headers = v
	return s
}

func (s *RefreshContainerAssetsResponse) SetStatusCode(v int32) *RefreshContainerAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshContainerAssetsResponse) SetBody(v *RefreshContainerAssetsResponseBody) *RefreshContainerAssetsResponse {
	s.Body = v
	return s
}

func (s *RefreshContainerAssetsResponse) Validate() error {
	return dara.Validate(s)
}
