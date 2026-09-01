// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailoverTestJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailoverTestJobModel(v *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) *DescribeFailoverTestJobResponseBody
	GetFailoverTestJobModel() *DescribeFailoverTestJobResponseBodyFailoverTestJobModel
	SetRequestId(v string) *DescribeFailoverTestJobResponseBody
	GetRequestId() *string
}

type DescribeFailoverTestJobResponseBody struct {
	// The failover test job.
	FailoverTestJobModel *DescribeFailoverTestJobResponseBodyFailoverTestJobModel `json:"FailoverTestJobModel,omitempty" xml:"FailoverTestJobModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFailoverTestJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobResponseBody) GetFailoverTestJobModel() *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	return s.FailoverTestJobModel
}

func (s *DescribeFailoverTestJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFailoverTestJobResponseBody) SetFailoverTestJobModel(v *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) *DescribeFailoverTestJobResponseBody {
	s.FailoverTestJobModel = v
	return s
}

func (s *DescribeFailoverTestJobResponseBody) SetRequestId(v string) *DescribeFailoverTestJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBody) Validate() error {
	if s.FailoverTestJobModel != nil {
		if err := s.FailoverTestJobModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFailoverTestJobResponseBodyFailoverTestJobModel struct {
	// The description of the failover test job.
	//
	// The description is 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The failover test duration. Unit: minutes. Valid values: **1 to 4320**.
	//
	// example:
	//
	// 60
	JobDuration *string `json:"JobDuration,omitempty" xml:"JobDuration,omitempty"`
	// The failover test job ID.
	//
	// example:
	//
	// ftj-bp1yh6mvi13aq3g8w****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The failover test type. Valid values:
	//
	// - **StartNow**: start immediately. The test starts immediately after the job is created.
	//
	// - **StartLater**: start later. Only the test job is created without starting the test.
	//
	// example:
	//
	// StartNow
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the failover test job.
	//
	// The name is 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of failover test resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The failover test resource type. Valid values: **PHYSICALCONNECTION**: Express Connect circuit.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time of the failover test job. The time is displayed in UTC in the YYYY-MM-DDThh:mm:ssZ format as defined by ISO 8601.
	//
	// example:
	//
	// 2023-11-21T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the failover test job. Valid values:
	//
	// - **Init**: pending.
	//
	// - **Starting**: starting.
	//
	// - **Testing**: in progress.
	//
	// - **Stopping**: stopping.
	//
	// - **Stopped**: completed.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The end time of the failover test job. The time is displayed in UTC in the YYYY-MM-DDThh:mm:ssZ format as defined by ISO 8601.
	//
	// example:
	//
	// 2023-11-21T15:00:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s DescribeFailoverTestJobResponseBodyFailoverTestJobModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetDescription() *string {
	return s.Description
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetJobDuration() *string {
	return s.JobDuration
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetJobId() *string {
	return s.JobId
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetJobType() *string {
	return s.JobType
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetName() *string {
	return s.Name
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetStatus() *string {
	return s.Status
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetDescription(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.Description = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetJobDuration(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.JobDuration = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetJobId(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.JobId = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetJobType(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.JobType = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetName(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.Name = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetResourceId(v []*string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.ResourceId = v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetResourceType(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.ResourceType = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetStartTime(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.StartTime = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetStatus(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.Status = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) SetStopTime(v string) *DescribeFailoverTestJobResponseBodyFailoverTestJobModel {
	s.StopTime = &v
	return s
}

func (s *DescribeFailoverTestJobResponseBodyFailoverTestJobModel) Validate() error {
	return dara.Validate(s)
}
