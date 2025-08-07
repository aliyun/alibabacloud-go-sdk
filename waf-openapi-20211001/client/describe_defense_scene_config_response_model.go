// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseSceneConfigResponseBody) *DescribeDefenseSceneConfigResponse
	GetBody() *DescribeDefenseSceneConfigResponseBody
}

type DescribeDefenseSceneConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseSceneConfigResponse) GetBody() *DescribeDefenseSceneConfigResponseBody {
	return s.Body
}

func (s *DescribeDefenseSceneConfigResponse) SetHeaders(v map[string]*string) *DescribeDefenseSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseSceneConfigResponse) SetStatusCode(v int32) *DescribeDefenseSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseSceneConfigResponse) SetBody(v *DescribeDefenseSceneConfigResponseBody) *DescribeDefenseSceneConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseSceneConfigResponse) Validate() error {
	return dara.Validate(s)
}
