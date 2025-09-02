// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupPagingResult(v *ListDataServiceGroupsResponseBodyGroupPagingResult) *ListDataServiceGroupsResponseBody
	GetGroupPagingResult() *ListDataServiceGroupsResponseBodyGroupPagingResult
	SetRequestId(v string) *ListDataServiceGroupsResponseBody
	GetRequestId() *string
}

type ListDataServiceGroupsResponseBody struct {
	// The paging result for the business processes.
	GroupPagingResult *ListDataServiceGroupsResponseBodyGroupPagingResult `json:"GroupPagingResult,omitempty" xml:"GroupPagingResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataServiceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceGroupsResponseBody) GetGroupPagingResult() *ListDataServiceGroupsResponseBodyGroupPagingResult {
	return s.GroupPagingResult
}

func (s *ListDataServiceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceGroupsResponseBody) SetGroupPagingResult(v *ListDataServiceGroupsResponseBodyGroupPagingResult) *ListDataServiceGroupsResponseBody {
	s.GroupPagingResult = v
	return s
}

func (s *ListDataServiceGroupsResponseBody) SetRequestId(v string) *ListDataServiceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceGroupsResponseBodyGroupPagingResult struct {
	// The business processes.
	Groups []*ListDataServiceGroupsResponseBodyGroupPagingResultGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The page number. The value of this parameter is the same as that of the PageNumber parameter in the request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceGroupsResponseBodyGroupPagingResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceGroupsResponseBodyGroupPagingResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) GetGroups() []*ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	return s.Groups
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) SetGroups(v []*ListDataServiceGroupsResponseBodyGroupPagingResultGroups) *ListDataServiceGroupsResponseBodyGroupPagingResult {
	s.Groups = v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) SetPageNumber(v int32) *ListDataServiceGroupsResponseBodyGroupPagingResult {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) SetPageSize(v int32) *ListDataServiceGroupsResponseBodyGroupPagingResult {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) SetTotalCount(v int32) *ListDataServiceGroupsResponseBodyGroupPagingResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResult) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceGroupsResponseBodyGroupPagingResultGroups struct {
	// The ID of the API Gateway group to which the workflow is bound.
	//
	// example:
	//
	// 100abc
	ApiGatewayGroupId *string `json:"ApiGatewayGroupId,omitempty" xml:"ApiGatewayGroupId,omitempty"`
	// The time when the business process was created.
	//
	// example:
	//
	// 2020-09-24T18:37:51+0800
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The user identifier (UID) of the creator of the business process. The value of this parameter may be empty for creators of some existing business processes.
	//
	// example:
	//
	// 10001
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the business process.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The business process ID.
	//
	// example:
	//
	// ds_123abc
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the business process.
	//
	// example:
	//
	// Test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the business process was modified.
	//
	// example:
	//
	// 2020-09-24T18:37:51+0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10003
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceGroupsResponseBodyGroupPagingResultGroups) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GoString() string {
	return s.String()
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetApiGatewayGroupId() *string {
	return s.ApiGatewayGroupId
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetDescription() *string {
	return s.Description
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetApiGatewayGroupId(v string) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.ApiGatewayGroupId = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetCreatedTime(v string) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.CreatedTime = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetCreatorId(v string) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.CreatorId = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetDescription(v string) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.Description = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetGroupId(v string) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.GroupId = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetGroupName(v string) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.GroupName = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetModifiedTime(v string) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.ModifiedTime = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetProjectId(v int64) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) SetTenantId(v int64) *ListDataServiceGroupsResponseBodyGroupPagingResultGroups {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceGroupsResponseBodyGroupPagingResultGroups) Validate() error {
	return dara.Validate(s)
}
