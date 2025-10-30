// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendSceneVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecommendSceneVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecommendSceneVariablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecommendSceneVariablesResponseBody) *DescribeRecommendSceneVariablesResponse
	GetBody() *DescribeRecommendSceneVariablesResponseBody
}

type DescribeRecommendSceneVariablesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecommendSceneVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecommendSceneVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendSceneVariablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendSceneVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecommendSceneVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecommendSceneVariablesResponse) GetBody() *DescribeRecommendSceneVariablesResponseBody {
	return s.Body
}

func (s *DescribeRecommendSceneVariablesResponse) SetHeaders(v map[string]*string) *DescribeRecommendSceneVariablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendSceneVariablesResponse) SetStatusCode(v int32) *DescribeRecommendSceneVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecommendSceneVariablesResponse) SetBody(v *DescribeRecommendSceneVariablesResponseBody) *DescribeRecommendSceneVariablesResponse {
	s.Body = v
	return s
}

func (s *DescribeRecommendSceneVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
