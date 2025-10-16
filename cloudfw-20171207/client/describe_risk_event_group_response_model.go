// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskEventGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskEventGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskEventGroupResponseBody) *DescribeRiskEventGroupResponse
	GetBody() *DescribeRiskEventGroupResponseBody
}

type DescribeRiskEventGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskEventGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskEventGroupResponse) GetBody() *DescribeRiskEventGroupResponseBody {
	return s.Body
}

func (s *DescribeRiskEventGroupResponse) SetHeaders(v map[string]*string) *DescribeRiskEventGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventGroupResponse) SetStatusCode(v int32) *DescribeRiskEventGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventGroupResponse) SetBody(v *DescribeRiskEventGroupResponseBody) *DescribeRiskEventGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskEventGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
