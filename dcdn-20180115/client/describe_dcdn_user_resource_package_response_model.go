// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserResourcePackageResponseBody) *DescribeDcdnUserResourcePackageResponse
	GetBody() *DescribeDcdnUserResourcePackageResponseBody
}

type DescribeDcdnUserResourcePackageResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserResourcePackageResponse) GetBody() *DescribeDcdnUserResourcePackageResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponse) SetStatusCode(v int32) *DescribeDcdnUserResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponse) SetBody(v *DescribeDcdnUserResourcePackageResponseBody) *DescribeDcdnUserResourcePackageResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponse) Validate() error {
	return dara.Validate(s)
}
