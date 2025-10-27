// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyScaItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyScaItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyScaItemResponseBody) *DescribePropertyScaItemResponse
	GetBody() *DescribePropertyScaItemResponseBody
}

type DescribePropertyScaItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyScaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyScaItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyScaItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyScaItemResponse) GetBody() *DescribePropertyScaItemResponseBody {
	return s.Body
}

func (s *DescribePropertyScaItemResponse) SetHeaders(v map[string]*string) *DescribePropertyScaItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyScaItemResponse) SetStatusCode(v int32) *DescribePropertyScaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyScaItemResponse) SetBody(v *DescribePropertyScaItemResponseBody) *DescribePropertyScaItemResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyScaItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
