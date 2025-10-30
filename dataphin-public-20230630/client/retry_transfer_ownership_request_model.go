// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryTransferOwnershipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RetryTransferOwnershipRequest
	GetOpTenantId() *int64
	SetPrivilegeTransferRecord(v *RetryTransferOwnershipRequestPrivilegeTransferRecord) *RetryTransferOwnershipRequest
	GetPrivilegeTransferRecord() *RetryTransferOwnershipRequestPrivilegeTransferRecord
}

type RetryTransferOwnershipRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId              *int64                                                `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	PrivilegeTransferRecord *RetryTransferOwnershipRequestPrivilegeTransferRecord `json:"PrivilegeTransferRecord,omitempty" xml:"PrivilegeTransferRecord,omitempty" type:"Struct"`
}

func (s RetryTransferOwnershipRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryTransferOwnershipRequest) GoString() string {
	return s.String()
}

func (s *RetryTransferOwnershipRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RetryTransferOwnershipRequest) GetPrivilegeTransferRecord() *RetryTransferOwnershipRequestPrivilegeTransferRecord {
	return s.PrivilegeTransferRecord
}

func (s *RetryTransferOwnershipRequest) SetOpTenantId(v int64) *RetryTransferOwnershipRequest {
	s.OpTenantId = &v
	return s
}

func (s *RetryTransferOwnershipRequest) SetPrivilegeTransferRecord(v *RetryTransferOwnershipRequestPrivilegeTransferRecord) *RetryTransferOwnershipRequest {
	s.PrivilegeTransferRecord = v
	return s
}

func (s *RetryTransferOwnershipRequest) Validate() error {
	if s.PrivilegeTransferRecord != nil {
		if err := s.PrivilegeTransferRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryTransferOwnershipRequestPrivilegeTransferRecord struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000001
	NewOwner *string `json:"NewOwner,omitempty" xml:"NewOwner,omitempty"`
	// example:
	//
	// comment
	TransferComment *string `json:"TransferComment,omitempty" xml:"TransferComment,omitempty"`
}

func (s RetryTransferOwnershipRequestPrivilegeTransferRecord) String() string {
	return dara.Prettify(s)
}

func (s RetryTransferOwnershipRequestPrivilegeTransferRecord) GoString() string {
	return s.String()
}

func (s *RetryTransferOwnershipRequestPrivilegeTransferRecord) GetId() *int64 {
	return s.Id
}

func (s *RetryTransferOwnershipRequestPrivilegeTransferRecord) GetNewOwner() *string {
	return s.NewOwner
}

func (s *RetryTransferOwnershipRequestPrivilegeTransferRecord) GetTransferComment() *string {
	return s.TransferComment
}

func (s *RetryTransferOwnershipRequestPrivilegeTransferRecord) SetId(v int64) *RetryTransferOwnershipRequestPrivilegeTransferRecord {
	s.Id = &v
	return s
}

func (s *RetryTransferOwnershipRequestPrivilegeTransferRecord) SetNewOwner(v string) *RetryTransferOwnershipRequestPrivilegeTransferRecord {
	s.NewOwner = &v
	return s
}

func (s *RetryTransferOwnershipRequestPrivilegeTransferRecord) SetTransferComment(v string) *RetryTransferOwnershipRequestPrivilegeTransferRecord {
	s.TransferComment = &v
	return s
}

func (s *RetryTransferOwnershipRequestPrivilegeTransferRecord) Validate() error {
	return dara.Validate(s)
}
