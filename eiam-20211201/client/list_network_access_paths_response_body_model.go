// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessPathsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAccessPaths(v []*ListNetworkAccessPathsResponseBodyNetworkAccessPaths) *ListNetworkAccessPathsResponseBody
	GetNetworkAccessPaths() []*ListNetworkAccessPathsResponseBodyNetworkAccessPaths
	SetRequestId(v string) *ListNetworkAccessPathsResponseBody
	GetRequestId() *string
}

type ListNetworkAccessPathsResponseBody struct {
	// The list of network endpoint access paths.
	NetworkAccessPaths []*ListNetworkAccessPathsResponseBodyNetworkAccessPaths `json:"NetworkAccessPaths,omitempty" xml:"NetworkAccessPaths,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetworkAccessPathsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessPathsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsResponseBody) GetNetworkAccessPaths() []*ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	return s.NetworkAccessPaths
}

func (s *ListNetworkAccessPathsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkAccessPathsResponseBody) SetNetworkAccessPaths(v []*ListNetworkAccessPathsResponseBodyNetworkAccessPaths) *ListNetworkAccessPathsResponseBody {
	s.NetworkAccessPaths = v
	return s
}

func (s *ListNetworkAccessPathsResponseBody) SetRequestId(v string) *ListNetworkAccessPathsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBody) Validate() error {
	if s.NetworkAccessPaths != nil {
		for _, item := range s.NetworkAccessPaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkAccessPathsResponseBodyNetworkAccessPaths struct {
	// The time when the dedicated network endpoint access path was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the dedicated network endpoint.
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The ID of the dedicated network endpoint access path.
	//
	// example:
	//
	// nap_examplexxx
	NetworkAccessPathId *string `json:"NetworkAccessPathId,omitempty" xml:"NetworkAccessPathId,omitempty"`
	// The ID of the ENI that is used by the dedicated network endpoint access path.
	//
	// example:
	//
	// eni-examplexxx
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The private IP address of the ENI that is used by the dedicated network endpoint access path.
	//
	// example:
	//
	// cn-hangzhou
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The status of the dedicated network endpoint access path. Valid values:
	//
	// - pending: The path is being initialized.
	//
	// - creating: The path is being created.
	//
	// - running: The path is running.
	//
	// - deleting: The path is being deleted.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the dedicated network endpoint access path was last updated. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the vSwitch to which the ENI of the dedicated network endpoint access path belongs.
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ListNetworkAccessPathsResponseBodyNetworkAccessPaths) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetNetworkAccessPathId() *string {
	return s.NetworkAccessPathId
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetStatus() *string {
	return s.Status
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetCreateTime(v int64) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetInstanceId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetNetworkAccessEndpointId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetNetworkAccessPathId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.NetworkAccessPathId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetNetworkInterfaceId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetPrivateIpAddress(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetStatus(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.Status = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetUpdateTime(v int64) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.UpdateTime = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetVSwitchId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.VSwitchId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) Validate() error {
	return dara.Validate(s)
}
