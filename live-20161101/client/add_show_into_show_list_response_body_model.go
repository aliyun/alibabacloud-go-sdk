// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShowIntoShowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddShowIntoShowListResponseBody
	GetRequestId() *string
	SetShowId(v string) *AddShowIntoShowListResponseBody
	GetShowId() *string
	SetFailedList(v string) *AddShowIntoShowListResponseBody
	GetFailedList() *string
	SetSuccessfulShowIds(v string) *AddShowIntoShowListResponseBody
	GetSuccessfulShowIds() *string
}

type AddShowIntoShowListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 987DA143-A39C-5B5D-AF5B-3B07944A0036
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the episode.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
	// The list of resources that failed to be added and the reason for failure.
	//
	// example:
	//
	// failedList[Show1, Show2...]
	FailedList *string `json:"failedList,omitempty" xml:"failedList,omitempty"`
	// The IDs of the episodes that were added.
	//
	// example:
	//
	// f1933f16-5467-4308-b3a9-e8d451a90999,547436b8-c839-4469-a2c0-704c1ce5ce00
	SuccessfulShowIds *string `json:"successfulShowIds,omitempty" xml:"successfulShowIds,omitempty"`
}

func (s AddShowIntoShowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddShowIntoShowListResponseBody) GoString() string {
	return s.String()
}

func (s *AddShowIntoShowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddShowIntoShowListResponseBody) GetShowId() *string {
	return s.ShowId
}

func (s *AddShowIntoShowListResponseBody) GetFailedList() *string {
	return s.FailedList
}

func (s *AddShowIntoShowListResponseBody) GetSuccessfulShowIds() *string {
	return s.SuccessfulShowIds
}

func (s *AddShowIntoShowListResponseBody) SetRequestId(v string) *AddShowIntoShowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddShowIntoShowListResponseBody) SetShowId(v string) *AddShowIntoShowListResponseBody {
	s.ShowId = &v
	return s
}

func (s *AddShowIntoShowListResponseBody) SetFailedList(v string) *AddShowIntoShowListResponseBody {
	s.FailedList = &v
	return s
}

func (s *AddShowIntoShowListResponseBody) SetSuccessfulShowIds(v string) *AddShowIntoShowListResponseBody {
	s.SuccessfulShowIds = &v
	return s
}

func (s *AddShowIntoShowListResponseBody) Validate() error {
	return dara.Validate(s)
}
