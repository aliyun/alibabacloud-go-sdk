// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDcdnObjectCacheByCacheTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshDcdnObjectCacheByCacheTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshDcdnObjectCacheByCacheTagResponse
	GetStatusCode() *int32
	SetBody(v *RefreshDcdnObjectCacheByCacheTagResponseBody) *RefreshDcdnObjectCacheByCacheTagResponse
	GetBody() *RefreshDcdnObjectCacheByCacheTagResponseBody
}

type RefreshDcdnObjectCacheByCacheTagResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDcdnObjectCacheByCacheTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDcdnObjectCacheByCacheTagResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshDcdnObjectCacheByCacheTagResponse) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCacheByCacheTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshDcdnObjectCacheByCacheTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshDcdnObjectCacheByCacheTagResponse) GetBody() *RefreshDcdnObjectCacheByCacheTagResponseBody {
	return s.Body
}

func (s *RefreshDcdnObjectCacheByCacheTagResponse) SetHeaders(v map[string]*string) *RefreshDcdnObjectCacheByCacheTagResponse {
	s.Headers = v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagResponse) SetStatusCode(v int32) *RefreshDcdnObjectCacheByCacheTagResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagResponse) SetBody(v *RefreshDcdnObjectCacheByCacheTagResponseBody) *RefreshDcdnObjectCacheByCacheTagResponse {
	s.Body = v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagResponse) Validate() error {
	return dara.Validate(s)
}
