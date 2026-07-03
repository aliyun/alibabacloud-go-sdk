// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *AddDataSourceRequest
	GetAccountId() *string
	SetCloudCode(v string) *AddDataSourceRequest
	GetCloudCode() *string
	SetDataSourceInstanceName(v string) *AddDataSourceRequest
	GetDataSourceInstanceName() *string
	SetDataSourceInstanceParams(v string) *AddDataSourceRequest
	GetDataSourceInstanceParams() *string
	SetDataSourceInstanceRemark(v string) *AddDataSourceRequest
	GetDataSourceInstanceRemark() *string
	SetDataSourceType(v string) *AddDataSourceRequest
	GetDataSourceType() *string
	SetRegionId(v string) *AddDataSourceRequest
	GetRegionId() *string
}

type AddDataSourceRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 123xxxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code for the multicloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// XX-Beijing-Kafka
	DataSourceInstanceName *string `json:"DataSourceInstanceName,omitempty" xml:"DataSourceInstanceName,omitempty"`
	// The parameters for the data source. The value must be a JSON array.
	//
	// example:
	//
	// [{"paraCode":"region_code","paraValue":"ap-guangzhou"}]
	DataSourceInstanceParams *string `json:"DataSourceInstanceParams,omitempty" xml:"DataSourceInstanceParams,omitempty"`
	// The remarks on the data source.
	//
	// example:
	//
	// XX Cloud Firewall Shanghai instance
	DataSourceInstanceRemark *string `json:"DataSourceInstanceRemark,omitempty" xml:"DataSourceInstanceRemark,omitempty"`
	// The type of the data source. Valid values:
	//
	// - obs: Huawei Cloud Object Storage Service (OBS).
	//
	// - wafApi: Tencent Cloud Web Application Firewall (WAF) download API.
	//
	// - ckafka: Tencent Cloud CKafka.
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The region of the data management center for threat analysis. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: For assets in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: For assets in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceRequest) GoString() string {
	return s.String()
}

func (s *AddDataSourceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *AddDataSourceRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *AddDataSourceRequest) GetDataSourceInstanceName() *string {
	return s.DataSourceInstanceName
}

func (s *AddDataSourceRequest) GetDataSourceInstanceParams() *string {
	return s.DataSourceInstanceParams
}

func (s *AddDataSourceRequest) GetDataSourceInstanceRemark() *string {
	return s.DataSourceInstanceRemark
}

func (s *AddDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *AddDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDataSourceRequest) SetAccountId(v string) *AddDataSourceRequest {
	s.AccountId = &v
	return s
}

func (s *AddDataSourceRequest) SetCloudCode(v string) *AddDataSourceRequest {
	s.CloudCode = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceInstanceName(v string) *AddDataSourceRequest {
	s.DataSourceInstanceName = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceInstanceParams(v string) *AddDataSourceRequest {
	s.DataSourceInstanceParams = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceInstanceRemark(v string) *AddDataSourceRequest {
	s.DataSourceInstanceRemark = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceType(v string) *AddDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *AddDataSourceRequest) SetRegionId(v string) *AddDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *AddDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
