// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardStrategyBindingItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardStrategyBindingItemsList(v []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) *ListForwardStrategyBindingItemsResponseBody
	GetForwardStrategyBindingItemsList() []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList
	SetItemType(v string) *ListForwardStrategyBindingItemsResponseBody
	GetItemType() *string
	SetRequestId(v string) *ListForwardStrategyBindingItemsResponseBody
	GetRequestId() *string
}

type ListForwardStrategyBindingItemsResponseBody struct {
	ForwardStrategyBindingItemsList []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList `json:"ForwardStrategyBindingItemsList,omitempty" xml:"ForwardStrategyBindingItemsList,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5F79AE39-6622-5292-87EF-DE45631DE4D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListForwardStrategyBindingItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategyBindingItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListForwardStrategyBindingItemsResponseBody) GetForwardStrategyBindingItemsList() []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList {
	return s.ForwardStrategyBindingItemsList
}

func (s *ListForwardStrategyBindingItemsResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *ListForwardStrategyBindingItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListForwardStrategyBindingItemsResponseBody) SetForwardStrategyBindingItemsList(v []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) *ListForwardStrategyBindingItemsResponseBody {
	s.ForwardStrategyBindingItemsList = v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBody) SetItemType(v string) *ListForwardStrategyBindingItemsResponseBody {
	s.ItemType = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBody) SetRequestId(v string) *ListForwardStrategyBindingItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBody) Validate() error {
	if s.ForwardStrategyBindingItemsList != nil {
		for _, item := range s.ForwardStrategyBindingItemsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList struct {
	// example:
	//
	// fs-41a7891ff6568421
	ForwardId *string                                                                            `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	Items     []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
}

func (s ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) GoString() string {
	return s.String()
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) GetForwardId() *string {
	return s.ForwardId
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) GetItems() []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems {
	return s.Items
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) SetForwardId(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList {
	s.ForwardId = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) SetItems(v []*ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList {
	s.Items = v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) SetMatchMode(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList {
	s.MatchMode = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsList) Validate() error {
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

type ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems struct {
	// example:
	//
	// pa-application-104b6b97b7f0c5d9
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 437008
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// tag-4c8b988bb0ffdfb3
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// nieshirui.nsr
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// example:
	//
	// ug-xxxxxx
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// example:
	//
	// aaaaa
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) GoString() string {
	return s.String()
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) GetTagId() *string {
	return s.TagId
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) GetTagName() *string {
	return s.TagName
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) SetApplicationId(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems {
	s.ApplicationId = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) SetApplicationName(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems {
	s.ApplicationName = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) SetTagId(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems {
	s.TagId = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) SetTagName(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems {
	s.TagName = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) SetUserGroupId(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems {
	s.UserGroupId = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) SetUserGroupName(v string) *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems {
	s.UserGroupName = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponseBodyForwardStrategyBindingItemsListItems) Validate() error {
	return dara.Validate(s)
}
