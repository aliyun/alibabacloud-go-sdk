// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnterpriseSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnterpriseSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnterpriseSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnterpriseSnapshotPolicyResponseBody) *DescribeEnterpriseSnapshotPolicyResponse
	GetBody() *DescribeEnterpriseSnapshotPolicyResponseBody
}

type DescribeEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnterpriseSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) GetBody() *DescribeEnterpriseSnapshotPolicyResponseBody {
	return s.Body
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DescribeEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *DescribeEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) SetBody(v *DescribeEnterpriseSnapshotPolicyResponseBody) *DescribeEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeEnterpriseSnapshotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
