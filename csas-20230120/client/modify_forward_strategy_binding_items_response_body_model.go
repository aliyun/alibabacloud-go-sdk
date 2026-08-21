// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardStrategyBindingItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardStrategyBindingItems(v *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) *ModifyForwardStrategyBindingItemsResponseBody
	GetForwardStrategyBindingItems() *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems
	SetRequestId(v string) *ModifyForwardStrategyBindingItemsResponseBody
	GetRequestId() *string
}

type ModifyForwardStrategyBindingItemsResponseBody struct {
	// The binding items of the forwarding rule after this modification.
	ForwardStrategyBindingItems *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems `json:"ForwardStrategyBindingItems,omitempty" xml:"ForwardStrategyBindingItems,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5F79AE39-6622-5292-87EF-DE45631DE4D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyForwardStrategyBindingItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyBindingItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyBindingItemsResponseBody) GetForwardStrategyBindingItems() *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems {
	return s.ForwardStrategyBindingItems
}

func (s *ModifyForwardStrategyBindingItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyForwardStrategyBindingItemsResponseBody) SetForwardStrategyBindingItems(v *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) *ModifyForwardStrategyBindingItemsResponseBody {
	s.ForwardStrategyBindingItems = v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBody) SetRequestId(v string) *ModifyForwardStrategyBindingItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBody) Validate() error {
	if s.ForwardStrategyBindingItems != nil {
		if err := s.ForwardStrategyBindingItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems struct {
	// The forwarding rule ID.
	//
	// example:
	//
	// fs-849ac29396d9ea98
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	// The binding content. This parameter is not returned when MatchMode is **UserGroupAll*	- or **ApplicationAll**.
	Items []*ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The policy matching target type. Valid values:
	//
	// - **UserGroupAll**: associates with all users.
	//
	// - **UserGroupNormal**: associates with specific user groups.
	//
	// - **ApplicationAll**: all private network applications.
	//
	// - **Application**: specific private network applications.
	//
	// - **Tag**: private network application tags.
	//
	// example:
	//
	// Application
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
}

func (s ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) GetForwardId() *string {
	return s.ForwardId
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) GetItems() []*ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems {
	return s.Items
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) SetForwardId(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems {
	s.ForwardId = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) SetItems(v []*ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems {
	s.Items = v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) SetMatchMode(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems {
	s.MatchMode = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItems) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems struct {
	// The private network access application ID.
	//
	// example:
	//
	// pa-application-104b6b97b7f0c5d9
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The private network access application name.
	//
	// example:
	//
	// OA System
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The private network access tag ID.
	//
	// example:
	//
	// tag-4c8b988bb0ffdfb3
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag name.
	//
	// example:
	//
	// Finance System
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// usergroup-3f9a2c7e10b4d856
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user group name.
	//
	// example:
	//
	// R&D Department
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) GetTagId() *string {
	return s.TagId
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) GetTagName() *string {
	return s.TagName
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) SetApplicationId(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems {
	s.ApplicationId = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) SetApplicationName(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems {
	s.ApplicationName = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) SetTagId(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems {
	s.TagId = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) SetTagName(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems {
	s.TagName = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) SetUserGroupId(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems {
	s.UserGroupId = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) SetUserGroupName(v string) *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems {
	s.UserGroupName = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsItems) Validate() error {
	return dara.Validate(s)
}
