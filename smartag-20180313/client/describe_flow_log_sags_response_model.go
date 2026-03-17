// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogSagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowLogSagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowLogSagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowLogSagsResponseBody) *DescribeFlowLogSagsResponse
	GetBody() *DescribeFlowLogSagsResponseBody
}

type DescribeFlowLogSagsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowLogSagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowLogSagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogSagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogSagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowLogSagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowLogSagsResponse) GetBody() *DescribeFlowLogSagsResponseBody {
	return s.Body
}

func (s *DescribeFlowLogSagsResponse) SetHeaders(v map[string]*string) *DescribeFlowLogSagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowLogSagsResponse) SetStatusCode(v int32) *DescribeFlowLogSagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowLogSagsResponse) SetBody(v *DescribeFlowLogSagsResponseBody) *DescribeFlowLogSagsResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowLogSagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
