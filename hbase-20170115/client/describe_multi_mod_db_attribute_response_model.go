// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiModDbAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMultiModDbAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMultiModDbAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMultiModDbAttributeResponseBody) *DescribeMultiModDbAttributeResponse
	GetBody() *DescribeMultiModDbAttributeResponseBody
}

type DescribeMultiModDbAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiModDbAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiModDbAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiModDbAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiModDbAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMultiModDbAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMultiModDbAttributeResponse) GetBody() *DescribeMultiModDbAttributeResponseBody {
	return s.Body
}

func (s *DescribeMultiModDbAttributeResponse) SetHeaders(v map[string]*string) *DescribeMultiModDbAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiModDbAttributeResponse) SetStatusCode(v int32) *DescribeMultiModDbAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiModDbAttributeResponse) SetBody(v *DescribeMultiModDbAttributeResponseBody) *DescribeMultiModDbAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeMultiModDbAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
