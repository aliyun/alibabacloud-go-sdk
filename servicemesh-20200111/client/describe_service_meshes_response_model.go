// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshesResponseBody) *DescribeServiceMeshesResponse
	GetBody() *DescribeServiceMeshesResponseBody
}

type DescribeServiceMeshesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshesResponse) GetBody() *DescribeServiceMeshesResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshesResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshesResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshesResponse) SetStatusCode(v int32) *DescribeServiceMeshesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshesResponse) SetBody(v *DescribeServiceMeshesResponseBody) *DescribeServiceMeshesResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
