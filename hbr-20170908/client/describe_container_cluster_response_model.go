// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerClusterResponseBody) *DescribeContainerClusterResponse
	GetBody() *DescribeContainerClusterResponseBody
}

type DescribeContainerClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerClusterResponse) GetBody() *DescribeContainerClusterResponseBody {
	return s.Body
}

func (s *DescribeContainerClusterResponse) SetHeaders(v map[string]*string) *DescribeContainerClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerClusterResponse) SetStatusCode(v int32) *DescribeContainerClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerClusterResponse) SetBody(v *DescribeContainerClusterResponseBody) *DescribeContainerClusterResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
