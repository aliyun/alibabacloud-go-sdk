// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanInstanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRatePlanInstanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRatePlanInstanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRatePlanInstanceStatusResponseBody) *DescribeRatePlanInstanceStatusResponse
	GetBody() *DescribeRatePlanInstanceStatusResponseBody
}

type DescribeRatePlanInstanceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRatePlanInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRatePlanInstanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanInstanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRatePlanInstanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRatePlanInstanceStatusResponse) GetBody() *DescribeRatePlanInstanceStatusResponseBody {
	return s.Body
}

func (s *DescribeRatePlanInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeRatePlanInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRatePlanInstanceStatusResponse) SetStatusCode(v int32) *DescribeRatePlanInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRatePlanInstanceStatusResponse) SetBody(v *DescribeRatePlanInstanceStatusResponseBody) *DescribeRatePlanInstanceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRatePlanInstanceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
