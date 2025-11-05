// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnRegionAndIspResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnRegionAndIspResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnRegionAndIspResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnRegionAndIspResponseBody) *DescribeCdnRegionAndIspResponse
	GetBody() *DescribeCdnRegionAndIspResponseBody
}

type DescribeCdnRegionAndIspResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnRegionAndIspResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnRegionAndIspResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnRegionAndIspResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnRegionAndIspResponse) GetBody() *DescribeCdnRegionAndIspResponseBody {
	return s.Body
}

func (s *DescribeCdnRegionAndIspResponse) SetHeaders(v map[string]*string) *DescribeCdnRegionAndIspResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnRegionAndIspResponse) SetStatusCode(v int32) *DescribeCdnRegionAndIspResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponse) SetBody(v *DescribeCdnRegionAndIspResponseBody) *DescribeCdnRegionAndIspResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnRegionAndIspResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
