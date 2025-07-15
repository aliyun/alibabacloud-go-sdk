// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionLOAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhysicalConnectionLOAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhysicalConnectionLOAResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhysicalConnectionLOAResponseBody) *DescribePhysicalConnectionLOAResponse
	GetBody() *DescribePhysicalConnectionLOAResponseBody
}

type DescribePhysicalConnectionLOAResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhysicalConnectionLOAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhysicalConnectionLOAResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionLOAResponse) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionLOAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhysicalConnectionLOAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhysicalConnectionLOAResponse) GetBody() *DescribePhysicalConnectionLOAResponseBody {
	return s.Body
}

func (s *DescribePhysicalConnectionLOAResponse) SetHeaders(v map[string]*string) *DescribePhysicalConnectionLOAResponse {
	s.Headers = v
	return s
}

func (s *DescribePhysicalConnectionLOAResponse) SetStatusCode(v int32) *DescribePhysicalConnectionLOAResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponse) SetBody(v *DescribePhysicalConnectionLOAResponseBody) *DescribePhysicalConnectionLOAResponse {
	s.Body = v
	return s
}

func (s *DescribePhysicalConnectionLOAResponse) Validate() error {
	return dara.Validate(s)
}
