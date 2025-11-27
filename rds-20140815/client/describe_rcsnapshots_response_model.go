// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCSnapshotsResponseBody) *DescribeRCSnapshotsResponse
	GetBody() *DescribeRCSnapshotsResponseBody
}

type DescribeRCSnapshotsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCSnapshotsResponse) GetBody() *DescribeRCSnapshotsResponseBody {
	return s.Body
}

func (s *DescribeRCSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeRCSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCSnapshotsResponse) SetStatusCode(v int32) *DescribeRCSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCSnapshotsResponse) SetBody(v *DescribeRCSnapshotsResponseBody) *DescribeRCSnapshotsResponse {
	s.Body = v
	return s
}

func (s *DescribeRCSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
