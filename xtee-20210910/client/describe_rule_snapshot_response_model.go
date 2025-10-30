// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleSnapshotResponseBody) *DescribeRuleSnapshotResponse
	GetBody() *DescribeRuleSnapshotResponseBody
}

type DescribeRuleSnapshotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleSnapshotResponse) GetBody() *DescribeRuleSnapshotResponseBody {
	return s.Body
}

func (s *DescribeRuleSnapshotResponse) SetHeaders(v map[string]*string) *DescribeRuleSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleSnapshotResponse) SetStatusCode(v int32) *DescribeRuleSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleSnapshotResponse) SetBody(v *DescribeRuleSnapshotResponseBody) *DescribeRuleSnapshotResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
