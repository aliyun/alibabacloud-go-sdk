// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcHoneyPotCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcHoneyPotCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcHoneyPotCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcHoneyPotCriteriaResponseBody) *DescribeVpcHoneyPotCriteriaResponse
	GetBody() *DescribeVpcHoneyPotCriteriaResponseBody
}

type DescribeVpcHoneyPotCriteriaResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcHoneyPotCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcHoneyPotCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcHoneyPotCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcHoneyPotCriteriaResponse) GetBody() *DescribeVpcHoneyPotCriteriaResponseBody {
	return s.Body
}

func (s *DescribeVpcHoneyPotCriteriaResponse) SetHeaders(v map[string]*string) *DescribeVpcHoneyPotCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponse) SetStatusCode(v int32) *DescribeVpcHoneyPotCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponse) SetBody(v *DescribeVpcHoneyPotCriteriaResponseBody) *DescribeVpcHoneyPotCriteriaResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
