// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateK8sAccessInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateK8sAccessInfoResponseBodyData) *GenerateK8sAccessInfoResponseBody
	GetData() *GenerateK8sAccessInfoResponseBodyData
	SetRequestId(v string) *GenerateK8sAccessInfoResponseBody
	GetRequestId() *string
}

type GenerateK8sAccessInfoResponseBody struct {
	// The returned data.
	Data *GenerateK8sAccessInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 061955B2-BC40-589F-AF63-C40A901EE279
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateK8sAccessInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateK8sAccessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateK8sAccessInfoResponseBody) GetData() *GenerateK8sAccessInfoResponseBodyData {
	return s.Data
}

func (s *GenerateK8sAccessInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateK8sAccessInfoResponseBody) SetData(v *GenerateK8sAccessInfoResponseBodyData) *GenerateK8sAccessInfoResponseBody {
	s.Data = v
	return s
}

func (s *GenerateK8sAccessInfoResponseBody) SetRequestId(v string) *GenerateK8sAccessInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateK8sAccessInfoResponseBodyData struct {
	// The aliuid of the user.
	//
	// example:
	//
	// 1766185894104***
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The SLS Logstore of the audit log.
	//
	// example:
	//
	// audit-cf6baf6afa106eca665296fdf68b65bf
	AuditLogStore *string `json:"AuditLogStore,omitempty" xml:"AuditLogStore,omitempty"`
	// The SLS project of the audit log.
	//
	// example:
	//
	// k8s-log-custom-your-project-sd89ehaaa
	AuditProject *string `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	// The region ID of the server.
	//
	// example:
	//
	// cn-hangzhou
	AuditRegionId *string `json:"AuditRegionId,omitempty" xml:"AuditRegionId,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The CPU architecture, which can be ARM or x86. Valid values:
	//
	// - arm64
	//
	// - x86
	//
	// example:
	//
	// x86
	CpuArch *string `json:"CpuArch,omitempty" xml:"CpuArch,omitempty"`
	// The expiration time, in milliseconds.
	//
	// example:
	//
	// 1804230578566
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// 11618788
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The installation key of the server.
	//
	// example:
	//
	// BC66185***
	InstallKey *string `json:"InstallKey,omitempty" xml:"InstallKey,omitempty"`
}

func (s GenerateK8sAccessInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateK8sAccessInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetAuditLogStore() *string {
	return s.AuditLogStore
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetAuditProject() *string {
	return s.AuditProject
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetAuditRegionId() *string {
	return s.AuditRegionId
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetCpuArch() *string {
	return s.CpuArch
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *GenerateK8sAccessInfoResponseBodyData) GetInstallKey() *string {
	return s.InstallKey
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetAliUid(v int64) *GenerateK8sAccessInfoResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetAuditLogStore(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.AuditLogStore = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetAuditProject(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.AuditProject = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetAuditRegionId(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.AuditRegionId = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetClusterId(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetClusterName(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetCpuArch(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.CpuArch = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetExpireDate(v int64) *GenerateK8sAccessInfoResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetGroupId(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) SetInstallKey(v string) *GenerateK8sAccessInfoResponseBodyData {
	s.InstallKey = &v
	return s
}

func (s *GenerateK8sAccessInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
