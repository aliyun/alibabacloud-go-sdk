// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWafUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWafUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWafUsageDataResponseBody) *DescribeWafUsageDataResponse
	GetBody() *DescribeWafUsageDataResponseBody
}

type DescribeWafUsageDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWafUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWafUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeWafUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWafUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWafUsageDataResponse) GetBody() *DescribeWafUsageDataResponseBody {
	return s.Body
}

func (s *DescribeWafUsageDataResponse) SetHeaders(v map[string]*string) *DescribeWafUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeWafUsageDataResponse) SetStatusCode(v int32) *DescribeWafUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWafUsageDataResponse) SetBody(v *DescribeWafUsageDataResponseBody) *DescribeWafUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeWafUsageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
