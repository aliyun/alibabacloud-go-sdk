// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerServiceK8sClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerServiceK8sClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerServiceK8sClustersResponseBody) *DescribeContainerServiceK8sClustersResponse
	GetBody() *DescribeContainerServiceK8sClustersResponseBody
}

type DescribeContainerServiceK8sClustersResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerServiceK8sClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerServiceK8sClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerServiceK8sClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerServiceK8sClustersResponse) GetBody() *DescribeContainerServiceK8sClustersResponseBody {
	return s.Body
}

func (s *DescribeContainerServiceK8sClustersResponse) SetHeaders(v map[string]*string) *DescribeContainerServiceK8sClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerServiceK8sClustersResponse) SetStatusCode(v int32) *DescribeContainerServiceK8sClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerServiceK8sClustersResponse) SetBody(v *DescribeContainerServiceK8sClustersResponseBody) *DescribeContainerServiceK8sClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerServiceK8sClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
