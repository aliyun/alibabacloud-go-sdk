// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSchedulerRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSchedulerRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSchedulerRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSchedulerRulesResponseBody) *DescribeSchedulerRulesResponse
	GetBody() *DescribeSchedulerRulesResponseBody
}

type DescribeSchedulerRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSchedulerRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSchedulerRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchedulerRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSchedulerRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSchedulerRulesResponse) GetBody() *DescribeSchedulerRulesResponseBody {
	return s.Body
}

func (s *DescribeSchedulerRulesResponse) SetHeaders(v map[string]*string) *DescribeSchedulerRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSchedulerRulesResponse) SetStatusCode(v int32) *DescribeSchedulerRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSchedulerRulesResponse) SetBody(v *DescribeSchedulerRulesResponseBody) *DescribeSchedulerRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeSchedulerRulesResponse) Validate() error {
	return dara.Validate(s)
}
