// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicaSetRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReplicaSetRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReplicaSetRoleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReplicaSetRoleResponseBody) *DescribeReplicaSetRoleResponse
	GetBody() *DescribeReplicaSetRoleResponseBody
}

type DescribeReplicaSetRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReplicaSetRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReplicaSetRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaSetRoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReplicaSetRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReplicaSetRoleResponse) GetBody() *DescribeReplicaSetRoleResponseBody {
	return s.Body
}

func (s *DescribeReplicaSetRoleResponse) SetHeaders(v map[string]*string) *DescribeReplicaSetRoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicaSetRoleResponse) SetStatusCode(v int32) *DescribeReplicaSetRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReplicaSetRoleResponse) SetBody(v *DescribeReplicaSetRoleResponseBody) *DescribeReplicaSetRoleResponse {
	s.Body = v
	return s
}

func (s *DescribeReplicaSetRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
