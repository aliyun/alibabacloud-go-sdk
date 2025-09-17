// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateSourceServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AssociateSourceServersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AssociateSourceServersRequest
	GetResourceOwnerAccount() *string
	SetSourceId(v []*string) *AssociateSourceServersRequest
	GetSourceId() []*string
	SetWorkgroupId(v string) *AssociateSourceServersRequest
	GetWorkgroupId() *string
}

type AssociateSourceServersRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the migration source. A workgroup can be associated with a maximum of 50 migration sources.
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

func (s AssociateSourceServersRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateSourceServersRequest) GoString() string {
	return s.String()
}

func (s *AssociateSourceServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateSourceServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateSourceServersRequest) GetSourceId() []*string {
	return s.SourceId
}

func (s *AssociateSourceServersRequest) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *AssociateSourceServersRequest) SetOwnerId(v int64) *AssociateSourceServersRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateSourceServersRequest) SetResourceOwnerAccount(v string) *AssociateSourceServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateSourceServersRequest) SetSourceId(v []*string) *AssociateSourceServersRequest {
	s.SourceId = v
	return s
}

func (s *AssociateSourceServersRequest) SetWorkgroupId(v string) *AssociateSourceServersRequest {
	s.WorkgroupId = &v
	return s
}

func (s *AssociateSourceServersRequest) Validate() error {
	return dara.Validate(s)
}
