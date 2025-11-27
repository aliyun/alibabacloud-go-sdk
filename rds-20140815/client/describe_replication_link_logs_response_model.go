// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationLinkLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReplicationLinkLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReplicationLinkLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReplicationLinkLogsResponseBody) *DescribeReplicationLinkLogsResponse
	GetBody() *DescribeReplicationLinkLogsResponseBody
}

type DescribeReplicationLinkLogsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReplicationLinkLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReplicationLinkLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationLinkLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicationLinkLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReplicationLinkLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReplicationLinkLogsResponse) GetBody() *DescribeReplicationLinkLogsResponseBody {
	return s.Body
}

func (s *DescribeReplicationLinkLogsResponse) SetHeaders(v map[string]*string) *DescribeReplicationLinkLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicationLinkLogsResponse) SetStatusCode(v int32) *DescribeReplicationLinkLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponse) SetBody(v *DescribeReplicationLinkLogsResponseBody) *DescribeReplicationLinkLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeReplicationLinkLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
