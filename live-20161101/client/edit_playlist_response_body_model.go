// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditPlaylistResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCasterId(v string) *EditPlaylistResponseBody
  GetCasterId() *string 
  SetItems(v *EditPlaylistResponseBodyItems) *EditPlaylistResponseBody
  GetItems() *EditPlaylistResponseBodyItems 
  SetProgramId(v string) *EditPlaylistResponseBody
  GetProgramId() *string 
  SetRequestId(v string) *EditPlaylistResponseBody
  GetRequestId() *string 
}

type EditPlaylistResponseBody struct {
  // The ID of the production studio. Use this ID as a request parameter to add a callback or a virtual studio layout.
  // 
  // example:
  // 
  // 0e94d1f4-1a65-445c-9dcf-de8b3b8d****
  CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
  // The list of added playlist items.
  Items *EditPlaylistResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
  // The playlist ID. Use this ID as a request parameter to delete, query, start, or stop the playlist.
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

func (s EditPlaylistResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditPlaylistResponseBody) GoString() string {
  return s.String()
}

func (s *EditPlaylistResponseBody) GetCasterId() *string  {
  return s.CasterId
}

func (s *EditPlaylistResponseBody) GetItems() *EditPlaylistResponseBodyItems  {
  return s.Items
}

func (s *EditPlaylistResponseBody) GetProgramId() *string  {
  return s.ProgramId
}

func (s *EditPlaylistResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditPlaylistResponseBody) SetCasterId(v string) *EditPlaylistResponseBody {
  s.CasterId = &v
  return s
}

func (s *EditPlaylistResponseBody) SetItems(v *EditPlaylistResponseBodyItems) *EditPlaylistResponseBody {
  s.Items = v
  return s
}

func (s *EditPlaylistResponseBody) SetProgramId(v string) *EditPlaylistResponseBody {
  s.ProgramId = &v
  return s
}

func (s *EditPlaylistResponseBody) SetRequestId(v string) *EditPlaylistResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditPlaylistResponseBody) Validate() error {
  if s.Items != nil {
    if err := s.Items.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EditPlaylistResponseBodyItems struct {
  // The list of playlist items that failed to be added.
  FailedItems []*EditPlaylistResponseBodyItemsFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
  // The list of playlist items that were successfully added.
  SuccessItems []*EditPlaylistResponseBodyItemsSuccessItems `json:"SuccessItems,omitempty" xml:"SuccessItems,omitempty" type:"Repeated"`
}

func (s EditPlaylistResponseBodyItems) String() string {
  return dara.Prettify(s)
}

func (s EditPlaylistResponseBodyItems) GoString() string {
  return s.String()
}

func (s *EditPlaylistResponseBodyItems) GetFailedItems() []*EditPlaylistResponseBodyItemsFailedItems  {
  return s.FailedItems
}

func (s *EditPlaylistResponseBodyItems) GetSuccessItems() []*EditPlaylistResponseBodyItemsSuccessItems  {
  return s.SuccessItems
}

func (s *EditPlaylistResponseBodyItems) SetFailedItems(v []*EditPlaylistResponseBodyItemsFailedItems) *EditPlaylistResponseBodyItems {
  s.FailedItems = v
  return s
}

func (s *EditPlaylistResponseBodyItems) SetSuccessItems(v []*EditPlaylistResponseBodyItemsSuccessItems) *EditPlaylistResponseBodyItems {
  s.SuccessItems = v
  return s
}

func (s *EditPlaylistResponseBodyItems) Validate() error {
  if s.FailedItems != nil {
    for _, item := range s.FailedItems {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.SuccessItems != nil {
    for _, item := range s.SuccessItems {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EditPlaylistResponseBodyItemsFailedItems struct {
  // The ID of the playlist item.
  // 
  // example:
  // 
  // c09f3d63-eacf-4fbf-bd48-a07a6ba7****
  ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
  // The name of the playlist item.
  // 
  // example:
  // 
  // item2
  ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
}

func (s EditPlaylistResponseBodyItemsFailedItems) String() string {
  return dara.Prettify(s)
}

func (s EditPlaylistResponseBodyItemsFailedItems) GoString() string {
  return s.String()
}

func (s *EditPlaylistResponseBodyItemsFailedItems) GetItemId() *string  {
  return s.ItemId
}

func (s *EditPlaylistResponseBodyItemsFailedItems) GetItemName() *string  {
  return s.ItemName
}

func (s *EditPlaylistResponseBodyItemsFailedItems) SetItemId(v string) *EditPlaylistResponseBodyItemsFailedItems {
  s.ItemId = &v
  return s
}

func (s *EditPlaylistResponseBodyItemsFailedItems) SetItemName(v string) *EditPlaylistResponseBodyItemsFailedItems {
  s.ItemName = &v
  return s
}

func (s *EditPlaylistResponseBodyItemsFailedItems) Validate() error {
  return dara.Validate(s)
}

type EditPlaylistResponseBodyItemsSuccessItems struct {
  // The ID of the playlist item.
  // 
  // example:
  // 
  // c09f3d63-eacf-4fbf-bd48-a07a6ba7****
  ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
  // The name of the playlist item.
  // 
  // example:
  // 
  // item1
  ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
}

func (s EditPlaylistResponseBodyItemsSuccessItems) String() string {
  return dara.Prettify(s)
}

func (s EditPlaylistResponseBodyItemsSuccessItems) GoString() string {
  return s.String()
}

func (s *EditPlaylistResponseBodyItemsSuccessItems) GetItemId() *string  {
  return s.ItemId
}

func (s *EditPlaylistResponseBodyItemsSuccessItems) GetItemName() *string  {
  return s.ItemName
}

func (s *EditPlaylistResponseBodyItemsSuccessItems) SetItemId(v string) *EditPlaylistResponseBodyItemsSuccessItems {
  s.ItemId = &v
  return s
}

func (s *EditPlaylistResponseBodyItemsSuccessItems) SetItemName(v string) *EditPlaylistResponseBodyItemsSuccessItems {
  s.ItemName = &v
  return s
}

func (s *EditPlaylistResponseBodyItemsSuccessItems) Validate() error {
  return dara.Validate(s)
}

