// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBIpWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsDBIpWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsDBIpWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsDBIpWhiteListResponseBody) *DescribeDrdsDBIpWhiteListResponse
	GetBody() *DescribeDrdsDBIpWhiteListResponseBody
}

type DescribeDrdsDBIpWhiteListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsDBIpWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsDBIpWhiteListResponse) GetBody() *DescribeDrdsDBIpWhiteListResponseBody {
	return s.Body
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetStatusCode(v int32) *DescribeDrdsDBIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponse) SetBody(v *DescribeDrdsDBIpWhiteListResponseBody) *DescribeDrdsDBIpWhiteListResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
