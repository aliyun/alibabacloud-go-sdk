// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLeniPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLeniPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetLeniPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *GetLeniPrivateIpAddressResponseBodyContent) *GetLeniPrivateIpAddressResponseBody
	GetContent() *GetLeniPrivateIpAddressResponseBodyContent
	SetMessage(v string) *GetLeniPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLeniPrivateIpAddressResponseBody
	GetRequestId() *string
}

type GetLeniPrivateIpAddressResponseBody struct {
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
	Content *GetLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLeniPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLeniPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetLeniPrivateIpAddressResponseBody) GetContent() *GetLeniPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *GetLeniPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLeniPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *GetLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetCode(v int32) *GetLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetContent(v *GetLeniPrivateIpAddressResponseBodyContent) *GetLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetMessage(v string) *GetLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetRequestId(v string) *GetLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLeniPrivateIpAddressResponseBodyContent struct {
	// The description.
	//
	// example:
	//
	// zhenyuan wdl workflow
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
	// 1663722356000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 1635231890000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP address.
	//
	// example:
	//
	// 10.42.****
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
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLeniPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetDescription() *string {
	return s.Description
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetIpName() *string {
	return s.IpName
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetDescription(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetGmtCreate(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetGmtModified(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetMessage(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetPrivateIpAddress(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetRegionId(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetResourceGroupId(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetStatus(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
