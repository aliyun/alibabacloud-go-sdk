// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookNodesOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlaybookNodesOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlaybookNodesOutputResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlaybookNodesOutputResponseBody) *DescribePlaybookNodesOutputResponse
	GetBody() *DescribePlaybookNodesOutputResponseBody
}

type DescribePlaybookNodesOutputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookNodesOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookNodesOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlaybookNodesOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlaybookNodesOutputResponse) GetBody() *DescribePlaybookNodesOutputResponseBody {
	return s.Body
}

func (s *DescribePlaybookNodesOutputResponse) SetHeaders(v map[string]*string) *DescribePlaybookNodesOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookNodesOutputResponse) SetStatusCode(v int32) *DescribePlaybookNodesOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponse) SetBody(v *DescribePlaybookNodesOutputResponseBody) *DescribePlaybookNodesOutputResponse {
	s.Body = v
	return s
}

func (s *DescribePlaybookNodesOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
