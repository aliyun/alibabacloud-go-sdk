// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormationCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerTaskId(v int64) *GetFormationCrawlerRequest
	GetCrawlerTaskId() *int64
	SetCrawlerTaskName(v string) *GetFormationCrawlerRequest
	GetCrawlerTaskName() *string
	SetDBClusterId(v string) *GetFormationCrawlerRequest
	GetDBClusterId() *string
	SetRegionId(v string) *GetFormationCrawlerRequest
	GetRegionId() *string
}

type GetFormationCrawlerRequest struct {
	// The task ID.
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
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the supported regions and zones, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetFormationCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFormationCrawlerRequest) GoString() string {
	return s.String()
}

func (s *GetFormationCrawlerRequest) GetCrawlerTaskId() *int64 {
	return s.CrawlerTaskId
}

func (s *GetFormationCrawlerRequest) GetCrawlerTaskName() *string {
	return s.CrawlerTaskName
}

func (s *GetFormationCrawlerRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetFormationCrawlerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetFormationCrawlerRequest) SetCrawlerTaskId(v int64) *GetFormationCrawlerRequest {
	s.CrawlerTaskId = &v
	return s
}

func (s *GetFormationCrawlerRequest) SetCrawlerTaskName(v string) *GetFormationCrawlerRequest {
	s.CrawlerTaskName = &v
	return s
}

func (s *GetFormationCrawlerRequest) SetDBClusterId(v string) *GetFormationCrawlerRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetFormationCrawlerRequest) SetRegionId(v string) *GetFormationCrawlerRequest {
	s.RegionId = &v
	return s
}

func (s *GetFormationCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
