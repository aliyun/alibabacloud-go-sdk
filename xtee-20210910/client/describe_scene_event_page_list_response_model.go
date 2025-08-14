// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneEventPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSceneEventPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSceneEventPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSceneEventPageListResponseBody) *DescribeSceneEventPageListResponse
	GetBody() *DescribeSceneEventPageListResponseBody
}

type DescribeSceneEventPageListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSceneEventPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSceneEventPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneEventPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneEventPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSceneEventPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSceneEventPageListResponse) GetBody() *DescribeSceneEventPageListResponseBody {
	return s.Body
}

func (s *DescribeSceneEventPageListResponse) SetHeaders(v map[string]*string) *DescribeSceneEventPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneEventPageListResponse) SetStatusCode(v int32) *DescribeSceneEventPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSceneEventPageListResponse) SetBody(v *DescribeSceneEventPageListResponseBody) *DescribeSceneEventPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeSceneEventPageListResponse) Validate() error {
	return dara.Validate(s)
}
