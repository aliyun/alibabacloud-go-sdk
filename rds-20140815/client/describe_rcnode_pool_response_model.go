// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCNodePoolResponseBody) *DescribeRCNodePoolResponse
	GetBody() *DescribeRCNodePoolResponseBody
}

type DescribeRCNodePoolResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNodePoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCNodePoolResponse) GetBody() *DescribeRCNodePoolResponseBody {
	return s.Body
}

func (s *DescribeRCNodePoolResponse) SetHeaders(v map[string]*string) *DescribeRCNodePoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCNodePoolResponse) SetStatusCode(v int32) *DescribeRCNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCNodePoolResponse) SetBody(v *DescribeRCNodePoolResponseBody) *DescribeRCNodePoolResponse {
	s.Body = v
	return s
}

func (s *DescribeRCNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
