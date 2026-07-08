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
	// The hybrid cloud cluster ID.
	//
	// > This parameter applies only to hybrid cloud scenarios. You can call [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) to obtain hybrid cloud cluster information.
	//
	// example:
	//
	// 993
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zxu3***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The extended parameters of the export task. You can filter the exported content by specifying conditions. The value is a JSON string constructed from a series of parameters.
	//
	// > The specific parameters vary depending on the specified **export task type*	- (**Type**). For more information, refer to **Export task parameter description**.
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
	// The language type. Valid values:
	//
	// - **cn*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the export task. Valid values:
	//
	// - **apisec_api*	- (default): API asset task.
	//
	// - **apisec_abnormal**: API risk task.
	//
	// - **apisec_event**: API security event task.
	//
	// example:
	//
	// apisec_api
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time zone ID.
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
