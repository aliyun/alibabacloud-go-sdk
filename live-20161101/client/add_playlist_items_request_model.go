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
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster operation](https://help.aliyun.com/document_detail/2848009.html), check the CasterId parameter value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, navigate to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the production studio name.
	//
	// > The production studio name in the production studio list on the Cloud Production Studio page of the ApsaraVideo Live console is the production studio ID.
	//
	//
	// The production studio must meet the following configurations:
	//
	// - **NormType**: **3**. Create a lightweight carousel production studio in advance. You can call the **CreateCaster*	- operation to create a production studio.
	//
	// - **CasterTemplate**: lp_noTranscode.
	//
	// - **channelEnable**: 0.
	//
	// - **programEffect**: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0e94d1f4-1a65-445c-9dcf-de8b3b8d****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The playlist item configuration. If this is the first time you add a playlist item, specify this parameter for initialization. For more information, see **ProgramConfig**.
	//
	// example:
	//
	// [{"RepeatNumber":"0","ProgramName":"my program"}]
	ProgramConfig *string `json:"ProgramConfig,omitempty" xml:"ProgramConfig,omitempty"`
	// The playlist ID. If the production studio already has a playlist, you must specify the corresponding ProgramId. If no playlist has been created, you can leave this parameter empty, and the system performs automatic creation.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The list of playlist item inputs. The value is a JSON string. For more information, see **InputProgramItem**.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ItemName":"item1","ResourceType":"vod","ResourceValue":"5f8809f2-3352-4d1f-a8f7-86f9429f****"}, {"ItemName": "item2","ResourceType": "vod","ResourceValue": "e7411c0b-dd98-4c61-a545-f8bfba6c****"}]
	ProgramItems *string `json:"ProgramItems,omitempty" xml:"ProgramItems,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
