// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBandwidthPackageResponseBody) *DescribeBandwidthPackageResponse
	GetBody() *DescribeBandwidthPackageResponseBody
}

type DescribeBandwidthPackageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBandwidthPackageResponse) GetBody() *DescribeBandwidthPackageResponseBody {
	return s.Body
}

func (s *DescribeBandwidthPackageResponse) SetHeaders(v map[string]*string) *DescribeBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwidthPackageResponse) SetStatusCode(v int32) *DescribeBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwidthPackageResponse) SetBody(v *DescribeBandwidthPackageResponseBody) *DescribeBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *DescribeBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
