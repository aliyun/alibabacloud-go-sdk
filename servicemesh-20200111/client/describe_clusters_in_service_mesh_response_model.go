// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersInServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClustersInServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClustersInServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClustersInServiceMeshResponseBody) *DescribeClustersInServiceMeshResponse
	GetBody() *DescribeClustersInServiceMeshResponseBody
}

type DescribeClustersInServiceMeshResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClustersInServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClustersInServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClustersInServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClustersInServiceMeshResponse) GetBody() *DescribeClustersInServiceMeshResponseBody {
	return s.Body
}

func (s *DescribeClustersInServiceMeshResponse) SetHeaders(v map[string]*string) *DescribeClustersInServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersInServiceMeshResponse) SetStatusCode(v int32) *DescribeClustersInServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponse) SetBody(v *DescribeClustersInServiceMeshResponseBody) *DescribeClustersInServiceMeshResponse {
	s.Body = v
	return s
}

func (s *DescribeClustersInServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
