// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResultExportConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeResultExportConfigRequest
	GetDBClusterId() *string
	SetExportType(v string) *DescribeResultExportConfigRequest
	GetExportType() *string
	SetRegionId(v string) *DescribeResultExportConfigRequest
	GetRegionId() *string
}

type DescribeResultExportConfigRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vbf80pjcz*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The export type. Valid values:
	//
	// 	- SLS: Indicates that the export destination is SLS.
	//
	// 	- OSS: Indicates that the export destination is OSS.
	//
	// example:
	//
	// SLS
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeResultExportConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultExportConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeResultExportConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeResultExportConfigRequest) GetExportType() *string {
	return s.ExportType
}

func (s *DescribeResultExportConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResultExportConfigRequest) SetDBClusterId(v string) *DescribeResultExportConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeResultExportConfigRequest) SetExportType(v string) *DescribeResultExportConfigRequest {
	s.ExportType = &v
	return s
}

func (s *DescribeResultExportConfigRequest) SetRegionId(v string) *DescribeResultExportConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResultExportConfigRequest) Validate() error {
	return dara.Validate(s)
}
