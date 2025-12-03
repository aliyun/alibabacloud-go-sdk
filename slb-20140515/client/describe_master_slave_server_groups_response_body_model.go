// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMasterSlaveServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMasterSlaveServerGroups(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) *DescribeMasterSlaveServerGroupsResponseBody
	GetMasterSlaveServerGroups() *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups
	SetRequestId(v string) *DescribeMasterSlaveServerGroupsResponseBody
	GetRequestId() *string
}

type DescribeMasterSlaveServerGroupsResponseBody struct {
	// The primary/secondary server groups.
	MasterSlaveServerGroups *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups `json:"MasterSlaveServerGroups,omitempty" xml:"MasterSlaveServerGroups,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBody) GetMasterSlaveServerGroups() *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups {
	return s.MasterSlaveServerGroups
}

func (s *DescribeMasterSlaveServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMasterSlaveServerGroupsResponseBody) SetMasterSlaveServerGroups(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) *DescribeMasterSlaveServerGroupsResponseBody {
	s.MasterSlaveServerGroups = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBody) SetRequestId(v string) *DescribeMasterSlaveServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBody) Validate() error {
	if s.MasterSlaveServerGroups != nil {
		if err := s.MasterSlaveServerGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups struct {
	MasterSlaveServerGroup []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup `json:"MasterSlaveServerGroup,omitempty" xml:"MasterSlaveServerGroup,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) GetMasterSlaveServerGroup() []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	return s.MasterSlaveServerGroup
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) SetMasterSlaveServerGroup(v []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups {
	s.MasterSlaveServerGroup = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) Validate() error {
	if s.MasterSlaveServerGroup != nil {
		for _, item := range s.MasterSlaveServerGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup struct {
	// The associated resources.
	AssociatedObjects *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects `json:"AssociatedObjects,omitempty" xml:"AssociatedObjects,omitempty" type:"Struct"`
	// The time when the CLB instance was created. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-12-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the primary/secondary server group.
	//
	// example:
	//
	// rsp-0bfuc******
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	// The name of the primary/secondary server group.
	//
	// example:
	//
	// Group3
	MasterSlaveServerGroupName *string `json:"MasterSlaveServerGroupName,omitempty" xml:"MasterSlaveServerGroupName,omitempty"`
	// The tags to add to the resource.
	Tags *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) GetAssociatedObjects() *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects {
	return s.AssociatedObjects
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) GetMasterSlaveServerGroupName() *string {
	return s.MasterSlaveServerGroupName
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) GetTags() *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags {
	return s.Tags
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetAssociatedObjects(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.AssociatedObjects = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetCreateTime(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetMasterSlaveServerGroupName(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.MasterSlaveServerGroupName = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetTags(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.Tags = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) Validate() error {
	if s.AssociatedObjects != nil {
		if err := s.AssociatedObjects.Validate(); err != nil {
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

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects struct {
	// The listeners.
	Listeners *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Struct"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) GetListeners() *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners {
	return s.Listeners
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) SetListeners(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects {
	s.Listeners = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) Validate() error {
	if s.Listeners != nil {
		if err := s.Listeners.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners struct {
	Listener []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener `json:"Listener,omitempty" xml:"Listener,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) GetListener() []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener {
	return s.Listener
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) SetListener(v []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners {
	s.Listener = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) Validate() error {
	if s.Listener != nil {
		for _, item := range s.Listener {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener struct {
	// The listener port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The listener protocol.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) GetPort() *int32 {
	return s.Port
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) SetPort(v int32) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener {
	s.Port = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) SetProtocol(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener {
	s.Protocol = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) Validate() error {
	return dara.Validate(s)
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags struct {
	Tag []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags) GetTag() []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag {
	return s.Tag
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags) SetTag(v []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags {
	s.Tag = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTags) Validate() error {
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

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// test_slb_yaochi_tag_key-0
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 000098dab00323fb
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) SetTagKey(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) SetTagValue(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupTagsTag) Validate() error {
	return dara.Validate(s)
}
