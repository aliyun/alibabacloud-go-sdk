// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListBySceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOperatorListBySceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOperatorListBySceneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOperatorListBySceneResponseBody) *DescribeOperatorListBySceneResponse
	GetBody() *DescribeOperatorListBySceneResponseBody
}

type DescribeOperatorListBySceneResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperatorListBySceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperatorListBySceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListBySceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListBySceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOperatorListBySceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOperatorListBySceneResponse) GetBody() *DescribeOperatorListBySceneResponseBody {
	return s.Body
}

func (s *DescribeOperatorListBySceneResponse) SetHeaders(v map[string]*string) *DescribeOperatorListBySceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorListBySceneResponse) SetStatusCode(v int32) *DescribeOperatorListBySceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorListBySceneResponse) SetBody(v *DescribeOperatorListBySceneResponseBody) *DescribeOperatorListBySceneResponse {
	s.Body = v
	return s
}

func (s *DescribeOperatorListBySceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
