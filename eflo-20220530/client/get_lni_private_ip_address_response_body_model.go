// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLniPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLniPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetLniPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *GetLniPrivateIpAddressResponseBodyContent) *GetLniPrivateIpAddressResponseBody
	GetContent() *GetLniPrivateIpAddressResponseBodyContent
	SetMessage(v string) *GetLniPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLniPrivateIpAddressResponseBody
	GetRequestId() *string
}

type GetLniPrivateIpAddressResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetLniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:GetLniPrivateIpAddress, arn=acs:eflo:cn-wulanchabu:1382782317087063:networkinterface/00
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLniPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLniPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetLniPrivateIpAddressResponseBody) GetContent() *GetLniPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *GetLniPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLniPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *GetLniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) SetCode(v int32) *GetLniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) SetContent(v *GetLniPrivateIpAddressResponseBodyContent) *GetLniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) SetMessage(v string) *GetLniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) SetRequestId(v string) *GetLniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLniPrivateIpAddressResponseBodyContent struct {
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 2022-12-26 20:16:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// MAC address of the secondary private network
	//
	// example:
	//
	// 00-ff-84-15-ba-67
	IpAddressMac *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	// IP unique identifier
	//
	// example:
	//
	// sip-xxxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:ListVpdRouteEntries, arn=acs:eflo:cn-wulanchabu:1263399219805497:vpd_rte/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-2ze4uww7n6hsfzrwq77y
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The secondary private IP address of the Lingjun network interface controller.
	//
	// example:
	//
	// 10.42.5.92
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subnet instance ID.
	//
	// example:
	//
	// subnet-aj93mko8
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s GetLniPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetLniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetDescription() *string {
	return s.Description
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetIpAddressMac() *string {
	return s.IpAddressMac
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetIpName() *string {
	return s.IpName
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetLniPrivateIpAddressResponseBodyContent) GetSubnetId() *string {
	return s.SubnetId
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetDescription(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetGmtCreate(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetIpAddressMac(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.IpAddressMac = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetIpName(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetMessage(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetNetworkInterfaceId(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetPrivateIpAddress(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetRegionId(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetResourceGroupId(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetStatus(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetSubnetId(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.SubnetId = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
