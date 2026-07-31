// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormationCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerTaskId(v int64) *DeleteFormationCrawlerRequest
	GetCrawlerTaskId() *int64
	SetCrawlerTaskName(v string) *DeleteFormationCrawlerRequest
	GetCrawlerTaskName() *string
	SetDBClusterId(v string) *DeleteFormationCrawlerRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DeleteFormationCrawlerRequest
	GetRegionId() *string
}

type DeleteFormationCrawlerRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// wz_log
	CrawlerTaskName *string `json:"CrawlerTaskName,omitempty" xml:"CrawlerTaskName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vbc***
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

func (s DeleteFormationCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormationCrawlerRequest) GoString() string {
	return s.String()
}

func (s *DeleteFormationCrawlerRequest) GetCrawlerTaskId() *int64 {
	return s.CrawlerTaskId
}

func (s *DeleteFormationCrawlerRequest) GetCrawlerTaskName() *string {
	return s.CrawlerTaskName
}

func (s *DeleteFormationCrawlerRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteFormationCrawlerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFormationCrawlerRequest) SetCrawlerTaskId(v int64) *DeleteFormationCrawlerRequest {
	s.CrawlerTaskId = &v
	return s
}

func (s *DeleteFormationCrawlerRequest) SetCrawlerTaskName(v string) *DeleteFormationCrawlerRequest {
	s.CrawlerTaskName = &v
	return s
}

func (s *DeleteFormationCrawlerRequest) SetDBClusterId(v string) *DeleteFormationCrawlerRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteFormationCrawlerRequest) SetRegionId(v string) *DeleteFormationCrawlerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFormationCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
