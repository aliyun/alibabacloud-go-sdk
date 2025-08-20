// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentCommodityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentCommodityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentCommodityResponse
	GetStatusCode() *int32
	SetBody(v *SegmentCommodityResponseBody) *SegmentCommodityResponse
	GetBody() *SegmentCommodityResponseBody
}

type SegmentCommodityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentCommodityResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommodityResponse) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentCommodityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentCommodityResponse) GetBody() *SegmentCommodityResponseBody {
	return s.Body
}

func (s *SegmentCommodityResponse) SetHeaders(v map[string]*string) *SegmentCommodityResponse {
	s.Headers = v
	return s
}

func (s *SegmentCommodityResponse) SetStatusCode(v int32) *SegmentCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentCommodityResponse) SetBody(v *SegmentCommodityResponseBody) *SegmentCommodityResponse {
	s.Body = v
	return s
}

func (s *SegmentCommodityResponse) Validate() error {
	return dara.Validate(s)
}
