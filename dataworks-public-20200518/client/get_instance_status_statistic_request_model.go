// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStatusStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizDate(v string) *GetInstanceStatusStatisticRequest
	GetBizDate() *string
	SetDagType(v string) *GetInstanceStatusStatisticRequest
	GetDagType() *string
	SetProjectEnv(v string) *GetInstanceStatusStatisticRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *GetInstanceStatusStatisticRequest
	GetProjectId() *int64
	SetSchedulerPeriod(v string) *GetInstanceStatusStatisticRequest
	GetSchedulerPeriod() *string
	SetSchedulerType(v string) *GetInstanceStatusStatisticRequest
	GetSchedulerType() *string
}

type GetInstanceStatusStatisticRequest struct {
	// The date on which the numbers of instances in different states are obtained. Specify the date in the yyyy-MM-dd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-01-01
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The type of the directed acyclic graph (DAG). Valid values:
	//
	// 	- MANUAL: DAG for a manually triggered workflow
	//
	// 	- SMOKE_TEST: DAG for a smoke testing workflow
	//
	// 	- SUPPLY_DATA: DAG for a data backfill instance
	//
	// 	- BUSINESS_PROCESS_DAG: DAG for a one-time workflow
	//
	// <!---->
	//
	// 	- DAILY
	//
	// 	- MANUAL
	//
	// 	- SMOKE_TEST
	//
	// 	- SUPPLY_DATA
	//
	// 	- BUSINESS_PROCESS_DAG
	//
	// example:
	//
	// MANUAL
	DagType *string `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// The runtime environment. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling cycle. Valid values:
	//
	// 	- MINUTE
	//
	// 	- HOUR
	//
	// 	- DAY
	//
	// 	- WEEK
	//
	// 	- MONTH
	//
	// example:
	//
	// DAY
	SchedulerPeriod *string `json:"SchedulerPeriod,omitempty" xml:"SchedulerPeriod,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL: auto triggered node
	//
	// 	- MANUAL: manually triggered node
	//
	// 	- PAUSE: paused node
	//
	// 	- SKIP: dry-run node
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s GetInstanceStatusStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusStatisticRequest) GetBizDate() *string {
	return s.BizDate
}

func (s *GetInstanceStatusStatisticRequest) GetDagType() *string {
	return s.DagType
}

func (s *GetInstanceStatusStatisticRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetInstanceStatusStatisticRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceStatusStatisticRequest) GetSchedulerPeriod() *string {
	return s.SchedulerPeriod
}

func (s *GetInstanceStatusStatisticRequest) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetInstanceStatusStatisticRequest) SetBizDate(v string) *GetInstanceStatusStatisticRequest {
	s.BizDate = &v
	return s
}

func (s *GetInstanceStatusStatisticRequest) SetDagType(v string) *GetInstanceStatusStatisticRequest {
	s.DagType = &v
	return s
}

func (s *GetInstanceStatusStatisticRequest) SetProjectEnv(v string) *GetInstanceStatusStatisticRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetInstanceStatusStatisticRequest) SetProjectId(v int64) *GetInstanceStatusStatisticRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceStatusStatisticRequest) SetSchedulerPeriod(v string) *GetInstanceStatusStatisticRequest {
	s.SchedulerPeriod = &v
	return s
}

func (s *GetInstanceStatusStatisticRequest) SetSchedulerType(v string) *GetInstanceStatusStatisticRequest {
	s.SchedulerType = &v
	return s
}

func (s *GetInstanceStatusStatisticRequest) Validate() error {
	return dara.Validate(s)
}
