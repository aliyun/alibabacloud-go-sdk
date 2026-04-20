// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterTaskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterTaskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterTaskAttributeResponseBody) *DescribeAIDBClusterTaskAttributeResponse
	GetBody() *DescribeAIDBClusterTaskAttributeResponseBody
}

type DescribeAIDBClusterTaskAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterTaskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterTaskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterTaskAttributeResponse) GetBody() *DescribeAIDBClusterTaskAttributeResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterTaskAttributeResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponse) SetStatusCode(v int32) *DescribeAIDBClusterTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponse) SetBody(v *DescribeAIDBClusterTaskAttributeResponseBody) *DescribeAIDBClusterTaskAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
