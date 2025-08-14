// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDecisionResultFluctuationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDecisionResultFluctuationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDecisionResultFluctuationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDecisionResultFluctuationResponseBody) *DescribeDecisionResultFluctuationResponse
	GetBody() *DescribeDecisionResultFluctuationResponseBody
}

type DescribeDecisionResultFluctuationResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDecisionResultFluctuationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDecisionResultFluctuationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultFluctuationResponse) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultFluctuationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDecisionResultFluctuationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDecisionResultFluctuationResponse) GetBody() *DescribeDecisionResultFluctuationResponseBody {
	return s.Body
}

func (s *DescribeDecisionResultFluctuationResponse) SetHeaders(v map[string]*string) *DescribeDecisionResultFluctuationResponse {
	s.Headers = v
	return s
}

func (s *DescribeDecisionResultFluctuationResponse) SetStatusCode(v int32) *DescribeDecisionResultFluctuationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponse) SetBody(v *DescribeDecisionResultFluctuationResponseBody) *DescribeDecisionResultFluctuationResponse {
	s.Body = v
	return s
}

func (s *DescribeDecisionResultFluctuationResponse) Validate() error {
	return dara.Validate(s)
}
