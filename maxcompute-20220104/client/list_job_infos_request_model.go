// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *ListJobInfosRequest
	GetAscOrder() *bool
	SetExtNodeIdList(v []*string) *ListJobInfosRequest
	GetExtNodeIdList() []*string
	SetExtNodeNameList(v []*string) *ListJobInfosRequest
	GetExtNodeNameList() []*string
	SetFrom(v int64) *ListJobInfosRequest
	GetFrom() *int64
	SetInstanceIdList(v []*string) *ListJobInfosRequest
	GetInstanceIdList() []*string
	SetJobOwnerList(v []*string) *ListJobInfosRequest
	GetJobOwnerList() []*string
	SetPriorityList(v []*int64) *ListJobInfosRequest
	GetPriorityList() []*int64
	SetProjectList(v []*string) *ListJobInfosRequest
	GetProjectList() []*string
	SetQuotaNickname(v string) *ListJobInfosRequest
	GetQuotaNickname() *string
	SetSceneTagList(v []*string) *ListJobInfosRequest
	GetSceneTagList() []*string
	SetSignatureList(v []*string) *ListJobInfosRequest
	GetSignatureList() []*string
	SetSortByList(v []*string) *ListJobInfosRequest
	GetSortByList() []*string
	SetSortOrderList(v []*string) *ListJobInfosRequest
	GetSortOrderList() []*string
	SetStatusList(v []*string) *ListJobInfosRequest
	GetStatusList() []*string
	SetTaskNameList(v []*string) *ListJobInfosRequest
	GetTaskNameList() []*string
	SetTo(v int64) *ListJobInfosRequest
	GetTo() *int64
	SetTypeList(v []*string) *ListJobInfosRequest
	GetTypeList() []*string
	SetOrderColumn(v string) *ListJobInfosRequest
	GetOrderColumn() *string
	SetPageNumber(v int64) *ListJobInfosRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListJobInfosRequest
	GetPageSize() *int64
	SetRegion(v string) *ListJobInfosRequest
	GetRegion() *string
	SetTenantId(v string) *ListJobInfosRequest
	GetTenantId() *string
}

type ListJobInfosRequest struct {
	// Specifies whether to sort the results in ascending or descending order.
	//
	// example:
	//
	// true
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// The upstream node ID.
	ExtNodeIdList   []*string `json:"extNodeIdList,omitempty" xml:"extNodeIdList,omitempty" type:"Repeated"`
	ExtNodeNameList []*string `json:"extNodeNameList,omitempty" xml:"extNodeNameList,omitempty" type:"Repeated"`
	// The start UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1672112000
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The job instance ID.
	InstanceIdList []*string `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty" type:"Repeated"`
	// The job owner.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The job priority.
	PriorityList []*int64 `json:"priorityList,omitempty" xml:"priorityList,omitempty" type:"Repeated"`
	// The project name.
	ProjectList []*string `json:"projectList,omitempty" xml:"projectList,omitempty" type:"Repeated"`
	// The nickname of the quota.
	//
	// example:
	//
	// quota_nickname
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The smart diagnosis tag.
	SceneTagList []*string `json:"sceneTagList,omitempty" xml:"sceneTagList,omitempty" type:"Repeated"`
	// The job signature.
	SignatureList []*string `json:"signatureList,omitempty" xml:"signatureList,omitempty" type:"Repeated"`
	// The fields for multi-column sorting.
	SortByList []*string `json:"sortByList,omitempty" xml:"sortByList,omitempty" type:"Repeated"`
	// The sort orders for multi-column sorting.
	SortOrderList []*string `json:"sortOrderList,omitempty" xml:"sortOrderList,omitempty" type:"Repeated"`
	// The job status.
	StatusList   []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
	TaskNameList []*string `json:"taskNameList,omitempty" xml:"taskNameList,omitempty" type:"Repeated"`
	// The end UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1672112130
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The job type.
	TypeList []*string `json:"typeList,omitempty" xml:"typeList,omitempty" type:"Repeated"`
	// The column to use for sorting.
	//
	// example:
	//
	// cuUsage
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListJobInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfosRequest) GoString() string {
	return s.String()
}

func (s *ListJobInfosRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *ListJobInfosRequest) GetExtNodeIdList() []*string {
	return s.ExtNodeIdList
}

func (s *ListJobInfosRequest) GetExtNodeNameList() []*string {
	return s.ExtNodeNameList
}

