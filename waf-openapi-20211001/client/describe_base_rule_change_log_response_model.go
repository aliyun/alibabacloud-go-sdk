// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBaseRuleChangeLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBaseRuleChangeLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBaseRuleChangeLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBaseRuleChangeLogResponseBody) *DescribeBaseRuleChangeLogResponse
	GetBody() *DescribeBaseRuleChangeLogResponseBody
}

type DescribeBaseRuleChangeLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBaseRuleChangeLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBaseRuleChangeLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseRuleChangeLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeBaseRuleChangeLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBaseRuleChangeLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBaseRuleChangeLogResponse) GetBody() *DescribeBaseRuleChangeLogResponseBody {
	return s.Body
}

func (s *DescribeBaseRuleChangeLogResponse) SetHeaders(v map[string]*string) *DescribeBaseRuleChangeLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeBaseRuleChangeLogResponse) SetStatusCode(v int32) *DescribeBaseRuleChangeLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponse) SetBody(v *DescribeBaseRuleChangeLogResponseBody) *DescribeBaseRuleChangeLogResponse {
	s.Body = v
	return s
}

func (s *DescribeBaseRuleChangeLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
