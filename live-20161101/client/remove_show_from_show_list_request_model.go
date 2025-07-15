// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveShowFromShowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *RemoveShowFromShowListRequest
	GetCasterId() *string
	SetOwnerId(v int64) *RemoveShowFromShowListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveShowFromShowListRequest
	GetRegionId() *string
	SetShowId(v string) *RemoveShowFromShowListRequest
	GetShowId() *string
	SetIsBatchMode(v bool) *RemoveShowFromShowListRequest
	GetIsBatchMode() *bool
	SetShowIdList(v []*string) *RemoveShowFromShowListRequest
	GetShowIdList() []*string
}

type RemoveShowFromShowListRequest struct {
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/69338.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the episode.
	//
	// >  You can obtain the ID by checking the value of the response parameter ShowId of the [AddShowIntoShowList](https://help.aliyun.com/document_detail/370861.html) operation.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
	// Specifies whether to remove multiple episodes at a time. Valid values:
	//
	// 	- true: removes multiple episodes at a time.
	//
	// 	- false: removes a single episode.
	//
	// >  If you do not configure this parameter or this parameter is left empty, a single episode is to be removed.
	//
	// example:
	//
	// false
	IsBatchMode *bool `json:"isBatchMode,omitempty" xml:"isBatchMode,omitempty"`
	// The IDs of episodes that you want to remove.
	ShowIdList []*string `json:"showIdList,omitempty" xml:"showIdList,omitempty" type:"Repeated"`
}

func (s RemoveShowFromShowListRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveShowFromShowListRequest) GoString() string {
	return s.String()
}

func (s *RemoveShowFromShowListRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *RemoveShowFromShowListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveShowFromShowListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveShowFromShowListRequest) GetShowId() *string {
	return s.ShowId
}

func (s *RemoveShowFromShowListRequest) GetIsBatchMode() *bool {
	return s.IsBatchMode
}

func (s *RemoveShowFromShowListRequest) GetShowIdList() []*string {
	return s.ShowIdList
}

func (s *RemoveShowFromShowListRequest) SetCasterId(v string) *RemoveShowFromShowListRequest {
	s.CasterId = &v
	return s
}

func (s *RemoveShowFromShowListRequest) SetOwnerId(v int64) *RemoveShowFromShowListRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveShowFromShowListRequest) SetRegionId(v string) *RemoveShowFromShowListRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveShowFromShowListRequest) SetShowId(v string) *RemoveShowFromShowListRequest {
	s.ShowId = &v
	return s
}

func (s *RemoveShowFromShowListRequest) SetIsBatchMode(v bool) *RemoveShowFromShowListRequest {
	s.IsBatchMode = &v
	return s
}

func (s *RemoveShowFromShowListRequest) SetShowIdList(v []*string) *RemoveShowFromShowListRequest {
	s.ShowIdList = v
	return s
}

func (s *RemoveShowFromShowListRequest) Validate() error {
	return dara.Validate(s)
}
