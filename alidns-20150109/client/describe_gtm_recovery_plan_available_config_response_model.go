// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlanAvailableConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlanAvailableConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmRecoveryPlanAvailableConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmRecoveryPlanAvailableConfigResponseBody) *DescribeGtmRecoveryPlanAvailableConfigResponse
	GetBody() *DescribeGtmRecoveryPlanAvailableConfigResponseBody
}

type DescribeGtmRecoveryPlanAvailableConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmRecoveryPlanAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) GetBody() *DescribeGtmRecoveryPlanAvailableConfigResponseBody {
	return s.Body
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmRecoveryPlanAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) SetStatusCode(v int32) *DescribeGtmRecoveryPlanAvailableConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) SetBody(v *DescribeGtmRecoveryPlanAvailableConfigResponseBody) *DescribeGtmRecoveryPlanAvailableConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigResponse) Validate() error {
	return dara.Validate(s)
}
