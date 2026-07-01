// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAgenticDBProjectsResponseBodyItems) *DescribeAgenticDBProjectsResponseBody
	GetItems() []*DescribeAgenticDBProjectsResponseBodyItems
	SetPageNumber(v int32) *DescribeAgenticDBProjectsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBProjectsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAgenticDBProjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAgenticDBProjectsResponseBody
	GetTotalCount() *int32
}

type DescribeAgenticDBProjectsResponseBody struct {
	Items []*DescribeAgenticDBProjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F6A7B8C9-D0E1-2345-FABC-456789012345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAgenticDBProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectsResponseBody) GetItems() []*DescribeAgenticDBProjectsResponseBodyItems {
	return s.Items
}

func (s *DescribeAgenticDBProjectsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBProjectsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBProjectsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAgenticDBProjectsResponseBody) SetItems(v []*DescribeAgenticDBProjectsResponseBodyItems) *DescribeAgenticDBProjectsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBody) SetPageNumber(v int32) *DescribeAgenticDBProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBody) SetPageSize(v int32) *DescribeAgenticDBProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBody) SetRequestId(v string) *DescribeAgenticDBProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBody) SetTotalCount(v int32) *DescribeAgenticDBProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAgenticDBProjectsResponseBodyItems struct {
	// example:
	//
	// 2026-06-10T11:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// br-1a2b3c4d5e6f
	DefaultBranchId *string `json:"DefaultBranchId,omitempty" xml:"DefaultBranchId,omitempty"`
	// example:
	//
	// main
	DefaultBranchName *string `json:"DefaultBranchName,omitempty" xml:"DefaultBranchName,omitempty"`
	// example:
	//
	// Production analytics database
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// analytics-prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBProjectsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetDefaultBranchId() *string {
	return s.DefaultBranchId
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetDefaultBranchName() *string {
	return s.DefaultBranchName
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetCreateTime(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetDBClusterId(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetDefaultBranchId(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.DefaultBranchId = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetDefaultBranchName(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.DefaultBranchName = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetDescription(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetProjectId(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetProjectName(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.ProjectName = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetStatus(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) SetTenantId(v string) *DescribeAgenticDBProjectsResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
