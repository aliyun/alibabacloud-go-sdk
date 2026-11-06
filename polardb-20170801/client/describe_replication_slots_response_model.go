// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationSlotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReplicationSlotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReplicationSlotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReplicationSlotsResponseBody) *DescribeReplicationSlotsResponse
	GetBody() *DescribeReplicationSlotsResponseBody
}

type DescribeReplicationSlotsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReplicationSlotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReplicationSlotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationSlotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicationSlotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReplicationSlotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReplicationSlotsResponse) GetBody() *DescribeReplicationSlotsResponseBody {
	return s.Body
}

func (s *DescribeReplicationSlotsResponse) SetHeaders(v map[string]*string) *DescribeReplicationSlotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicationSlotsResponse) SetStatusCode(v int32) *DescribeReplicationSlotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReplicationSlotsResponse) SetBody(v *DescribeReplicationSlotsResponseBody) *DescribeReplicationSlotsResponse {
	s.Body = v
	return s
}

func (s *DescribeReplicationSlotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
