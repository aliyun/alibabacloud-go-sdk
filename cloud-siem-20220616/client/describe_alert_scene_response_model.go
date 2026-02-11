// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertSceneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertSceneResponseBody) *DescribeAlertSceneResponse
	GetBody() *DescribeAlertSceneResponseBody
}

type DescribeAlertSceneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertSceneResponse) GetBody() *DescribeAlertSceneResponseBody {
	return s.Body
}

func (s *DescribeAlertSceneResponse) SetHeaders(v map[string]*string) *DescribeAlertSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSceneResponse) SetStatusCode(v int32) *DescribeAlertSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSceneResponse) SetBody(v *DescribeAlertSceneResponseBody) *DescribeAlertSceneResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
