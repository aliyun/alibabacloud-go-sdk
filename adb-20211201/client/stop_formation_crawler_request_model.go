// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopFormationCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerTaskId(v int64) *StopFormationCrawlerRequest
	GetCrawlerTaskId() *int64
	SetCrawlerTaskName(v string) *StopFormationCrawlerRequest
	GetCrawlerTaskName() *string
	SetDBClusterId(v string) *StopFormationCrawlerRequest
	GetDBClusterId() *string
	SetRegionId(v string) *StopFormationCrawlerRequest
	GetRegionId() *string
}

type StopFormationCrawlerRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21
	CrawlerTaskId *int64 `json:"CrawlerTaskId,omitempty" xml:"CrawlerTaskId,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-crawler-dbname
	CrawlerTaskName *string `json:"CrawlerTaskName,omitempty" xml:"CrawlerTaskName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopFormationCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s StopFormationCrawlerRequest) GoString() string {
	return s.String()
}

func (s *StopFormationCrawlerRequest) GetCrawlerTaskId() *int64 {
	return s.CrawlerTaskId
}

func (s *StopFormationCrawlerRequest) GetCrawlerTaskName() *string {
	return s.CrawlerTaskName
}

func (s *StopFormationCrawlerRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *StopFormationCrawlerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopFormationCrawlerRequest) SetCrawlerTaskId(v int64) *StopFormationCrawlerRequest {
	s.CrawlerTaskId = &v
	return s
}

func (s *StopFormationCrawlerRequest) SetCrawlerTaskName(v string) *StopFormationCrawlerRequest {
	s.CrawlerTaskName = &v
	return s
}

func (s *StopFormationCrawlerRequest) SetDBClusterId(v string) *StopFormationCrawlerRequest {
	s.DBClusterId = &v
	return s
}

func (s *StopFormationCrawlerRequest) SetRegionId(v string) *StopFormationCrawlerRequest {
	s.RegionId = &v
	return s
}

func (s *StopFormationCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
