// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceRegionIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceRegionIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceRegionIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceRegionIdResponseBody) *DescribeResourceRegionIdResponse
	GetBody() *DescribeResourceRegionIdResponseBody
}

type DescribeResourceRegionIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceRegionIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceRegionIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceRegionIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceRegionIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceRegionIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceRegionIdResponse) GetBody() *DescribeResourceRegionIdResponseBody {
	return s.Body
}

func (s *DescribeResourceRegionIdResponse) SetHeaders(v map[string]*string) *DescribeResourceRegionIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceRegionIdResponse) SetStatusCode(v int32) *DescribeResourceRegionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceRegionIdResponse) SetBody(v *DescribeResourceRegionIdResponseBody) *DescribeResourceRegionIdResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceRegionIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
