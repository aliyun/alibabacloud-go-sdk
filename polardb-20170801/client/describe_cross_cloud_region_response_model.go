// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossCloudRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossCloudRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossCloudRegionResponseBody) *DescribeCrossCloudRegionResponse
	GetBody() *DescribeCrossCloudRegionResponseBody
}

type DescribeCrossCloudRegionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossCloudRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossCloudRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossCloudRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossCloudRegionResponse) GetBody() *DescribeCrossCloudRegionResponseBody {
	return s.Body
}

func (s *DescribeCrossCloudRegionResponse) SetHeaders(v map[string]*string) *DescribeCrossCloudRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossCloudRegionResponse) SetStatusCode(v int32) *DescribeCrossCloudRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossCloudRegionResponse) SetBody(v *DescribeCrossCloudRegionResponseBody) *DescribeCrossCloudRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossCloudRegionResponse) Validate() error {
	return dara.Validate(s)
}
