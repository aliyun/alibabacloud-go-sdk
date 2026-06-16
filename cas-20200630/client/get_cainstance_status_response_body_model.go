// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCAInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceStatusList(v []*GetCAInstanceStatusResponseBodyInstanceStatusList) *GetCAInstanceStatusResponseBody
	GetInstanceStatusList() []*GetCAInstanceStatusResponseBodyInstanceStatusList
	SetRequestId(v string) *GetCAInstanceStatusResponseBody
	GetRequestId() *string
}

type GetCAInstanceStatusResponseBody struct {
	// The status details of the private CA instance.
	InstanceStatusList []*GetCAInstanceStatusResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25589516-2A56-5159-AB88-4A1D9824E183
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCAInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCAInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusResponseBody) GetInstanceStatusList() []*GetCAInstanceStatusResponseBodyInstanceStatusList {
	return s.InstanceStatusList
}

func (s *GetCAInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCAInstanceStatusResponseBody) SetInstanceStatusList(v []*GetCAInstanceStatusResponseBodyInstanceStatusList) *GetCAInstanceStatusResponseBody {
	s.InstanceStatusList = v
	return s
}

func (s *GetCAInstanceStatusResponseBody) SetRequestId(v string) *GetCAInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCAInstanceStatusResponseBody) Validate() error {
	if s.InstanceStatusList != nil {
		for _, item := range s.InstanceStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCAInstanceStatusResponseBodyInstanceStatusList struct {
	// The expiration date of the private CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// > This parameter is returned only if **Status*	- is **USED*	- (the private CA instance is enabled) or **REVOKE*	- (the private CA instance is revoked).
	//
	// example:
	//
	// 1792944000000
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The issuance date of the private CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// > This parameter is returned only if **Status*	- is **USED*	- (the private CA instance is enabled) or **REVOKE*	- (the private CA instance is revoked).
	//
	// example:
	//
	// 1635177600000
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The number of certificates that the private CA instance has issued.
	//
	// example:
	//
	// 1
	CertIssuedCount *int32 `json:"CertIssuedCount,omitempty" xml:"CertIssuedCount,omitempty"`
	// The number of certificates that the private CA instance can issue.
	//
	// If the private CA is a root CA (**Type*	- is **ROOT**), this parameter indicates the number of intermediate CA certificates that can be issued. If the private CA is an intermediate CA (**Type*	- is **SUB_ROOT**), this parameter indicates the total number of client certificates and server-side certificates that can be issued.
	//
	// example:
	//
	// 10
	CertTotalCount *int32 `json:"CertTotalCount,omitempty" xml:"CertTotalCount,omitempty"`
	// The unique identifier of the private CA certificate.
	//
	// > This parameter is returned only if **Status*	- is **USED*	- (the private CA instance is enabled) or **REVOKE*	- (the private CA instance is revoked).
	//
	// example:
	//
	// a7bb2dd212a2112128cd5cc9b753****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the private CA instance.
	//
	// example:
	//
	// cas-member-0hmi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the private CA instance. Valid values:
	//
	// - **BUY**: The instance is purchased but not enabled.
	//
	// - **USED**: The instance is enabled.
	//
	// - **REFUND**: A refund has been issued for the instance.
	//
	// - **REVOKE**: The instance is revoked.
	//
	// example:
	//
	// USED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the private CA instance. Valid values:
	//
	// - **ROOT**: Root CA instance.
	//
	// - **SUB_ROOT**: Intermediate CA instance.
	//
	// example:
	//
	// ROOT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The expiration date of the private CA instance. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// > This parameter corresponds to the subscription duration that you selected for the Private Certificate Authority (PCA) service when you purchased the instance.
	//
	// example:
	//
	// 1637251200000
	UseExpireTime *int64 `json:"UseExpireTime,omitempty" xml:"UseExpireTime,omitempty"`
}

func (s GetCAInstanceStatusResponseBodyInstanceStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetCAInstanceStatusResponseBodyInstanceStatusList) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetCertIssuedCount() *int32 {
	return s.CertIssuedCount
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetCertTotalCount() *int32 {
	return s.CertTotalCount
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetStatus() *string {
	return s.Status
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetType() *string {
	return s.Type
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) GetUseExpireTime() *int64 {
	return s.UseExpireTime
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetAfterTime(v int64) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.AfterTime = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetBeforeTime(v int64) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.BeforeTime = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetCertIssuedCount(v int32) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.CertIssuedCount = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetCertTotalCount(v int32) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.CertTotalCount = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetIdentifier(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.Identifier = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetInstanceId(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.InstanceId = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetStatus(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.Status = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetType(v string) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.Type = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) SetUseExpireTime(v int64) *GetCAInstanceStatusResponseBodyInstanceStatusList {
	s.UseExpireTime = &v
	return s
}

func (s *GetCAInstanceStatusResponseBodyInstanceStatusList) Validate() error {
	return dara.Validate(s)
}
