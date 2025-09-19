// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupedSecurityEventMarkMissListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGroupedSecurityEventMarkMissListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGroupedSecurityEventMarkMissListResponse
	GetStatusCode() *int32
	SetBody(v *QueryGroupedSecurityEventMarkMissListResponseBody) *QueryGroupedSecurityEventMarkMissListResponse
	GetBody() *QueryGroupedSecurityEventMarkMissListResponseBody
}

type QueryGroupedSecurityEventMarkMissListResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupedSecurityEventMarkMissListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupedSecurityEventMarkMissListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupedSecurityEventMarkMissListResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupedSecurityEventMarkMissListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGroupedSecurityEventMarkMissListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGroupedSecurityEventMarkMissListResponse) GetBody() *QueryGroupedSecurityEventMarkMissListResponseBody {
	return s.Body
}

func (s *QueryGroupedSecurityEventMarkMissListResponse) SetHeaders(v map[string]*string) *QueryGroupedSecurityEventMarkMissListResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponse) SetStatusCode(v int32) *QueryGroupedSecurityEventMarkMissListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponse) SetBody(v *QueryGroupedSecurityEventMarkMissListResponseBody) *QueryGroupedSecurityEventMarkMissListResponse {
	s.Body = v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponse) Validate() error {
	return dara.Validate(s)
}
