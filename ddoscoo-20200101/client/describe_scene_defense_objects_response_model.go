// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneDefenseObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSceneDefenseObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSceneDefenseObjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSceneDefenseObjectsResponseBody) *DescribeSceneDefenseObjectsResponse
	GetBody() *DescribeSceneDefenseObjectsResponseBody
}

type DescribeSceneDefenseObjectsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSceneDefenseObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSceneDefenseObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefenseObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSceneDefenseObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSceneDefenseObjectsResponse) GetBody() *DescribeSceneDefenseObjectsResponseBody {
	return s.Body
}

func (s *DescribeSceneDefenseObjectsResponse) SetHeaders(v map[string]*string) *DescribeSceneDefenseObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneDefenseObjectsResponse) SetStatusCode(v int32) *DescribeSceneDefenseObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponse) SetBody(v *DescribeSceneDefenseObjectsResponseBody) *DescribeSceneDefenseObjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeSceneDefenseObjectsResponse) Validate() error {
	return dara.Validate(s)
}
