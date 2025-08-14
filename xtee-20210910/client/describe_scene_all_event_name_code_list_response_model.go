// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneAllEventNameCodeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSceneAllEventNameCodeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSceneAllEventNameCodeListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSceneAllEventNameCodeListResponseBody) *DescribeSceneAllEventNameCodeListResponse
	GetBody() *DescribeSceneAllEventNameCodeListResponseBody
}

type DescribeSceneAllEventNameCodeListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSceneAllEventNameCodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSceneAllEventNameCodeListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneAllEventNameCodeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneAllEventNameCodeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSceneAllEventNameCodeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSceneAllEventNameCodeListResponse) GetBody() *DescribeSceneAllEventNameCodeListResponseBody {
	return s.Body
}

func (s *DescribeSceneAllEventNameCodeListResponse) SetHeaders(v map[string]*string) *DescribeSceneAllEventNameCodeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponse) SetStatusCode(v int32) *DescribeSceneAllEventNameCodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponse) SetBody(v *DescribeSceneAllEventNameCodeListResponseBody) *DescribeSceneAllEventNameCodeListResponse {
	s.Body = v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponse) Validate() error {
	return dara.Validate(s)
}
