// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersForRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClustersForRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClustersForRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClustersForRegionResponseBody) *DescribeClustersForRegionResponse
	GetBody() *DescribeClustersForRegionResponseBody
}

type DescribeClustersForRegionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClustersForRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClustersForRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersForRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersForRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClustersForRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClustersForRegionResponse) GetBody() *DescribeClustersForRegionResponseBody {
	return s.Body
}

func (s *DescribeClustersForRegionResponse) SetHeaders(v map[string]*string) *DescribeClustersForRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersForRegionResponse) SetStatusCode(v int32) *DescribeClustersForRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClustersForRegionResponse) SetBody(v *DescribeClustersForRegionResponseBody) *DescribeClustersForRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeClustersForRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
