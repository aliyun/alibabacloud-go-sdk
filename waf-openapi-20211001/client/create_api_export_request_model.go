// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateApiExportRequest
	GetClusterId() *string
	SetInstanceId(v string) *CreateApiExportRequest
	GetInstanceId() *string
	SetParam(v string) *CreateApiExportRequest
	GetParam() *string
	SetRegion(v string) *CreateApiExportRequest
	GetRegion() *string
	SetRegionId(v string) *CreateApiExportRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateApiExportRequest
	GetResourceManagerResourceGroupId() *string
	SetType(v string) *CreateApiExportRequest
	GetType() *string
	SetZoneId(v string) *CreateApiExportRequest
	GetZoneId() *string
}

type CreateApiExportRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter is applicable only to hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query hybrid cloud clusters.
	//
	// example:
	//
	// 993
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zxu3***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The filter conditions for the export task. The value is a JSON string.
	//
	// > The filter conditions vary based on the export task type specified by **Type**. For more information, see **Export task parameters**.
	//
	// example:
	//
	// {
	//
	//     "startTime": 1741449600,
	//
	//     "endTime": 1744079820,
	//
	//     "sensitiveLevel": "L1"
	//
	// }
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The language of the response. Valid values:
	//
	// - **cn*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the export task. Valid values:
	//
	// - **apisec_api*	- (default): Exports API asset data.
	//
	// - **apisec_abnormal**: Exports API threat data.
	//
	// - **apisec_event**: Exports API security event data.
	//
	// example:
	//
	// apisec_api
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time zone of the export data, such as **Asia/Shanghai**.
	//
	// example:
	//
	// Asia/Shanghai
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateApiExportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiExportRequest) GoString() string {
	return s.String()
}

func (s *CreateApiExportRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateApiExportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApiExportRequest) GetParam() *string {
	return s.Param
}

func (s *CreateApiExportRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateApiExportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApiExportRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateApiExportRequest) GetType() *string {
	return s.Type
}

func (s *CreateApiExportRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateApiExportRequest) SetClusterId(v string) *CreateApiExportRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateApiExportRequest) SetInstanceId(v string) *CreateApiExportRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApiExportRequest) SetParam(v string) *CreateApiExportRequest {
	s.Param = &v
	return s
}

func (s *CreateApiExportRequest) SetRegion(v string) *CreateApiExportRequest {
	s.Region = &v
	return s
}

func (s *CreateApiExportRequest) SetRegionId(v string) *CreateApiExportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApiExportRequest) SetResourceManagerResourceGroupId(v string) *CreateApiExportRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateApiExportRequest) SetType(v string) *CreateApiExportRequest {
	s.Type = &v
	return s
}

func (s *CreateApiExportRequest) SetZoneId(v string) *CreateApiExportRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateApiExportRequest) Validate() error {
	return dara.Validate(s)
}
