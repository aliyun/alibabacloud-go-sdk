// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshVMsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshVMsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshVMsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshVMsResponseBody) *DescribeServiceMeshVMsResponse
	GetBody() *DescribeServiceMeshVMsResponseBody
}

type DescribeServiceMeshVMsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshVMsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshVMsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshVMsResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshVMsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshVMsResponse) GetBody() *DescribeServiceMeshVMsResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshVMsResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshVMsResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshVMsResponse) SetStatusCode(v int32) *DescribeServiceMeshVMsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshVMsResponse) SetBody(v *DescribeServiceMeshVMsResponseBody) *DescribeServiceMeshVMsResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshVMsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
