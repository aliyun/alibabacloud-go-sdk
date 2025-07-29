// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *GetJobInstanceListRequest
	GetEndTimestamp() *int64
	SetGroupId(v string) *GetJobInstanceListRequest
	GetGroupId() *string
	SetJobId(v int64) *GetJobInstanceListRequest
	GetJobId() *int64
	SetNamespace(v string) *GetJobInstanceListRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetJobInstanceListRequest
	GetNamespaceSource() *string
	SetPageNum(v int32) *GetJobInstanceListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetJobInstanceListRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetJobInstanceListRequest
	GetRegionId() *string
	SetStartTimestamp(v int64) *GetJobInstanceListRequest
	GetStartTimestamp() *int64
	SetStatus(v int32) *GetJobInstanceListRequest
	GetStatus() *int32
}

type GetJobInstanceListRequest struct {
	// The end of the time range to query. Specify a UNIX timestamp.
	//
	// example:
	//
	// 1684202400000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The application group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp.
	//
	// example:
	//
	// 1684116000000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The status of the job instance. Valid values:
	//
	// 1: The job instance is pending. 3: The job instance is running. 4: The job instance is run. 5: The job instance fails. 9: The request for running the job instance is rejected. To specify this parameter, you must declare the following enumeration class: com.alibaba.schedulerx.common.domain.InstanceStatus.
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *GetJobInstanceListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetJobInstanceListRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *GetJobInstanceListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetJobInstanceListRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetJobInstanceListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetJobInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetJobInstanceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetJobInstanceListRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *GetJobInstanceListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetJobInstanceListRequest) SetEndTimestamp(v int64) *GetJobInstanceListRequest {
	s.EndTimestamp = &v
	return s
}

func (s *GetJobInstanceListRequest) SetGroupId(v string) *GetJobInstanceListRequest {
	s.GroupId = &v
	return s
}

func (s *GetJobInstanceListRequest) SetJobId(v int64) *GetJobInstanceListRequest {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceListRequest) SetNamespace(v string) *GetJobInstanceListRequest {
	s.Namespace = &v
	return s
}

func (s *GetJobInstanceListRequest) SetNamespaceSource(v string) *GetJobInstanceListRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetJobInstanceListRequest) SetPageNum(v int32) *GetJobInstanceListRequest {
	s.PageNum = &v
	return s
}

func (s *GetJobInstanceListRequest) SetPageSize(v int32) *GetJobInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobInstanceListRequest) SetRegionId(v string) *GetJobInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobInstanceListRequest) SetStartTimestamp(v int64) *GetJobInstanceListRequest {
	s.StartTimestamp = &v
	return s
}

func (s *GetJobInstanceListRequest) SetStatus(v int32) *GetJobInstanceListRequest {
	s.Status = &v
	return s
}

func (s *GetJobInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
