// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnbindFlowLogSagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUnbindFlowLogSagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUnbindFlowLogSagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUnbindFlowLogSagsResponseBody) *DescribeUnbindFlowLogSagsResponse
	GetBody() *DescribeUnbindFlowLogSagsResponseBody
}

type DescribeUnbindFlowLogSagsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUnbindFlowLogSagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUnbindFlowLogSagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnbindFlowLogSagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnbindFlowLogSagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUnbindFlowLogSagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUnbindFlowLogSagsResponse) GetBody() *DescribeUnbindFlowLogSagsResponseBody {
	return s.Body
}

func (s *DescribeUnbindFlowLogSagsResponse) SetHeaders(v map[string]*string) *DescribeUnbindFlowLogSagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponse) SetStatusCode(v int32) *DescribeUnbindFlowLogSagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponse) SetBody(v *DescribeUnbindFlowLogSagsResponseBody) *DescribeUnbindFlowLogSagsResponse {
	s.Body = v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
