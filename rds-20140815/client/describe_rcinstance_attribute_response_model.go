// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstanceAttributeResponseBody) *DescribeRCInstanceAttributeResponse
	GetBody() *DescribeRCInstanceAttributeResponseBody
}

type DescribeRCInstanceAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstanceAttributeResponse) GetBody() *DescribeRCInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeRCInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeRCInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstanceAttributeResponse) SetStatusCode(v int32) *DescribeRCInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstanceAttributeResponse) SetBody(v *DescribeRCInstanceAttributeResponseBody) *DescribeRCInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
