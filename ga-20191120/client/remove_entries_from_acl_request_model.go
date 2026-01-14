// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntriesFromAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v []*RemoveEntriesFromAclRequestAclEntries) *RemoveEntriesFromAclRequest
	GetAclEntries() []*RemoveEntriesFromAclRequestAclEntries
	SetAclId(v string) *RemoveEntriesFromAclRequest
	GetAclId() *string
	SetClientToken(v string) *RemoveEntriesFromAclRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveEntriesFromAclRequest
	GetDryRun() *bool
	SetRegionId(v string) *RemoveEntriesFromAclRequest
	GetRegionId() *string
}

type RemoveEntriesFromAclRequest struct {
	// The entries that you want to delete. You can delete up to 50 entries at a time.
	//
	// This parameter is required.
	AclEntries []*RemoveEntriesFromAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
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
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
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

func (s RemoveEntriesFromAclRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntriesFromAclRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclRequest) GetAclEntries() []*RemoveEntriesFromAclRequestAclEntries {
	return s.AclEntries
}

func (s *RemoveEntriesFromAclRequest) GetAclId() *string {
	return s.AclId
}

func (s *RemoveEntriesFromAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveEntriesFromAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveEntriesFromAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveEntriesFromAclRequest) SetAclEntries(v []*RemoveEntriesFromAclRequestAclEntries) *RemoveEntriesFromAclRequest {
	s.AclEntries = v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetAclId(v string) *RemoveEntriesFromAclRequest {
	s.AclId = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetClientToken(v string) *RemoveEntriesFromAclRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetDryRun(v bool) *RemoveEntriesFromAclRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetRegionId(v string) *RemoveEntriesFromAclRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) Validate() error {
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

type RemoveEntriesFromAclRequestAclEntries struct {
	// The entry (IP address or CIDR block) that you want to delete. You can delete up to 50 entries at a time.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s RemoveEntriesFromAclRequestAclEntries) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntriesFromAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclRequestAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *RemoveEntriesFromAclRequestAclEntries) SetEntry(v string) *RemoveEntriesFromAclRequestAclEntries {
	s.Entry = &v
	return s
}

func (s *RemoveEntriesFromAclRequestAclEntries) Validate() error {
	return dara.Validate(s)
}
