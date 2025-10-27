// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyTypeScaItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyTypeScaItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyTypeScaItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyTypeScaItemResponseBody) *DescribePropertyTypeScaItemResponse
	GetBody() *DescribePropertyTypeScaItemResponseBody
}

type DescribePropertyTypeScaItemResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyTypeScaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyTypeScaItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyTypeScaItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyTypeScaItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyTypeScaItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyTypeScaItemResponse) GetBody() *DescribePropertyTypeScaItemResponseBody {
	return s.Body
}

func (s *DescribePropertyTypeScaItemResponse) SetHeaders(v map[string]*string) *DescribePropertyTypeScaItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyTypeScaItemResponse) SetStatusCode(v int32) *DescribePropertyTypeScaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyTypeScaItemResponse) SetBody(v *DescribePropertyTypeScaItemResponseBody) *DescribePropertyTypeScaItemResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyTypeScaItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
