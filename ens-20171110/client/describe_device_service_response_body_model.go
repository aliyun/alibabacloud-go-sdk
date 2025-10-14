// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppMetaData(v *DescribeDeviceServiceResponseBodyAppMetaData) *DescribeDeviceServiceResponseBody
	GetAppMetaData() *DescribeDeviceServiceResponseBodyAppMetaData
	SetAppStatus(v *DescribeDeviceServiceResponseBodyAppStatus) *DescribeDeviceServiceResponseBody
	GetAppStatus() *DescribeDeviceServiceResponseBodyAppStatus
	SetRequestId(v string) *DescribeDeviceServiceResponseBody
	GetRequestId() *string
	SetResourceDetailInfos(v []*DescribeDeviceServiceResponseBodyResourceDetailInfos) *DescribeDeviceServiceResponseBody
	GetResourceDetailInfos() []*DescribeDeviceServiceResponseBodyResourceDetailInfos
	SetResourceInfos(v []*DescribeDeviceServiceResponseBodyResourceInfos) *DescribeDeviceServiceResponseBody
	GetResourceInfos() []*DescribeDeviceServiceResponseBodyResourceInfos
}

type DescribeDeviceServiceResponseBody struct {
	// The basic properties of the application.
	AppMetaData *DescribeDeviceServiceResponseBodyAppMetaData `json:"AppMetaData,omitempty" xml:"AppMetaData,omitempty" type:"Struct"`
	// The status information of the application.
	AppStatus *DescribeDeviceServiceResponseBodyAppStatus `json:"AppStatus,omitempty" xml:"AppStatus,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3A535110-3EE3-5EC5-B1ED-10B7067A1FC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the devices.
	ResourceDetailInfos []*DescribeDeviceServiceResponseBodyResourceDetailInfos `json:"ResourceDetailInfos,omitempty" xml:"ResourceDetailInfos,omitempty" type:"Repeated"`
	// The information about the instances.
	ResourceInfos []*DescribeDeviceServiceResponseBodyResourceInfos `json:"ResourceInfos,omitempty" xml:"ResourceInfos,omitempty" type:"Repeated"`
}

func (s DescribeDeviceServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBody) GetAppMetaData() *DescribeDeviceServiceResponseBodyAppMetaData {
	return s.AppMetaData
}

func (s *DescribeDeviceServiceResponseBody) GetAppStatus() *DescribeDeviceServiceResponseBodyAppStatus {
	return s.AppStatus
}

func (s *DescribeDeviceServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceServiceResponseBody) GetResourceDetailInfos() []*DescribeDeviceServiceResponseBodyResourceDetailInfos {
	return s.ResourceDetailInfos
}

func (s *DescribeDeviceServiceResponseBody) GetResourceInfos() []*DescribeDeviceServiceResponseBodyResourceInfos {
	return s.ResourceInfos
}

func (s *DescribeDeviceServiceResponseBody) SetAppMetaData(v *DescribeDeviceServiceResponseBodyAppMetaData) *DescribeDeviceServiceResponseBody {
	s.AppMetaData = v
	return s
}

func (s *DescribeDeviceServiceResponseBody) SetAppStatus(v *DescribeDeviceServiceResponseBodyAppStatus) *DescribeDeviceServiceResponseBody {
	s.AppStatus = v
	return s
}

func (s *DescribeDeviceServiceResponseBody) SetRequestId(v string) *DescribeDeviceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceServiceResponseBody) SetResourceDetailInfos(v []*DescribeDeviceServiceResponseBodyResourceDetailInfos) *DescribeDeviceServiceResponseBody {
	s.ResourceDetailInfos = v
	return s
}

func (s *DescribeDeviceServiceResponseBody) SetResourceInfos(v []*DescribeDeviceServiceResponseBodyResourceInfos) *DescribeDeviceServiceResponseBody {
	s.ResourceInfos = v
	return s
}

func (s *DescribeDeviceServiceResponseBody) Validate() error {
	if s.AppMetaData != nil {
		if err := s.AppMetaData.Validate(); err != nil {
			return err
		}
	}
	if s.AppStatus != nil {
		if err := s.AppStatus.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceDetailInfos != nil {
		for _, item := range s.ResourceDetailInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceInfos != nil {
		for _, item := range s.ResourceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeviceServiceResponseBodyAppMetaData struct {
	// The ID of the application.
	//
	// example:
	//
	// 97a32f2a-aa2c-436a-b19c-05b20d258618
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// iotx-api-admin
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The stable version number of the application.
	//
	// example:
	//
	// v1
	AppStableVersion *string `json:"AppStableVersion,omitempty" xml:"AppStableVersion,omitempty"`
	// The type of the application. The value is of the enumeration type. Valid values:
	//
	// 	- Common
	//
	// 	- Scheduler
	//
	// example:
	//
	// Common
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The name of the application cluster.
	//
	// example:
	//
	// poc
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2022-03-03T03:42:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyAppMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyAppMetaData) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) GetAppStableVersion() *string {
	return s.AppStableVersion
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) GetAppType() *string {
	return s.AppType
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) SetAppId(v string) *DescribeDeviceServiceResponseBodyAppMetaData {
	s.AppId = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) SetAppName(v string) *DescribeDeviceServiceResponseBodyAppMetaData {
	s.AppName = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) SetAppStableVersion(v string) *DescribeDeviceServiceResponseBodyAppMetaData {
	s.AppStableVersion = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) SetAppType(v string) *DescribeDeviceServiceResponseBodyAppMetaData {
	s.AppType = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) SetClusterName(v string) *DescribeDeviceServiceResponseBodyAppMetaData {
	s.ClusterName = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) SetCreateTime(v string) *DescribeDeviceServiceResponseBodyAppMetaData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) SetDescription(v string) *DescribeDeviceServiceResponseBodyAppMetaData {
	s.Description = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppMetaData) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceServiceResponseBodyAppStatus struct {
	// The status of the application. The value is of the enumeration type. Valid values:
	//
	// Three intermediate states:
	//
	// 	- CREATING
	//
	// 	- UPDATING
	//
	// 	- DELETING
	//
	// Four final states:
	//
	// 	- CREATE_FAILED
	//
	// 	- UPDATE_FAILED
	//
	// 	- DELETE_FAILED
	//
	// 	- RUNNING
	//
	// example:
	//
	// CREATING
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The description of the application status.
	//
	// example:
	//
	// Creating in progress
	StatusDescrip *string `json:"StatusDescrip,omitempty" xml:"StatusDescrip,omitempty"`
	// The time when the status was last updated.
	//
	// example:
	//
	// 2021-01-26T05:04Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyAppStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyAppStatus) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyAppStatus) GetPhase() *string {
	return s.Phase
}

func (s *DescribeDeviceServiceResponseBodyAppStatus) GetStatusDescrip() *string {
	return s.StatusDescrip
}

func (s *DescribeDeviceServiceResponseBodyAppStatus) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDeviceServiceResponseBodyAppStatus) SetPhase(v string) *DescribeDeviceServiceResponseBodyAppStatus {
	s.Phase = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppStatus) SetStatusDescrip(v string) *DescribeDeviceServiceResponseBodyAppStatus {
	s.StatusDescrip = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppStatus) SetUpdateTime(v string) *DescribeDeviceServiceResponseBodyAppStatus {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyAppStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceServiceResponseBodyResourceDetailInfos struct {
	// The name of the device.
	//
	// example:
	//
	// 5JhF100NEgdBcpNren32
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the cloud device.
	//
	// example:
	//
	// h-uf6009zoaexs5pefypbo
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 120.27.219.219
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The Internet service provider (ISP).
	//
	// example:
	//
	// cmcc
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-2ze40oyhjorpyw61k7be
	ImageID *string `json:"ImageID,omitempty" xml:"ImageID,omitempty"`
	// The media access control (MAC) address of the device.
	//
	// example:
	//
	// AA:BB:77:88:99:03
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-jiaozuo-2
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// The server name of the ENS node.
	//
	// example:
	//
	// ens-nc2
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The status of the device.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the device.
	//
	// example:
	//
	// ens.ac6.large
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyResourceDetailInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyResourceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetID() *string {
	return s.ID
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetIP() *string {
	return s.IP
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetISP() *string {
	return s.ISP
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetImageID() *string {
	return s.ImageID
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetMac() *string {
	return s.Mac
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetRegionID() *string {
	return s.RegionID
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetServer() *string {
	return s.Server
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) GetType() *string {
	return s.Type
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetDeviceName(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.DeviceName = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetID(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.ID = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetIP(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.IP = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetISP(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.ISP = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetImageID(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.ImageID = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetMac(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.Mac = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetRegionID(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.RegionID = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetServer(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.Server = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetStatus(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.Status = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) SetType(v string) *DescribeDeviceServiceResponseBodyResourceDetailInfos {
	s.Type = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceDetailInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceServiceResponseBodyResourceInfos struct {
	// The version of the application.
	//
	// example:
	//
	// v1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The area code.
	//
	// example:
	//
	// 410800
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// The region name.
	//
	// example:
	//
	// Jiaozuo City, Henan Province, Central China
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2019-10-02T08:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the devices.
	DeviceInfos []*DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5s9boobrmh5000kv4jmi0oeai
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The internal IP addresses.
	InternalIps []*DescribeDeviceServiceResponseBodyResourceInfosInternalIps `json:"InternalIps,omitempty" xml:"InternalIps,omitempty" type:"Repeated"`
	// The public IP addresses.
	PublicIps []*DescribeDeviceServiceResponseBodyResourceInfosPublicIps `json:"PublicIps,omitempty" xml:"PublicIps,omitempty" type:"Repeated"`
	// The ID of the region.
	//
	// example:
	//
	// cn-jiaozuo-2
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-jiaozuo-2
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China Jiaozuo-2
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyResourceInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyResourceInfos) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetAreaCode() *string {
	return s.AreaCode
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetAreaName() *string {
	return s.AreaName
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetDeviceInfos() []*DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos {
	return s.DeviceInfos
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetInternalIps() []*DescribeDeviceServiceResponseBodyResourceInfosInternalIps {
	return s.InternalIps
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetPublicIps() []*DescribeDeviceServiceResponseBodyResourceInfosPublicIps {
	return s.PublicIps
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetAppVersion(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.AppVersion = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetAreaCode(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.AreaCode = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetAreaName(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.AreaName = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetCreateTime(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.CreateTime = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetDeviceInfos(v []*DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.DeviceInfos = v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetInstanceId(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetInstanceStatus(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetInternalIps(v []*DescribeDeviceServiceResponseBodyResourceInfosInternalIps) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.InternalIps = v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetPublicIps(v []*DescribeDeviceServiceResponseBodyResourceInfosPublicIps) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.PublicIps = v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetRegionCode(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.RegionCode = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetRegionId(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) SetRegionName(v string) *DescribeDeviceServiceResponseBodyResourceInfos {
	s.RegionName = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfos) Validate() error {
	if s.DeviceInfos != nil {
		for _, item := range s.DeviceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InternalIps != nil {
		for _, item := range s.InternalIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PublicIps != nil {
		for _, item := range s.PublicIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos struct {
	// The name of the device.
	//
	// example:
	//
	// test-api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network information.
	Network []*DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Repeated"`
	// The status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) GetName() *string {
	return s.Name
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) GetNetwork() []*DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork {
	return s.Network
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) SetName(v string) *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos {
	s.Name = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) SetNetwork(v []*DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos {
	s.Network = v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) SetStatus(v string) *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos {
	s.Status = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfos) Validate() error {
	if s.Network != nil {
		for _, item := range s.Network {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork struct {
	// The port of the container.
	//
	// example:
	//
	// 10000-10010
	ContainerPorts *string `json:"ContainerPorts,omitempty" xml:"ContainerPorts,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 39.105.62.120
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The port range.
	//
	// example:
	//
	// 80-8080
	HostPorts *string `json:"HostPorts,omitempty" xml:"HostPorts,omitempty"`
	// The protocol of the gateway. The value is of the enumeration type. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) GetContainerPorts() *string {
	return s.ContainerPorts
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) GetHostPorts() *string {
	return s.HostPorts
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) SetContainerPorts(v string) *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork {
	s.ContainerPorts = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) SetExternalIp(v string) *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork {
	s.ExternalIp = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) SetHostPorts(v string) *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork {
	s.HostPorts = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) SetProtocol(v string) *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork {
	s.Protocol = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosDeviceInfosNetwork) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceServiceResponseBodyResourceInfosInternalIps struct {
	// The internal IP address.
	//
	// example:
	//
	// 10.0.2.3
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyResourceInfosInternalIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyResourceInfosInternalIps) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosInternalIps) GetIp() *string {
	return s.Ip
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosInternalIps) SetIp(v string) *DescribeDeviceServiceResponseBodyResourceInfosInternalIps {
	s.Ip = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosInternalIps) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceServiceResponseBodyResourceInfosPublicIps struct {
	// The public IP address.
	//
	// example:
	//
	// 122.13.173.137
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDeviceServiceResponseBodyResourceInfosPublicIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponseBodyResourceInfosPublicIps) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosPublicIps) GetIp() *string {
	return s.Ip
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosPublicIps) SetIp(v string) *DescribeDeviceServiceResponseBodyResourceInfosPublicIps {
	s.Ip = &v
	return s
}

func (s *DescribeDeviceServiceResponseBodyResourceInfosPublicIps) Validate() error {
	return dara.Validate(s)
}
