// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeGroupContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItemIds(v *AddCasterEpisodeGroupContentResponseBodyItemIds) *AddCasterEpisodeGroupContentResponseBody
	GetItemIds() *AddCasterEpisodeGroupContentResponseBodyItemIds
	SetProgramId(v string) *AddCasterEpisodeGroupContentResponseBody
	GetProgramId() *string
	SetRequestId(v string) *AddCasterEpisodeGroupContentResponseBody
	GetRequestId() *string
}

type AddCasterEpisodeGroupContentResponseBody struct {
	// The IDs of the episodes.
	ItemIds *AddCasterEpisodeGroupContentResponseBodyItemIds `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Struct"`
	// The ID of the episode list. You can use the ID as a request parameter in the API operation that is used to create or add episodes, remove episodes, query episodes, edit an episode list, delete an episode list, query the information about an episode list, start playing an episode list, or stop playing an episode list.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68X****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCasterEpisodeGroupContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeGroupContentResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupContentResponseBody) GetItemIds() *AddCasterEpisodeGroupContentResponseBodyItemIds {
	return s.ItemIds
}

func (s *AddCasterEpisodeGroupContentResponseBody) GetProgramId() *string {
	return s.ProgramId
}

func (s *AddCasterEpisodeGroupContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasterEpisodeGroupContentResponseBody) SetItemIds(v *AddCasterEpisodeGroupContentResponseBodyItemIds) *AddCasterEpisodeGroupContentResponseBody {
	s.ItemIds = v
	return s
}

func (s *AddCasterEpisodeGroupContentResponseBody) SetProgramId(v string) *AddCasterEpisodeGroupContentResponseBody {
	s.ProgramId = &v
	return s
}

func (s *AddCasterEpisodeGroupContentResponseBody) SetRequestId(v string) *AddCasterEpisodeGroupContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasterEpisodeGroupContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddCasterEpisodeGroupContentResponseBodyItemIds struct {
	ItemId []*string `json:"ItemId,omitempty" xml:"ItemId,omitempty" type:"Repeated"`
}

func (s AddCasterEpisodeGroupContentResponseBodyItemIds) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeGroupContentResponseBodyItemIds) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupContentResponseBodyItemIds) GetItemId() []*string {
	return s.ItemId
}

func (s *AddCasterEpisodeGroupContentResponseBodyItemIds) SetItemId(v []*string) *AddCasterEpisodeGroupContentResponseBodyItemIds {
	s.ItemId = v
	return s
}

func (s *AddCasterEpisodeGroupContentResponseBodyItemIds) Validate() error {
	return dara.Validate(s)
}
