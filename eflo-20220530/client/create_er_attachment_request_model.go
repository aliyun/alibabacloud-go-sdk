// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReceiveAllRoute(v bool) *CreateErAttachmentRequest
	GetAutoReceiveAllRoute() *bool
	SetErAttachmentName(v string) *CreateErAttachmentRequest
	GetErAttachmentName() *string
	SetErId(v string) *CreateErAttachmentRequest
	GetErId() *string
	SetInstanceId(v string) *CreateErAttachmentRequest
	GetInstanceId() *string
	SetInstanceType(v string) *CreateErAttachmentRequest
	GetInstanceType() *string
	SetRegionId(v string) *CreateErAttachmentRequest
	GetRegionId() *string
	SetResourceTenantId(v string) *CreateErAttachmentRequest
	GetResourceTenantId() *string
}

type CreateErAttachmentRequest struct {
	// Indicates whether to automatically receive all routes from all instances under the Lingjun HUB. Valid values:
	//
	// 	- **true**: received automatically.
	//
	// 	- **false**: Not received.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// The name of the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachemnt-vpd-xksd2obl
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// Lingjun HUB ID.
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
	// You can query **Lingjun CIDR Block*	- and **Lingjun Connection*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html?) respectively.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xksd2obl
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The category of the instance. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPD
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the tenant to which the resource belongs. This parameter is required for cross-account resources.
	//
	// example:
	//
	// 1620939556166277
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
}

func (s CreateErAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentRequest) GetAutoReceiveAllRoute() *bool {
	return s.AutoReceiveAllRoute
}

func (s *CreateErAttachmentRequest) GetErAttachmentName() *string {
	return s.ErAttachmentName
}

func (s *CreateErAttachmentRequest) GetErId() *string {
	return s.ErId
}

func (s *CreateErAttachmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateErAttachmentRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateErAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateErAttachmentRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *CreateErAttachmentRequest) SetAutoReceiveAllRoute(v bool) *CreateErAttachmentRequest {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *CreateErAttachmentRequest) SetErAttachmentName(v string) *CreateErAttachmentRequest {
	s.ErAttachmentName = &v
	return s
}

func (s *CreateErAttachmentRequest) SetErId(v string) *CreateErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *CreateErAttachmentRequest) SetInstanceId(v string) *CreateErAttachmentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateErAttachmentRequest) SetInstanceType(v string) *CreateErAttachmentRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateErAttachmentRequest) SetRegionId(v string) *CreateErAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateErAttachmentRequest) SetResourceTenantId(v string) *CreateErAttachmentRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *CreateErAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
