// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntries(v []*GetAclResponseBodyAclEntries) *GetAclResponseBody
	GetAclEntries() []*GetAclResponseBodyAclEntries
	SetAclId(v string) *GetAclResponseBody
	GetAclId() *string
	SetAclName(v string) *GetAclResponseBody
	GetAclName() *string
	SetAclStatus(v string) *GetAclResponseBody
	GetAclStatus() *string
	SetAddressIPVersion(v string) *GetAclResponseBody
	GetAddressIPVersion() *string
	SetRelatedListeners(v []*GetAclResponseBodyRelatedListeners) *GetAclResponseBody
	GetRelatedListeners() []*GetAclResponseBodyRelatedListeners
	SetRequestId(v string) *GetAclResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetAclResponseBody
	GetResourceGroupId() *string
	SetTags(v []*GetAclResponseBodyTags) *GetAclResponseBody
	GetTags() []*GetAclResponseBodyTags
}

type GetAclResponseBody struct {
	// The entries of the ACL.
	AclEntries []*GetAclResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the GA instance.
	//
	// example:
	//
	// test-acl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP version of the network ACL. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// example:
	//
	// active
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The ID of the network ACL.
	//
	// example:
	//
	// IPv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The listeners that are associated with the ACL.
	RelatedListeners []*GetAclResponseBodyRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
	// The ID of the network ACL.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the network ACL.
	//
	// example:
	//
	// rg-acfmx7itmygzsza
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag of the ACL.
	Tags []*GetAclResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetAclResponseBody) GetAclEntries() []*GetAclResponseBodyAclEntries {
	return s.AclEntries
}

func (s *GetAclResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *GetAclResponseBody) GetAclName() *string {
	return s.AclName
}

func (s *GetAclResponseBody) GetAclStatus() *string {
	return s.AclStatus
}

func (s *GetAclResponseBody) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *GetAclResponseBody) GetRelatedListeners() []*GetAclResponseBodyRelatedListeners {
	return s.RelatedListeners
}

func (s *GetAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAclResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetAclResponseBody) GetTags() []*GetAclResponseBodyTags {
	return s.Tags
}

func (s *GetAclResponseBody) SetAclEntries(v []*GetAclResponseBodyAclEntries) *GetAclResponseBody {
	s.AclEntries = v
	return s
}

func (s *GetAclResponseBody) SetAclId(v string) *GetAclResponseBody {
	s.AclId = &v
	return s
}

func (s *GetAclResponseBody) SetAclName(v string) *GetAclResponseBody {
	s.AclName = &v
	return s
}

func (s *GetAclResponseBody) SetAclStatus(v string) *GetAclResponseBody {
	s.AclStatus = &v
	return s
}

func (s *GetAclResponseBody) SetAddressIPVersion(v string) *GetAclResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *GetAclResponseBody) SetRelatedListeners(v []*GetAclResponseBodyRelatedListeners) *GetAclResponseBody {
	s.RelatedListeners = v
	return s
}

func (s *GetAclResponseBody) SetRequestId(v string) *GetAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAclResponseBody) SetResourceGroupId(v string) *GetAclResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetAclResponseBody) SetTags(v []*GetAclResponseBodyTags) *GetAclResponseBody {
	s.Tags = v
	return s
}

func (s *GetAclResponseBody) Validate() error {
	if s.AclEntries != nil {
		for _, item := range s.AclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedListeners != nil {
		for _, item := range s.RelatedListeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAclResponseBodyAclEntries struct {
	// An IP address entry (192.168.XX.XX) or a CIDR block entry (10.0.XX.XX/24).
	//
	// example:
	//
	// 10.0.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The description of the ACL entry.
	//
	// example:
	//
	// test-entry
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
}

func (s GetAclResponseBodyAclEntries) String() string {
	return dara.Prettify(s)
}

func (s GetAclResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *GetAclResponseBodyAclEntries) GetEntry() *string {
	return s.Entry
}

func (s *GetAclResponseBodyAclEntries) GetEntryDescription() *string {
	return s.EntryDescription
}

func (s *GetAclResponseBodyAclEntries) SetEntry(v string) *GetAclResponseBodyAclEntries {
	s.Entry = &v
	return s
}

func (s *GetAclResponseBodyAclEntries) SetEntryDescription(v string) *GetAclResponseBodyAclEntries {
	s.EntryDescription = &v
	return s
}

func (s *GetAclResponseBodyAclEntries) Validate() error {
	return dara.Validate(s)
}

type GetAclResponseBodyRelatedListeners struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The type of the ACL. Valid values:
	//
	// 	- **white**: Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. Whitelists are suitable for scenarios in which you want to allow access from specific IP addresses to an application. If a whitelist is improperly configured, risks may arise. After a whitelist is configured for a listener, only requests from the IP addresses that are added to the whitelist are distributed by the listener. If a whitelist is enabled but no IP address is added to the whitelist, the listener forwards all requests.
	//
	// 	- **black**: All requests from the IP addresses or CIDR blocks in the ACL are rejected. Blacklists are suitable for scenarios in which you want to deny access from specific IP addresses to an application. If a blacklist is enabled but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// example:
	//
	// White
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s GetAclResponseBodyRelatedListeners) String() string {
	return dara.Prettify(s)
}

func (s GetAclResponseBodyRelatedListeners) GoString() string {
	return s.String()
}

func (s *GetAclResponseBodyRelatedListeners) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetAclResponseBodyRelatedListeners) GetAclType() *string {
	return s.AclType
}

func (s *GetAclResponseBodyRelatedListeners) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetAclResponseBodyRelatedListeners) SetAcceleratorId(v string) *GetAclResponseBodyRelatedListeners {
	s.AcceleratorId = &v
	return s
}

func (s *GetAclResponseBodyRelatedListeners) SetAclType(v string) *GetAclResponseBodyRelatedListeners {
	s.AclType = &v
	return s
}

func (s *GetAclResponseBodyRelatedListeners) SetListenerId(v string) *GetAclResponseBodyRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *GetAclResponseBodyRelatedListeners) Validate() error {
	return dara.Validate(s)
}

type GetAclResponseBodyTags struct {
	// The key of tag N that is add to the ACL.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is add to the ACL.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAclResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetAclResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetAclResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetAclResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetAclResponseBodyTags) SetKey(v string) *GetAclResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetAclResponseBodyTags) SetValue(v string) *GetAclResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetAclResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
