// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEntriesToAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v []*AddEntriesToAclRequestAclEntries) *AddEntriesToAclRequest
	GetAclEntries() []*AddEntriesToAclRequestAclEntries
	SetAclId(v string) *AddEntriesToAclRequest
	GetAclId() *string
	SetClientToken(v string) *AddEntriesToAclRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddEntriesToAclRequest
	GetDryRun() *bool
	SetRegionId(v string) *AddEntriesToAclRequest
	GetRegionId() *string
}

type AddEntriesToAclRequest struct {
	// The access control policy group entries to add. An entry can be an IP address or a CIDR block.
	//
	// You can add up to 50 entries at a time.
	//
	// This parameter is required.
	AclEntries []*AddEntriesToAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The ID of the access control policy group.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The client token that is used to ensure the idempotence of a request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value is different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request format, and business limitations without actually adding IP entries to the access control policy group. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddEntriesToAclRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEntriesToAclRequest) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequest) GetAclEntries() []*AddEntriesToAclRequestAclEntries {
	return s.AclEntries
}

func (s *AddEntriesToAclRequest) GetAclId() *string {
	return s.AclId
}

func (s *AddEntriesToAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddEntriesToAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddEntriesToAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddEntriesToAclRequest) SetAclEntries(v []*AddEntriesToAclRequestAclEntries) *AddEntriesToAclRequest {
	s.AclEntries = v
	return s
}

func (s *AddEntriesToAclRequest) SetAclId(v string) *AddEntriesToAclRequest {
	s.AclId = &v
	return s
}

func (s *AddEntriesToAclRequest) SetClientToken(v string) *AddEntriesToAclRequest {
	s.ClientToken = &v
	return s
}

func (s *AddEntriesToAclRequest) SetDryRun(v bool) *AddEntriesToAclRequest {
	s.DryRun = &v
	return s
}

func (s *AddEntriesToAclRequest) SetRegionId(v string) *AddEntriesToAclRequest {
	s.RegionId = &v
	return s
}

func (s *AddEntriesToAclRequest) Validate() error {
	if s.AclEntries != nil {
		for _, item := range s.AclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddEntriesToAclRequestAclEntries struct {
	// The access control policy group entry to add. An entry can be an IP address or a CIDR block. You can add up to 50 entries at a time.
	//
	// > This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The description of the access control policy group entry.
	//
	// You can add descriptions for up to 50 entries at a time.
	//
	// The description must be 1 to 256 characters in length and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_). Chinese characters are supported.
	//
	// example:
	//
	// test-entry
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
}

func (s AddEntriesToAclRequestAclEntries) String() string {
	return dara.Prettify(s)
}

func (s AddEntriesToAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequestAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *AddEntriesToAclRequestAclEntries) GetEntryDescription() *string {
	return s.EntryDescription
}

func (s *AddEntriesToAclRequestAclEntries) SetEntry(v string) *AddEntriesToAclRequestAclEntries {
	s.Entry = &v
	return s
}

func (s *AddEntriesToAclRequestAclEntries) SetEntryDescription(v string) *AddEntriesToAclRequestAclEntries {
	s.EntryDescription = &v
	return s
}

func (s *AddEntriesToAclRequestAclEntries) Validate() error {
	return dara.Validate(s)
}
