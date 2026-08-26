// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeGroupContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddCasterEpisodeGroupContentRequest
	GetClientToken() *string
	SetContent(v string) *AddCasterEpisodeGroupContentRequest
	GetContent() *string
	SetOwnerId(v int64) *AddCasterEpisodeGroupContentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddCasterEpisodeGroupContentRequest
	GetRegionId() *string
}

type AddCasterEpisodeGroupContentRequest struct {
	// A client-generated token that is used to ensure the idempotence of the request.
	//
	// > The client generates this value. Make sure that the value is unique among different requests. The value can be up to 64 ASCII characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8751ad99-2ddb-4aac-ad44-84b21102****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The properties of the episode in the production studio. This parameter is a JSON string. The parameter names are in upper camel case. The properties are described as follows:
	//
	// - **CallbackUrl**: The webhook address.
	//
	// - **SideOutputUrl**: The custom bypass output URL.
	//
	// - **RepeatNum**: The number of times to loop the episode. A value of 0 means the episode does not loop. A value of -1 means the episode loops indefinitely.
	//
	// - **StartTime**: The start time in UTC. The format is *yyyy-MM-dd*T*HH:mm:ss*Z.
	//
	// - **DomainName**: The domain name.
	//
	// - **Items**
	//
	//   : The list of items in the episode.
	//
	//   - **ItemName**: The item name.
	//
	//   - **VodUrl**: The URL of the video-on-demand (VOD) file. This parameter is required only when the resource is a video file that has not been imported to the Material Library. The MP4, FLV, and TS formats are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"CallbackUrl":"http://example.aliyundoc.com/callBackLive","SideOutputUrl":"rtmp://guide.aliyundoc.com/caster/4a82a3d1b7f0462ea37348366201****?auth_key=1608953344-0-0-ac8c628078541d7055a170ec59a5****","DomainName":"developer.aliyundoc.com ","StartTime":"2018-03-26T16:00:00Z","RepeatNum":-1,"Items":[{"ItemName":"program1","VodUrl":"http://learn.aliyundoc.com"},{"ItemName":"program2","VodUrl":"http://demo.aliyundoc.com"}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddCasterEpisodeGroupContentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeGroupContentRequest) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupContentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddCasterEpisodeGroupContentRequest) GetContent() *string {
	return s.Content
}

func (s *AddCasterEpisodeGroupContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCasterEpisodeGroupContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCasterEpisodeGroupContentRequest) SetClientToken(v string) *AddCasterEpisodeGroupContentRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCasterEpisodeGroupContentRequest) SetContent(v string) *AddCasterEpisodeGroupContentRequest {
	s.Content = &v
	return s
}

func (s *AddCasterEpisodeGroupContentRequest) SetOwnerId(v int64) *AddCasterEpisodeGroupContentRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCasterEpisodeGroupContentRequest) SetRegionId(v string) *AddCasterEpisodeGroupContentRequest {
	s.RegionId = &v
	return s
}

func (s *AddCasterEpisodeGroupContentRequest) Validate() error {
	return dara.Validate(s)
}
