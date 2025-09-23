// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCronJobPolicyServerlessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateCronJobPolicyServerlessResponseBody
	GetAction() *string
	SetCronExpression(v string) *CreateCronJobPolicyServerlessResponseBody
	GetCronExpression() *string
	SetDBClusterId(v string) *CreateCronJobPolicyServerlessResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *CreateCronJobPolicyServerlessResponseBody
	GetEndTime() *string
	SetJobId(v string) *CreateCronJobPolicyServerlessResponseBody
	GetJobId() *string
	SetRegionId(v string) *CreateCronJobPolicyServerlessResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateCronJobPolicyServerlessResponseBody
	GetRequestId() *string
	SetStartTime(v string) *CreateCronJobPolicyServerlessResponseBody
	GetStartTime() *string
	SetStatus(v string) *CreateCronJobPolicyServerlessResponseBody
	GetStatus() *string
}

type CreateCronJobPolicyServerlessResponseBody struct {
	// example:
	//
	// ModifyDBClusterServerlessConf
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// 0 0 8 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2021-04-07T04:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1fa3c0e7-b568-4f41-b703-463c96a91bd8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 45D24263-7E3A-4140-9472-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2020-05-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateCronJobPolicyServerlessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCronJobPolicyServerlessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetAction() *string {
	return s.Action
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCronJobPolicyServerlessResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetAction(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.Action = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetCronExpression(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.CronExpression = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetDBClusterId(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetEndTime(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetJobId(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetRegionId(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetRequestId(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetStartTime(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) SetStatus(v string) *CreateCronJobPolicyServerlessResponseBody {
	s.Status = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponseBody) Validate() error {
	return dara.Validate(s)
}
