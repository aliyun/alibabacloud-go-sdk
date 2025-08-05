// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOtsTableSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOtsTableSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOtsTableSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOtsTableSnapshotsResponseBody) *DescribeOtsTableSnapshotsResponse
	GetBody() *DescribeOtsTableSnapshotsResponseBody
}

type DescribeOtsTableSnapshotsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOtsTableSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOtsTableSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOtsTableSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOtsTableSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOtsTableSnapshotsResponse) GetBody() *DescribeOtsTableSnapshotsResponseBody {
	return s.Body
}

func (s *DescribeOtsTableSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeOtsTableSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOtsTableSnapshotsResponse) SetStatusCode(v int32) *DescribeOtsTableSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponse) SetBody(v *DescribeOtsTableSnapshotsResponseBody) *DescribeOtsTableSnapshotsResponse {
	s.Body = v
	return s
}

func (s *DescribeOtsTableSnapshotsResponse) Validate() error {
	return dara.Validate(s)
}
