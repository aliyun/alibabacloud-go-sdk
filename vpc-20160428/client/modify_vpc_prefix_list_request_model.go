// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcPrefixListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddPrefixListEntry(v []*ModifyVpcPrefixListRequestAddPrefixListEntry) *ModifyVpcPrefixListRequest
	GetAddPrefixListEntry() []*ModifyVpcPrefixListRequestAddPrefixListEntry
	SetClientToken(v string) *ModifyVpcPrefixListRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyVpcPrefixListRequest
	GetDryRun() *bool
	SetMaxEntries(v int32) *ModifyVpcPrefixListRequest
	GetMaxEntries() *int32
	SetOwnerAccount(v string) *ModifyVpcPrefixListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpcPrefixListRequest
	GetOwnerId() *int64
	SetPrefixListDescription(v string) *ModifyVpcPrefixListRequest
	GetPrefixListDescription() *string
	SetPrefixListId(v string) *ModifyVpcPrefixListRequest
	GetPrefixListId() *string
	SetPrefixListName(v string) *ModifyVpcPrefixListRequest
	GetPrefixListName() *string
	SetRegionId(v string) *ModifyVpcPrefixListRequest
	GetRegionId() *string
	SetRemovePrefixListEntry(v []*ModifyVpcPrefixListRequestRemovePrefixListEntry) *ModifyVpcPrefixListRequest
	GetRemovePrefixListEntry() []*ModifyVpcPrefixListRequestRemovePrefixListEntry
	SetResourceOwnerAccount(v string) *ModifyVpcPrefixListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpcPrefixListRequest
	GetResourceOwnerId() *int64
}

type ModifyVpcPrefixListRequest struct {
	// The information about CIDR blocks to be added to the prefix list.
	AddPrefixListEntry []*ModifyVpcPrefixListRequestAddPrefixListEntry `json:"AddPrefixListEntry,omitempty" xml:"AddPrefixListEntry,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The maximum number of CIDR blocks supported by the prefix list after the configuration of the prefix list is modified.
	//
	// example:
	//
	// 20
	MaxEntries   *int32  `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new description of the prefix list.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newdescription
	PrefixListDescription *string `json:"PrefixListDescription,omitempty" xml:"PrefixListDescription,omitempty"`
	// The ID of the prefix list.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-0b7hwu67****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The new name of the prefix list.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newname
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The region ID of the prefix list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about CIDR blocks to be deleted to the prefix list.
	RemovePrefixListEntry []*ModifyVpcPrefixListRequestRemovePrefixListEntry `json:"RemovePrefixListEntry,omitempty" xml:"RemovePrefixListEntry,omitempty" type:"Repeated"`
	ResourceOwnerAccount  *string                                            `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64                                             `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyVpcPrefixListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPrefixListRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcPrefixListRequest) GetAddPrefixListEntry() []*ModifyVpcPrefixListRequestAddPrefixListEntry {
	return s.AddPrefixListEntry
}

func (s *ModifyVpcPrefixListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpcPrefixListRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyVpcPrefixListRequest) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *ModifyVpcPrefixListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpcPrefixListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpcPrefixListRequest) GetPrefixListDescription() *string {
	return s.PrefixListDescription
}

func (s *ModifyVpcPrefixListRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *ModifyVpcPrefixListRequest) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *ModifyVpcPrefixListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpcPrefixListRequest) GetRemovePrefixListEntry() []*ModifyVpcPrefixListRequestRemovePrefixListEntry {
	return s.RemovePrefixListEntry
}

func (s *ModifyVpcPrefixListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpcPrefixListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpcPrefixListRequest) SetAddPrefixListEntry(v []*ModifyVpcPrefixListRequestAddPrefixListEntry) *ModifyVpcPrefixListRequest {
	s.AddPrefixListEntry = v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetClientToken(v string) *ModifyVpcPrefixListRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetDryRun(v bool) *ModifyVpcPrefixListRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetMaxEntries(v int32) *ModifyVpcPrefixListRequest {
	s.MaxEntries = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetOwnerAccount(v string) *ModifyVpcPrefixListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetOwnerId(v int64) *ModifyVpcPrefixListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetPrefixListDescription(v string) *ModifyVpcPrefixListRequest {
	s.PrefixListDescription = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetPrefixListId(v string) *ModifyVpcPrefixListRequest {
	s.PrefixListId = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetPrefixListName(v string) *ModifyVpcPrefixListRequest {
	s.PrefixListName = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetRegionId(v string) *ModifyVpcPrefixListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetRemovePrefixListEntry(v []*ModifyVpcPrefixListRequestRemovePrefixListEntry) *ModifyVpcPrefixListRequest {
	s.RemovePrefixListEntry = v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetResourceOwnerAccount(v string) *ModifyVpcPrefixListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) SetResourceOwnerId(v int64) *ModifyVpcPrefixListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpcPrefixListRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyVpcPrefixListRequestAddPrefixListEntry struct {
	// The CIDR block to be added to the prefix list.
	//
	// >  If the CIDR block already exists in the prefix list, you can only modify the description of the CIDR block by setting the **AddPrefixListEntry.N.Description*	- parameter.
	//
	// example:
	//
	// 172.16.0.0/12
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the CIDR block to be added to the prefix list.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newcidr
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyVpcPrefixListRequestAddPrefixListEntry) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPrefixListRequestAddPrefixListEntry) GoString() string {
	return s.String()
}

func (s *ModifyVpcPrefixListRequestAddPrefixListEntry) GetCidr() *string {
	return s.Cidr
}

func (s *ModifyVpcPrefixListRequestAddPrefixListEntry) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpcPrefixListRequestAddPrefixListEntry) SetCidr(v string) *ModifyVpcPrefixListRequestAddPrefixListEntry {
	s.Cidr = &v
	return s
}

func (s *ModifyVpcPrefixListRequestAddPrefixListEntry) SetDescription(v string) *ModifyVpcPrefixListRequestAddPrefixListEntry {
	s.Description = &v
	return s
}

func (s *ModifyVpcPrefixListRequestAddPrefixListEntry) Validate() error {
	return dara.Validate(s)
}

type ModifyVpcPrefixListRequestRemovePrefixListEntry struct {
	// The CIDR block that you want to delete from the prefix list.
	//
	// example:
	//
	// 192.168.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the CIDR block that you want to delete.
	//
	// example:
	//
	// cidr
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyVpcPrefixListRequestRemovePrefixListEntry) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPrefixListRequestRemovePrefixListEntry) GoString() string {
	return s.String()
}

func (s *ModifyVpcPrefixListRequestRemovePrefixListEntry) GetCidr() *string {
	return s.Cidr
}

func (s *ModifyVpcPrefixListRequestRemovePrefixListEntry) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpcPrefixListRequestRemovePrefixListEntry) SetCidr(v string) *ModifyVpcPrefixListRequestRemovePrefixListEntry {
	s.Cidr = &v
	return s
}

func (s *ModifyVpcPrefixListRequestRemovePrefixListEntry) SetDescription(v string) *ModifyVpcPrefixListRequestRemovePrefixListEntry {
	s.Description = &v
	return s
}

func (s *ModifyVpcPrefixListRequestRemovePrefixListEntry) Validate() error {
	return dara.Validate(s)
}
