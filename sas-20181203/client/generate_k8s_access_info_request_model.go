// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateK8sAccessInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunYundunGatewayApiName(v string) *GenerateK8sAccessInfoRequest
	GetAliyunYundunGatewayApiName() *string
	SetAliyunYundunGatewayPopName(v string) *GenerateK8sAccessInfoRequest
	GetAliyunYundunGatewayPopName() *string
	SetAliyunYundunGatewayProjectName(v string) *GenerateK8sAccessInfoRequest
	GetAliyunYundunGatewayProjectName() *string
	SetAuditLogStore(v string) *GenerateK8sAccessInfoRequest
	GetAuditLogStore() *string
	SetAuditProject(v string) *GenerateK8sAccessInfoRequest
	GetAuditProject() *string
	SetAuditRegionId(v string) *GenerateK8sAccessInfoRequest
	GetAuditRegionId() *string
	SetClusterName(v string) *GenerateK8sAccessInfoRequest
	GetClusterName() *string
	SetCpuArch(v string) *GenerateK8sAccessInfoRequest
	GetCpuArch() *string
	SetExpireDate(v int64) *GenerateK8sAccessInfoRequest
	GetExpireDate() *int64
	SetGroupId(v int64) *GenerateK8sAccessInfoRequest
	GetGroupId() *int64
	SetVendor(v string) *GenerateK8sAccessInfoRequest
	GetVendor() *string
}

type GenerateK8sAccessInfoRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayApiName *string `json:"AliyunYundunGatewayApiName,omitempty" xml:"AliyunYundunGatewayApiName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayPopName *string `json:"AliyunYundunGatewayPopName,omitempty" xml:"AliyunYundunGatewayPopName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayProjectName *string `json:"AliyunYundunGatewayProjectName,omitempty" xml:"AliyunYundunGatewayProjectName,omitempty"`
	// The Simple Log Service Logstore that is used to store the audit logs.
	//
	// example:
	//
	// audit-cf6baf6afa106eca665296fdf68b65bf
	AuditLogStore *string `json:"AuditLogStore,omitempty" xml:"AuditLogStore,omitempty"`
	// The Simple Log Service project that is used to store the audit logs.
	//
	// example:
	//
	// k8s-log-custom-huxintest1018-2
	AuditProject *string `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	// The ID of the region in which the audit logs are stored.
	//
	// example:
	//
	// cn-hangzhou
	AuditRegionId *string `json:"AuditRegionId,omitempty" xml:"AuditRegionId,omitempty"`
	// The name of the Kubernetes cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// CPU architectures are divided into ARM architecture and x86 architecture.
	//
	// example:
	//
	// arm
	CpuArch *string `json:"CpuArch,omitempty" xml:"CpuArch,omitempty"`
	// The time at which the container ends to be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1711951508388
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11341690
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The service provider of the cloud asset. Valid values:
	//
	// 	- **Tencent**
	//
	// 	- **HUAWEICLOUD**
	//
	// 	- **Azure**
	//
	// 	- **AWS**
	//
	// 	- **Others**
	//
	// This parameter is required.
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GenerateK8sAccessInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateK8sAccessInfoRequest) GoString() string {
	return s.String()
}

func (s *GenerateK8sAccessInfoRequest) GetAliyunYundunGatewayApiName() *string {
	return s.AliyunYundunGatewayApiName
}

func (s *GenerateK8sAccessInfoRequest) GetAliyunYundunGatewayPopName() *string {
	return s.AliyunYundunGatewayPopName
}

func (s *GenerateK8sAccessInfoRequest) GetAliyunYundunGatewayProjectName() *string {
	return s.AliyunYundunGatewayProjectName
}

func (s *GenerateK8sAccessInfoRequest) GetAuditLogStore() *string {
	return s.AuditLogStore
}

func (s *GenerateK8sAccessInfoRequest) GetAuditProject() *string {
	return s.AuditProject
}

func (s *GenerateK8sAccessInfoRequest) GetAuditRegionId() *string {
	return s.AuditRegionId
}

func (s *GenerateK8sAccessInfoRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *GenerateK8sAccessInfoRequest) GetCpuArch() *string {
	return s.CpuArch
}

func (s *GenerateK8sAccessInfoRequest) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *GenerateK8sAccessInfoRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GenerateK8sAccessInfoRequest) GetVendor() *string {
	return s.Vendor
}

func (s *GenerateK8sAccessInfoRequest) SetAliyunYundunGatewayApiName(v string) *GenerateK8sAccessInfoRequest {
	s.AliyunYundunGatewayApiName = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetAliyunYundunGatewayPopName(v string) *GenerateK8sAccessInfoRequest {
	s.AliyunYundunGatewayPopName = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetAliyunYundunGatewayProjectName(v string) *GenerateK8sAccessInfoRequest {
	s.AliyunYundunGatewayProjectName = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetAuditLogStore(v string) *GenerateK8sAccessInfoRequest {
	s.AuditLogStore = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetAuditProject(v string) *GenerateK8sAccessInfoRequest {
	s.AuditProject = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetAuditRegionId(v string) *GenerateK8sAccessInfoRequest {
	s.AuditRegionId = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetClusterName(v string) *GenerateK8sAccessInfoRequest {
	s.ClusterName = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetCpuArch(v string) *GenerateK8sAccessInfoRequest {
	s.CpuArch = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetExpireDate(v int64) *GenerateK8sAccessInfoRequest {
	s.ExpireDate = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetGroupId(v int64) *GenerateK8sAccessInfoRequest {
	s.GroupId = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) SetVendor(v string) *GenerateK8sAccessInfoRequest {
	s.Vendor = &v
	return s
}

func (s *GenerateK8sAccessInfoRequest) Validate() error {
	return dara.Validate(s)
}
