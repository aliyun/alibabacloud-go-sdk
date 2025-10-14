// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionIdResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsRegionIdResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsRegionIdResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsRegionIdResourceResponseBody) *DescribeEnsRegionIdResourceResponse
	GetBody() *DescribeEnsRegionIdResourceResponseBody
}

type DescribeEnsRegionIdResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsRegionIdResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsRegionIdResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsRegionIdResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsRegionIdResourceResponse) GetBody() *DescribeEnsRegionIdResourceResponseBody {
	return s.Body
}

func (s *DescribeEnsRegionIdResourceResponse) SetHeaders(v map[string]*string) *DescribeEnsRegionIdResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) SetStatusCode(v int32) *DescribeEnsRegionIdResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) SetBody(v *DescribeEnsRegionIdResourceResponseBody) *DescribeEnsRegionIdResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
