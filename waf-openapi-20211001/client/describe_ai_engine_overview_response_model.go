// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineOverviewResponseBody) *DescribeAiEngineOverviewResponse
	GetBody() *DescribeAiEngineOverviewResponseBody
}

type DescribeAiEngineOverviewResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineOverviewResponse) GetBody() *DescribeAiEngineOverviewResponseBody {
	return s.Body
}

func (s *DescribeAiEngineOverviewResponse) SetHeaders(v map[string]*string) *DescribeAiEngineOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineOverviewResponse) SetStatusCode(v int32) *DescribeAiEngineOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineOverviewResponse) SetBody(v *DescribeAiEngineOverviewResponseBody) *DescribeAiEngineOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
