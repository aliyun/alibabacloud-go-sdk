// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectHistoryRetentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody
	GetDBClusterId() *string
	SetHistoryRetentionSeconds(v int64) *DescribeAgenticDBProjectHistoryRetentionResponseBody
	GetHistoryRetentionSeconds() *int64
	SetProjectId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody
	GetProjectId() *string
	SetQueriedAt(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody
	GetQueriedAt() *string
	SetRequestId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody
	GetRequestId() *string
	SetTenantId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody
	GetTenantId() *string
}

type DescribeAgenticDBProjectHistoryRetentionResponseBody struct {
	// DBClusterId
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The effective history data retention period for the project. For the actual recoverable range, query the branch recovery window.
	//
	// example:
	//
	// 604800
	HistoryRetentionSeconds *int64 `json:"HistoryRetentionSeconds,omitempty" xml:"HistoryRetentionSeconds,omitempty"`
	// ProjectId
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The time when the query was completed, in UTC RFC3339 format. This does not represent the latest recoverable time.
	//
	// example:
	//
	// 2026-10-16T07:32:00Z
	QueriedAt *string `json:"QueriedAt,omitempty" xml:"QueriedAt,omitempty"`
	// The request ID.
	//
	// example:
	//
	// req-123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TenantId
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBProjectHistoryRetentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectHistoryRetentionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) GetHistoryRetentionSeconds() *int64 {
	return s.HistoryRetentionSeconds
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) GetQueriedAt() *string {
	return s.QueriedAt
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) SetDBClusterId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) SetHistoryRetentionSeconds(v int64) *DescribeAgenticDBProjectHistoryRetentionResponseBody {
	s.HistoryRetentionSeconds = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) SetProjectId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) SetQueriedAt(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody {
	s.QueriedAt = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) SetRequestId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) SetTenantId(v string) *DescribeAgenticDBProjectHistoryRetentionResponseBody {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponseBody) Validate() error {
	return dara.Validate(s)
}
