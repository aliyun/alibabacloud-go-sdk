// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchRestoreWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody
	GetBranchId() *string
	SetDBClusterId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody
	GetDBClusterId() *string
	SetEarliestRestorableTime(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody
	GetEarliestRestorableTime() *string
	SetProjectId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody
	GetProjectId() *string
	SetQueriedAt(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody
	GetQueriedAt() *string
	SetRequestId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody
	GetRequestId() *string
	SetTenantId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody
	GetTenantId() *string
}

type DescribeAgenticDBBranchRestoreWindowResponseBody struct {
	// The ID of the source branch.
	//
	// example:
	//
	// br-1a2b3c4d5e6f
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// DBClusterId
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The earliest time available for creating a historical branch, in UTC RFC3339 format. This time point can be selected. The available range changes over time and is revalidated during creation.
	//
	// example:
	//
	// 2026-10-15T07:32:00Z
	EarliestRestorableTime *string `json:"EarliestRestorableTime,omitempty" xml:"EarliestRestorableTime,omitempty"`
	// ProjectId
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The time when the query was completed, in UTC RFC3339 format. This does not represent the latest restorable time.
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

func (s DescribeAgenticDBBranchRestoreWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchRestoreWindowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) GetEarliestRestorableTime() *string {
	return s.EarliestRestorableTime
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) GetQueriedAt() *string {
	return s.QueriedAt
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) SetBranchId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) SetDBClusterId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) SetEarliestRestorableTime(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody {
	s.EarliestRestorableTime = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) SetProjectId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) SetQueriedAt(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody {
	s.QueriedAt = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) SetRequestId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) SetTenantId(v string) *DescribeAgenticDBBranchRestoreWindowResponseBody {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
