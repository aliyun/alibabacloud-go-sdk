// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryTransferOwnershipShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RetryTransferOwnershipShrinkRequest
	GetOpTenantId() *int64
	SetPrivilegeTransferRecordShrink(v string) *RetryTransferOwnershipShrinkRequest
	GetPrivilegeTransferRecordShrink() *string
}

type RetryTransferOwnershipShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId                    *int64  `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	PrivilegeTransferRecordShrink *string `json:"PrivilegeTransferRecord,omitempty" xml:"PrivilegeTransferRecord,omitempty"`
}

func (s RetryTransferOwnershipShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryTransferOwnershipShrinkRequest) GoString() string {
	return s.String()
}

func (s *RetryTransferOwnershipShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RetryTransferOwnershipShrinkRequest) GetPrivilegeTransferRecordShrink() *string {
	return s.PrivilegeTransferRecordShrink
}

func (s *RetryTransferOwnershipShrinkRequest) SetOpTenantId(v int64) *RetryTransferOwnershipShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RetryTransferOwnershipShrinkRequest) SetPrivilegeTransferRecordShrink(v string) *RetryTransferOwnershipShrinkRequest {
	s.PrivilegeTransferRecordShrink = &v
	return s
}

func (s *RetryTransferOwnershipShrinkRequest) Validate() error {
	return dara.Validate(s)
}
