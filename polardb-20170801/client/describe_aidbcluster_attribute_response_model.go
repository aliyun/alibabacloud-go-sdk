// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterAttributeResponseBody) *DescribeAIDBClusterAttributeResponse
	GetBody() *DescribeAIDBClusterAttributeResponseBody
}

type DescribeAIDBClusterAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterAttributeResponse) GetBody() *DescribeAIDBClusterAttributeResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterAttributeResponse) SetStatusCode(v int32) *DescribeAIDBClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponse) SetBody(v *DescribeAIDBClusterAttributeResponseBody) *DescribeAIDBClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
