// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressRecoverSuspendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecord(v *AddAddressRecoverSuspendRequestAuditRecord) *AddAddressRecoverSuspendRequest
	GetAuditRecord() *AddAddressRecoverSuspendRequestAuditRecord
	SetCustSpaceId(v string) *AddAddressRecoverSuspendRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *AddAddressRecoverSuspendRequest
	GetOwnerId() *int64
	SetRequestType(v string) *AddAddressRecoverSuspendRequest
	GetRequestType() *string
	SetResourceOwnerAccount(v string) *AddAddressRecoverSuspendRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAddressRecoverSuspendRequest
	GetResourceOwnerId() *int64
}

type AddAddressRecoverSuspendRequest struct {
	AuditRecord *AddAddressRecoverSuspendRequestAuditRecord `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值
	RequestType          *string `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddAddressRecoverSuspendRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAddressRecoverSuspendRequest) GoString() string {
	return s.String()
}

func (s *AddAddressRecoverSuspendRequest) GetAuditRecord() *AddAddressRecoverSuspendRequestAuditRecord {
	return s.AuditRecord
}

func (s *AddAddressRecoverSuspendRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddAddressRecoverSuspendRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAddressRecoverSuspendRequest) GetRequestType() *string {
	return s.RequestType
}

func (s *AddAddressRecoverSuspendRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAddressRecoverSuspendRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAddressRecoverSuspendRequest) SetAuditRecord(v *AddAddressRecoverSuspendRequestAuditRecord) *AddAddressRecoverSuspendRequest {
	s.AuditRecord = v
	return s
}

func (s *AddAddressRecoverSuspendRequest) SetCustSpaceId(v string) *AddAddressRecoverSuspendRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddAddressRecoverSuspendRequest) SetOwnerId(v int64) *AddAddressRecoverSuspendRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAddressRecoverSuspendRequest) SetRequestType(v string) *AddAddressRecoverSuspendRequest {
	s.RequestType = &v
	return s
}

func (s *AddAddressRecoverSuspendRequest) SetResourceOwnerAccount(v string) *AddAddressRecoverSuspendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAddressRecoverSuspendRequest) SetResourceOwnerId(v int64) *AddAddressRecoverSuspendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAddressRecoverSuspendRequest) Validate() error {
	if s.AuditRecord != nil {
		if err := s.AuditRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAddressRecoverSuspendRequestAuditRecord struct {
	// example:
	//
	// 示例值示例值示例值
	ApplyReason                            *string   `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	MessageDestinationCountry              []*string `json:"MessageDestinationCountry,omitempty" xml:"MessageDestinationCountry,omitempty" type:"Repeated"`
	MessageDestinationInternationalCountry []*string `json:"MessageDestinationInternationalCountry,omitempty" xml:"MessageDestinationInternationalCountry,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	RecoveryDate *string `json:"RecoveryDate,omitempty" xml:"RecoveryDate,omitempty"`
	// example:
	//
	// 示例值示例值
	SuspensionDate *string `json:"SuspensionDate,omitempty" xml:"SuspensionDate,omitempty"`
}

func (s AddAddressRecoverSuspendRequestAuditRecord) String() string {
	return dara.Prettify(s)
}

func (s AddAddressRecoverSuspendRequestAuditRecord) GoString() string {
	return s.String()
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) GetMessageDestinationCountry() []*string {
	return s.MessageDestinationCountry
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) GetMessageDestinationInternationalCountry() []*string {
	return s.MessageDestinationInternationalCountry
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) GetRecoveryDate() *string {
	return s.RecoveryDate
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) GetSuspensionDate() *string {
	return s.SuspensionDate
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) SetApplyReason(v string) *AddAddressRecoverSuspendRequestAuditRecord {
	s.ApplyReason = &v
	return s
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) SetMessageDestinationCountry(v []*string) *AddAddressRecoverSuspendRequestAuditRecord {
	s.MessageDestinationCountry = v
	return s
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) SetMessageDestinationInternationalCountry(v []*string) *AddAddressRecoverSuspendRequestAuditRecord {
	s.MessageDestinationInternationalCountry = v
	return s
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) SetRecoveryDate(v string) *AddAddressRecoverSuspendRequestAuditRecord {
	s.RecoveryDate = &v
	return s
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) SetSuspensionDate(v string) *AddAddressRecoverSuspendRequestAuditRecord {
	s.SuspensionDate = &v
	return s
}

func (s *AddAddressRecoverSuspendRequestAuditRecord) Validate() error {
	return dara.Validate(s)
}
