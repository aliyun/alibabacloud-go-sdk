// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestorePlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestorePlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestorePlansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestorePlansResponseBody) *DescribeRestorePlansResponse
	GetBody() *DescribeRestorePlansResponseBody
}

type DescribeRestorePlansResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestorePlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestorePlansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestorePlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestorePlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestorePlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestorePlansResponse) GetBody() *DescribeRestorePlansResponseBody {
	return s.Body
}

func (s *DescribeRestorePlansResponse) SetHeaders(v map[string]*string) *DescribeRestorePlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestorePlansResponse) SetStatusCode(v int32) *DescribeRestorePlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestorePlansResponse) SetBody(v *DescribeRestorePlansResponseBody) *DescribeRestorePlansResponse {
	s.Body = v
	return s
}

func (s *DescribeRestorePlansResponse) Validate() error {
	return dara.Validate(s)
}
