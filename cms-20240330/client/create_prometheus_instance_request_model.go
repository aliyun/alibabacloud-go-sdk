// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveDuration(v int32) *CreatePrometheusInstanceRequest
	GetArchiveDuration() *int32
	SetAuthFreeReadPolicy(v string) *CreatePrometheusInstanceRequest
	GetAuthFreeReadPolicy() *string
	SetAuthFreeWritePolicy(v string) *CreatePrometheusInstanceRequest
	GetAuthFreeWritePolicy() *string
	SetEnableAuthFreeRead(v bool) *CreatePrometheusInstanceRequest
	GetEnableAuthFreeRead() *bool
	SetEnableAuthFreeWrite(v bool) *CreatePrometheusInstanceRequest
	GetEnableAuthFreeWrite() *bool
	SetEnableAuthToken(v bool) *CreatePrometheusInstanceRequest
	GetEnableAuthToken() *bool
	SetPaymentType(v string) *CreatePrometheusInstanceRequest
	GetPaymentType() *string
	SetPrometheusInstanceName(v string) *CreatePrometheusInstanceRequest
	GetPrometheusInstanceName() *string
	SetResourceGroupId(v string) *CreatePrometheusInstanceRequest
	GetResourceGroupId() *string
	SetStatus(v string) *CreatePrometheusInstanceRequest
	GetStatus() *string
	SetStorageDuration(v int32) *CreatePrometheusInstanceRequest
	GetStorageDuration() *int32
	SetTags(v []*CreatePrometheusInstanceRequestTags) *CreatePrometheusInstanceRequest
	GetTags() []*CreatePrometheusInstanceRequestTags
	SetWorkspace(v string) *CreatePrometheusInstanceRequest
	GetWorkspace() *string
}

type CreatePrometheusInstanceRequest struct {
	// The number of days that data is automatically archived after the storage period expires. A value of 0 indicates that data is not archived. Valid values for the archive duration:
	//
	// 	- V1: 60 to 365 days.
	//
	// 	- V2: 60 to 3650 days (3650 indicates permanent retention).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 60
	ArchiveDuration *int32 `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
	// The authentication-free read policy. IP CIDR blocks and VPC IDs are supported.
	//
	// example:
	//
	// {
	//
	//   "SourceIp": [
	//
	//     "192.168.1.0/24",
	//
	//     "172.168.2.22"
	//
	//   ],
	//
	//   "SourceVpc": [
	//
	//     "vpc-xx1",
	//
	//     "vpc-xx2"
	//
	//   ]
	//
	// }
	AuthFreeReadPolicy *string `json:"authFreeReadPolicy,omitempty" xml:"authFreeReadPolicy,omitempty"`
	// The authentication-free write policy.
	//
	// example:
	//
	// {
	//
	//   "SourceIp": [
	//
	//     "192.168.1.0/24",
	//
	//     "172.168.2.22"
	//
	//   ],
	//
	//   "SourceVpc": [
	//
	//     "vpc-xx1",
	//
	//     "vpc-xx2"
	//
	//   ]
	//
	// }
	AuthFreeWritePolicy *string `json:"authFreeWritePolicy,omitempty" xml:"authFreeWritePolicy,omitempty"`
	// Specifies whether to enable authentication-free read. This parameter is supported only for V2 instances.
	//
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// Specifies whether to enable authentication-free write. This parameter is supported only for V2 instances.
	//
	// example:
	//
	// true
	EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
	// Specifies whether to enable the authorization token. This parameter is supported only for V1 instances.
	//
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// The billable methods. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go by metric reporting volume.
	//
	// 	- POSTPAY_GB: pay-as-you-go by metric write volume.
	//
	// If this parameter is left empty, the default billing method configured by the user is used. If the user has not configured a default billing method, the system uses pay-as-you-go by metric reporting volume.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// name1
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz5qqvjyatgoy
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The storage duration (in days):
	//
	// 	- By write volume: 90 or 180.
	//
	// 	- By metric reporting volume: 15, 30, 60, 90, or 180.
	//
	// example:
	//
	// 90
	StorageDuration *int32 `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
	// The tags.
	Tags []*CreatePrometheusInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace to which the instance belongs. Default value: default-cms-{userId}-{regionId}.
	//
	// example:
	//
	// wokspace1
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreatePrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceRequest) GetArchiveDuration() *int32 {
	return s.ArchiveDuration
}

func (s *CreatePrometheusInstanceRequest) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *CreatePrometheusInstanceRequest) GetAuthFreeWritePolicy() *string {
	return s.AuthFreeWritePolicy
}

func (s *CreatePrometheusInstanceRequest) GetEnableAuthFreeRead() *bool {
	return s.EnableAuthFreeRead
}

func (s *CreatePrometheusInstanceRequest) GetEnableAuthFreeWrite() *bool {
	return s.EnableAuthFreeWrite
}

func (s *CreatePrometheusInstanceRequest) GetEnableAuthToken() *bool {
	return s.EnableAuthToken
}

func (s *CreatePrometheusInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreatePrometheusInstanceRequest) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *CreatePrometheusInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrometheusInstanceRequest) GetStatus() *string {
	return s.Status
}

func (s *CreatePrometheusInstanceRequest) GetStorageDuration() *int32 {
	return s.StorageDuration
}

func (s *CreatePrometheusInstanceRequest) GetTags() []*CreatePrometheusInstanceRequestTags {
	return s.Tags
}

func (s *CreatePrometheusInstanceRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreatePrometheusInstanceRequest) SetArchiveDuration(v int32) *CreatePrometheusInstanceRequest {
	s.ArchiveDuration = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetAuthFreeReadPolicy(v string) *CreatePrometheusInstanceRequest {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetAuthFreeWritePolicy(v string) *CreatePrometheusInstanceRequest {
	s.AuthFreeWritePolicy = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetEnableAuthFreeRead(v bool) *CreatePrometheusInstanceRequest {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetEnableAuthFreeWrite(v bool) *CreatePrometheusInstanceRequest {
	s.EnableAuthFreeWrite = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetEnableAuthToken(v bool) *CreatePrometheusInstanceRequest {
	s.EnableAuthToken = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetPaymentType(v string) *CreatePrometheusInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetPrometheusInstanceName(v string) *CreatePrometheusInstanceRequest {
	s.PrometheusInstanceName = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetResourceGroupId(v string) *CreatePrometheusInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetStatus(v string) *CreatePrometheusInstanceRequest {
	s.Status = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetStorageDuration(v int32) *CreatePrometheusInstanceRequest {
	s.StorageDuration = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetTags(v []*CreatePrometheusInstanceRequestTags) *CreatePrometheusInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetWorkspace(v string) *CreatePrometheusInstanceRequest {
	s.Workspace = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePrometheusInstanceRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 110109200001214284
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreatePrometheusInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreatePrometheusInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusInstanceRequestTags) SetKey(v string) *CreatePrometheusInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreatePrometheusInstanceRequestTags) SetValue(v string) *CreatePrometheusInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreatePrometheusInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
