// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReceiveAllRoute(v bool) *ListErAttachmentsRequest
	GetAutoReceiveAllRoute() *bool
	SetEnablePage(v bool) *ListErAttachmentsRequest
	GetEnablePage() *bool
	SetErAttachmentId(v string) *ListErAttachmentsRequest
	GetErAttachmentId() *string
	SetErAttachmentName(v string) *ListErAttachmentsRequest
	GetErAttachmentName() *string
	SetErId(v string) *ListErAttachmentsRequest
	GetErId() *string
	SetInstanceId(v string) *ListErAttachmentsRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ListErAttachmentsRequest
	GetInstanceType() *string
	SetPageNumber(v int32) *ListErAttachmentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListErAttachmentsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListErAttachmentsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListErAttachmentsRequest
	GetResourceGroupId() *string
	SetResourceTenantId(v string) *ListErAttachmentsRequest
	GetResourceTenantId() *string
	SetStatus(v string) *ListErAttachmentsRequest
	GetStatus() *string
}

type ListErAttachmentsRequest struct {
	// Whether to automatically receive all routes from all instances under this Lingjun HUB. Valid values:
	//
	// 	- **true**: received automatically.
	//
	// 	- **false**: Not received.
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// Specifies whether to enable paged query. Valid values:
	//
	// 	- **true**: enables paged query.
	//
	// 	- **false**: Paged query is not enabled.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// The ID of the network instance connection
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// The name of the network instance connection.
	//
	// example:
	//
	// vcc-cn-209300qha01
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the network instance. Valid values: **VPD*	- and **VCC**.
	//
	// For more information, see [What is Lingjun?](https://help.aliyun.com/document_detail/444430.html)
	//
	// You can query **Lingjun CIDR blocks*	- and **Lingjun connections*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html?) respectively.
	//
	// example:
	//
	// vcc-cn-209300qha01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mitigation plan of the instance. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// example:
	//
	// VCC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The page number to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the instance belongs.
	//
	// example:
	//
	// 1111156667137893
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListErAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListErAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsRequest) GetAutoReceiveAllRoute() *bool {
	return s.AutoReceiveAllRoute
}

func (s *ListErAttachmentsRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListErAttachmentsRequest) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *ListErAttachmentsRequest) GetErAttachmentName() *string {
	return s.ErAttachmentName
}

func (s *ListErAttachmentsRequest) GetErId() *string {
	return s.ErId
}

func (s *ListErAttachmentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListErAttachmentsRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListErAttachmentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListErAttachmentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListErAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErAttachmentsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErAttachmentsRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListErAttachmentsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListErAttachmentsRequest) SetAutoReceiveAllRoute(v bool) *ListErAttachmentsRequest {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *ListErAttachmentsRequest) SetEnablePage(v bool) *ListErAttachmentsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErAttachmentsRequest) SetErAttachmentId(v string) *ListErAttachmentsRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetErAttachmentName(v string) *ListErAttachmentsRequest {
	s.ErAttachmentName = &v
	return s
}

func (s *ListErAttachmentsRequest) SetErId(v string) *ListErAttachmentsRequest {
	s.ErId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetInstanceId(v string) *ListErAttachmentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetInstanceType(v string) *ListErAttachmentsRequest {
	s.InstanceType = &v
	return s
}

func (s *ListErAttachmentsRequest) SetPageNumber(v int32) *ListErAttachmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErAttachmentsRequest) SetPageSize(v int32) *ListErAttachmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListErAttachmentsRequest) SetRegionId(v string) *ListErAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetResourceGroupId(v string) *ListErAttachmentsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetResourceTenantId(v string) *ListErAttachmentsRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetStatus(v string) *ListErAttachmentsRequest {
	s.Status = &v
	return s
}

func (s *ListErAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}