func (s *ListJobInfosRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListJobInfosRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *ListJobInfosRequest) GetJobOwnerList() []*string {
	return s.JobOwnerList
}

func (s *ListJobInfosRequest) GetPriorityList() []*int64 {
	return s.PriorityList
}

func (s *ListJobInfosRequest) GetProjectList() []*string {
	return s.ProjectList
}

func (s *ListJobInfosRequest) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *ListJobInfosRequest) GetSceneTagList() []*string {
	return s.SceneTagList
}

func (s *ListJobInfosRequest) GetSignatureList() []*string {
	return s.SignatureList
}

func (s *ListJobInfosRequest) GetSortByList() []*string {
	return s.SortByList
}

func (s *ListJobInfosRequest) GetSortOrderList() []*string {
	return s.SortOrderList
}

func (s *ListJobInfosRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListJobInfosRequest) GetTaskNameList() []*string {
	return s.TaskNameList
}

func (s *ListJobInfosRequest) GetTo() *int64 {
	return s.To
}

func (s *ListJobInfosRequest) GetTypeList() []*string {
	return s.TypeList
}

func (s *ListJobInfosRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListJobInfosRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListJobInfosRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListJobInfosRequest) GetRegion() *string {
	return s.Region
}

func (s *ListJobInfosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListJobInfosRequest) SetAscOrder(v bool) *ListJobInfosRequest {
	s.AscOrder = &v
	return s
}

func (s *ListJobInfosRequest) SetExtNodeIdList(v []*string) *ListJobInfosRequest {
	s.ExtNodeIdList = v
	return s
}

func (s *ListJobInfosRequest) SetExtNodeNameList(v []*string) *ListJobInfosRequest {
	s.ExtNodeNameList = v
	return s
}

func (s *ListJobInfosRequest) SetFrom(v int64) *ListJobInfosRequest {
	s.From = &v
	return s
}

func (s *ListJobInfosRequest) SetInstanceIdList(v []*string) *ListJobInfosRequest {
	s.InstanceIdList = v
	return s
}

func (s *ListJobInfosRequest) SetJobOwnerList(v []*string) *ListJobInfosRequest {
	s.JobOwnerList = v
	return s
}

func (s *ListJobInfosRequest) SetPriorityList(v []*int64) *ListJobInfosRequest {
	s.PriorityList = v
	return s
}

func (s *ListJobInfosRequest) SetProjectList(v []*string) *ListJobInfosRequest {
	s.ProjectList = v
	return s
}

func (s *ListJobInfosRequest) SetQuotaNickname(v string) *ListJobInfosRequest {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobInfosRequest) SetSceneTagList(v []*string) *ListJobInfosRequest {
	s.SceneTagList = v
	return s
}

func (s *ListJobInfosRequest) SetSignatureList(v []*string) *ListJobInfosRequest {
	s.SignatureList = v
	return s
}

func (s *ListJobInfosRequest) SetSortByList(v []*string) *ListJobInfosRequest {
	s.SortByList = v
	return s
}

func (s *ListJobInfosRequest) SetSortOrderList(v []*string) *ListJobInfosRequest {
	s.SortOrderList = v
	return s
}

func (s *ListJobInfosRequest) SetStatusList(v []*string) *ListJobInfosRequest {
	s.StatusList = v
	return s
}

func (s *ListJobInfosRequest) SetTaskNameList(v []*string) *ListJobInfosRequest {
	s.TaskNameList = v
	return s
}

func (s *ListJobInfosRequest) SetTo(v int64) *ListJobInfosRequest {
	s.To = &v
	return s
}

func (s *ListJobInfosRequest) SetTypeList(v []*string) *ListJobInfosRequest {
	s.TypeList = v
	return s
}

func (s *ListJobInfosRequest) SetOrderColumn(v string) *ListJobInfosRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListJobInfosRequest) SetPageNumber(v int64) *ListJobInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobInfosRequest) SetPageSize(v int64) *ListJobInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobInfosRequest) SetRegion(v string) *ListJobInfosRequest {
	s.Region = &v
	return s
}

func (s *ListJobInfosRequest) SetTenantId(v string) *ListJobInfosRequest {
	s.TenantId = &v
	return s
}

func (s *ListJobInfosRequest) Validate() error {
	return dara.Validate(s)
}
