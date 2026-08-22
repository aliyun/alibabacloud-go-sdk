// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchTopologyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchTopologyResponseBody) *DescribeOpenSearchTopologyResponse
	GetBody() *DescribeOpenSearchTopologyResponseBody
}

type DescribeOpenSearchTopologyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchTopologyResponse) GetBody() *DescribeOpenSearchTopologyResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchTopologyResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchTopologyResponse) SetStatusCode(v int32) *DescribeOpenSearchTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponse) SetBody(v *DescribeOpenSearchTopologyResponseBody) *DescribeOpenSearchTopologyResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
