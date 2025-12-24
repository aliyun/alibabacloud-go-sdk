// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlaybookResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlaybookResponseBody) *DescribePlaybookResponse
	GetBody() *DescribePlaybookResponseBody
}

type DescribePlaybookResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlaybookResponse) GetBody() *DescribePlaybookResponseBody {
	return s.Body
}

func (s *DescribePlaybookResponse) SetHeaders(v map[string]*string) *DescribePlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookResponse) SetStatusCode(v int32) *DescribePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookResponse) SetBody(v *DescribePlaybookResponseBody) *DescribePlaybookResponse {
	s.Body = v
	return s
}

func (s *DescribePlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
