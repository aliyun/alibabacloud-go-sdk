// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupedFields(v *GetCloudAssetSummaryResponseBodyGroupedFields) *GetCloudAssetSummaryResponseBody
	GetGroupedFields() *GetCloudAssetSummaryResponseBodyGroupedFields
	SetRequestId(v string) *GetCloudAssetSummaryResponseBody
	GetRequestId() *string
}

type GetCloudAssetSummaryResponseBody struct {
	// Summary information of cloud assets.
	GroupedFields *GetCloudAssetSummaryResponseBodyGroupedFields `json:"GroupedFields,omitempty" xml:"GroupedFields,omitempty" type:"Struct"`
	// 本次调用请求的ID，是由阿里云为该请求生成的唯一标识符，可用于排查和定位问题。
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudAssetSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryResponseBody) GetGroupedFields() *GetCloudAssetSummaryResponseBodyGroupedFields {
	return s.GroupedFields
}

func (s *GetCloudAssetSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudAssetSummaryResponseBody) SetGroupedFields(v *GetCloudAssetSummaryResponseBodyGroupedFields) *GetCloudAssetSummaryResponseBody {
	s.GroupedFields = v
	return s
}

func (s *GetCloudAssetSummaryResponseBody) SetRequestId(v string) *GetCloudAssetSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBody) Validate() error {
	if s.GroupedFields != nil {
		if err := s.GroupedFields.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAssetSummaryResponseBodyGroupedFields struct {
	// List of cloud product statistics
	CloudAssetSummaryMetas []*GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas `json:"CloudAssetSummaryMetas,omitempty" xml:"CloudAssetSummaryMetas,omitempty" type:"Repeated"`
	// Total number of cloud product instances.
	//
	// example:
	//
	// 919
	InstanceCountTotal *int32 `json:"InstanceCountTotal,omitempty" xml:"InstanceCountTotal,omitempty"`
	// Total number of cloud product instances at risk
	//
	// example:
	//
	// 544
	InstanceRiskCountTotal *int32 `json:"InstanceRiskCountTotal,omitempty" xml:"InstanceRiskCountTotal,omitempty"`
}

func (s GetCloudAssetSummaryResponseBodyGroupedFields) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryResponseBodyGroupedFields) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) GetCloudAssetSummaryMetas() []*GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	return s.CloudAssetSummaryMetas
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) GetInstanceCountTotal() *int32 {
	return s.InstanceCountTotal
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) GetInstanceRiskCountTotal() *int32 {
	return s.InstanceRiskCountTotal
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) SetCloudAssetSummaryMetas(v []*GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) *GetCloudAssetSummaryResponseBodyGroupedFields {
	s.CloudAssetSummaryMetas = v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) SetInstanceCountTotal(v int32) *GetCloudAssetSummaryResponseBodyGroupedFields {
	s.InstanceCountTotal = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) SetInstanceRiskCountTotal(v int32) *GetCloudAssetSummaryResponseBodyGroupedFields {
	s.InstanceRiskCountTotal = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) Validate() error {
	if s.CloudAssetSummaryMetas != nil {
		for _, item := range s.CloudAssetSummaryMetas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas struct {
	// Subtype of the cloud product
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// 云产品的类型。取值：
	//
	// - **0**：云服务器 ECS
	//
	// - **1**：负载均衡
	//
	// - **3**：云数据库 RDS
	//
	// - **4**：云数据库 MongoDB 版
	//
	// - **5**：云数据库 Tair（兼容 Redis）
	//
	// - **6**：容器镜像服务
	//
	// - **8**：容器服务Kubernetes版
	//
	// - **9**：专有网络VPC
	//
	// - **11**：操作审计
	//
	// - **12**：CDN
	//
	// - **13**：数字证书管理服务（原SSL证书）
	//
	// - **14**：云效
	//
	// - **15**：访问控制
	//
	// - **16**：DDoS防护
	//
	// - **17**：Web应用防火墙
	//
	// - **18**：对象存储
	//
	// - **19**：云原生关系型数据库 PolarDB
	//
	// - **20**：云数据库 PostgreSQL 版
	//
	// - **21**：微服务引擎
	//
	// - **22**：文件存储NAS
	//
	// - **23**：数据安全中心
	//
	// - **24**：弹性公网IP
	//
	// - **25**：云身份服务-EIAM
	//
	// - **26**：PolarDB-X
	//
	// - **27**：Elasticsearch
	//
	// example:
	//
	// 16
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Total number of this type of cloud product instances.
	//
	// example:
	//
	// 16
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// Total number of risky instances for this type of cloud product.
	//
	// example:
	//
	// 5
	InstanceRiskCount *int32 `json:"InstanceRiskCount,omitempty" xml:"InstanceRiskCount,omitempty"`
	// 服务器厂商。取值：
	//
	// - **0**：阿里云资产
	//
	// - **1**：云外资产
	//
	// - **2**：IDC资产
	//
	// - **3**、**4**、**5**、**7**：其它云资产
	//
	// - **8**：轻量级资产
	//
	// example:
	//
	// 3
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetInstanceRiskCount() *int32 {
	return s.InstanceRiskCount
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetVendor() *int32 {
	return s.Vendor
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetAssetSubType(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.AssetSubType = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetAssetType(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.AssetType = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetInstanceCount(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.InstanceCount = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetInstanceRiskCount(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.InstanceRiskCount = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetVendor(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.Vendor = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) Validate() error {
	return dara.Validate(s)
}
