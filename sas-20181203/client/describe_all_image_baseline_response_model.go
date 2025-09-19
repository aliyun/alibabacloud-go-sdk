// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllImageBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllImageBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllImageBaselineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllImageBaselineResponseBody) *DescribeAllImageBaselineResponse
	GetBody() *DescribeAllImageBaselineResponseBody
}

type DescribeAllImageBaselineResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllImageBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllImageBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllImageBaselineResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllImageBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllImageBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllImageBaselineResponse) GetBody() *DescribeAllImageBaselineResponseBody {
	return s.Body
}

func (s *DescribeAllImageBaselineResponse) SetHeaders(v map[string]*string) *DescribeAllImageBaselineResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllImageBaselineResponse) SetStatusCode(v int32) *DescribeAllImageBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllImageBaselineResponse) SetBody(v *DescribeAllImageBaselineResponseBody) *DescribeAllImageBaselineResponse {
	s.Body = v
	return s
}

func (s *DescribeAllImageBaselineResponse) Validate() error {
	return dara.Validate(s)
}
