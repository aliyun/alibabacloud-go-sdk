// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeployGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDeployGroupResponseBody
	GetCode() *int32
	SetDeployGroupList(v *ListDeployGroupResponseBodyDeployGroupList) *ListDeployGroupResponseBody
	GetDeployGroupList() *ListDeployGroupResponseBodyDeployGroupList
	SetMessage(v string) *ListDeployGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeployGroupResponseBody
	GetRequestId() *string
}

type ListDeployGroupResponseBody struct {
	// The status code of the request or a POP error code.
	//
	// example:
	//
	// 200
	Code            *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	DeployGroupList *ListDeployGroupResponseBodyDeployGroupList `json:"DeployGroupList,omitempty" xml:"DeployGroupList,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3FDE-DS9R-*********************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeployGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDeployGroupResponseBody) GetDeployGroupList() *ListDeployGroupResponseBodyDeployGroupList {
	return s.DeployGroupList
}

func (s *ListDeployGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeployGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeployGroupResponseBody) SetCode(v int32) *ListDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeployGroupResponseBody) SetDeployGroupList(v *ListDeployGroupResponseBodyDeployGroupList) *ListDeployGroupResponseBody {
	s.DeployGroupList = v
	return s
}

func (s *ListDeployGroupResponseBody) SetMessage(v string) *ListDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeployGroupResponseBody) SetRequestId(v string) *ListDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployGroupResponseBody) Validate() error {
	if s.DeployGroupList != nil {
		if err := s.DeployGroupList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeployGroupResponseBodyDeployGroupList struct {
	DeployGroup []*ListDeployGroupResponseBodyDeployGroupListDeployGroup `json:"DeployGroup,omitempty" xml:"DeployGroup,omitempty" type:"Repeated"`
}

func (s ListDeployGroupResponseBodyDeployGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupResponseBodyDeployGroupList) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBodyDeployGroupList) GetDeployGroup() []*ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	return s.DeployGroup
}

func (s *ListDeployGroupResponseBodyDeployGroupList) SetDeployGroup(v []*ListDeployGroupResponseBodyDeployGroupListDeployGroup) *ListDeployGroupResponseBodyDeployGroupList {
	s.DeployGroup = v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupList) Validate() error {
	if s.DeployGroup != nil {
		for _, item := range s.DeployGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeployGroupResponseBodyDeployGroupListDeployGroup struct {
	AppId                   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId            *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	BaseComponentMetaName   *string `json:"BaseComponentMetaName,omitempty" xml:"BaseComponentMetaName,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName             *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CpuLimit                *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	CpuRequest              *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	CreateTime              *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CsClusterId             *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty"`
	DeploymentName          *string `json:"DeploymentName,omitempty" xml:"DeploymentName,omitempty"`
	Env                     *string `json:"Env,omitempty" xml:"Env,omitempty"`
	EphemeralStorageLimit   *string `json:"EphemeralStorageLimit,omitempty" xml:"EphemeralStorageLimit,omitempty"`
	EphemeralStorageRequest *string `json:"EphemeralStorageRequest,omitempty" xml:"EphemeralStorageRequest,omitempty"`
	GroupId                 *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName               *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType               *int32  `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	Labels                  *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	LastUpdateTime          *int64  `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	MemoryLimit             *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	MemoryRequest           *string `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty"`
	NameSpace               *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	PackagePublicUrl        *string `json:"PackagePublicUrl,omitempty" xml:"PackagePublicUrl,omitempty"`
	PackageUrl              *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion          *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	PackageVersionId        *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	PostStart               *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop                 *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	Reversion               *string `json:"Reversion,omitempty" xml:"Reversion,omitempty"`
	Selector                *string `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Strategy                *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	UpdateTime              *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VExtServerGroupId       *string `json:"VExtServerGroupId,omitempty" xml:"VExtServerGroupId,omitempty"`
	VServerGroupId          *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s ListDeployGroupResponseBodyDeployGroupListDeployGroup) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupResponseBodyDeployGroupListDeployGroup) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetAppId() *string {
	return s.AppId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetBaseComponentMetaName() *string {
	return s.BaseComponentMetaName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCpuRequest() *string {
	return s.CpuRequest
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetEnv() *string {
	return s.Env
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetEphemeralStorageLimit() *string {
	return s.EphemeralStorageLimit
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetEphemeralStorageRequest() *string {
	return s.EphemeralStorageRequest
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetGroupType() *int32 {
	return s.GroupType
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetLabels() *string {
	return s.Labels
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetMemoryRequest() *string {
	return s.MemoryRequest
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetNameSpace() *string {
	return s.NameSpace
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackagePublicUrl() *string {
	return s.PackagePublicUrl
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPostStart() *string {
	return s.PostStart
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPreStop() *string {
	return s.PreStop
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetReversion() *string {
	return s.Reversion
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetSelector() *string {
	return s.Selector
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetStatus() *string {
	return s.Status
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetStrategy() *string {
	return s.Strategy
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetVExtServerGroupId() *string {
	return s.VExtServerGroupId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetAppId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.AppId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetAppVersionId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.AppVersionId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetBaseComponentMetaName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.BaseComponentMetaName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetClusterId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.ClusterId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetClusterName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.ClusterName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCpuLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CpuLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCpuRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CpuRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCreateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CreateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCsClusterId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CsClusterId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetDeploymentName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.DeploymentName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetEnv(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Env = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetEphemeralStorageLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.EphemeralStorageLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetEphemeralStorageRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.EphemeralStorageRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupType(v int32) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupType = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetLabels(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Labels = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetLastUpdateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.LastUpdateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetMemoryLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.MemoryLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetMemoryRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.MemoryRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetNameSpace(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.NameSpace = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackagePublicUrl(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackagePublicUrl = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageUrl(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageUrl = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageVersion(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageVersion = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageVersionId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageVersionId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPostStart(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PostStart = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPreStop(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PreStop = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetReversion(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Reversion = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetSelector(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Selector = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetStatus(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Status = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetStrategy(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Strategy = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetUpdateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.UpdateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetVExtServerGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.VExtServerGroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetVServerGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.VServerGroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) Validate() error {
	return dara.Validate(s)
}
