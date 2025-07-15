// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFotaTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFotaStatus(v string) *DescribeFotaTasksRequest
	GetFotaStatus() *string
	SetLang(v string) *DescribeFotaTasksRequest
	GetLang() *string
	SetMaxResults(v int32) *DescribeFotaTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFotaTasksRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeFotaTasksRequest
	GetRegionId() *string
	SetTaskUid(v []*string) *DescribeFotaTasksRequest
	GetTaskUid() []*string
	SetUserStatus(v string) *DescribeFotaTasksRequest
	GetUserStatus() *string
}

type DescribeFotaTasksRequest struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden.
	FotaStatus *string `json:"FotaStatus,omitempty" xml:"FotaStatus,omitempty"`
	// The language of the image version to update.
	//
	// Valid values:
	//
	// 	- en: English.
	//
	// 	- zh: Simplified Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: 1 to 100
	//
	// 	- Default value: 20
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the NextToken parameter is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the image update tasks.
	TaskUid []*string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty" type:"Repeated"`
	// Specifies whether to automatically push the image update task.
	//
	// Valid values:
	//
	// 	- Running: automatically pushes the image update task.
	//
	// 	- Pending: does not automatically push the image update task.
	//
	// example:
	//
	// Pending
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeFotaTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksRequest) GetFotaStatus() *string {
	return s.FotaStatus
}

func (s *DescribeFotaTasksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFotaTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFotaTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFotaTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFotaTasksRequest) GetTaskUid() []*string {
	return s.TaskUid
}

func (s *DescribeFotaTasksRequest) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeFotaTasksRequest) SetFotaStatus(v string) *DescribeFotaTasksRequest {
	s.FotaStatus = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetLang(v string) *DescribeFotaTasksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetMaxResults(v int32) *DescribeFotaTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetNextToken(v string) *DescribeFotaTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetRegionId(v string) *DescribeFotaTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFotaTasksRequest) SetTaskUid(v []*string) *DescribeFotaTasksRequest {
	s.TaskUid = v
	return s
}

func (s *DescribeFotaTasksRequest) SetUserStatus(v string) *DescribeFotaTasksRequest {
	s.UserStatus = &v
	return s
}

func (s *DescribeFotaTasksRequest) Validate() error {
	return dara.Validate(s)
}
