// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchNodesResponseBody) *DescribeOpenSearchNodesResponse
	GetBody() *DescribeOpenSearchNodesResponseBody
}

type DescribeOpenSearchNodesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchNodesResponse) GetBody() *DescribeOpenSearchNodesResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchNodesResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchNodesResponse) SetStatusCode(v int32) *DescribeOpenSearchNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchNodesResponse) SetBody(v *DescribeOpenSearchNodesResponseBody) *DescribeOpenSearchNodesResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
