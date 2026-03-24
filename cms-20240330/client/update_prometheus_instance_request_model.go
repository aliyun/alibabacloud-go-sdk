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
	// The number of days to store archived data after the storage duration expires. A value of 0 disables archiving. For V1 instances, the valid values are 1 to 365. This is supported only for the pay-by-data-write billing method. For V2 instances, the valid values are 1 to 3650. A value of 3650 indicates permanent storage.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 365
	ArchiveDuration *int `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
	// The policy for password-free read access. The policy supports IP address segments and VPC IDs.
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
	// The policy for password-free write access. The policy supports IP address segments and VPC IDs.
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
	// Specifies whether to enable password-free read access.
	//
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// Specifies whether to enable password-free write access.
	//
	// example:
	//
	// true
	EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
	// Specifies whether to enable authentication with an access token.
	//
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// The billing method. You can change the billing method only once during the instance lifecycle. Valid values: \\`POSTPAY\\` (pay-as-you-go based on reported metrics) and \\`POSTPAY_GB\\` (pay-as-you-go based on data writes).
	//
	// example:
	//
	// POSTPAY_GB
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test-prom-name
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// The status of the instance storage database. Only RUNNING is supported. If this parameter is left empty, the status of the storage database is not changed.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The storage duration in days. If the instance is billed by data writes, valid values are 90 and 180. If the instance is billed by reported metrics, valid values are 15, 30, 60, 90, and 180.
	//
	// example:
	//
	// 90
	StorageDuration *int `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
	// The workspace to which the instance belongs.
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
