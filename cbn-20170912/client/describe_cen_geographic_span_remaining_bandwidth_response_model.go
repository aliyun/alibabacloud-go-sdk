// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenGeographicSpanRemainingBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenGeographicSpanRemainingBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenGeographicSpanRemainingBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenGeographicSpanRemainingBandwidthResponseBody) *DescribeCenGeographicSpanRemainingBandwidthResponse
	GetBody() *DescribeCenGeographicSpanRemainingBandwidthResponseBody
}

type DescribeCenGeographicSpanRemainingBandwidthResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenGeographicSpanRemainingBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpanRemainingBandwidthResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) GetBody() *DescribeCenGeographicSpanRemainingBandwidthResponseBody {
	return s.Body
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) SetHeaders(v map[string]*string) *DescribeCenGeographicSpanRemainingBandwidthResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) SetStatusCode(v int32) *DescribeCenGeographicSpanRemainingBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) SetBody(v *DescribeCenGeographicSpanRemainingBandwidthResponseBody) *DescribeCenGeographicSpanRemainingBandwidthResponse {
	s.Body = v
	return s
}

func (s *DescribeCenGeographicSpanRemainingBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
