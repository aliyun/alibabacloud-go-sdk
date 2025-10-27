// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterCnnfStatusDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListClusterCnnfStatusDetailResponseBodyData) *ListClusterCnnfStatusDetailResponseBody
	GetData() []*ListClusterCnnfStatusDetailResponseBodyData
	SetRequestId(v string) *ListClusterCnnfStatusDetailResponseBody
	GetRequestId() *string
}

type ListClusterCnnfStatusDetailResponseBody struct {
	// An array that consists of the protection status of the container firewall.
	Data []*ListClusterCnnfStatusDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 64329F40-5C94-51D3-A400-37AA7BAC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterCnnfStatusDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCnnfStatusDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterCnnfStatusDetailResponseBody) GetData() []*ListClusterCnnfStatusDetailResponseBodyData {
	return s.Data
}

func (s *ListClusterCnnfStatusDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterCnnfStatusDetailResponseBody) SetData(v []*ListClusterCnnfStatusDetailResponseBodyData) *ListClusterCnnfStatusDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBody) SetRequestId(v string) *ListClusterCnnfStatusDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterCnnfStatusDetailResponseBodyData struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Indicates whether the container firewall plug-in is installed.
	//
	// example:
	//
	// true
	Installed *bool `json:"Installed,omitempty" xml:"Installed,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-bp180bogui4fc0z4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public IP address of the associated instance.
	//
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the associated instance.
	//
	// example:
	//
	// 10.42.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The cause why the plug-in is invalid. Valid values:
	//
	// 	- **PLUGIN_OFFLINE**: The plug-in is offline.
	//
	// 	- **PLUGIN_NOT_INSTALLED**: The plug-in is not installed.
	//
	// 	- **PLUGIN_INVALID_VERSION**: The version of the plug-in is invalid.
	//
	// example:
	//
	// PLUGIN_OFFLINE
	InvalidType *string `json:"InvalidType,omitempty" xml:"InvalidType,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// security
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The machine type of the instance. The value is fixed as **ecs**.
	//
	// example:
	//
	// ecs
	MachineType *int32 `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The name of the plug-in. The value is fixed as **alinet**.
	//
	// example:
	//
	// alinet
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The version of the plug-in.
	//
	// example:
	//
	// 3.3
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
	// The online status of the plug-in. Valid values:
	//
	// 	- **false**: The plug-in is offline.
	//
	// 	- **true**: The plug-in is online.
	//
	// example:
	//
	// false
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListClusterCnnfStatusDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCnnfStatusDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetInstalled() *bool {
	return s.Installed
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetInvalidType() *string {
	return s.InvalidType
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetMachineName() *string {
	return s.MachineName
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetMachineType() *int32 {
	return s.MachineType
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetPluginName() *string {
	return s.PluginName
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetClusterId(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetInstalled(v bool) *ListClusterCnnfStatusDetailResponseBodyData {
	s.Installed = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetInstanceId(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetInternetIp(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetIntranetIp(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetInvalidType(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.InvalidType = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetMachineName(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.MachineName = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetMachineType(v int32) *ListClusterCnnfStatusDetailResponseBodyData {
	s.MachineType = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetPluginName(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.PluginName = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetPluginVersion(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.PluginVersion = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetStatus(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) SetUuid(v string) *ListClusterCnnfStatusDetailResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListClusterCnnfStatusDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
