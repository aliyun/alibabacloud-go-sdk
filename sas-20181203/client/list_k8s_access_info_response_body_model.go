// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sAccessInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetK8sAccessInfos(v []*ListK8sAccessInfoResponseBodyK8sAccessInfos) *ListK8sAccessInfoResponseBody
	GetK8sAccessInfos() []*ListK8sAccessInfoResponseBodyK8sAccessInfos
	SetRequestId(v string) *ListK8sAccessInfoResponseBody
	GetRequestId() *string
}

type ListK8sAccessInfoResponseBody struct {
	// The information about the Kubernetes clusters.
	K8sAccessInfos []*ListK8sAccessInfoResponseBodyK8sAccessInfos `json:"K8sAccessInfos,omitempty" xml:"K8sAccessInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListK8sAccessInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8sAccessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8sAccessInfoResponseBody) GetK8sAccessInfos() []*ListK8sAccessInfoResponseBodyK8sAccessInfos {
	return s.K8sAccessInfos
}

func (s *ListK8sAccessInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8sAccessInfoResponseBody) SetK8sAccessInfos(v []*ListK8sAccessInfoResponseBodyK8sAccessInfos) *ListK8sAccessInfoResponseBody {
	s.K8sAccessInfos = v
	return s
}

func (s *ListK8sAccessInfoResponseBody) SetRequestId(v string) *ListK8sAccessInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8sAccessInfoResponseBody) Validate() error {
	if s.K8sAccessInfos != nil {
		for _, item := range s.K8sAccessInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sAccessInfoResponseBodyK8sAccessInfos struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1960721413485****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The Simple Log Service Logstore that is used to store the audit logs.
	//
	// example:
	//
	// audit-cf6baf6afa106eca665296fdf68b****
	AuditLogStore *string `json:"AuditLogStore,omitempty" xml:"AuditLogStore,omitempty"`
	// The Simple Log Service project that is used to store the audit logs.
	//
	// example:
	//
	// k8s-log-custom-your-project-sd89eh****
	AuditProject *string `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	// The ID of the region in which the server is deployed.
	//
	// example:
	//
	// cn-hangzhou
	AuditRegionId *string `json:"AuditRegionId,omitempty" xml:"AuditRegionId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c0da5e4cb82a848c4a57c4dc9f49a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// arm
	CpuArch *string `json:"CpuArch,omitempty" xml:"CpuArch,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1690596321613
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// 11088522
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the server group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The UUID of the access information.
	//
	// example:
	//
	// 67070
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The installation key of the Kubernetes cluster.
	//
	// example:
	//
	// xxx
	InstallKey *string `json:"InstallKey,omitempty" xml:"InstallKey,omitempty"`
	// The service provider.
	//
	// example:
	//
	// ALIYUN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListK8sAccessInfoResponseBodyK8sAccessInfos) String() string {
	return dara.Prettify(s)
}

func (s ListK8sAccessInfoResponseBodyK8sAccessInfos) GoString() string {
	return s.String()
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetAuditLogStore() *string {
	return s.AuditLogStore
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetAuditProject() *string {
	return s.AuditProject
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetAuditRegionId() *string {
	return s.AuditRegionId
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetCpuArch() *string {
	return s.CpuArch
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetGroupId() *string {
	return s.GroupId
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetGroupName() *string {
	return s.GroupName
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetId() *int64 {
	return s.Id
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetInstallKey() *string {
	return s.InstallKey
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) GetVendor() *string {
	return s.Vendor
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetAliUid(v int64) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.AliUid = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetAuditLogStore(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.AuditLogStore = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetAuditProject(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.AuditProject = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetAuditRegionId(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.AuditRegionId = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetClusterId(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.ClusterId = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetClusterName(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.ClusterName = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetCpuArch(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.CpuArch = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetExpireDate(v int64) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.ExpireDate = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetGroupId(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.GroupId = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetGroupName(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.GroupName = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetId(v int64) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.Id = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetInstallKey(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.InstallKey = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) SetVendor(v string) *ListK8sAccessInfoResponseBodyK8sAccessInfos {
	s.Vendor = &v
	return s
}

func (s *ListK8sAccessInfoResponseBodyK8sAccessInfos) Validate() error {
	return dara.Validate(s)
}
