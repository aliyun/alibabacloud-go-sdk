// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentPlaybookResponseBody) *DescribeComponentPlaybookResponse
	GetBody() *DescribeComponentPlaybookResponseBody
}

type DescribeComponentPlaybookResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentPlaybookResponse) GetBody() *DescribeComponentPlaybookResponseBody {
	return s.Body
}

func (s *DescribeComponentPlaybookResponse) SetHeaders(v map[string]*string) *DescribeComponentPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentPlaybookResponse) SetStatusCode(v int32) *DescribeComponentPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentPlaybookResponse) SetBody(v *DescribeComponentPlaybookResponseBody) *DescribeComponentPlaybookResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
