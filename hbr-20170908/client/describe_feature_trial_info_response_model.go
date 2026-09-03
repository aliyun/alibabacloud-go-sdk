// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureTrialInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFeatureTrialInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFeatureTrialInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFeatureTrialInfoResponseBody) *DescribeFeatureTrialInfoResponse
	GetBody() *DescribeFeatureTrialInfoResponseBody
}

type DescribeFeatureTrialInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFeatureTrialInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFeatureTrialInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureTrialInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeFeatureTrialInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFeatureTrialInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFeatureTrialInfoResponse) GetBody() *DescribeFeatureTrialInfoResponseBody {
	return s.Body
}

func (s *DescribeFeatureTrialInfoResponse) SetHeaders(v map[string]*string) *DescribeFeatureTrialInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeFeatureTrialInfoResponse) SetStatusCode(v int32) *DescribeFeatureTrialInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponse) SetBody(v *DescribeFeatureTrialInfoResponseBody) *DescribeFeatureTrialInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeFeatureTrialInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
