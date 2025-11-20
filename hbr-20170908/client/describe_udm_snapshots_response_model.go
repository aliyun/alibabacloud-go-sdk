// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUdmSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUdmSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUdmSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUdmSnapshotsResponseBody) *DescribeUdmSnapshotsResponse
	GetBody() *DescribeUdmSnapshotsResponseBody
}

type DescribeUdmSnapshotsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUdmSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUdmSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdmSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUdmSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUdmSnapshotsResponse) GetBody() *DescribeUdmSnapshotsResponseBody {
	return s.Body
}

func (s *DescribeUdmSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeUdmSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUdmSnapshotsResponse) SetStatusCode(v int32) *DescribeUdmSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUdmSnapshotsResponse) SetBody(v *DescribeUdmSnapshotsResponseBody) *DescribeUdmSnapshotsResponse {
	s.Body = v
	return s
}

func (s *DescribeUdmSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
