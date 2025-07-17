// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *CreateEnvironmentRequest
	GetAliyunLang() *string
	SetBindResourceId(v string) *CreateEnvironmentRequest
	GetBindResourceId() *string
	SetEnvironmentName(v string) *CreateEnvironmentRequest
	GetEnvironmentName() *string
	SetEnvironmentSubType(v string) *CreateEnvironmentRequest
	GetEnvironmentSubType() *string
	SetEnvironmentType(v string) *CreateEnvironmentRequest
	GetEnvironmentType() *string
	SetFeePackage(v string) *CreateEnvironmentRequest
	GetFeePackage() *string
	SetGrafanaWorkspaceId(v string) *CreateEnvironmentRequest
	GetGrafanaWorkspaceId() *string
	SetInitEnvironment(v bool) *CreateEnvironmentRequest
	GetInitEnvironment() *bool
	SetManagedType(v string) *CreateEnvironmentRequest
	GetManagedType() *string
	SetPrometheusInstanceId(v string) *CreateEnvironmentRequest
	GetPrometheusInstanceId() *string
	SetRegionId(v string) *CreateEnvironmentRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEnvironmentRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateEnvironmentRequestTags) *CreateEnvironmentRequest
	GetTags() []*CreateEnvironmentRequestTags
}

type CreateEnvironmentRequest struct {
	// The language. Default value: zh.
	//
	// Valid values:
	//
	// 	- en: English
	//
	// 	- zh: Chinese
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The ID of the resource bound to the environment, such as the container ID or VPC ID. For a Cloud environment, specify the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c6e9dec475dca4a50a188411d8cbxxx
	BindResourceId *string `json:"BindResourceId,omitempty" xml:"BindResourceId,omitempty"`
	// The name of the environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// env1
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// The subtype of the environment. Valid values:
	//
	// 	- CS: Container Service for Kubernetes (ACK) or Distributed Cloud Container Platform for Kubernetes (ACK One)
	//
	// 	- ECS: ECS
	//
	// 	- Cloud: cloud service
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS, ACK, etc.
	EnvironmentSubType *string `json:"EnvironmentSubType,omitempty" xml:"EnvironmentSubType,omitempty"`
	// The type of the environment. Valid values:
	//
	// 	- CS: Container Service
	//
	// 	- ECS: Elastic Compute Service
	//
	// 	- Cloud: cloud service
	//
	// This parameter is required.
	//
	// example:
	//
	// CS
	EnvironmentType *string `json:"EnvironmentType,omitempty" xml:"EnvironmentType,omitempty"`
	// The payable resource plan.
	//
	// 	- If the EnvironmentType parameter is set to CS, set the value to CS_Basic or CS_Pro. Default value: CS_Basic.
	//
	// 	- Otherwise, leave the parameter empty.
	//
	// example:
	//
	// CS_Basic
	FeePackage *string `json:"FeePackage,omitempty" xml:"FeePackage,omitempty"`
	// The ID of the Grafana workspace associated with the environment. If this parameter is left empty, the default shared Grafana workspace is used.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// grafana-rnglkcdrntlhk0****
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// Specifies whether to initialize the environment.
	//
	// example:
	//
	// false
	InitEnvironment *bool `json:"InitEnvironment,omitempty" xml:"InitEnvironment,omitempty"`
	// Specifies whether agents or exporters are managed. Valid values:
	//
	// 	- none: No. By default, no managed agents or exporters are provided for ACK clusters.
	//
	// 	- agent: Agents are managed. By default, managed agents are provided for ASK clusters, ACS clusters, and ACK One clusters.
	//
	// 	- agent-exporter: Agents and exporters are managed. By default, managed agents and exporters are provided for cloud services.
	//
	// example:
	//
	// none
	ManagedType *string `json:"ManagedType,omitempty" xml:"ManagedType,omitempty"`
	// The ID of the Prometheus instance. If no Prometheus instance is created, call the InitEnvironment operation.
	//
	// example:
	//
	// c6e9dec475dca4a50a188411d8cbxxx
	PrometheusInstanceId *string `json:"PrometheusInstanceId,omitempty" xml:"PrometheusInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the instance. You can specify this parameter to manage tags for the instance.
	Tags []*CreateEnvironmentRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateEnvironmentRequest) GetBindResourceId() *string {
	return s.BindResourceId
}

func (s *CreateEnvironmentRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *CreateEnvironmentRequest) GetEnvironmentSubType() *string {
	return s.EnvironmentSubType
}

func (s *CreateEnvironmentRequest) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *CreateEnvironmentRequest) GetFeePackage() *string {
	return s.FeePackage
}

func (s *CreateEnvironmentRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *CreateEnvironmentRequest) GetInitEnvironment() *bool {
	return s.InitEnvironment
}

func (s *CreateEnvironmentRequest) GetManagedType() *string {
	return s.ManagedType
}

func (s *CreateEnvironmentRequest) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *CreateEnvironmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnvironmentRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEnvironmentRequest) GetTags() []*CreateEnvironmentRequestTags {
	return s.Tags
}

func (s *CreateEnvironmentRequest) SetAliyunLang(v string) *CreateEnvironmentRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateEnvironmentRequest) SetBindResourceId(v string) *CreateEnvironmentRequest {
	s.BindResourceId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetEnvironmentName(v string) *CreateEnvironmentRequest {
	s.EnvironmentName = &v
	return s
}

func (s *CreateEnvironmentRequest) SetEnvironmentSubType(v string) *CreateEnvironmentRequest {
	s.EnvironmentSubType = &v
	return s
}

func (s *CreateEnvironmentRequest) SetEnvironmentType(v string) *CreateEnvironmentRequest {
	s.EnvironmentType = &v
	return s
}

func (s *CreateEnvironmentRequest) SetFeePackage(v string) *CreateEnvironmentRequest {
	s.FeePackage = &v
	return s
}

func (s *CreateEnvironmentRequest) SetGrafanaWorkspaceId(v string) *CreateEnvironmentRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetInitEnvironment(v bool) *CreateEnvironmentRequest {
	s.InitEnvironment = &v
	return s
}

func (s *CreateEnvironmentRequest) SetManagedType(v string) *CreateEnvironmentRequest {
	s.ManagedType = &v
	return s
}

func (s *CreateEnvironmentRequest) SetPrometheusInstanceId(v string) *CreateEnvironmentRequest {
	s.PrometheusInstanceId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetRegionId(v string) *CreateEnvironmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetResourceGroupId(v string) *CreateEnvironmentRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetTags(v []*CreateEnvironmentRequestTags) *CreateEnvironmentRequest {
	s.Tags = v
	return s
}

func (s *CreateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}

type CreateEnvironmentRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEnvironmentRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentRequestTags) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateEnvironmentRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateEnvironmentRequestTags) SetKey(v string) *CreateEnvironmentRequestTags {
	s.Key = &v
	return s
}

func (s *CreateEnvironmentRequestTags) SetValue(v string) *CreateEnvironmentRequestTags {
	s.Value = &v
	return s
}

func (s *CreateEnvironmentRequestTags) Validate() error {
	return dara.Validate(s)
}
