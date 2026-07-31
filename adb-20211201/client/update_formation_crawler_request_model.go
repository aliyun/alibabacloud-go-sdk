// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormationCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerInfo(v string) *UpdateFormationCrawlerRequest
	GetCrawlerInfo() *string
	SetCrawlerTaskId(v string) *UpdateFormationCrawlerRequest
	GetCrawlerTaskId() *string
	SetDBClusterId(v string) *UpdateFormationCrawlerRequest
	GetDBClusterId() *string
	SetRegionId(v string) *UpdateFormationCrawlerRequest
	GetRegionId() *string
}

type UpdateFormationCrawlerRequest struct {
	// The JSON string that contains the complete configuration of the crawler. This is the most important parameter. For more information about the internal JSON structure, see the CrawlerInfo structure definition section.
	//
	// This parameter is required.
	CrawlerInfo *string `json:"CrawlerInfo,omitempty" xml:"CrawlerInfo,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 21
	CrawlerTaskId *string `json:"CrawlerTaskId,omitempty" xml:"CrawlerTaskId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateFormationCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormationCrawlerRequest) GoString() string {
	return s.String()
}

func (s *UpdateFormationCrawlerRequest) GetCrawlerInfo() *string {
	return s.CrawlerInfo
}

func (s *UpdateFormationCrawlerRequest) GetCrawlerTaskId() *string {
	return s.CrawlerTaskId
}

func (s *UpdateFormationCrawlerRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateFormationCrawlerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateFormationCrawlerRequest) SetCrawlerInfo(v string) *UpdateFormationCrawlerRequest {
	s.CrawlerInfo = &v
	return s
}

func (s *UpdateFormationCrawlerRequest) SetCrawlerTaskId(v string) *UpdateFormationCrawlerRequest {
	s.CrawlerTaskId = &v
	return s
}

func (s *UpdateFormationCrawlerRequest) SetDBClusterId(v string) *UpdateFormationCrawlerRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateFormationCrawlerRequest) SetRegionId(v string) *UpdateFormationCrawlerRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateFormationCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
