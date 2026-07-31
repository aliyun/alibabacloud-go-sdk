// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFormationCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerInfo(v string) *CreateFormationCrawlerRequest
	GetCrawlerInfo() *string
	SetDBClusterId(v string) *CreateFormationCrawlerRequest
	GetDBClusterId() *string
	SetRegionId(v string) *CreateFormationCrawlerRequest
	GetRegionId() *string
}

type CreateFormationCrawlerRequest struct {
	// The JSON string that contains the complete crawler configuration. This is the most important parameter. For the internal JSON structure, see the CrawlerInfo structure definition section.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "sourceType": "OSSWAREHOUSE",
	//
	//   "dbName": "your_target_db",
	//
	//   "sourceInfo": {
	//
	//     "ossSourceInfo": {
	//
	//       "sourceMode": "WAREHOUSE",
	//
	//       "ossLocations": ["oss://your-bucket/your-path/"],
	//
	//       "exclusions": [],
	//
	//       "inclusions": []
	//
	//     }
	//
	//   },
	//
	//   "classifiers": ["csv"],
	//
	//   "schemaChangePolicy": {
	//
	//     "updateRule": "ONLY_ADD_COLUMN",
	//
	//     "deleteRule": "IGNORE"
	//
	//   },
	//
	//   "frequency": {
	//
	//     "type": "monthly",
	//
	//     "cron": "0+00+00+1+*+?+*"
	//
	//   },
	//
	//   "configuration": "adb.crawler.csv.columns.specify.delimiter.char=auto\\nadb.crawler.csv.columns.specify.quote.char=auto\\n"
	//
	// }
	CrawlerInfo *string `json:"CrawlerInfo,omitempty" xml:"CrawlerInfo,omitempty"`
	// The ADB instance ID. This specifies the resource-level scope of the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateFormationCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFormationCrawlerRequest) GoString() string {
	return s.String()
}

func (s *CreateFormationCrawlerRequest) GetCrawlerInfo() *string {
	return s.CrawlerInfo
}

func (s *CreateFormationCrawlerRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateFormationCrawlerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFormationCrawlerRequest) SetCrawlerInfo(v string) *CreateFormationCrawlerRequest {
	s.CrawlerInfo = &v
	return s
}

func (s *CreateFormationCrawlerRequest) SetDBClusterId(v string) *CreateFormationCrawlerRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateFormationCrawlerRequest) SetRegionId(v string) *CreateFormationCrawlerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFormationCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
