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
}

type AddEntriesToAclRequest struct {
	// The ACL entries that you want to add. You can add at most 20 entries in each call.
	//
	// This parameter is required.
	AclEntries []*AddEntriesToAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The ID of the ACL.
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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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

func (s *AddEntriesToAclRequest) Validate() error {
	return dara.Validate(s)
}

type AddEntriesToAclRequestAclEntries struct {
	// The description of the ACL entry. The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// You can add at most 20 entries in each call.
	//
	// example:
	//
	// test-entry
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The CIDR block in the ACL entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.1.0/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s AddEntriesToAclRequestAclEntries) String() string {
	return dara.Prettify(s)
}

func (s AddEntriesToAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequestAclEntries) GetDescription() *string {
	return s.Description
}

func (s *AddEntriesToAclRequestAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *AddEntriesToAclRequestAclEntries) SetDescription(v string) *AddEntriesToAclRequestAclEntries {
	s.Description = &v
	return s
}

func (s *AddEntriesToAclRequestAclEntries) SetEntry(v string) *AddEntriesToAclRequestAclEntries {
	s.Entry = &v
	return s
}

func (s *AddEntriesToAclRequestAclEntries) Validate() error {
	return dara.Validate(s)
}
