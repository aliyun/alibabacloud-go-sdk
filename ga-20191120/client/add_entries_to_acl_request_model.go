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
	// The entries (IP addresses or CIDR blocks) that you want to add to the ACL.
	//
	// You can add at most 50 entries at a time.
	//
	// This parameter is required.
	AclEntries []*AddEntriesToAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The ACL ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
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
	// The entry (IP address or CIDR block) that you want to add. You can add at most 50 entries at a time.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The description of the entry that you want to add to the ACL.
	//
	// You can add at most 50 descriptions in each request.
	//
	// The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
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
