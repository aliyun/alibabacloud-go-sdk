// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeACLAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeACLAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeACLAttributeResponseBody) *DescribeACLAttributeResponse
	GetBody() *DescribeACLAttributeResponseBody
}

type DescribeACLAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeACLAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeACLAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeACLAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeACLAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeACLAttributeResponse) GetBody() *DescribeACLAttributeResponseBody {
	return s.Body
}

func (s *DescribeACLAttributeResponse) SetHeaders(v map[string]*string) *DescribeACLAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeACLAttributeResponse) SetStatusCode(v int32) *DescribeACLAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeACLAttributeResponse) SetBody(v *DescribeACLAttributeResponseBody) *DescribeACLAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeACLAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
