// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesFlatPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvicesFlatPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvicesFlatPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvicesFlatPageResponseBody) *DescribeAdvicesFlatPageResponse
	GetBody() *DescribeAdvicesFlatPageResponseBody
}

type DescribeAdvicesFlatPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvicesFlatPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvicesFlatPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvicesFlatPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvicesFlatPageResponse) GetBody() *DescribeAdvicesFlatPageResponseBody {
	return s.Body
}

func (s *DescribeAdvicesFlatPageResponse) SetHeaders(v map[string]*string) *DescribeAdvicesFlatPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvicesFlatPageResponse) SetStatusCode(v int32) *DescribeAdvicesFlatPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponse) SetBody(v *DescribeAdvicesFlatPageResponseBody) *DescribeAdvicesFlatPageResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvicesFlatPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
