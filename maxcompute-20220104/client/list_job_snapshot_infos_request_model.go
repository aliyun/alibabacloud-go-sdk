// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobSnapshotInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *ListJobSnapshotInfosRequest
	GetAscOrder() *bool
	SetExtNodeIdList(v []*string) *ListJobSnapshotInfosRequest
	GetExtNodeIdList() []*string
	SetFrom(v int64) *ListJobSnapshotInfosRequest
	GetFrom() *int64
	SetInstanceIdList(v []*string) *ListJobSnapshotInfosRequest
	GetInstanceIdList() []*string
	SetJobOwnerList(v []*string) *ListJobSnapshotInfosRequest
	GetJobOwnerList() []*string
	SetPriorityList(v []*int64) *ListJobSnapshotInfosRequest
	GetPriorityList() []*int64
	SetProjectList(v []*string) *ListJobSnapshotInfosRequest
	GetProjectList() []*string
	SetQuotaNickname(v string) *ListJobSnapshotInfosRequest
	GetQuotaNickname() *string
	SetSignatureList(v []*string) *ListJobSnapshotInfosRequest
	GetSignatureList() []*string
	SetSortByList(v []*string) *ListJobSnapshotInfosRequest
	GetSortByList() []*string
	SetSortOrderList(v []*string) *ListJobSnapshotInfosRequest
	GetSortOrderList() []*string
	SetStatusList(v []*string) *ListJobSnapshotInfosRequest
	GetStatusList() []*string
	SetTo(v int64) *ListJobSnapshotInfosRequest
	GetTo() *int64
	SetTypeList(v []*string) *ListJobSnapshotInfosRequest
	GetTypeList() []*string
	SetOrderColumn(v string) *ListJobSnapshotInfosRequest
	GetOrderColumn() *string
	SetPageNumber(v int64) *ListJobSnapshotInfosRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListJobSnapshotInfosRequest
	GetPageSize() *int64
	SetRegion(v string) *ListJobSnapshotInfosRequest
	GetRegion() *string
	SetTenantId(v string) *ListJobSnapshotInfosRequest
	GetTenantId() *string
}

type ListJobSnapshotInfosRequest struct {
	// Specifies whether to sort data in ascending order.
	//
	// example:
	//
	// true
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The ID of the upstream node.
	ExtNodeIdList []*string `json:"extNodeIdList,omitempty" xml:"extNodeIdList,omitempty" type:"Repeated"`
	// Start timestamp.
	//
	// > This parameter is invalid. The end timestamp should be the time point for the snapshot you want to view.
	//
	// example:
	//
	// 1706840714
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The instance ID.
	InstanceIdList []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
	// The account that commits the job.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The priority of the job.
	PriorityList []*int64 `json:"priorityList,omitempty" xml:"priorityList,omitempty" type:"Repeated"`
	// The name of project.
	ProjectList []*string `json:"projectList,omitempty" xml:"projectList,omitempty" type:"Repeated"`
	// The nickname of the compute Quota used by the job.
	//
	// example:
	//
	// quota_A
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The signature of the SQL job.
	SignatureList []*string `json:"signatureList,omitempty" xml:"signatureList,omitempty" type:"Repeated"`
	// The sorting columns.
	SortByList []*string `json:"sortByList,omitempty" xml:"sortByList,omitempty" type:"Repeated"`
	// The orders for the sorting columns.
	SortOrderList []*string `json:"sortOrderList,omitempty" xml:"sortOrderList,omitempty" type:"Repeated"`
	// The status of jobs.
	StatusList []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
	// End timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1706840714
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The type of the job.
	TypeList []*string `json:"typeList,omitempty" xml:"typeList,omitempty" type:"Repeated"`
	// The sorting column.
	//
	// example:
	//
	// cpuUsage
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose Tenants > Tenant Property from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListJobSnapshotInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobSnapshotInfosRequest) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *ListJobSnapshotInfosRequest) GetExtNodeIdList() []*string {
	return s.ExtNodeIdList
}

func (s *ListJobSnapshotInfosRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListJobSnapshotInfosRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *ListJobSnapshotInfosRequest) GetJobOwnerList() []*string {
	return s.JobOwnerList
}

func (s *ListJobSnapshotInfosRequest) GetPriorityList() []*int64 {
	return s.PriorityList
}

func (s *ListJobSnapshotInfosRequest) GetProjectList() []*string {
	return s.ProjectList
}

func (s *ListJobSnapshotInfosRequest) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *ListJobSnapshotInfosRequest) GetSignatureList() []*string {
	return s.SignatureList
}

func (s *ListJobSnapshotInfosRequest) GetSortByList() []*string {
	return s.SortByList
}

func (s *ListJobSnapshotInfosRequest) GetSortOrderList() []*string {
	return s.SortOrderList
}

func (s *ListJobSnapshotInfosRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListJobSnapshotInfosRequest) GetTo() *int64 {
	return s.To
}

func (s *ListJobSnapshotInfosRequest) GetTypeList() []*string {
	return s.TypeList
}

func (s *ListJobSnapshotInfosRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListJobSnapshotInfosRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListJobSnapshotInfosRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListJobSnapshotInfosRequest) GetRegion() *string {
	return s.Region
}

func (s *ListJobSnapshotInfosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListJobSnapshotInfosRequest) SetAscOrder(v bool) *ListJobSnapshotInfosRequest {
	s.AscOrder = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetExtNodeIdList(v []*string) *ListJobSnapshotInfosRequest {
	s.ExtNodeIdList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetFrom(v int64) *ListJobSnapshotInfosRequest {
	s.From = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetInstanceIdList(v []*string) *ListJobSnapshotInfosRequest {
	s.InstanceIdList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetJobOwnerList(v []*string) *ListJobSnapshotInfosRequest {
	s.JobOwnerList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetPriorityList(v []*int64) *ListJobSnapshotInfosRequest {
	s.PriorityList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetProjectList(v []*string) *ListJobSnapshotInfosRequest {
	s.ProjectList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetQuotaNickname(v string) *ListJobSnapshotInfosRequest {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetSignatureList(v []*string) *ListJobSnapshotInfosRequest {
	s.SignatureList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetSortByList(v []*string) *ListJobSnapshotInfosRequest {
	s.SortByList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetSortOrderList(v []*string) *ListJobSnapshotInfosRequest {
	s.SortOrderList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetStatusList(v []*string) *ListJobSnapshotInfosRequest {
	s.StatusList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetTo(v int64) *ListJobSnapshotInfosRequest {
	s.To = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetTypeList(v []*string) *ListJobSnapshotInfosRequest {
	s.TypeList = v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetOrderColumn(v string) *ListJobSnapshotInfosRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetPageNumber(v int64) *ListJobSnapshotInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetPageSize(v int64) *ListJobSnapshotInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetRegion(v string) *ListJobSnapshotInfosRequest {
	s.Region = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) SetTenantId(v string) *ListJobSnapshotInfosRequest {
	s.TenantId = &v
	return s
}

func (s *ListJobSnapshotInfosRequest) Validate() error {
	return dara.Validate(s)
}
