// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppInstanceAttributeResponseBody) *DescribeAppInstanceAttributeResponse
	GetBody() *DescribeAppInstanceAttributeResponseBody
}

type DescribeAppInstanceAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppInstanceAttributeResponse) GetBody() *DescribeAppInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeAppInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeAppInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppInstanceAttributeResponse) SetStatusCode(v int32) *DescribeAppInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponse) SetBody(v *DescribeAppInstanceAttributeResponseBody) *DescribeAppInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAppInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
