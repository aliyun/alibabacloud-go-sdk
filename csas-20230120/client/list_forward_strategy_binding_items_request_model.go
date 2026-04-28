// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardStrategyBindingItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardIds(v []*string) *ListForwardStrategyBindingItemsRequest
	GetForwardIds() []*string
	SetItemType(v string) *ListForwardStrategyBindingItemsRequest
	GetItemType() *string
}

type ListForwardStrategyBindingItemsRequest struct {
	// This parameter is required.
	ForwardIds []*string `json:"ForwardIds,omitempty" xml:"ForwardIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
}

func (s ListForwardStrategyBindingItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategyBindingItemsRequest) GoString() string {
	return s.String()
}

func (s *ListForwardStrategyBindingItemsRequest) GetForwardIds() []*string {
	return s.ForwardIds
}

func (s *ListForwardStrategyBindingItemsRequest) GetItemType() *string {
	return s.ItemType
}

func (s *ListForwardStrategyBindingItemsRequest) SetForwardIds(v []*string) *ListForwardStrategyBindingItemsRequest {
	s.ForwardIds = v
	return s
}

func (s *ListForwardStrategyBindingItemsRequest) SetItemType(v string) *ListForwardStrategyBindingItemsRequest {
	s.ItemType = &v
	return s
}

func (s *ListForwardStrategyBindingItemsRequest) Validate() error {
	return dara.Validate(s)
}
