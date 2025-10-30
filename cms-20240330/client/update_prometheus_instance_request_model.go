// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveDuration(v int) *UpdatePrometheusInstanceRequest
	GetArchiveDuration() *int
	SetAuthFreeReadPolicy(v string) *UpdatePrometheusInstanceRequest
	GetAuthFreeReadPolicy() *string
	SetAuthFreeWritePolicy(v string) *UpdatePrometheusInstanceRequest
	GetAuthFreeWritePolicy() *string
	SetEnableAuthFreeRead(v bool) *UpdatePrometheusInstanceRequest
	GetEnableAuthFreeRead() *bool
	SetEnableAuthFreeWrite(v bool) *UpdatePrometheusInstanceRequest
	GetEnableAuthFreeWrite() *bool
	SetEnableAuthToken(v bool) *UpdatePrometheusInstanceRequest
	GetEnableAuthToken() *bool
	SetPaymentType(v string) *UpdatePrometheusInstanceRequest
	GetPaymentType() *string
	SetPrometheusInstanceName(v string) *UpdatePrometheusInstanceRequest
	GetPrometheusInstanceName() *string
	SetStatus(v string) *UpdatePrometheusInstanceRequest
	GetStatus() *string
	SetStorageDuration(v int) *UpdatePrometheusInstanceRequest
	GetStorageDuration() *int
	SetWorkspace(v string) *UpdatePrometheusInstanceRequest
	GetWorkspace() *string
}

type UpdatePrometheusInstanceRequest struct {
	// The number of days to automatically archive and save after the storage expires, 0 means no archiving. The range of archiving days:
	//
	// V1: 1~365 days. Only supported for metric write volume.
	//
	// V2: 1~3650 days (3650 indicates permanent storage).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 365
	ArchiveDuration *int `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
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
	// Password-free write policy (supports IP segments and VpcId).
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
	// Whether to enable password-free read.
	//
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// Whether to enable password-free write.
	//
	// example:
	//
	// true
	EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
	// Whether to enable access token authentication.
	//
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// Billing method (can only be modified once during the instance\\"s lifecycle):
	//
	// POSTPAY: Postpaid by metric reporting volume.
	//
	// POSTPAY_GB: Postpaid by metric write volume.
	//
	// example:
	//
	// POSTPAY_GB
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// Instance name.
	//
	// example:
	//
	// test-prom-name
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// Instance storage DB status (only supports RUNNING). If empty, the storage DB status will not be changed.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Storage duration (days):
	//
	// By write volume: 90, 180.
	//
	// By metric reporting volume: 15, 30, 60, 90, 180.
	//
	// example:
	//
	// 90
	StorageDuration *int `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
	// Belonging workspace.
	//
	// example:
	//
	// default-cms-1500199863951574-cn-shanghai
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdatePrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusInstanceRequest) GetArchiveDuration() *int {
	return s.ArchiveDuration
}

func (s *UpdatePrometheusInstanceRequest) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *UpdatePrometheusInstanceRequest) GetAuthFreeWritePolicy() *string {
	return s.AuthFreeWritePolicy
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

func (s *UpdatePrometheusInstanceRequest) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *UpdatePrometheusInstanceRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrometheusInstanceRequest) GetStorageDuration() *int {
	return s.StorageDuration
}

func (s *UpdatePrometheusInstanceRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePrometheusInstanceRequest) SetArchiveDuration(v int) *UpdatePrometheusInstanceRequest {
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

func (s *UpdatePrometheusInstanceRequest) SetPrometheusInstanceName(v string) *UpdatePrometheusInstanceRequest {
	s.PrometheusInstanceName = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetStatus(v string) *UpdatePrometheusInstanceRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetStorageDuration(v int) *UpdatePrometheusInstanceRequest {
	s.StorageDuration = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) SetWorkspace(v string) *UpdatePrometheusInstanceRequest {
	s.Workspace = &v
	return s
}

func (s *UpdatePrometheusInstanceRequest) Validate() error {
	return dara.Validate(s)
}
