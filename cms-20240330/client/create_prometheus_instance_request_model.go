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
	// The number of days to automatically archive and save after the storage expires, 0 means no archiving. The range of archiving days is as follows:
	//
	// 	- V1: 60~365 days.
	//
	// 	- V2: 60~3650 days (3650 indicates permanent storage).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 60
	ArchiveDuration *int32 `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
	// Password-free read policy (supports IP segments and VpcId).
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
	// Password-free write policy.
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
	// Whether to enable password-free read (only supported in V2 version).
	//
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// Whether to enable password-free write (only supported in V2 version).
	//
	// example:
	//
	// true
	EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
	// Whether to enable authorization Token (only supported in V1 version).
	//
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// Billing method:
	//
	// 	- POSTPAY: Postpaid by metric reporting volume.
	//
	// 	- POSTPAY_GB: Postpaid by metric write volume.
	//
	// Note, if left blank, the user\\"s default billing method configuration will be used. If the user has not configured a default, the system defaults to billing by metric reporting volume.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// Instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// name1
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// Instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Storage duration (days):
	//
	// 	- By write volume: 90, 180.
	//
	// 	- By metric reporting volume: 15, 30, 60, 90, 180.
	//
	// example:
	//
	// 90
	StorageDuration *int32 `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
	// Tag values.
	Tags []*CreatePrometheusInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Belonging workspace, default value: default-cms-{userId}-{regionId}.
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
	// Tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Tag value.
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
