// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToEcrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGrantRulesToEcrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGrantRulesToEcrResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGrantRulesToEcrResponseBody) *DescribeGrantRulesToEcrResponse
	GetBody() *DescribeGrantRulesToEcrResponseBody
}

type DescribeGrantRulesToEcrResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGrantRulesToEcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGrantRulesToEcrResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToEcrResponse) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToEcrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGrantRulesToEcrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGrantRulesToEcrResponse) GetBody() *DescribeGrantRulesToEcrResponseBody {
	return s.Body
}

func (s *DescribeGrantRulesToEcrResponse) SetHeaders(v map[string]*string) *DescribeGrantRulesToEcrResponse {
	s.Headers = v
	return s
}

func (s *DescribeGrantRulesToEcrResponse) SetStatusCode(v int32) *DescribeGrantRulesToEcrResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponse) SetBody(v *DescribeGrantRulesToEcrResponseBody) *DescribeGrantRulesToEcrResponse {
	s.Body = v
	return s
}

func (s *DescribeGrantRulesToEcrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
