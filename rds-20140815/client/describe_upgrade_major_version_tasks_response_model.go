// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeMajorVersionTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpgradeMajorVersionTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpgradeMajorVersionTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpgradeMajorVersionTasksResponseBody) *DescribeUpgradeMajorVersionTasksResponse
	GetBody() *DescribeUpgradeMajorVersionTasksResponseBody
}

type DescribeUpgradeMajorVersionTasksResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpgradeMajorVersionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpgradeMajorVersionTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpgradeMajorVersionTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpgradeMajorVersionTasksResponse) GetBody() *DescribeUpgradeMajorVersionTasksResponseBody {
	return s.Body
}

func (s *DescribeUpgradeMajorVersionTasksResponse) SetHeaders(v map[string]*string) *DescribeUpgradeMajorVersionTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponse) SetStatusCode(v int32) *DescribeUpgradeMajorVersionTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponse) SetBody(v *DescribeUpgradeMajorVersionTasksResponseBody) *DescribeUpgradeMajorVersionTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
