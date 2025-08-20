// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCompactionServiceSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCompactionServiceSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCompactionServiceSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCompactionServiceSwitchResponseBody) *DescribeCompactionServiceSwitchResponse
	GetBody() *DescribeCompactionServiceSwitchResponseBody
}

type DescribeCompactionServiceSwitchResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCompactionServiceSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCompactionServiceSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCompactionServiceSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeCompactionServiceSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCompactionServiceSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCompactionServiceSwitchResponse) GetBody() *DescribeCompactionServiceSwitchResponseBody {
	return s.Body
}

func (s *DescribeCompactionServiceSwitchResponse) SetHeaders(v map[string]*string) *DescribeCompactionServiceSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeCompactionServiceSwitchResponse) SetStatusCode(v int32) *DescribeCompactionServiceSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCompactionServiceSwitchResponse) SetBody(v *DescribeCompactionServiceSwitchResponseBody) *DescribeCompactionServiceSwitchResponse {
	s.Body = v
	return s
}

func (s *DescribeCompactionServiceSwitchResponse) Validate() error {
	return dara.Validate(s)
}
