// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshErObjectCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshErObjectCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshErObjectCachesResponse
	GetStatusCode() *int32
	SetBody(v *RefreshErObjectCachesResponseBody) *RefreshErObjectCachesResponse
	GetBody() *RefreshErObjectCachesResponseBody
}

type RefreshErObjectCachesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshErObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshErObjectCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshErObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshErObjectCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshErObjectCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshErObjectCachesResponse) GetBody() *RefreshErObjectCachesResponseBody {
	return s.Body
}

func (s *RefreshErObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshErObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshErObjectCachesResponse) SetStatusCode(v int32) *RefreshErObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshErObjectCachesResponse) SetBody(v *RefreshErObjectCachesResponseBody) *RefreshErObjectCachesResponse {
	s.Body = v
	return s
}

func (s *RefreshErObjectCachesResponse) Validate() error {
	return dara.Validate(s)
}
