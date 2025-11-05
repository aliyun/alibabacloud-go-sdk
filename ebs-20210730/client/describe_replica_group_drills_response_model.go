// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicaGroupDrillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReplicaGroupDrillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReplicaGroupDrillsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReplicaGroupDrillsResponseBody) *DescribeReplicaGroupDrillsResponse
	GetBody() *DescribeReplicaGroupDrillsResponseBody
}

type DescribeReplicaGroupDrillsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReplicaGroupDrillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReplicaGroupDrillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReplicaGroupDrillsResponse) GetBody() *DescribeReplicaGroupDrillsResponseBody {
	return s.Body
}

func (s *DescribeReplicaGroupDrillsResponse) SetHeaders(v map[string]*string) *DescribeReplicaGroupDrillsResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicaGroupDrillsResponse) SetStatusCode(v int32) *DescribeReplicaGroupDrillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponse) SetBody(v *DescribeReplicaGroupDrillsResponseBody) *DescribeReplicaGroupDrillsResponse {
	s.Body = v
	return s
}

func (s *DescribeReplicaGroupDrillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
