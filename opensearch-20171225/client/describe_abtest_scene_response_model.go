// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeABTestSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeABTestSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeABTestSceneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeABTestSceneResponseBody) *DescribeABTestSceneResponse
	GetBody() *DescribeABTestSceneResponseBody
}

type DescribeABTestSceneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeABTestSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeABTestSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeABTestSceneResponse) GetBody() *DescribeABTestSceneResponseBody {
	return s.Body
}

func (s *DescribeABTestSceneResponse) SetHeaders(v map[string]*string) *DescribeABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestSceneResponse) SetStatusCode(v int32) *DescribeABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeABTestSceneResponse) SetBody(v *DescribeABTestSceneResponseBody) *DescribeABTestSceneResponse {
	s.Body = v
	return s
}

func (s *DescribeABTestSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
