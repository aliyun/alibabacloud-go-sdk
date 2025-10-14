// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsResourceUsageResponseBody) *DescribeEnsResourceUsageResponse
	GetBody() *DescribeEnsResourceUsageResponseBody
}

type DescribeEnsResourceUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsResourceUsageResponse) GetBody() *DescribeEnsResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeEnsResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeEnsResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsResourceUsageResponse) SetStatusCode(v int32) *DescribeEnsResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsResourceUsageResponse) SetBody(v *DescribeEnsResourceUsageResponseBody) *DescribeEnsResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
