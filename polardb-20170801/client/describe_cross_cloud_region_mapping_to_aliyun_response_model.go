// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudRegionMappingToAliyunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossCloudRegionMappingToAliyunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossCloudRegionMappingToAliyunResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossCloudRegionMappingToAliyunResponseBody) *DescribeCrossCloudRegionMappingToAliyunResponse
	GetBody() *DescribeCrossCloudRegionMappingToAliyunResponseBody
}

type DescribeCrossCloudRegionMappingToAliyunResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossCloudRegionMappingToAliyunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossCloudRegionMappingToAliyunResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionMappingToAliyunResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponse) GetBody() *DescribeCrossCloudRegionMappingToAliyunResponseBody {
	return s.Body
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponse) SetHeaders(v map[string]*string) *DescribeCrossCloudRegionMappingToAliyunResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponse) SetStatusCode(v int32) *DescribeCrossCloudRegionMappingToAliyunResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponse) SetBody(v *DescribeCrossCloudRegionMappingToAliyunResponseBody) *DescribeCrossCloudRegionMappingToAliyunResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossCloudRegionMappingToAliyunResponse) Validate() error {
	return dara.Validate(s)
}
