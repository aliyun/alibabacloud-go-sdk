// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferOwnershipForAllObjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *TransferOwnershipForAllObjectShrinkRequest
	GetOpTenantId() *int64
	SetPrivilegeTransferRecordShrink(v string) *TransferOwnershipForAllObjectShrinkRequest
	GetPrivilegeTransferRecordShrink() *string
}

type TransferOwnershipForAllObjectShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId                    *int64  `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	PrivilegeTransferRecordShrink *string `json:"PrivilegeTransferRecord,omitempty" xml:"PrivilegeTransferRecord,omitempty"`
}

func (s TransferOwnershipForAllObjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferOwnershipForAllObjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *TransferOwnershipForAllObjectShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *TransferOwnershipForAllObjectShrinkRequest) GetPrivilegeTransferRecordShrink() *string {
	return s.PrivilegeTransferRecordShrink
}

func (s *TransferOwnershipForAllObjectShrinkRequest) SetOpTenantId(v int64) *TransferOwnershipForAllObjectShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *TransferOwnershipForAllObjectShrinkRequest) SetPrivilegeTransferRecordShrink(v string) *TransferOwnershipForAllObjectShrinkRequest {
	s.PrivilegeTransferRecordShrink = &v
	return s
}

func (s *TransferOwnershipForAllObjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
