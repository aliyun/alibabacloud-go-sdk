// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRayJobs(v []*ListRayJobResponseBodyRayJobs) *ListRayJobResponseBody
	GetRayJobs() []*ListRayJobResponseBodyRayJobs
	SetRequestId(v string) *ListRayJobResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRayJobResponseBody
	GetTotalCount() *int32
}

type ListRayJobResponseBody struct {
	RayJobs []*ListRayJobResponseBodyRayJobs `json:"rayJobs,omitempty" xml:"rayJobs,omitempty" type:"Repeated"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRayJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRayJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListRayJobResponseBody) GetRayJobs() []*ListRayJobResponseBodyRayJobs {
	return s.RayJobs
}

func (s *ListRayJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRayJobResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRayJobResponseBody) SetRayJobs(v []*ListRayJobResponseBodyRayJobs) *ListRayJobResponseBody {
	s.RayJobs = v
	return s
}

func (s *ListRayJobResponseBody) SetRequestId(v string) *ListRayJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRayJobResponseBody) SetTotalCount(v int32) *ListRayJobResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRayJobResponseBody) Validate() error {
	if s.RayJobs != nil {
		for _, item := range s.RayJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRayJobResponseBodyRayJobs struct {
	// example:
	//
	// Running
	ClusterState *string `json:"clusterState,omitempty" xml:"clusterState,omitempty"`
	// example:
	//
	// Alice
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 120
	CuHours *float64 `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
	// example:
	//
	// https://emr-ray-gateway.aliyuncs.com?token=xxxxxxxxx
	DashboardUrl *string `json:"dashboardUrl,omitempty" xml:"dashboardUrl,omitempty"`
	// example:
	//
	// 3564
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1776945509000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// testRayJob
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// err-1.2.0 (Ray 2.55.1, Python 3.12)
	ResourceQueue *string `json:"resourceQueue,omitempty" xml:"resourceQueue,omitempty"`
	// example:
	//
	// 1776945499000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// rj-uiulpgow9xljimm1
	SubmissionId *string `json:"submissionId,omitempty" xml:"submissionId,omitempty"`
	// example:
	//
	// 1776945399000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
}

func (s ListRayJobResponseBodyRayJobs) String() string {
	return dara.Prettify(s)
}

func (s ListRayJobResponseBodyRayJobs) GoString() string {
	return s.String()
}

func (s *ListRayJobResponseBodyRayJobs) GetClusterState() *string {
	return s.ClusterState
}

func (s *ListRayJobResponseBodyRayJobs) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListRayJobResponseBodyRayJobs) GetCuHours() *float64 {
	return s.CuHours
}

func (s *ListRayJobResponseBodyRayJobs) GetDashboardUrl() *string {
	return s.DashboardUrl
}

func (s *ListRayJobResponseBodyRayJobs) GetDuration() *int64 {
	return s.Duration
}

func (s *ListRayJobResponseBodyRayJobs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListRayJobResponseBodyRayJobs) GetName() *string {
	return s.Name
}

func (s *ListRayJobResponseBodyRayJobs) GetResourceQueue() *string {
	return s.ResourceQueue
}

func (s *ListRayJobResponseBodyRayJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListRayJobResponseBodyRayJobs) GetStatus() *string {
	return s.Status
}

func (s *ListRayJobResponseBodyRayJobs) GetSubmissionId() *string {
	return s.SubmissionId
}

func (s *ListRayJobResponseBodyRayJobs) GetSubmitTime() *int64 {
	return s.SubmitTime
}

func (s *ListRayJobResponseBodyRayJobs) SetClusterState(v string) *ListRayJobResponseBodyRayJobs {
	s.ClusterState = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetCreatorName(v string) *ListRayJobResponseBodyRayJobs {
	s.CreatorName = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetCuHours(v float64) *ListRayJobResponseBodyRayJobs {
	s.CuHours = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetDashboardUrl(v string) *ListRayJobResponseBodyRayJobs {
	s.DashboardUrl = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetDuration(v int64) *ListRayJobResponseBodyRayJobs {
	s.Duration = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetEndTime(v int64) *ListRayJobResponseBodyRayJobs {
	s.EndTime = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetName(v string) *ListRayJobResponseBodyRayJobs {
	s.Name = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetResourceQueue(v string) *ListRayJobResponseBodyRayJobs {
	s.ResourceQueue = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetStartTime(v int64) *ListRayJobResponseBodyRayJobs {
	s.StartTime = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetStatus(v string) *ListRayJobResponseBodyRayJobs {
	s.Status = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetSubmissionId(v string) *ListRayJobResponseBodyRayJobs {
	s.SubmissionId = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) SetSubmitTime(v int64) *ListRayJobResponseBodyRayJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListRayJobResponseBodyRayJobs) Validate() error {
	return dara.Validate(s)
}
