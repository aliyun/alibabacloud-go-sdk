// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRegionBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkRegionBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkRegionBlockResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkRegionBlockResponseBody) *DescribeNetworkRegionBlockResponse
	GetBody() *DescribeNetworkRegionBlockResponseBody
}

type DescribeNetworkRegionBlockResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkRegionBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkRegionBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRegionBlockResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkRegionBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkRegionBlockResponse) GetBody() *DescribeNetworkRegionBlockResponseBody {
	return s.Body
}

func (s *DescribeNetworkRegionBlockResponse) SetHeaders(v map[string]*string) *DescribeNetworkRegionBlockResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRegionBlockResponse) SetStatusCode(v int32) *DescribeNetworkRegionBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkRegionBlockResponse) SetBody(v *DescribeNetworkRegionBlockResponseBody) *DescribeNetworkRegionBlockResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkRegionBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
