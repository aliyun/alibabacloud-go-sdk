// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutePlaybooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExecutePlaybooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExecutePlaybooksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExecutePlaybooksResponseBody) *DescribeExecutePlaybooksResponse
	GetBody() *DescribeExecutePlaybooksResponseBody
}

type DescribeExecutePlaybooksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExecutePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExecutePlaybooksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExecutePlaybooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExecutePlaybooksResponse) GetBody() *DescribeExecutePlaybooksResponseBody {
	return s.Body
}

func (s *DescribeExecutePlaybooksResponse) SetHeaders(v map[string]*string) *DescribeExecutePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeExecutePlaybooksResponse) SetStatusCode(v int32) *DescribeExecutePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExecutePlaybooksResponse) SetBody(v *DescribeExecutePlaybooksResponseBody) *DescribeExecutePlaybooksResponse {
	s.Body = v
	return s
}

func (s *DescribeExecutePlaybooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
