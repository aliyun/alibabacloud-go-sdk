// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferOwnershipForAllObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *TransferOwnershipForAllObjectRequest
	GetOpTenantId() *int64
	SetPrivilegeTransferRecord(v *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) *TransferOwnershipForAllObjectRequest
	GetPrivilegeTransferRecord() *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord
}

type TransferOwnershipForAllObjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId              *int64                                                       `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	PrivilegeTransferRecord *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord `json:"PrivilegeTransferRecord,omitempty" xml:"PrivilegeTransferRecord,omitempty" type:"Struct"`
}

func (s TransferOwnershipForAllObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferOwnershipForAllObjectRequest) GoString() string {
	return s.String()
}

func (s *TransferOwnershipForAllObjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *TransferOwnershipForAllObjectRequest) GetPrivilegeTransferRecord() *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord {
	return s.PrivilegeTransferRecord
}

func (s *TransferOwnershipForAllObjectRequest) SetOpTenantId(v int64) *TransferOwnershipForAllObjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *TransferOwnershipForAllObjectRequest) SetPrivilegeTransferRecord(v *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) *TransferOwnershipForAllObjectRequest {
	s.PrivilegeTransferRecord = v
	return s
}

func (s *TransferOwnershipForAllObjectRequest) Validate() error {
	return dara.Validate(s)
}

type TransferOwnershipForAllObjectRequestPrivilegeTransferRecord struct {
	// This parameter is required.
	//
	// example:
	//
	// 30000002
	NewOwner *string `json:"NewOwner,omitempty" xml:"NewOwner,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000001
	OldOwner *string `json:"OldOwner,omitempty" xml:"OldOwner,omitempty"`
	// example:
	//
	// comment
	TransferComment *string `json:"TransferComment,omitempty" xml:"TransferComment,omitempty"`
}

func (s TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) String() string {
	return dara.Prettify(s)
}

func (s TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) GoString() string {
	return s.String()
}

func (s *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) GetNewOwner() *string {
	return s.NewOwner
}

func (s *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) GetOldOwner() *string {
	return s.OldOwner
}

func (s *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) GetTransferComment() *string {
	return s.TransferComment
}

func (s *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) SetNewOwner(v string) *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord {
	s.NewOwner = &v
	return s
}

func (s *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) SetOldOwner(v string) *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord {
	s.OldOwner = &v
	return s
}

func (s *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) SetTransferComment(v string) *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord {
	s.TransferComment = &v
	return s
}

func (s *TransferOwnershipForAllObjectRequestPrivilegeTransferRecord) Validate() error {
	return dara.Validate(s)
}
