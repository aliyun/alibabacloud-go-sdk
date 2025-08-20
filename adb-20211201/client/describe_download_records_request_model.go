// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDownloadRecordsRequest
	GetDBClusterId() *string
	SetLang(v string) *DescribeDownloadRecordsRequest
	GetLang() *string
	SetRegionId(v string) *DescribeDownloadRecordsRequest
	GetRegionId() *string
}

type DescribeDownloadRecordsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vb6ha79k6e****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of the returned data. Valid values:
	//
	// 	- **zh**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// 	- **ja**: Japanese.
	//
	// 	- **zh-tw**: traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDownloadRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDownloadRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDownloadRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDownloadRecordsRequest) SetDBClusterId(v string) *DescribeDownloadRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetLang(v string) *DescribeDownloadRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetRegionId(v string) *DescribeDownloadRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) Validate() error {
	return dara.Validate(s)
}
