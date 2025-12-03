// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntrys(v *DescribeAccessControlListAttributeResponseBodyAclEntrys) *DescribeAccessControlListAttributeResponseBody
	GetAclEntrys() *DescribeAccessControlListAttributeResponseBodyAclEntrys
	SetAclId(v string) *DescribeAccessControlListAttributeResponseBody
	GetAclId() *string
	SetAclName(v string) *DescribeAccessControlListAttributeResponseBody
	GetAclName() *string
	SetAddressIPVersion(v string) *DescribeAccessControlListAttributeResponseBody
	GetAddressIPVersion() *string
	SetCreateTime(v string) *DescribeAccessControlListAttributeResponseBody
	GetCreateTime() *string
	SetRelatedListeners(v *DescribeAccessControlListAttributeResponseBodyRelatedListeners) *DescribeAccessControlListAttributeResponseBody
	GetRelatedListeners() *DescribeAccessControlListAttributeResponseBodyRelatedListeners
	SetRequestId(v string) *DescribeAccessControlListAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeAccessControlListAttributeResponseBody
	GetResourceGroupId() *string
	SetTags(v *DescribeAccessControlListAttributeResponseBodyTags) *DescribeAccessControlListAttributeResponseBody
	GetTags() *DescribeAccessControlListAttributeResponseBodyTags
	SetTotalAclEntry(v int32) *DescribeAccessControlListAttributeResponseBody
	GetTotalAclEntry() *int32
}

type DescribeAccessControlListAttributeResponseBody struct {
	// The information about the access control policy.
	AclEntrys *DescribeAccessControlListAttributeResponseBodyAclEntrys `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty" type:"Struct"`
	// The ACL ID.
	//
	// example:
	//
	// acl-bp1ut10zzvh1y8dfs****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ACL name.
	//
	// example:
	//
	// doctest
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP version. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The time when the ACL was created. The time follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-08-31T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The listeners with which the ACL is associated.
	RelatedListeners *DescribeAccessControlListAttributeResponseBodyRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C9906A1D-86F7-4C9C-A369-54DA42EF206A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmz3jksig****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the ACL.
	Tags *DescribeAccessControlListAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The total number of ACL entries.
	//
	// example:
	//
	// 200
	TotalAclEntry *int32 `json:"TotalAclEntry,omitempty" xml:"TotalAclEntry,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAclEntrys() *DescribeAccessControlListAttributeResponseBodyAclEntrys {
	return s.AclEntrys
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAclName() *string {
	return s.AclName
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAccessControlListAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAccessControlListAttributeResponseBody) GetRelatedListeners() *DescribeAccessControlListAttributeResponseBodyRelatedListeners {
	return s.RelatedListeners
}

func (s *DescribeAccessControlListAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessControlListAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAccessControlListAttributeResponseBody) GetTags() *DescribeAccessControlListAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeAccessControlListAttributeResponseBody) GetTotalAclEntry() *int32 {
	return s.TotalAclEntry
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclEntrys(v *DescribeAccessControlListAttributeResponseBodyAclEntrys) *DescribeAccessControlListAttributeResponseBody {
	s.AclEntrys = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclName(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAddressIPVersion(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetCreateTime(v string) *DescribeAccessControlListAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetRelatedListeners(v *DescribeAccessControlListAttributeResponseBodyRelatedListeners) *DescribeAccessControlListAttributeResponseBody {
	s.RelatedListeners = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetRequestId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetResourceGroupId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetTags(v *DescribeAccessControlListAttributeResponseBodyTags) *DescribeAccessControlListAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetTotalAclEntry(v int32) *DescribeAccessControlListAttributeResponseBody {
	s.TotalAclEntry = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) Validate() error {
	if s.AclEntrys != nil {
		if err := s.AclEntrys.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedListeners != nil {
		if err := s.RelatedListeners.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessControlListAttributeResponseBodyAclEntrys struct {
	AclEntry []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry `json:"AclEntry,omitempty" xml:"AclEntry,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrys) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrys) GetAclEntry() []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	return s.AclEntry
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrys) SetAclEntry(v []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) *DescribeAccessControlListAttributeResponseBodyAclEntrys {
	s.AclEntry = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrys) Validate() error {
	if s.AclEntry != nil {
		for _, item := range s.AclEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry struct {
	// The remarks of the ACL entry.
	//
	// example:
	//
	// test
	AclEntryComment *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	// The IP entry in the ACL.
	//
	// example:
	//
	// 192.168.0.1
	AclEntryIP *string `json:"AclEntryIP,omitempty" xml:"AclEntryIP,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) GetAclEntryComment() *string {
	return s.AclEntryComment
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) GetAclEntryIP() *string {
	return s.AclEntryIP
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) SetAclEntryComment(v string) *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	s.AclEntryComment = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) SetAclEntryIP(v string) *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	s.AclEntryIP = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessControlListAttributeResponseBodyRelatedListeners struct {
	RelatedListener []*DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener `json:"RelatedListener,omitempty" xml:"RelatedListener,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListeners) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListeners) GetRelatedListener() []*DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	return s.RelatedListener
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListeners) SetRelatedListener(v []*DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) *DescribeAccessControlListAttributeResponseBodyRelatedListeners {
	s.RelatedListener = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListeners) Validate() error {
	if s.RelatedListener != nil {
		for _, item := range s.RelatedListener {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener struct {
	// The type of ACL. Valid values:
	//
	// 	- **black**
	//
	// 	- **white**
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The frontend port of the listener with which the ACL is associated.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The CLB instance ID.
	//
	// example:
	//
	// lb-bp1qpzldlm38bnexl****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The type of protocol that the associated listener uses.
	//
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) GetAclType() *string {
	return s.AclType
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetAclType(v string) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.AclType = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetListenerPort(v int32) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.ListenerPort = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetLoadBalancerId(v string) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetProtocol(v string) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.Protocol = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessControlListAttributeResponseBodyTags struct {
	Tag []*DescribeAccessControlListAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyTags) GetTag() []*DescribeAccessControlListAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeAccessControlListAttributeResponseBodyTags) SetTag(v []*DescribeAccessControlListAttributeResponseBodyTagsTag) *DescribeAccessControlListAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyTags) Validate() error {
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

type DescribeAccessControlListAttributeResponseBodyTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeAccessControlListAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeAccessControlListAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeAccessControlListAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeAccessControlListAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
