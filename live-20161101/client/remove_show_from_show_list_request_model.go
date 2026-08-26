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
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value that is returned in the response.
	//
	// - If you created the production studio in the LIVE console, find the production studio name in the LIVE console by choosing **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio**.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The show ID.
	//
	// > Obtain the ShowId from the response of the [AddShowIntoShowList](https://help.aliyun.com/document_detail/2848051.html) operation.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
	// Specifies whether to delete shows in a batch. Valid values:
	//
	// - true: Deletes shows in a batch.
	//
	// - false: Deletes a single show.
	//
	// > If you do not specify this parameter or leave it empty, a single show is deleted.
	//
	// example:
	//
	// false
	IsBatchMode *bool `json:"isBatchMode,omitempty" xml:"isBatchMode,omitempty"`
	// The IDs of the shows to delete.
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
