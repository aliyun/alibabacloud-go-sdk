// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExistSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExistSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExistSceneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExistSceneResponseBody) *DescribeExistSceneResponse
	GetBody() *DescribeExistSceneResponseBody
}

type DescribeExistSceneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExistSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExistSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExistSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeExistSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExistSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExistSceneResponse) GetBody() *DescribeExistSceneResponseBody {
	return s.Body
}

func (s *DescribeExistSceneResponse) SetHeaders(v map[string]*string) *DescribeExistSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeExistSceneResponse) SetStatusCode(v int32) *DescribeExistSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExistSceneResponse) SetBody(v *DescribeExistSceneResponseBody) *DescribeExistSceneResponse {
	s.Body = v
	return s
}

func (s *DescribeExistSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
