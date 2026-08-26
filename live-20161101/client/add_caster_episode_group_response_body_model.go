// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItemIds(v *AddCasterEpisodeGroupResponseBodyItemIds) *AddCasterEpisodeGroupResponseBody
	GetItemIds() *AddCasterEpisodeGroupResponseBodyItemIds
	SetProgramId(v string) *AddCasterEpisodeGroupResponseBody
	GetProgramId() *string
	SetRequestId(v string) *AddCasterEpisodeGroupResponseBody
	GetRequestId() *string
}

type AddCasterEpisodeGroupResponseBody struct {
	ItemIds *AddCasterEpisodeGroupResponseBodyItemIds `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Struct"`
	// The program list ID. This ID can be used as a request parameter for deleting a production studio program list, creating or adding program list items, deleting program list items, querying program list items, editing a program list, deleting a program list, querying a program list, starting a program list, or stopping a program list.
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

func (s AddCasterEpisodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupResponseBody) GetItemIds() *AddCasterEpisodeGroupResponseBodyItemIds {
	return s.ItemIds
}

func (s *AddCasterEpisodeGroupResponseBody) GetProgramId() *string {
	return s.ProgramId
}

func (s *AddCasterEpisodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasterEpisodeGroupResponseBody) SetItemIds(v *AddCasterEpisodeGroupResponseBodyItemIds) *AddCasterEpisodeGroupResponseBody {
	s.ItemIds = v
	return s
}

func (s *AddCasterEpisodeGroupResponseBody) SetProgramId(v string) *AddCasterEpisodeGroupResponseBody {
	s.ProgramId = &v
	return s
}

func (s *AddCasterEpisodeGroupResponseBody) SetRequestId(v string) *AddCasterEpisodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasterEpisodeGroupResponseBody) Validate() error {
	if s.ItemIds != nil {
		if err := s.ItemIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddCasterEpisodeGroupResponseBodyItemIds struct {
	ItemId []*string `json:"ItemId,omitempty" xml:"ItemId,omitempty" type:"Repeated"`
}

func (s AddCasterEpisodeGroupResponseBodyItemIds) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeGroupResponseBodyItemIds) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupResponseBodyItemIds) GetItemId() []*string {
	return s.ItemId
}

func (s *AddCasterEpisodeGroupResponseBodyItemIds) SetItemId(v []*string) *AddCasterEpisodeGroupResponseBodyItemIds {
	s.ItemId = v
	return s
}

func (s *AddCasterEpisodeGroupResponseBodyItemIds) Validate() error {
	return dara.Validate(s)
}
