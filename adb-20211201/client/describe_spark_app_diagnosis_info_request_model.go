// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAppDiagnosisInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeSparkAppDiagnosisInfoRequest
	GetAppId() *string
	SetDBClusterId(v string) *DescribeSparkAppDiagnosisInfoRequest
	GetDBClusterId() *string
	SetLanguage(v string) *DescribeSparkAppDiagnosisInfoRequest
	GetLanguage() *string
	SetRegionId(v string) *DescribeSparkAppDiagnosisInfoRequest
	GetRegionId() *string
}

type DescribeSparkAppDiagnosisInfoRequest struct {
	// The application ID.
	//
	// > You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query all application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202411061017sh0ad564b000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view the information of all clusters in the specified region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf663j39b0jd2***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of the response. Valid values:
	//
	// - en: English.
	//
	// - zh: Chinese. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to view the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSparkAppDiagnosisInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAppDiagnosisInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkAppDiagnosisInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSparkAppDiagnosisInfoRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkAppDiagnosisInfoRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeSparkAppDiagnosisInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkAppDiagnosisInfoRequest) SetAppId(v string) *DescribeSparkAppDiagnosisInfoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoRequest) SetDBClusterId(v string) *DescribeSparkAppDiagnosisInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoRequest) SetLanguage(v string) *DescribeSparkAppDiagnosisInfoRequest {
	s.Language = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoRequest) SetRegionId(v string) *DescribeSparkAppDiagnosisInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoRequest) Validate() error {
	return dara.Validate(s)
}
