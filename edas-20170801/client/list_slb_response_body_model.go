// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSlbResponseBody
	GetCode() *int32
	SetMessage(v string) *ListSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSlbResponseBody
	GetRequestId() *string
	SetSlbList(v *ListSlbResponseBodySlbList) *ListSlbResponseBody
	GetSlbList() *ListSlbResponseBodySlbList
}

type ListSlbResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of SLB instances.
	SlbList *ListSlbResponseBodySlbList `json:"SlbList,omitempty" xml:"SlbList,omitempty" type:"Struct"`
}

func (s ListSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSlbResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSlbResponseBody) GetSlbList() *ListSlbResponseBodySlbList {
	return s.SlbList
}

func (s *ListSlbResponseBody) SetCode(v int32) *ListSlbResponseBody {
	s.Code = &v
	return s
}

func (s *ListSlbResponseBody) SetMessage(v string) *ListSlbResponseBody {
	s.Message = &v
	return s
}

func (s *ListSlbResponseBody) SetRequestId(v string) *ListSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlbResponseBody) SetSlbList(v *ListSlbResponseBodySlbList) *ListSlbResponseBody {
	s.SlbList = v
	return s
}

func (s *ListSlbResponseBody) Validate() error {
	if s.SlbList != nil {
		if err := s.SlbList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSlbResponseBodySlbList struct {
	SlbEntity []*ListSlbResponseBodySlbListSlbEntity `json:"SlbEntity,omitempty" xml:"SlbEntity,omitempty" type:"Repeated"`
}

func (s ListSlbResponseBodySlbList) String() string {
	return dara.Prettify(s)
}

func (s ListSlbResponseBodySlbList) GoString() string {
	return s.String()
}

func (s *ListSlbResponseBodySlbList) GetSlbEntity() []*ListSlbResponseBodySlbListSlbEntity {
	return s.SlbEntity
}

func (s *ListSlbResponseBodySlbList) SetSlbEntity(v []*ListSlbResponseBodySlbListSlbEntity) *ListSlbResponseBodySlbList {
	s.SlbEntity = v
	return s
}

func (s *ListSlbResponseBodySlbList) Validate() error {
	if s.SlbEntity != nil {
		for _, item := range s.SlbEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSlbResponseBodySlbListSlbEntity struct {
	// The IP address of the SLB instance.
	//
	// example:
	//
	// 39.176.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The type of the IP addresses. Valid values:
	//
	// 	- internet: Users can connect to the SLB instance over the Internet.
	//
	// 	- intranet: Users can connect to the SLB instance over the internal network.
	//
	// example:
	//
	// internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// Indicates whether the SLB instance has expired. Valid values:
	//
	// 	- true: The SLB instance has expired.
	//
	// 	- false: The SLB instance has not expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The ID of the resource group in Enterprise Distributed Application Service (EDAS).
	//
	// example:
	//
	// 0
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The type of the network.
	//
	// example:
	//
	// classic
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether Kubernetes applications can be reused. Valid values:
	//
	// 	- true: Kubernetes applications can be reused.
	//
	// 	- false: Kubernetes applications cannot be reused.
	//
	// example:
	//
	// true
	Reusable *bool `json:"Reusable,omitempty" xml:"Reusable,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-2ze055t3xv7s8****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The name of the SLB instance.
	//
	// example:
	//
	// adce
	SlbName *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
	// The status of the SLB instance.
	//
	// example:
	//
	// active
	SlbStatus *string `json:"SlbStatus,omitempty" xml:"SlbStatus,omitempty"`
	// The tag of the SLB instance.
	//
	// example:
	//
	// [{"tagKey":"tag","tagValue":"value"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// edas_****_**st@aliyun-****.com
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1f90rfybszjogyw****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch in the VPC.
	//
	// example:
	//
	// vsw-bp156w1gpbv0o50hs****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListSlbResponseBodySlbListSlbEntity) String() string {
	return dara.Prettify(s)
}

func (s ListSlbResponseBodySlbListSlbEntity) GoString() string {
	return s.String()
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetAddress() *string {
	return s.Address
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetAddressType() *string {
	return s.AddressType
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetExpired() *bool {
	return s.Expired
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetGroupId() *int32 {
	return s.GroupId
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetReusable() *bool {
	return s.Reusable
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetSlbId() *string {
	return s.SlbId
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetSlbName() *string {
	return s.SlbName
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetSlbStatus() *string {
	return s.SlbStatus
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetTags() *string {
	return s.Tags
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListSlbResponseBodySlbListSlbEntity) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetAddress(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.Address = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetAddressType(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.AddressType = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetExpired(v bool) *ListSlbResponseBodySlbListSlbEntity {
	s.Expired = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetGroupId(v int32) *ListSlbResponseBodySlbListSlbEntity {
	s.GroupId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetNetworkType(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.NetworkType = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetRegionId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.RegionId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetReusable(v bool) *ListSlbResponseBodySlbListSlbEntity {
	s.Reusable = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetSlbId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.SlbId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetSlbName(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.SlbName = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetSlbStatus(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.SlbStatus = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetTags(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.Tags = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetUserId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.UserId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetVpcId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.VpcId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetVswitchId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.VswitchId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) Validate() error {
	return dara.Validate(s)
}
