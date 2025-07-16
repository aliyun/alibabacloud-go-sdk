// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCacheByCacheTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshObjectCacheByCacheTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshObjectCacheByCacheTagResponse
	GetStatusCode() *int32
	SetBody(v *RefreshObjectCacheByCacheTagResponseBody) *RefreshObjectCacheByCacheTagResponse
	GetBody() *RefreshObjectCacheByCacheTagResponseBody
}

type RefreshObjectCacheByCacheTagResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshObjectCacheByCacheTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshObjectCacheByCacheTagResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCacheByCacheTagResponse) GoString() string {
	return s.String()
}

func (s *RefreshObjectCacheByCacheTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshObjectCacheByCacheTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshObjectCacheByCacheTagResponse) GetBody() *RefreshObjectCacheByCacheTagResponseBody {
	return s.Body
}

func (s *RefreshObjectCacheByCacheTagResponse) SetHeaders(v map[string]*string) *RefreshObjectCacheByCacheTagResponse {
	s.Headers = v
	return s
}

func (s *RefreshObjectCacheByCacheTagResponse) SetStatusCode(v int32) *RefreshObjectCacheByCacheTagResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshObjectCacheByCacheTagResponse) SetBody(v *RefreshObjectCacheByCacheTagResponseBody) *RefreshObjectCacheByCacheTagResponse {
	s.Body = v
	return s
}

func (s *RefreshObjectCacheByCacheTagResponse) Validate() error {
	return dara.Validate(s)
}
