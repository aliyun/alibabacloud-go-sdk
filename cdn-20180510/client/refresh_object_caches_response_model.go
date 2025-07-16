// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshObjectCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshObjectCachesResponse
	GetStatusCode() *int32
	SetBody(v *RefreshObjectCachesResponseBody) *RefreshObjectCachesResponse
	GetBody() *RefreshObjectCachesResponseBody
}

type RefreshObjectCachesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshObjectCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshObjectCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshObjectCachesResponse) GetBody() *RefreshObjectCachesResponseBody {
	return s.Body
}

func (s *RefreshObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshObjectCachesResponse) SetStatusCode(v int32) *RefreshObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshObjectCachesResponse) SetBody(v *RefreshObjectCachesResponseBody) *RefreshObjectCachesResponse {
	s.Body = v
	return s
}

func (s *RefreshObjectCachesResponse) Validate() error {
	return dara.Validate(s)
}
