// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePlayingListResponseBody
	GetCode() *int32
	SetMessage(v string) *CreatePlayingListResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePlayingListResponseBody
	GetRequestId() *string
	SetResult(v *CreatePlayingListResponseBodyResult) *CreatePlayingListResponseBody
	GetResult() *CreatePlayingListResponseBodyResult
	SetSuccess(v string) *CreatePlayingListResponseBody
	GetSuccess() *string
}

type CreatePlayingListResponseBody struct {
	// Return code of the invocation
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Additional information, typically used to briefly describe a failed invocation to help the caller troubleshoot the issue.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10002398812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Actual return result of the service
	Result *CreatePlayingListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the invocation succeeded. true indicates success, and false indicates failure. When the value is false, check the Message field for details.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePlayingListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePlayingListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePlayingListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePlayingListResponseBody) GetResult() *CreatePlayingListResponseBodyResult {
	return s.Result
}

func (s *CreatePlayingListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreatePlayingListResponseBody) SetCode(v int32) *CreatePlayingListResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePlayingListResponseBody) SetMessage(v string) *CreatePlayingListResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePlayingListResponseBody) SetRequestId(v string) *CreatePlayingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePlayingListResponseBody) SetResult(v *CreatePlayingListResponseBodyResult) *CreatePlayingListResponseBody {
	s.Result = v
	return s
}

func (s *CreatePlayingListResponseBody) SetSuccess(v string) *CreatePlayingListResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePlayingListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePlayingListResponseBodyResult struct {
	// Third-party album name
	//
	// example:
	//
	// 晚安妈妈睡前故事
	AlbumName *string `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	// Third-party album ID
	//
	// example:
	//
	// 260744
	AlbumRawId *string `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	// Length
	//
	// example:
	//
	// 190
	AudioLength *int32 `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	// The copyright field is upgraded to indicate whether the content is playable: 0 means playable, 1 or 2 means not playable.
	//
	// example:
	//
	// 0
	Copyright *int32 `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	// thumbnail image object
	Cover *CreatePlayingListResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// Default playback order of the package: 0 for sequential, 1 for reverse.
	//
	// example:
	//
	// 1
	DefaultPlayOrder *int32 `json:"DefaultPlayOrder,omitempty" xml:"DefaultPlayOrder,omitempty"`
	// Playback URL
	//
	// example:
	//
	// https://openaudio.cos.tx.xmcdn.com/storages/587f-audiofreehighqps/15/CE/GKwRIJIGnb11ABc6SwF59DNb.mp3
	ItemUrl *string `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	// is collected
	//
	// example:
	//
	// false
	Liked *bool `json:"Liked,omitempty" xml:"Liked,omitempty"`
	// Lyrics URL
	//
	// example:
	//
	// https://aicontent.alibabausercontent.com/lyric/thirdsource/6f4c8408073db134b0d097c122b5a1a1.lrc
	LyricUrl *string `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	// Playback pattern (Repeat, Shuffle, RepeatOne, Normal)
	//
	// example:
	//
	// Repeat
	PlayMode *string `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	// Position of the item in the playlist
	//
	// example:
	//
	// 0
	Pos *int32 `json:"Pos,omitempty" xml:"Pos,omitempty"`
	// Playback progress of the song
	//
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// third-party ID
	//
	// example:
	//
	// 550144364
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// Author
	//
	// example:
	//
	// 晚安妈妈
	Singer *string `json:"Singer,omitempty" xml:"Singer,omitempty"`
	// Source
	//
	// example:
	//
	// ximalayav2
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// title
	//
	// example:
	//
	// 足球鞋赛车 晚安妈妈
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// type (such as music, program, joke, news, children_song, radio, etc.)
	//
	// example:
	//
	// story
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Playback availability: VALID(10), UNKNOWN(20), NOT_VALID(30)
	//
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s CreatePlayingListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBodyResult) GetAlbumName() *string {
	return s.AlbumName
}

func (s *CreatePlayingListResponseBodyResult) GetAlbumRawId() *string {
	return s.AlbumRawId
}

func (s *CreatePlayingListResponseBodyResult) GetAudioLength() *int32 {
	return s.AudioLength
}

func (s *CreatePlayingListResponseBodyResult) GetCopyright() *int32 {
	return s.Copyright
}

func (s *CreatePlayingListResponseBodyResult) GetCover() *CreatePlayingListResponseBodyResultCover {
	return s.Cover
}

func (s *CreatePlayingListResponseBodyResult) GetDefaultPlayOrder() *int32 {
	return s.DefaultPlayOrder
}

func (s *CreatePlayingListResponseBodyResult) GetItemUrl() *string {
	return s.ItemUrl
}

func (s *CreatePlayingListResponseBodyResult) GetLiked() *bool {
	return s.Liked
}

func (s *CreatePlayingListResponseBodyResult) GetLyricUrl() *string {
	return s.LyricUrl
}

func (s *CreatePlayingListResponseBodyResult) GetPlayMode() *string {
	return s.PlayMode
}

func (s *CreatePlayingListResponseBodyResult) GetPos() *int32 {
	return s.Pos
}

func (s *CreatePlayingListResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *CreatePlayingListResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *CreatePlayingListResponseBodyResult) GetSinger() *string {
	return s.Singer
}

func (s *CreatePlayingListResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *CreatePlayingListResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *CreatePlayingListResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreatePlayingListResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *CreatePlayingListResponseBodyResult) SetAlbumName(v string) *CreatePlayingListResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetAlbumRawId(v string) *CreatePlayingListResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetAudioLength(v int32) *CreatePlayingListResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetCopyright(v int32) *CreatePlayingListResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetCover(v *CreatePlayingListResponseBodyResultCover) *CreatePlayingListResponseBodyResult {
	s.Cover = v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetDefaultPlayOrder(v int32) *CreatePlayingListResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetItemUrl(v string) *CreatePlayingListResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetLiked(v bool) *CreatePlayingListResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetLyricUrl(v string) *CreatePlayingListResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetPlayMode(v string) *CreatePlayingListResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetPos(v int32) *CreatePlayingListResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetProgress(v int32) *CreatePlayingListResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetRawId(v string) *CreatePlayingListResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetSinger(v string) *CreatePlayingListResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetSource(v string) *CreatePlayingListResponseBodyResult {
	s.Source = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetTitle(v string) *CreatePlayingListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetType(v string) *CreatePlayingListResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) SetValid(v string) *CreatePlayingListResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *CreatePlayingListResponseBodyResult) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePlayingListResponseBodyResultCover struct {
	// Indicates whether OSS rules can be used for cropping.
	//
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// default image
	//
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Img *string `json:"Img,omitempty" xml:"Img,omitempty"`
	// Large image
	//
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// Medium image
	//
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Mediam *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	// medium image
	//
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// small image
	//
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s CreatePlayingListResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *CreatePlayingListResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *CreatePlayingListResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *CreatePlayingListResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *CreatePlayingListResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *CreatePlayingListResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *CreatePlayingListResponseBodyResultCover) SetCanResize(v bool) *CreatePlayingListResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetImg(v string) *CreatePlayingListResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetLarge(v string) *CreatePlayingListResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetMediam(v string) *CreatePlayingListResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetMedium(v string) *CreatePlayingListResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) SetSmall(v string) *CreatePlayingListResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *CreatePlayingListResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
