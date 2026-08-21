// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardStrategyBindingItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardId(v string) *ModifyForwardStrategyBindingItemsRequest
	GetForwardId() *string
	SetItemIds(v []*string) *ModifyForwardStrategyBindingItemsRequest
	GetItemIds() []*string
	SetMatchMode(v string) *ModifyForwardStrategyBindingItemsRequest
	GetMatchMode() *string
	SetModifyType(v string) *ModifyForwardStrategyBindingItemsRequest
	GetModifyType() *string
}

type ModifyForwardStrategyBindingItemsRequest struct {
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fs-849ac29396d9ea98
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	// The list of binding item IDs. Must be empty when MatchMode is **UserGroupAll*	- or **ApplicationAll**. Required for other values. Duplicates are not allowed in the list, and the specified objects must already exist.
	ItemIds []*string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Repeated"`
	// The policy matching target type. Required. Valid values:
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
	// When the value is **UserGroupAll*	- or **ApplicationAll**, ItemIds must be empty. When the value is **UserGroupNormal**, **Application**, or **Tag**, ItemIds is required.
	//
	// When ModifyType is not **Cover**, switching the matching target type is not allowed: **Application**, **Tag**, and **ApplicationAll*	- are mutually exclusive, and **UserGroupNormal*	- and **UserGroupAll*	- are mutually exclusive. If a binding item of a mutually exclusive type already exists on the same forwarding rule, the request is rejected.
	//
	// example:
	//
	// Application
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The modification method. Required. Valid values:
	//
	// - **Append**: appends to existing binding items. ItemIds cannot contain objects that are already bound.
	//
	// - **Delete**: deletes specified binding items. All objects in ItemIds must be already bound.
	//
	// - **Cover**: overwrites binding items of the same category by clearing all existing binding items of the same category on the forwarding rule and then writing ItemIds. The same category refers to **ApplicationAll**, **Application**, and **Tag**, or **UserGroupAll*	- and **UserGroupNormal**.
	//
	// When the value is **Append*	- or **Delete**, MatchMode cannot be **UserGroupAll*	- or **ApplicationAll**.
	//
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
}

func (s ModifyForwardStrategyBindingItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyBindingItemsRequest) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyBindingItemsRequest) GetForwardId() *string {
	return s.ForwardId
}

func (s *ModifyForwardStrategyBindingItemsRequest) GetItemIds() []*string {
	return s.ItemIds
}

func (s *ModifyForwardStrategyBindingItemsRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ModifyForwardStrategyBindingItemsRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyForwardStrategyBindingItemsRequest) SetForwardId(v string) *ModifyForwardStrategyBindingItemsRequest {
	s.ForwardId = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsRequest) SetItemIds(v []*string) *ModifyForwardStrategyBindingItemsRequest {
	s.ItemIds = v
	return s
}

func (s *ModifyForwardStrategyBindingItemsRequest) SetMatchMode(v string) *ModifyForwardStrategyBindingItemsRequest {
	s.MatchMode = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsRequest) SetModifyType(v string) *ModifyForwardStrategyBindingItemsRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsRequest) Validate() error {
	return dara.Validate(s)
}
