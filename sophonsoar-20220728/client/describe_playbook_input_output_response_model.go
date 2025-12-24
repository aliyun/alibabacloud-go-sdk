// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookInputOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlaybookInputOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlaybookInputOutputResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlaybookInputOutputResponseBody) *DescribePlaybookInputOutputResponse
	GetBody() *DescribePlaybookInputOutputResponseBody
}

type DescribePlaybookInputOutputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookInputOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookInputOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookInputOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlaybookInputOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlaybookInputOutputResponse) GetBody() *DescribePlaybookInputOutputResponseBody {
	return s.Body
}

func (s *DescribePlaybookInputOutputResponse) SetHeaders(v map[string]*string) *DescribePlaybookInputOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookInputOutputResponse) SetStatusCode(v int32) *DescribePlaybookInputOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookInputOutputResponse) SetBody(v *DescribePlaybookInputOutputResponseBody) *DescribePlaybookInputOutputResponse {
	s.Body = v
	return s
}

func (s *DescribePlaybookInputOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
