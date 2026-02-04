// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshErObjectCachesByCacheTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshErObjectCachesByCacheTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshErObjectCachesByCacheTagResponse
	GetStatusCode() *int32
	SetBody(v *RefreshErObjectCachesByCacheTagResponseBody) *RefreshErObjectCachesByCacheTagResponse
	GetBody() *RefreshErObjectCachesByCacheTagResponseBody
}

type RefreshErObjectCachesByCacheTagResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshErObjectCachesByCacheTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshErObjectCachesByCacheTagResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshErObjectCachesByCacheTagResponse) GoString() string {
	return s.String()
}

func (s *RefreshErObjectCachesByCacheTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshErObjectCachesByCacheTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshErObjectCachesByCacheTagResponse) GetBody() *RefreshErObjectCachesByCacheTagResponseBody {
	return s.Body
}

func (s *RefreshErObjectCachesByCacheTagResponse) SetHeaders(v map[string]*string) *RefreshErObjectCachesByCacheTagResponse {
	s.Headers = v
	return s
}

func (s *RefreshErObjectCachesByCacheTagResponse) SetStatusCode(v int32) *RefreshErObjectCachesByCacheTagResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshErObjectCachesByCacheTagResponse) SetBody(v *RefreshErObjectCachesByCacheTagResponseBody) *RefreshErObjectCachesByCacheTagResponse {
	s.Body = v
	return s
}

func (s *RefreshErObjectCachesByCacheTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
