// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessControlListAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessControlListAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessControlListAttributeResponseBody) *DescribeAccessControlListAttributeResponse
	GetBody() *DescribeAccessControlListAttributeResponseBody
}

type DescribeAccessControlListAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessControlListAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessControlListAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessControlListAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessControlListAttributeResponse) GetBody() *DescribeAccessControlListAttributeResponseBody {
	return s.Body
}

func (s *DescribeAccessControlListAttributeResponse) SetHeaders(v map[string]*string) *DescribeAccessControlListAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessControlListAttributeResponse) SetStatusCode(v int32) *DescribeAccessControlListAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponse) SetBody(v *DescribeAccessControlListAttributeResponseBody) *DescribeAccessControlListAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessControlListAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
