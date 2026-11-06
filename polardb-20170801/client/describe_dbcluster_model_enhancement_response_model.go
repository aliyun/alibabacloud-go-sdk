// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterModelEnhancementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterModelEnhancementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterModelEnhancementResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterModelEnhancementResponseBody) *DescribeDBClusterModelEnhancementResponse
	GetBody() *DescribeDBClusterModelEnhancementResponseBody
}

type DescribeDBClusterModelEnhancementResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterModelEnhancementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterModelEnhancementResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterModelEnhancementResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterModelEnhancementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterModelEnhancementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterModelEnhancementResponse) GetBody() *DescribeDBClusterModelEnhancementResponseBody {
	return s.Body
}

func (s *DescribeDBClusterModelEnhancementResponse) SetHeaders(v map[string]*string) *DescribeDBClusterModelEnhancementResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponse) SetStatusCode(v int32) *DescribeDBClusterModelEnhancementResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponse) SetBody(v *DescribeDBClusterModelEnhancementResponseBody) *DescribeDBClusterModelEnhancementResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
