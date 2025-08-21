// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainSnapshotDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainSnapshotDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainSnapshotDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainSnapshotDataResponseBody) *DescribeVsDomainSnapshotDataResponse
	GetBody() *DescribeVsDomainSnapshotDataResponseBody
}

type DescribeVsDomainSnapshotDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainSnapshotDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainSnapshotDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainSnapshotDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainSnapshotDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainSnapshotDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainSnapshotDataResponse) GetBody() *DescribeVsDomainSnapshotDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainSnapshotDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainSnapshotDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponse) SetStatusCode(v int32) *DescribeVsDomainSnapshotDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponse) SetBody(v *DescribeVsDomainSnapshotDataResponseBody) *DescribeVsDomainSnapshotDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainSnapshotDataResponse) Validate() error {
	return dara.Validate(s)
}
