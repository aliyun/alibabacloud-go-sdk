// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecEventsResponseBody) *DescribeApisecEventsResponse
	GetBody() *DescribeApisecEventsResponseBody
}

type DescribeApisecEventsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecEventsResponse) GetBody() *DescribeApisecEventsResponseBody {
	return s.Body
}

func (s *DescribeApisecEventsResponse) SetHeaders(v map[string]*string) *DescribeApisecEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecEventsResponse) SetStatusCode(v int32) *DescribeApisecEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecEventsResponse) SetBody(v *DescribeApisecEventsResponseBody) *DescribeApisecEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecEventsResponse) Validate() error {
	return dara.Validate(s)
}
