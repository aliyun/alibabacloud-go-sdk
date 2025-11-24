// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshLogsResponseBody) *DescribeServiceMeshLogsResponse
	GetBody() *DescribeServiceMeshLogsResponseBody
}

type DescribeServiceMeshLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshLogsResponse) GetBody() *DescribeServiceMeshLogsResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshLogsResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshLogsResponse) SetStatusCode(v int32) *DescribeServiceMeshLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshLogsResponse) SetBody(v *DescribeServiceMeshLogsResponseBody) *DescribeServiceMeshLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
