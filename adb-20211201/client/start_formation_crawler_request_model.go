// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartFormationCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerTaskId(v int64) *StartFormationCrawlerRequest
	GetCrawlerTaskId() *int64
	SetCrawlerTaskName(v string) *StartFormationCrawlerRequest
	GetCrawlerTaskName() *string
	SetDBClusterId(v string) *StartFormationCrawlerRequest
	GetDBClusterId() *string
	SetRegionId(v string) *StartFormationCrawlerRequest
	GetRegionId() *string
}

type StartFormationCrawlerRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 247
	CrawlerTaskId *int64 `json:"CrawlerTaskId,omitempty" xml:"CrawlerTaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// test-crawler-dbname
	CrawlerTaskName *string `json:"CrawlerTaskName,omitempty" xml:"CrawlerTaskName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// am-bp1pke2pcfavw****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartFormationCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s StartFormationCrawlerRequest) GoString() string {
	return s.String()
}

func (s *StartFormationCrawlerRequest) GetCrawlerTaskId() *int64 {
	return s.CrawlerTaskId
}

func (s *StartFormationCrawlerRequest) GetCrawlerTaskName() *string {
	return s.CrawlerTaskName
}

func (s *StartFormationCrawlerRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *StartFormationCrawlerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartFormationCrawlerRequest) SetCrawlerTaskId(v int64) *StartFormationCrawlerRequest {
	s.CrawlerTaskId = &v
	return s
}

func (s *StartFormationCrawlerRequest) SetCrawlerTaskName(v string) *StartFormationCrawlerRequest {
	s.CrawlerTaskName = &v
	return s
}

func (s *StartFormationCrawlerRequest) SetDBClusterId(v string) *StartFormationCrawlerRequest {
	s.DBClusterId = &v
	return s
}

func (s *StartFormationCrawlerRequest) SetRegionId(v string) *StartFormationCrawlerRequest {
	s.RegionId = &v
	return s
}

func (s *StartFormationCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
