// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPlaylistItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *AddPlaylistItemsRequest
	GetCasterId() *string
	SetOwnerId(v int64) *AddPlaylistItemsRequest
	GetOwnerId() *int64
	SetProgramConfig(v string) *AddPlaylistItemsRequest
	GetProgramConfig() *string
	SetProgramId(v string) *AddPlaylistItemsRequest
	GetProgramId() *string
	SetProgramItems(v string) *AddPlaylistItemsRequest
	GetProgramItems() *string
	SetRegionId(v string) *AddPlaylistItemsRequest
	GetRegionId() *string
}

type AddPlaylistItemsRequest struct {
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// The production studio must use the following configurations:
	//
	// 	- **NormType**: 3****. You need to call the **CreateCaster*	- operation to create a production studio for lightweight carousel playback in advance.
	//
	// 	- **CasterTemplate**: lp_noTranscode.
	//
	// 	- **channelEnable**: 0.
	//
	// 	- **programEffect**: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0e94d1f4-1a65-445c-9dcf-de8b3b8d****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The configurations of the episode list. If the episode list is added to the production studio for the first time, specify this parameter to pass in the initial configurations. For more information, see the **ProgramConfig*	- section of this topic.
	//
	// example:
	//
	// [{"RepeatNumber":"0","ProgramName":"my program"}]
	ProgramConfig *string `json:"ProgramConfig,omitempty" xml:"ProgramConfig,omitempty"`
	// The ID of the episode list. If you do not specify this parameter, an episode list is created by default.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The episodes that you want to add to the production studio. The value is a JSON string. For more information, see the **InputProgramItem*	- section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ItemName":"item1","ResourceType":"vod","ResourceValue":"5f8809f2-3352-4d1f-a8f7-86f9429f****"}, {"ItemName": "item2","ResourceType": "vod","ResourceValue": "e7411c0b-dd98-4c61-a545-f8bfba6c****"}]
	ProgramItems *string `json:"ProgramItems,omitempty" xml:"ProgramItems,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddPlaylistItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPlaylistItemsRequest) GoString() string {
	return s.String()
}

func (s *AddPlaylistItemsRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddPlaylistItemsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddPlaylistItemsRequest) GetProgramConfig() *string {
	return s.ProgramConfig
}

func (s *AddPlaylistItemsRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *AddPlaylistItemsRequest) GetProgramItems() *string {
	return s.ProgramItems
}

func (s *AddPlaylistItemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPlaylistItemsRequest) SetCasterId(v string) *AddPlaylistItemsRequest {
	s.CasterId = &v
	return s
}

func (s *AddPlaylistItemsRequest) SetOwnerId(v int64) *AddPlaylistItemsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddPlaylistItemsRequest) SetProgramConfig(v string) *AddPlaylistItemsRequest {
	s.ProgramConfig = &v
	return s
}

func (s *AddPlaylistItemsRequest) SetProgramId(v string) *AddPlaylistItemsRequest {
	s.ProgramId = &v
	return s
}

func (s *AddPlaylistItemsRequest) SetProgramItems(v string) *AddPlaylistItemsRequest {
	s.ProgramItems = &v
	return s
}

func (s *AddPlaylistItemsRequest) SetRegionId(v string) *AddPlaylistItemsRequest {
	s.RegionId = &v
	return s
}

func (s *AddPlaylistItemsRequest) Validate() error {
	return dara.Validate(s)
}
