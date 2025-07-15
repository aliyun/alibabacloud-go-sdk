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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8751ad99-2ddb-4aac-ad44-84b21102****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the episode list. The value is a JSON string. Use upper camel case for fields of the string. This parameter contains the following fields:
	//
	// 	- **CallbackUrl**: the callback URL.
	//
	// 	- **SideOutputUrl**: the custom standby URL.
	//
	// 	- **RepeatNum**: the number of times the episode list repeats after the first playback is complete. A value of 0 indicates that the episode list is played only once. A value of -1 indicates that the episode list is played in loop mode.
	//
	// 	- **DomainName**: the domain name.
	//
	// 	- **StartTime**: the time when the episode list starts to play. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// 	- **Items**: the information about the episode list. It is an array of ItemName and VodUrl.
	//
	//     	- **ItemName**: the name of the episode.
	//
	//     	- **VodUrl**: the URL of the VOD file. This field takes effect only when the video resource is a video file that is not from the media library. The video file must be in the MP4, FLV, or TS format.
	//
	// This parameter is required.
	//
	// example:
	//
	// CallbackUrl
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
