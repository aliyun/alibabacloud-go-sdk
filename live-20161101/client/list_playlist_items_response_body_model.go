// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaylistItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramItems(v []*ListPlaylistItemsResponseBodyProgramItems) *ListPlaylistItemsResponseBody
	GetProgramItems() []*ListPlaylistItemsResponseBodyProgramItems
	SetRequestId(v string) *ListPlaylistItemsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListPlaylistItemsResponseBody
	GetTotal() *int32
}

type ListPlaylistItemsResponseBody struct {
	// The episodes.
	ProgramItems []*ListPlaylistItemsResponseBodyProgramItems `json:"ProgramItems,omitempty" xml:"ProgramItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of episodes.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPlaylistItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlaylistItemsResponseBody) GetProgramItems() []*ListPlaylistItemsResponseBodyProgramItems {
	return s.ProgramItems
}

func (s *ListPlaylistItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPlaylistItemsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListPlaylistItemsResponseBody) SetProgramItems(v []*ListPlaylistItemsResponseBodyProgramItems) *ListPlaylistItemsResponseBody {
	s.ProgramItems = v
	return s
}

func (s *ListPlaylistItemsResponseBody) SetRequestId(v string) *ListPlaylistItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPlaylistItemsResponseBody) SetTotal(v int32) *ListPlaylistItemsResponseBody {
	s.Total = &v
	return s
}

func (s *ListPlaylistItemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPlaylistItemsResponseBodyProgramItems struct {
	// The sequence number of the episode in the query result.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The ID of the episode list. You can use the ID as a request parameter in the API operation that is used to edit the episode list, delete the episode list, query the information about the episode list, start the episode list, or stop the episode list.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The ID of the episode.
	//
	// example:
	//
	// c10f3d63-eacf-4fbf-bd48-a07a6ba7****
	ProgramItemId *string `json:"ProgramItemId,omitempty" xml:"ProgramItemId,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// playlistItem1
	ProgramItemName *string `json:"ProgramItemName,omitempty" xml:"ProgramItemName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// vod
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// asdfasdf8as9df8sa9df89****
	ResourceValue *string `json:"ResourceValue,omitempty" xml:"ResourceValue,omitempty"`
}

func (s ListPlaylistItemsResponseBodyProgramItems) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistItemsResponseBodyProgramItems) GoString() string {
	return s.String()
}

func (s *ListPlaylistItemsResponseBodyProgramItems) GetIndex() *int32 {
	return s.Index
}

func (s *ListPlaylistItemsResponseBodyProgramItems) GetProgramId() *string {
	return s.ProgramId
}

func (s *ListPlaylistItemsResponseBodyProgramItems) GetProgramItemId() *string {
	return s.ProgramItemId
}

func (s *ListPlaylistItemsResponseBodyProgramItems) GetProgramItemName() *string {
	return s.ProgramItemName
}

func (s *ListPlaylistItemsResponseBodyProgramItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPlaylistItemsResponseBodyProgramItems) GetResourceValue() *string {
	return s.ResourceValue
}

func (s *ListPlaylistItemsResponseBodyProgramItems) SetIndex(v int32) *ListPlaylistItemsResponseBodyProgramItems {
	s.Index = &v
	return s
}

func (s *ListPlaylistItemsResponseBodyProgramItems) SetProgramId(v string) *ListPlaylistItemsResponseBodyProgramItems {
	s.ProgramId = &v
	return s
}

func (s *ListPlaylistItemsResponseBodyProgramItems) SetProgramItemId(v string) *ListPlaylistItemsResponseBodyProgramItems {
	s.ProgramItemId = &v
	return s
}

func (s *ListPlaylistItemsResponseBodyProgramItems) SetProgramItemName(v string) *ListPlaylistItemsResponseBodyProgramItems {
	s.ProgramItemName = &v
	return s
}

func (s *ListPlaylistItemsResponseBodyProgramItems) SetResourceType(v string) *ListPlaylistItemsResponseBodyProgramItems {
	s.ResourceType = &v
	return s
}

func (s *ListPlaylistItemsResponseBodyProgramItems) SetResourceValue(v string) *ListPlaylistItemsResponseBodyProgramItems {
	s.ResourceValue = &v
	return s
}

func (s *ListPlaylistItemsResponseBodyProgramItems) Validate() error {
	return dara.Validate(s)
}
