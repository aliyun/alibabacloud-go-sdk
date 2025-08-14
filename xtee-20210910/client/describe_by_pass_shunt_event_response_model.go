// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeByPassShuntEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeByPassShuntEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeByPassShuntEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeByPassShuntEventResponseBody) *DescribeByPassShuntEventResponse
	GetBody() *DescribeByPassShuntEventResponseBody
}

type DescribeByPassShuntEventResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeByPassShuntEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeByPassShuntEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeByPassShuntEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeByPassShuntEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeByPassShuntEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeByPassShuntEventResponse) GetBody() *DescribeByPassShuntEventResponseBody {
	return s.Body
}

func (s *DescribeByPassShuntEventResponse) SetHeaders(v map[string]*string) *DescribeByPassShuntEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeByPassShuntEventResponse) SetStatusCode(v int32) *DescribeByPassShuntEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeByPassShuntEventResponse) SetBody(v *DescribeByPassShuntEventResponseBody) *DescribeByPassShuntEventResponse {
	s.Body = v
	return s
}

func (s *DescribeByPassShuntEventResponse) Validate() error {
	return dara.Validate(s)
}
