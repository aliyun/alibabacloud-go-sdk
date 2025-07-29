// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataTime(v string) *RerunJobRequest
	GetDataTime() *string
	SetEndDate(v int64) *RerunJobRequest
	GetEndDate() *int64
	SetGroupId(v string) *RerunJobRequest
	GetGroupId() *string
	SetJobId(v int64) *RerunJobRequest
	GetJobId() *int64
	SetNamespace(v string) *RerunJobRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *RerunJobRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *RerunJobRequest
	GetRegionId() *string
	SetStartDate(v int64) *RerunJobRequest
	GetStartDate() *int64
}

type RerunJobRequest struct {
	// The data timestamp of the job. Specify a string in the HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the job stops running. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645718400000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the job starts to rerun. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645459200000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s RerunJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunJobRequest) GoString() string {
	return s.String()
}

func (s *RerunJobRequest) GetDataTime() *string {
	return s.DataTime
}

func (s *RerunJobRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *RerunJobRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RerunJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *RerunJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *RerunJobRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *RerunJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RerunJobRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *RerunJobRequest) SetDataTime(v string) *RerunJobRequest {
	s.DataTime = &v
	return s
}

func (s *RerunJobRequest) SetEndDate(v int64) *RerunJobRequest {
	s.EndDate = &v
	return s
}

func (s *RerunJobRequest) SetGroupId(v string) *RerunJobRequest {
	s.GroupId = &v
	return s
}

func (s *RerunJobRequest) SetJobId(v int64) *RerunJobRequest {
	s.JobId = &v
	return s
}

func (s *RerunJobRequest) SetNamespace(v string) *RerunJobRequest {
	s.Namespace = &v
	return s
}

func (s *RerunJobRequest) SetNamespaceSource(v string) *RerunJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *RerunJobRequest) SetRegionId(v string) *RerunJobRequest {
	s.RegionId = &v
	return s
}

func (s *RerunJobRequest) SetStartDate(v int64) *RerunJobRequest {
	s.StartDate = &v
	return s
}

func (s *RerunJobRequest) Validate() error {
	return dara.Validate(s)
}
