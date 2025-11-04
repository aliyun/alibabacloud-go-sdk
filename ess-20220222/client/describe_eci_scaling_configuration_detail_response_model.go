// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEciScalingConfigurationDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEciScalingConfigurationDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEciScalingConfigurationDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEciScalingConfigurationDetailResponseBody) *DescribeEciScalingConfigurationDetailResponse
	GetBody() *DescribeEciScalingConfigurationDetailResponseBody
}

type DescribeEciScalingConfigurationDetailResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEciScalingConfigurationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEciScalingConfigurationDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEciScalingConfigurationDetailResponse) GetBody() *DescribeEciScalingConfigurationDetailResponseBody {
	return s.Body
}

func (s *DescribeEciScalingConfigurationDetailResponse) SetHeaders(v map[string]*string) *DescribeEciScalingConfigurationDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponse) SetStatusCode(v int32) *DescribeEciScalingConfigurationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponse) SetBody(v *DescribeEciScalingConfigurationDetailResponseBody) *DescribeEciScalingConfigurationDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
