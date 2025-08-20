// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkSQLDiagnosisAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeSparkSQLDiagnosisAttributeRequest
	GetAppId() *string
	SetDBClusterId(v string) *DescribeSparkSQLDiagnosisAttributeRequest
	GetDBClusterId() *string
	SetInnerQueryId(v int64) *DescribeSparkSQLDiagnosisAttributeRequest
	GetInnerQueryId() *int64
	SetLanguage(v string) *DescribeSparkSQLDiagnosisAttributeRequest
	GetLanguage() *string
	SetRegionId(v string) *DescribeSparkSQLDiagnosisAttributeRequest
	GetRegionId() *string
}

type DescribeSparkSQLDiagnosisAttributeRequest struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/612475.html) operation to query a list of Spark application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202411071444hzdvk486d9d2001****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The cluster ID.
	//
	// >
	//
	// 	- You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-2zeq4788qyy7k662
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the query executed within the Spark application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	InnerQueryId *int64 `json:"InnerQueryId,omitempty" xml:"InnerQueryId,omitempty"`
	// The language in which to return the query results. Valid values:
	//
	// 	- en: English.
	//
	// 	- zh: Chinese.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSparkSQLDiagnosisAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) GetInnerQueryId() *int64 {
	return s.InnerQueryId
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) SetAppId(v string) *DescribeSparkSQLDiagnosisAttributeRequest {
	s.AppId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) SetDBClusterId(v string) *DescribeSparkSQLDiagnosisAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) SetInnerQueryId(v int64) *DescribeSparkSQLDiagnosisAttributeRequest {
	s.InnerQueryId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) SetLanguage(v string) *DescribeSparkSQLDiagnosisAttributeRequest {
	s.Language = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) SetRegionId(v string) *DescribeSparkSQLDiagnosisAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeRequest) Validate() error {
	return dara.Validate(s)
}
