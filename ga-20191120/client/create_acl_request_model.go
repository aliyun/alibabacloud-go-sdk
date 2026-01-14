// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v []*CreateAclRequestAclEntries) *CreateAclRequest
	GetAclEntries() []*CreateAclRequestAclEntries
	SetAclName(v string) *CreateAclRequest
	GetAclName() *string
	SetAddressIPVersion(v string) *CreateAclRequest
	GetAddressIPVersion() *string
	SetClientToken(v string) *CreateAclRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateAclRequest
	GetDryRun() *bool
	SetRegionId(v string) *CreateAclRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateAclRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateAclRequestTag) *CreateAclRequest
	GetTag() []*CreateAclRequestTag
}

type CreateAclRequest struct {
	// The entries of IP addresses or CIDR blocks to add to the ACL.
	//
	// You can add a maximum of 50 entries at a time.
	AclEntries []*CreateAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The ACL name.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test-acl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP version of the ACL. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// This parameter is required.
	//
	// example:
	//
	// IPv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
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
	// Specifies whether to only precheck the request. Default value: false. Valid values:
	//
	// 	- **true**: prechecks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. If the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmwj7wvng3jbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the ACL.
	Tag []*CreateAclRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateAclRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) GetAclEntries() []*CreateAclRequestAclEntries {
	return s.AclEntries
}

func (s *CreateAclRequest) GetAclName() *string {
	return s.AclName
}

func (s *CreateAclRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *CreateAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAclRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAclRequest) GetTag() []*CreateAclRequestTag {
	return s.Tag
}

func (s *CreateAclRequest) SetAclEntries(v []*CreateAclRequestAclEntries) *CreateAclRequest {
	s.AclEntries = v
	return s
}

func (s *CreateAclRequest) SetAclName(v string) *CreateAclRequest {
	s.AclName = &v
	return s
}

func (s *CreateAclRequest) SetAddressIPVersion(v string) *CreateAclRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateAclRequest) SetClientToken(v string) *CreateAclRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAclRequest) SetDryRun(v bool) *CreateAclRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAclRequest) SetRegionId(v string) *CreateAclRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAclRequest) SetResourceGroupId(v string) *CreateAclRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAclRequest) SetTag(v []*CreateAclRequestTag) *CreateAclRequest {
	s.Tag = v
	return s
}

func (s *CreateAclRequest) Validate() error {
	if s.AclEntries != nil {
		for _, item := range s.AclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAclRequestAclEntries struct {
	// The IP addresses (192.168.XX.XX) or CIDR blocks (10.0.XX.XX/24) that you want to add to the ACL.
	//
	// You can add a maximum of 50 entries at a time.
	//
	// example:
	//
	// 10.0.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The description of the entry that you want to add to the ACL.
	//
	// You can add a maximum of 50 entries at a time.
	//
	// The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// test-entry
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
}

func (s CreateAclRequestAclEntries) String() string {
	return dara.Prettify(s)
}

func (s CreateAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *CreateAclRequestAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *CreateAclRequestAclEntries) GetEntryDescription() *string {
	return s.EntryDescription
}

func (s *CreateAclRequestAclEntries) SetEntry(v string) *CreateAclRequestAclEntries {
	s.Entry = &v
	return s
}

func (s *CreateAclRequestAclEntries) SetEntryDescription(v string) *CreateAclRequestAclEntries {
	s.EntryDescription = &v
	return s
}

func (s *CreateAclRequestAclEntries) Validate() error {
	return dara.Validate(s)
}

type CreateAclRequestTag struct {
	// The tag key of the ACL. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the ACL. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAclRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAclRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAclRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAclRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAclRequestTag) SetKey(v string) *CreateAclRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAclRequestTag) SetValue(v string) *CreateAclRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAclRequestTag) Validate() error {
	return dara.Validate(s)
}
