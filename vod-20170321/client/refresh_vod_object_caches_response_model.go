// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshVodObjectCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshVodObjectCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshVodObjectCachesResponse
	GetStatusCode() *int32
	SetBody(v *RefreshVodObjectCachesResponseBody) *RefreshVodObjectCachesResponse
	GetBody() *RefreshVodObjectCachesResponseBody
}

type RefreshVodObjectCachesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshVodObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshVodObjectCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshVodObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshVodObjectCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshVodObjectCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshVodObjectCachesResponse) GetBody() *RefreshVodObjectCachesResponseBody {
	return s.Body
}

func (s *RefreshVodObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshVodObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshVodObjectCachesResponse) SetStatusCode(v int32) *RefreshVodObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshVodObjectCachesResponse) SetBody(v *RefreshVodObjectCachesResponseBody) *RefreshVodObjectCachesResponse {
	s.Body = v
	return s
}

func (s *RefreshVodObjectCachesResponse) Validate() error {
	return dara.Validate(s)
}
