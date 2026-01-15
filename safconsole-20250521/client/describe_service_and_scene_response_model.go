// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAndSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceAndSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceAndSceneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceAndSceneResponseBody) *DescribeServiceAndSceneResponse
	GetBody() *DescribeServiceAndSceneResponseBody
}

type DescribeServiceAndSceneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceAndSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceAndSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAndSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceAndSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceAndSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceAndSceneResponse) GetBody() *DescribeServiceAndSceneResponseBody {
	return s.Body
}

func (s *DescribeServiceAndSceneResponse) SetHeaders(v map[string]*string) *DescribeServiceAndSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceAndSceneResponse) SetStatusCode(v int32) *DescribeServiceAndSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceAndSceneResponse) SetBody(v *DescribeServiceAndSceneResponseBody) *DescribeServiceAndSceneResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceAndSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
