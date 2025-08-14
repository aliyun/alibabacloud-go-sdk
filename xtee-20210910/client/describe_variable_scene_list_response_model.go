// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableSceneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVariableSceneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVariableSceneListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVariableSceneListResponseBody) *DescribeVariableSceneListResponse
	GetBody() *DescribeVariableSceneListResponseBody
}

type DescribeVariableSceneListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVariableSceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVariableSceneListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableSceneListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVariableSceneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVariableSceneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVariableSceneListResponse) GetBody() *DescribeVariableSceneListResponseBody {
	return s.Body
}

func (s *DescribeVariableSceneListResponse) SetHeaders(v map[string]*string) *DescribeVariableSceneListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVariableSceneListResponse) SetStatusCode(v int32) *DescribeVariableSceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVariableSceneListResponse) SetBody(v *DescribeVariableSceneListResponseBody) *DescribeVariableSceneListResponse {
	s.Body = v
	return s
}

func (s *DescribeVariableSceneListResponse) Validate() error {
	return dara.Validate(s)
}
