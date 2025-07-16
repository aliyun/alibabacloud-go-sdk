// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDeliverListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDeliverListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDeliverListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDeliverListResponseBody) *DescribeCdnDeliverListResponse
	GetBody() *DescribeCdnDeliverListResponseBody
}

type DescribeCdnDeliverListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDeliverListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDeliverListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDeliverListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDeliverListResponse) GetBody() *DescribeCdnDeliverListResponseBody {
	return s.Body
}

func (s *DescribeCdnDeliverListResponse) SetHeaders(v map[string]*string) *DescribeCdnDeliverListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDeliverListResponse) SetStatusCode(v int32) *DescribeCdnDeliverListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDeliverListResponse) SetBody(v *DescribeCdnDeliverListResponseBody) *DescribeCdnDeliverListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDeliverListResponse) Validate() error {
	return dara.Validate(s)
}
