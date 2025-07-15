// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveShowFromShowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveShowFromShowListResponseBody
	GetRequestId() *string
	SetShowId(v string) *RemoveShowFromShowListResponseBody
	GetShowId() *string
	SetFailedList(v string) *RemoveShowFromShowListResponseBody
	GetFailedList() *string
	SetSuccessfulShowIds(v string) *RemoveShowFromShowListResponseBody
	GetSuccessfulShowIds() *string
}

type RemoveShowFromShowListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the episode.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
	// The IDs of episodes that failed to be removed and the relevant failure information.
	//
	// example:
	//
	// failedList[Show1, Show2...]
	FailedList *string `json:"failedList,omitempty" xml:"failedList,omitempty"`
	// The IDs of episodes that were removed.
	//
	// example:
	//
	// f1933f16-5467-4308-b3a9-e8d451a90999
	SuccessfulShowIds *string `json:"successfulShowIds,omitempty" xml:"successfulShowIds,omitempty"`
}

func (s RemoveShowFromShowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveShowFromShowListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveShowFromShowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveShowFromShowListResponseBody) GetShowId() *string {
	return s.ShowId
}

func (s *RemoveShowFromShowListResponseBody) GetFailedList() *string {
	return s.FailedList
}

func (s *RemoveShowFromShowListResponseBody) GetSuccessfulShowIds() *string {
	return s.SuccessfulShowIds
}

func (s *RemoveShowFromShowListResponseBody) SetRequestId(v string) *RemoveShowFromShowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveShowFromShowListResponseBody) SetShowId(v string) *RemoveShowFromShowListResponseBody {
	s.ShowId = &v
	return s
}

func (s *RemoveShowFromShowListResponseBody) SetFailedList(v string) *RemoveShowFromShowListResponseBody {
	s.FailedList = &v
	return s
}

func (s *RemoveShowFromShowListResponseBody) SetSuccessfulShowIds(v string) *RemoveShowFromShowListResponseBody {
	s.SuccessfulShowIds = &v
	return s
}

func (s *RemoveShowFromShowListResponseBody) Validate() error {
	return dara.Validate(s)
}
