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
	// The access control policy group entries, which are IP address entries or CIDR block entries.
	//
	// You can add up to 50 entries at a time.
	AclEntries []*CreateAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The name of the access control policy group.
	//
	// The name must be 1 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test-acl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP version of the access control policy group. Valid values:
	//
	// - **IPv4**
	//
	// - **IPv6**
	//
	// This parameter is required.
	//
	// example:
	//
	// IPv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value is different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the access control policy group. The system checks the required parameters, request format, and business limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a Normal request, passes the dry run, and returns an HTTP 2xx status code and directly performs the operation.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the Alibaba Cloud Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwj7wvng3jbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The label information of the access control policy group.
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
	// The access control policy group entry, which is an IP address entry (192.168.XX.XX) or a CIDR block entry (10.0.XX.XX/24).
	//
	// You can add up to 50 entries at a time.
	//
	// example:
	//
	// 10.0.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The description of the access control policy group entry.
	//
	// You can add descriptions for up to 50 entries at a time.
	//
	// The description must be 1 to 256 characters in length and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), underscores (_), and Chinese characters.
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
	// The label key of the access control policy group. Once specified, the label key cannot be an empty string.
	//
	// The label key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 label keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The label value of the access control policy group. Once specified, the label value can be an empty string.
	//
	// The label value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 label values.
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
