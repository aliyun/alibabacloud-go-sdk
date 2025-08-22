// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDcdnObjectCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshDcdnObjectCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshDcdnObjectCachesResponse
	GetStatusCode() *int32
	SetBody(v *RefreshDcdnObjectCachesResponseBody) *RefreshDcdnObjectCachesResponse
	GetBody() *RefreshDcdnObjectCachesResponseBody
}

type RefreshDcdnObjectCachesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDcdnObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDcdnObjectCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshDcdnObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshDcdnObjectCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshDcdnObjectCachesResponse) GetBody() *RefreshDcdnObjectCachesResponseBody {
	return s.Body
}

func (s *RefreshDcdnObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshDcdnObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshDcdnObjectCachesResponse) SetStatusCode(v int32) *RefreshDcdnObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDcdnObjectCachesResponse) SetBody(v *RefreshDcdnObjectCachesResponseBody) *RefreshDcdnObjectCachesResponse {
	s.Body = v
	return s
}

func (s *RefreshDcdnObjectCachesResponse) Validate() error {
	return dara.Validate(s)
}
