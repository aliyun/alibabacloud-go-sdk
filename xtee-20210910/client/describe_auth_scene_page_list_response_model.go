// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthScenePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthScenePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthScenePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthScenePageListResponseBody) *DescribeAuthScenePageListResponse
	GetBody() *DescribeAuthScenePageListResponseBody
}

type DescribeAuthScenePageListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthScenePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthScenePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthScenePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthScenePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthScenePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthScenePageListResponse) GetBody() *DescribeAuthScenePageListResponseBody {
	return s.Body
}

func (s *DescribeAuthScenePageListResponse) SetHeaders(v map[string]*string) *DescribeAuthScenePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthScenePageListResponse) SetStatusCode(v int32) *DescribeAuthScenePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthScenePageListResponse) SetBody(v *DescribeAuthScenePageListResponseBody) *DescribeAuthScenePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthScenePageListResponse) Validate() error {
	return dara.Validate(s)
}
