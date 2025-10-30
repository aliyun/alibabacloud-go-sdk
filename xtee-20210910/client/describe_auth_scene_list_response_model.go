// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthSceneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthSceneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthSceneListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthSceneListResponseBody) *DescribeAuthSceneListResponse
	GetBody() *DescribeAuthSceneListResponseBody
}

type DescribeAuthSceneListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthSceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthSceneListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthSceneListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthSceneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthSceneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthSceneListResponse) GetBody() *DescribeAuthSceneListResponseBody {
	return s.Body
}

func (s *DescribeAuthSceneListResponse) SetHeaders(v map[string]*string) *DescribeAuthSceneListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthSceneListResponse) SetStatusCode(v int32) *DescribeAuthSceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthSceneListResponse) SetBody(v *DescribeAuthSceneListResponseBody) *DescribeAuthSceneListResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthSceneListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
