// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoSnapshotPolicyResponseBody) *DescribeAutoSnapshotPolicyResponse
	GetBody() *DescribeAutoSnapshotPolicyResponseBody
}

type DescribeAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoSnapshotPolicyResponse) GetBody() *DescribeAutoSnapshotPolicyResponseBody {
	return s.Body
}

func (s *DescribeAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponse) SetStatusCode(v int32) *DescribeAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponse) SetBody(v *DescribeAutoSnapshotPolicyResponseBody) *DescribeAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
