// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlaybooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlaybooksResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlaybooksResponseBody) *DescribePlaybooksResponse
	GetBody() *DescribePlaybooksResponseBody
}

type DescribePlaybooksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybooksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlaybooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlaybooksResponse) GetBody() *DescribePlaybooksResponseBody {
	return s.Body
}

func (s *DescribePlaybooksResponse) SetHeaders(v map[string]*string) *DescribePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybooksResponse) SetStatusCode(v int32) *DescribePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybooksResponse) SetBody(v *DescribePlaybooksResponseBody) *DescribePlaybooksResponse {
	s.Body = v
	return s
}

func (s *DescribePlaybooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
