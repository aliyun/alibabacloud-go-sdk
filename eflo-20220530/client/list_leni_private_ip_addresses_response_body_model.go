// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLeniPrivateIpAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListLeniPrivateIpAddressesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListLeniPrivateIpAddressesResponseBody
	GetCode() *int32
	SetContent(v *ListLeniPrivateIpAddressesResponseBodyContent) *ListLeniPrivateIpAddressesResponseBody
	GetContent() *ListLeniPrivateIpAddressesResponseBodyContent
	SetMessage(v string) *ListLeniPrivateIpAddressesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLeniPrivateIpAddressesResponseBody
	GetRequestId() *string
}

type ListLeniPrivateIpAddressesResponseBody struct {
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
	// The response data.
	Content *ListLeniPrivateIpAddressesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListLeniPrivateIpAddressesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListLeniPrivateIpAddressesResponseBody) GetContent() *ListLeniPrivateIpAddressesResponseBodyContent {
	return s.Content
}

func (s *ListLeniPrivateIpAddressesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLeniPrivateIpAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetAccessDeniedDetail(v string) *ListLeniPrivateIpAddressesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetCode(v int32) *ListLeniPrivateIpAddressesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetContent(v *ListLeniPrivateIpAddressesResponseBodyContent) *ListLeniPrivateIpAddressesResponseBody {
	s.Content = v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetMessage(v string) *ListLeniPrivateIpAddressesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetRequestId(v string) *ListLeniPrivateIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLeniPrivateIpAddressesResponseBodyContent struct {
	// The response parameters.
	Data []*ListLeniPrivateIpAddressesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) GetData() []*ListLeniPrivateIpAddressesResponseBodyContentData {
	return s.Data
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) SetData(v []*ListLeniPrivateIpAddressesResponseBodyContentData) *ListLeniPrivateIpAddressesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) SetResourceGroupId(v string) *ListLeniPrivateIpAddressesResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) SetTotal(v int64) *ListLeniPrivateIpAddressesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) Validate() error {
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

type ListLeniPrivateIpAddressesResponseBodyContentData struct {
	// The description.
	//
	// example:
	//
	// test_vpn1_pbr_route_54
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1675929918000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 1675929918000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP address.
	//
	// example:
	//
	// 10.0.****
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
	// rg-acfmzaq3ypaqkdy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetDescription() *string {
	return s.Description
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetIpName() *string {
	return s.IpName
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetDescription(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetElasticNetworkInterfaceId(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetGmtCreate(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.GmtCreate = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetGmtModified(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetIpName(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.IpName = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetMessage(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetPrivateIpAddress(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetRegionId(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetResourceGroupId(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetStatus(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
