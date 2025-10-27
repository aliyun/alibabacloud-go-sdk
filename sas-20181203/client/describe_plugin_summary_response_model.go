// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginSummaryResponseBody) *DescribePluginSummaryResponse
	GetBody() *DescribePluginSummaryResponseBody
}

type DescribePluginSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginSummaryResponse) GetBody() *DescribePluginSummaryResponseBody {
	return s.Body
}

func (s *DescribePluginSummaryResponse) SetHeaders(v map[string]*string) *DescribePluginSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginSummaryResponse) SetStatusCode(v int32) *DescribePluginSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginSummaryResponse) SetBody(v *DescribePluginSummaryResponseBody) *DescribePluginSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribePluginSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
