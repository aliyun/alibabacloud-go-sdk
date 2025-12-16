// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeABTestGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeABTestGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeABTestGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeABTestGroupResponseBody) *DescribeABTestGroupResponse
	GetBody() *DescribeABTestGroupResponseBody
}

type DescribeABTestGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeABTestGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeABTestGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeABTestGroupResponse) GetBody() *DescribeABTestGroupResponseBody {
	return s.Body
}

func (s *DescribeABTestGroupResponse) SetHeaders(v map[string]*string) *DescribeABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestGroupResponse) SetStatusCode(v int32) *DescribeABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeABTestGroupResponse) SetBody(v *DescribeABTestGroupResponseBody) *DescribeABTestGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeABTestGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
