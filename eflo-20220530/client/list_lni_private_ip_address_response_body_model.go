// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLniPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListLniPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListLniPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *ListLniPrivateIpAddressResponseBodyContent) *ListLniPrivateIpAddressResponseBody
	GetContent() *ListLniPrivateIpAddressResponseBodyContent
	SetMessage(v string) *ListLniPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLniPrivateIpAddressResponseBody
	GetRequestId() *string
}

type ListLniPrivateIpAddressResponseBody struct {
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
	Content *ListLniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListLniPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListLniPrivateIpAddressResponseBody) GetContent() *ListLniPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *ListLniPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLniPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *ListLniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) SetCode(v int32) *ListLniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) SetContent(v *ListLniPrivateIpAddressResponseBodyContent) *ListLniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) SetMessage(v string) *ListLniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) SetRequestId(v string) *ListLniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLniPrivateIpAddressResponseBodyContent struct {
	// The returned result.
	Data []*ListLniPrivateIpAddressResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBodyContent) GetData() []*ListLniPrivateIpAddressResponseBodyContentData {
	return s.Data
}

func (s *ListLniPrivateIpAddressResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLniPrivateIpAddressResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListLniPrivateIpAddressResponseBodyContent) SetData(v []*ListLniPrivateIpAddressResponseBodyContentData) *ListLniPrivateIpAddressResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContent) SetResourceGroupId(v string) *ListLniPrivateIpAddressResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContent) SetTotal(v int64) *ListLniPrivateIpAddressResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContent) Validate() error {
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

type ListLniPrivateIpAddressResponseBodyContentData struct {
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
	// 1651734291000
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
	// sip-1hq1ql7vz
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-bp11hq1ql7vza3k4xz7q
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// Secondary private IP address of Lingjun network interface controller
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
	// rg-aekzt452sjgqm2y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetDescription() *string {
	return s.Description
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetIpAddressMac() *string {
	return s.IpAddressMac
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetIpName() *string {
	return s.IpName
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetDescription(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetGmtCreate(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.GmtCreate = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetIpAddressMac(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.IpAddressMac = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetIpName(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.IpName = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetMessage(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetNetworkInterfaceId(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetPrivateIpAddress(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetRegionId(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetResourceGroupId(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetStatus(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
