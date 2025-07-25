// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmRecoveryPlanResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmRecoveryPlanResponseBody) *DescribeGtmRecoveryPlanResponse
	GetBody() *DescribeGtmRecoveryPlanResponseBody
}

type DescribeGtmRecoveryPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmRecoveryPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmRecoveryPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmRecoveryPlanResponse) GetBody() *DescribeGtmRecoveryPlanResponseBody {
	return s.Body
}

func (s *DescribeGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponse) SetStatusCode(v int32) *DescribeGtmRecoveryPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponse) SetBody(v *DescribeGtmRecoveryPlanResponseBody) *DescribeGtmRecoveryPlanResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponse) Validate() error {
	return dara.Validate(s)
}
