// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormationCrawlerScheduleStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerTaskId(v int64) *UpdateFormationCrawlerScheduleStateRequest
	GetCrawlerTaskId() *int64
	SetCrawlerTaskName(v string) *UpdateFormationCrawlerScheduleStateRequest
	GetCrawlerTaskName() *string
	SetDBClusterId(v string) *UpdateFormationCrawlerScheduleStateRequest
	GetDBClusterId() *string
	SetRegionId(v string) *UpdateFormationCrawlerScheduleStateRequest
	GetRegionId() *string
	SetScheduleState(v string) *UpdateFormationCrawlerScheduleStateRequest
	GetScheduleState() *string
}

type UpdateFormationCrawlerScheduleStateRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 247
	CrawlerTaskId *int64 `json:"CrawlerTaskId,omitempty" xml:"CrawlerTaskId,omitempty"`
	// The name of the crawler task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-crawler-dbname
	CrawlerTaskName *string `json:"CrawlerTaskName,omitempty" xml:"CrawlerTaskName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1565u55p32****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling state. Valid values:
	//
	// - NORMAL: resume.
	//
	// - DISABLED: pause.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISABLED
	ScheduleState *string `json:"ScheduleState,omitempty" xml:"ScheduleState,omitempty"`
}

func (s UpdateFormationCrawlerScheduleStateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormationCrawlerScheduleStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateFormationCrawlerScheduleStateRequest) GetCrawlerTaskId() *int64 {
	return s.CrawlerTaskId
}

func (s *UpdateFormationCrawlerScheduleStateRequest) GetCrawlerTaskName() *string {
	return s.CrawlerTaskName
}

func (s *UpdateFormationCrawlerScheduleStateRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateFormationCrawlerScheduleStateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateFormationCrawlerScheduleStateRequest) GetScheduleState() *string {
	return s.ScheduleState
}

func (s *UpdateFormationCrawlerScheduleStateRequest) SetCrawlerTaskId(v int64) *UpdateFormationCrawlerScheduleStateRequest {
	s.CrawlerTaskId = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateRequest) SetCrawlerTaskName(v string) *UpdateFormationCrawlerScheduleStateRequest {
	s.CrawlerTaskName = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateRequest) SetDBClusterId(v string) *UpdateFormationCrawlerScheduleStateRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateRequest) SetRegionId(v string) *UpdateFormationCrawlerScheduleStateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateRequest) SetScheduleState(v string) *UpdateFormationCrawlerScheduleStateRequest {
	s.ScheduleState = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateRequest) Validate() error {
	return dara.Validate(s)
}
