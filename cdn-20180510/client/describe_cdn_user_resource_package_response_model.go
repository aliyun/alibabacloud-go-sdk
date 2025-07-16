// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnUserResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnUserResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnUserResourcePackageResponseBody) *DescribeCdnUserResourcePackageResponse
	GetBody() *DescribeCdnUserResourcePackageResponseBody
}

type DescribeCdnUserResourcePackageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnUserResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnUserResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnUserResourcePackageResponse) GetBody() *DescribeCdnUserResourcePackageResponseBody {
	return s.Body
}

func (s *DescribeCdnUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeCdnUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserResourcePackageResponse) SetStatusCode(v int32) *DescribeCdnUserResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponse) SetBody(v *DescribeCdnUserResourcePackageResponseBody) *DescribeCdnUserResourcePackageResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnUserResourcePackageResponse) Validate() error {
	return dara.Validate(s)
}
