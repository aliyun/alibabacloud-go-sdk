// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPlaylistItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *AddPlaylistItemsResponseBodyItems) *AddPlaylistItemsResponseBody
	GetItems() *AddPlaylistItemsResponseBodyItems
	SetProgramId(v string) *AddPlaylistItemsResponseBody
	GetProgramId() *string
	SetRequestId(v string) *AddPlaylistItemsResponseBody
	GetRequestId() *string
}

type AddPlaylistItemsResponseBody struct {
	// The information about the episodes.
	Items *AddPlaylistItemsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the episode list. You can use the ID as a request parameter in the API operation that is used to remove episodes, query episodes, edit an episode list, delete an episode list, query the information about an episode list, start playing an episode list, or stop playing an episode list.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPlaylistItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPlaylistItemsResponseBody) GoString() string {
	return s.String()
}

func (s *AddPlaylistItemsResponseBody) GetItems() *AddPlaylistItemsResponseBodyItems {
	return s.Items
}

func (s *AddPlaylistItemsResponseBody) GetProgramId() *string {
	return s.ProgramId
}

func (s *AddPlaylistItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPlaylistItemsResponseBody) SetItems(v *AddPlaylistItemsResponseBodyItems) *AddPlaylistItemsResponseBody {
	s.Items = v
	return s
}

func (s *AddPlaylistItemsResponseBody) SetProgramId(v string) *AddPlaylistItemsResponseBody {
	s.ProgramId = &v
	return s
}

func (s *AddPlaylistItemsResponseBody) SetRequestId(v string) *AddPlaylistItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPlaylistItemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddPlaylistItemsResponseBodyItems struct {
	// The episodes that failed to be added.
	FailedItems []*AddPlaylistItemsResponseBodyItemsFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The episodes that were added.
	SuccessItems []*AddPlaylistItemsResponseBodyItemsSuccessItems `json:"SuccessItems,omitempty" xml:"SuccessItems,omitempty" type:"Repeated"`
}

func (s AddPlaylistItemsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s AddPlaylistItemsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *AddPlaylistItemsResponseBodyItems) GetFailedItems() []*AddPlaylistItemsResponseBodyItemsFailedItems {
	return s.FailedItems
}

func (s *AddPlaylistItemsResponseBodyItems) GetSuccessItems() []*AddPlaylistItemsResponseBodyItemsSuccessItems {
	return s.SuccessItems
}

func (s *AddPlaylistItemsResponseBodyItems) SetFailedItems(v []*AddPlaylistItemsResponseBodyItemsFailedItems) *AddPlaylistItemsResponseBodyItems {
	s.FailedItems = v
	return s
}

func (s *AddPlaylistItemsResponseBodyItems) SetSuccessItems(v []*AddPlaylistItemsResponseBodyItemsSuccessItems) *AddPlaylistItemsResponseBodyItems {
	s.SuccessItems = v
	return s
}

func (s *AddPlaylistItemsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type AddPlaylistItemsResponseBodyItemsFailedItems struct {
	// The ID of the episode.
	//
	// example:
	//
	// c09f3d63-eacf-4fbf-bd48-a07a6ba7****
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// item1
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
}

func (s AddPlaylistItemsResponseBodyItemsFailedItems) String() string {
	return dara.Prettify(s)
}

func (s AddPlaylistItemsResponseBodyItemsFailedItems) GoString() string {
	return s.String()
}

func (s *AddPlaylistItemsResponseBodyItemsFailedItems) GetItemId() *string {
	return s.ItemId
}

func (s *AddPlaylistItemsResponseBodyItemsFailedItems) GetItemName() *string {
	return s.ItemName
}

func (s *AddPlaylistItemsResponseBodyItemsFailedItems) SetItemId(v string) *AddPlaylistItemsResponseBodyItemsFailedItems {
	s.ItemId = &v
	return s
}

func (s *AddPlaylistItemsResponseBodyItemsFailedItems) SetItemName(v string) *AddPlaylistItemsResponseBodyItemsFailedItems {
	s.ItemName = &v
	return s
}

func (s *AddPlaylistItemsResponseBodyItemsFailedItems) Validate() error {
	return dara.Validate(s)
}

type AddPlaylistItemsResponseBodyItemsSuccessItems struct {
	// The ID of the episode.
	//
	// example:
	//
	// c09f3d63-eacf-4fbf-bd48-a07a6ba7****
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// item2
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
}

func (s AddPlaylistItemsResponseBodyItemsSuccessItems) String() string {
	return dara.Prettify(s)
}

func (s AddPlaylistItemsResponseBodyItemsSuccessItems) GoString() string {
	return s.String()
}

func (s *AddPlaylistItemsResponseBodyItemsSuccessItems) GetItemId() *string {
	return s.ItemId
}

func (s *AddPlaylistItemsResponseBodyItemsSuccessItems) GetItemName() *string {
	return s.ItemName
}

func (s *AddPlaylistItemsResponseBodyItemsSuccessItems) SetItemId(v string) *AddPlaylistItemsResponseBodyItemsSuccessItems {
	s.ItemId = &v
	return s
}

func (s *AddPlaylistItemsResponseBodyItemsSuccessItems) SetItemName(v string) *AddPlaylistItemsResponseBodyItemsSuccessItems {
	s.ItemName = &v
	return s
}

func (s *AddPlaylistItemsResponseBodyItemsSuccessItems) Validate() error {
	return dara.Validate(s)
}
