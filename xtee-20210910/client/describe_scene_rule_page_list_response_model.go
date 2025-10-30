// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneRulePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSceneRulePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSceneRulePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSceneRulePageListResponseBody) *DescribeSceneRulePageListResponse
	GetBody() *DescribeSceneRulePageListResponseBody
}

type DescribeSceneRulePageListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSceneRulePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSceneRulePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneRulePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneRulePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSceneRulePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSceneRulePageListResponse) GetBody() *DescribeSceneRulePageListResponseBody {
	return s.Body
}

func (s *DescribeSceneRulePageListResponse) SetHeaders(v map[string]*string) *DescribeSceneRulePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneRulePageListResponse) SetStatusCode(v int32) *DescribeSceneRulePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSceneRulePageListResponse) SetBody(v *DescribeSceneRulePageListResponseBody) *DescribeSceneRulePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeSceneRulePageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
