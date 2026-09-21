// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSkillsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSkillsResponseBody) *DescribeSkillsResponse
	GetBody() *DescribeSkillsResponseBody
}

type DescribeSkillsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSkillsResponse) GetBody() *DescribeSkillsResponseBody {
	return s.Body
}

func (s *DescribeSkillsResponse) SetHeaders(v map[string]*string) *DescribeSkillsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSkillsResponse) SetStatusCode(v int32) *DescribeSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSkillsResponse) SetBody(v *DescribeSkillsResponseBody) *DescribeSkillsResponse {
	s.Body = v
	return s
}

func (s *DescribeSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
