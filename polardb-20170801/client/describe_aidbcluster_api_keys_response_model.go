// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterApiKeysResponseBody) *DescribeAIDBClusterApiKeysResponse
	GetBody() *DescribeAIDBClusterApiKeysResponseBody
}

type DescribeAIDBClusterApiKeysResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterApiKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterApiKeysResponse) GetBody() *DescribeAIDBClusterApiKeysResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterApiKeysResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterApiKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponse) SetStatusCode(v int32) *DescribeAIDBClusterApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponse) SetBody(v *DescribeAIDBClusterApiKeysResponseBody) *DescribeAIDBClusterApiKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
