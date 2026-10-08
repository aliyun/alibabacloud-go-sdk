// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationList(v *ListApplicationResponseBodyApplicationList) *ListApplicationResponseBody
	GetApplicationList() *ListApplicationResponseBodyApplicationList
	SetCode(v int32) *ListApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *ListApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApplicationResponseBody
	GetRequestId() *string
}

type ListApplicationResponseBody struct {
	ApplicationList *ListApplicationResponseBodyApplicationList `json:"ApplicationList,omitempty" xml:"ApplicationList,omitempty" type:"Struct"`
	// The status code of the response.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5d6fa0bc-cc3**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBody) GetApplicationList() *ListApplicationResponseBodyApplicationList {
	return s.ApplicationList
}

func (s *ListApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationResponseBody) SetApplicationList(v *ListApplicationResponseBodyApplicationList) *ListApplicationResponseBody {
	s.ApplicationList = v
	return s
}

func (s *ListApplicationResponseBody) SetCode(v int32) *ListApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationResponseBody) SetMessage(v string) *ListApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationResponseBody) SetRequestId(v string) *ListApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationResponseBody) Validate() error {
	if s.ApplicationList != nil {
		if err := s.ApplicationList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationResponseBodyApplicationList struct {
	Application []*ListApplicationResponseBodyApplicationListApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s ListApplicationResponseBodyApplicationList) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBodyApplicationList) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyApplicationList) GetApplication() []*ListApplicationResponseBodyApplicationListApplication {
	return s.Application
}

func (s *ListApplicationResponseBodyApplicationList) SetApplication(v []*ListApplicationResponseBodyApplicationListApplication) *ListApplicationResponseBodyApplicationList {
	s.Application = v
	return s
}

func (s *ListApplicationResponseBodyApplicationList) Validate() error {
	if s.Application != nil {
		for _, item := range s.Application {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationResponseBodyApplicationListApplication struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ApplicationType      *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	BuildPackageId       *int64  `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType          *int32  `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtSlbIp             *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty"`
	ExtSlbListenerPort   *int32  `json:"ExtSlbListenerPort,omitempty" xml:"ExtSlbListenerPort,omitempty"`
	Instances            *int32  `json:"Instances,omitempty" xml:"Instances,omitempty"`
	K8sNamespace         *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId          *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RunningInstanceCount *int32  `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty"`
	SlbIp                *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	SlbListenerPort      *int32  `json:"SlbListenerPort,omitempty" xml:"SlbListenerPort,omitempty"`
	SlbPort              *int32  `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListApplicationResponseBodyApplicationListApplication) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBodyApplicationListApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetBuildPackageId() *int64 {
	return s.BuildPackageId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetExtSlbIp() *string {
	return s.ExtSlbIp
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetExtSlbListenerPort() *int32 {
	return s.ExtSlbListenerPort
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetInstances() *int32 {
	return s.Instances
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetName() *string {
	return s.Name
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetPort() *int32 {
	return s.Port
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetRunningInstanceCount() *int32 {
	return s.RunningInstanceCount
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetSlbIp() *string {
	return s.SlbIp
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetSlbListenerPort() *int32 {
	return s.SlbListenerPort
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetSlbPort() *int32 {
	return s.SlbPort
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetState() *string {
	return s.State
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetAppId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.AppId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetApplicationType(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetBuildPackageId(v int64) *ListApplicationResponseBodyApplicationListApplication {
	s.BuildPackageId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetClusterId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ClusterId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetClusterType(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.ClusterType = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetCreateTime(v int64) *ListApplicationResponseBodyApplicationListApplication {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetExtSlbIp(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ExtSlbIp = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetExtSlbListenerPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.ExtSlbListenerPort = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetInstances(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.Instances = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetK8sNamespace(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.K8sNamespace = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetName(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.Name = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetNamespaceId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.Port = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetRegionId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.RegionId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetResourceGroupId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetRunningInstanceCount(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.RunningInstanceCount = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetSlbIp(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.SlbIp = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetSlbListenerPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.SlbListenerPort = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetSlbPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.SlbPort = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetState(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.State = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) Validate() error {
	return dara.Validate(s)
}
