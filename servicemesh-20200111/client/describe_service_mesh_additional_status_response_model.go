// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshAdditionalStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshAdditionalStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshAdditionalStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshAdditionalStatusResponseBody) *DescribeServiceMeshAdditionalStatusResponse
	GetBody() *DescribeServiceMeshAdditionalStatusResponseBody
}

type DescribeServiceMeshAdditionalStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshAdditionalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshAdditionalStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshAdditionalStatusResponse) GetBody() *DescribeServiceMeshAdditionalStatusResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshAdditionalStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshAdditionalStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponse) SetStatusCode(v int32) *DescribeServiceMeshAdditionalStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponse) SetBody(v *DescribeServiceMeshAdditionalStatusResponseBody) *DescribeServiceMeshAdditionalStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
