// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateSourceServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DisassociateSourceServersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisassociateSourceServersRequest
	GetResourceOwnerAccount() *string
	SetSourceId(v []*string) *DisassociateSourceServersRequest
	GetSourceId() []*string
	SetWorkgroupId(v string) *DisassociateSourceServersRequest
	GetWorkgroupId() *string
}

type DisassociateSourceServersRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The IDs of migration sources that you want to disassociate from the workgroup. You can specify up to 50 migration sources.
	//
	// This parameter is required.
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	// The ID of the workgroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-***
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s DisassociateSourceServersRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateSourceServersRequest) GoString() string {
	return s.String()
}

func (s *DisassociateSourceServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisassociateSourceServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisassociateSourceServersRequest) GetSourceId() []*string {
	return s.SourceId
}

func (s *DisassociateSourceServersRequest) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *DisassociateSourceServersRequest) SetOwnerId(v int64) *DisassociateSourceServersRequest {
	s.OwnerId = &v
	return s
}

func (s *DisassociateSourceServersRequest) SetResourceOwnerAccount(v string) *DisassociateSourceServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisassociateSourceServersRequest) SetSourceId(v []*string) *DisassociateSourceServersRequest {
	s.SourceId = v
	return s
}

func (s *DisassociateSourceServersRequest) SetWorkgroupId(v string) *DisassociateSourceServersRequest {
	s.WorkgroupId = &v
	return s
}

func (s *DisassociateSourceServersRequest) Validate() error {
	return dara.Validate(s)
}
