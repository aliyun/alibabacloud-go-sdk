// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveDuration(v int32) *UpdatePrometheusInstanceRequest
	GetArchiveDuration() *int32
	SetAuthFreeReadPolicy(v string) *UpdatePrometheusInstanceRequest
	GetAuthFreeReadPolicy() *string
	SetAuthFreeWritePolicy(v string) *UpdatePrometheusInstanceRequest
	GetAuthFreeWritePolicy() *string
	SetClusterId(v string) *UpdatePrometheusInstanceRequest
	GetClusterId() *string
	SetEnableAuthFreeRead(v bool) *UpdatePrometheusInstanceRequest
	GetEnableAuthFreeRead() *bool
	SetEnableAuthFreeWrite(v bool) *UpdatePrometheusInstanceRequest
	GetEnableAuthFreeWrite() *bool
	SetEnableAuthToken(v bool) *UpdatePrometheusInstanceRequest
	GetEnableAuthToken() *bool
	SetPaymentType(v string) *UpdatePrometheusInstanceRequest
	GetPaymentType() *string
	SetRegionId(v string) *UpdatePrometheusInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdatePrometheusInstanceRequest
	GetResourceGroupId() *string
	SetStorageDuration(v int32) *UpdatePrometheusInstanceRequest
	GetStorageDuration() *int32
}

type UpdatePrometheusInstanceRequest struct {
	// The number of days for which data is automatically archived after the storage expires. Valid values: 60, 90, 180, and 365. 0 indicates that the data is not archived.
	//
	// example:
	//
	// 90
	ArchiveDuration *int32 `json:"ArchiveDuration,omitempty" xml:"ArchiveDuration,omitempty"`
	// The IP addresses or CIDR blocks for which password-free read is enabled. Separate multiple IP addresses with line breaks.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0.0.0.0/0
	AuthFreeReadPolicy *string `json:"AuthFreeReadPolicy,omitempty" xml:"AuthFreeReadPolicy,omitempty"`
	// The IP addresses or CIDR blocks for which password-free write is enabled. Separate multiple IP addresses with line breaks.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0.0.0.0/0
	AuthFreeWritePolicy *string `json:"AuthFreeWritePolicy,omitempty" xml:"AuthFreeWritePolicy,omitempty"`
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable password-free read.
	//
	// if can be null:
	// true
	EnableAuthFreeRead *bool `json:"EnableAuthFreeRead,omitempty" xml:"EnableAuthFreeRead,omitempty"`
	// Specifies whether to enable password-free write.
	//
	// if can be null:
	// true
	EnableAuthFreeWrite *bool `json:"EnableAuthFreeWrite,omitempty" xml:"EnableAuthFreeWrite,omitempty"`
	// Specifies whether to enable access token authentication.
	//
	// if can be null:
	// true
	EnableAuthToken *bool `json:"EnableAuthToken,omitempty" xml:"EnableAuthToken,omitempty"`
	// The billing mode. Valid values: POSTPAY: charges fees based on the amount of reported metric data. POSTPAY_GB: charges fees based on the amount of written metric data.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Prometheus resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The data storage duration. Unit: days.
	//
	// example:
	//
	// 90
	StorageDuration *int32 `json:"StorageDuration,omitempty" xml:"StorageDuration,omitempty"`
}

func (s UpdatePrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusInstanceRequest) GetArchiveDuration() *int32 {
	return s.ArchiveDuration
}

func (s *UpdatePrometheusInstanceRequest) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *UpdatePrometheusInstanceRequest) GetAuthFreeWritePolicy() *string {
	return s.AuthFreeWritePolicy
}

func (s *UpdatePrometheusInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusInstanceRequest) GetEnableAuthFreeRead() *bool {
	return s.EnableAuthFreeRead
}

func (s *UpdatePrometheusInstanceRequest) GetEnableAuthFreeWrite() *bool {
	return s.EnableAuthFreeWrite
}

func (s *UpdatePrometheusInstanceRequest) GetEnableAuthToken() *bool {
	return s.EnableAuthToken
}

func (s *UpdatePrometheusInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *UpdatePrometheusInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrometheusInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdatePrometheusInstanceRequest) GetStorageDuration() *int32 {
	return s.StorageDuration
}

func (s *UpdatePrometheusInstanceRequest) SetArchiveDuration(v int32) *UpdatePrometheusInstanceRequest {
	s.ArchiveDuration = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetAuthFreeReadPolicy(v string) *UpdatePrometheusInstanceRequest {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetAuthFreeWritePolicy(v string) *UpdatePrometheusInstanceRequest {
	s.AuthFreeWritePolicy = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetClusterId(v string) *UpdatePrometheusInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetEnableAuthFreeRead(v bool) *UpdatePrometheusInstanceRequest {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetEnableAuthFreeWrite(v bool) *UpdatePrometheusInstanceRequest {
	s.EnableAuthFreeWrite = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetEnableAuthToken(v bool) *UpdatePrometheusInstanceRequest {
	s.EnableAuthToken = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetPaymentType(v string) *UpdatePrometheusInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetRegionId(v string) *UpdatePrometheusInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetResourceGroupId(v string) *UpdatePrometheusInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetStorageDuration(v int32) *UpdatePrometheusInstanceRequest {
	s.StorageDuration = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) Validate() error {
	return dara.Validate(s)
}
