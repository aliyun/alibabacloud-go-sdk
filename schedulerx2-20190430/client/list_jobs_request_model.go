// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListJobsRequest
	GetGroupId() *string
	SetJobName(v string) *ListJobsRequest
	GetJobName() *string
	SetNamespace(v string) *ListJobsRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ListJobsRequest
	GetNamespaceSource() *string
	SetPageNum(v int32) *ListJobsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListJobsRequest
	GetRegionId() *string
	SetStatus(v string) *ListJobsRequest
	GetStatus() *string
}

type ListJobsRequest struct {
	// The application ID. You can obtain the ID on the **Application Management*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// DocTest.Group
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The node name.
	//
	// example:
	//
	// helloword
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The namespace. You can obtain the namespace on the **Namespace*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Required only for special third-party users.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The page number.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of records per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The node status.
	//
	// - **0**: disabled
	//
	// - **1**: enabled
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListJobsRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListJobsRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ListJobsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobsRequest) SetGroupId(v string) *ListJobsRequest {
	s.GroupId = &v
	return s
}

func (s *ListJobsRequest) SetJobName(v string) *ListJobsRequest {
	s.JobName = &v
	return s
}

func (s *ListJobsRequest) SetNamespace(v string) *ListJobsRequest {
	s.Namespace = &v
	return s
}

func (s *ListJobsRequest) SetNamespaceSource(v string) *ListJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListJobsRequest) SetPageNum(v int32) *ListJobsRequest {
	s.PageNum = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetRegionId(v string) *ListJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
