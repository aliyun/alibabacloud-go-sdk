// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexControlPlayingListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *IndexControlPlayingListResponseBody
	GetCode() *int32
	SetMessage(v string) *IndexControlPlayingListResponseBody
	GetMessage() *string
	SetRequestId(v string) *IndexControlPlayingListResponseBody
	GetRequestId() *string
	SetResult(v *IndexControlPlayingListResponseBodyResult) *IndexControlPlayingListResponseBody
	GetResult() *IndexControlPlayingListResponseBodyResult
	SetSuccess(v string) *IndexControlPlayingListResponseBody
	GetSuccess() *string
}

type IndexControlPlayingListResponseBody struct {
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
	// Actual return result from the service
	Result *IndexControlPlayingListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the invocation succeeded. true indicates success, and false indicates failure. If the invocation fails, check the Message field for details.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IndexControlPlayingListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *IndexControlPlayingListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IndexControlPlayingListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IndexControlPlayingListResponseBody) GetResult() *IndexControlPlayingListResponseBodyResult {
	return s.Result
}

func (s *IndexControlPlayingListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *IndexControlPlayingListResponseBody) SetCode(v int32) *IndexControlPlayingListResponseBody {
	s.Code = &v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetMessage(v string) *IndexControlPlayingListResponseBody {
	s.Message = &v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetRequestId(v string) *IndexControlPlayingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetResult(v *IndexControlPlayingListResponseBodyResult) *IndexControlPlayingListResponseBody {
	s.Result = v
	return s
}

func (s *IndexControlPlayingListResponseBody) SetSuccess(v string) *IndexControlPlayingListResponseBody {
	s.Success = &v
	return s
}

func (s *IndexControlPlayingListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IndexControlPlayingListResponseBodyResult struct {
	// Third-party album Name
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
	// The copyright field is upgraded to indicate whether playback is allowed: 0 means playable, 1 or 2 means not playable.
	//
	// example:
	//
	// 0
	Copyright *int32 `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	// Thumbnail image object
	Cover *IndexControlPlayingListResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// Default playback order for the package: 0 for sequential order, 1 for reverse order.
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
	// Indicates whether the item is collected
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
	// Playback mode (Repeat, Shuffle, RepeatOne, Normal)
	//
	// example:
	//
	// Repeat
	PlayMode *string `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	// The position of this item in the playlist
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
	// Third-party ID
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
	// Title
	//
	// example:
	//
	// 足球鞋赛车 晚安妈妈
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Type (such as music, program, joke, news, children_song, radio, etc.)
	//
	// example:
	//
	// story
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether playback is allowed: VALID(10), UNKNOWN(20), NOT_VALID(30)
	//
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s IndexControlPlayingListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponseBodyResult) GetAlbumName() *string {
	return s.AlbumName
}

func (s *IndexControlPlayingListResponseBodyResult) GetAlbumRawId() *string {
	return s.AlbumRawId
}

func (s *IndexControlPlayingListResponseBodyResult) GetAudioLength() *int32 {
	return s.AudioLength
}

func (s *IndexControlPlayingListResponseBodyResult) GetCopyright() *int32 {
	return s.Copyright
}

func (s *IndexControlPlayingListResponseBodyResult) GetCover() *IndexControlPlayingListResponseBodyResultCover {
	return s.Cover
}

func (s *IndexControlPlayingListResponseBodyResult) GetDefaultPlayOrder() *int32 {
	return s.DefaultPlayOrder
}

func (s *IndexControlPlayingListResponseBodyResult) GetItemUrl() *string {
	return s.ItemUrl
}

func (s *IndexControlPlayingListResponseBodyResult) GetLiked() *bool {
	return s.Liked
}

func (s *IndexControlPlayingListResponseBodyResult) GetLyricUrl() *string {
	return s.LyricUrl
}

func (s *IndexControlPlayingListResponseBodyResult) GetPlayMode() *string {
	return s.PlayMode
}

func (s *IndexControlPlayingListResponseBodyResult) GetPos() *int32 {
	return s.Pos
}

func (s *IndexControlPlayingListResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *IndexControlPlayingListResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *IndexControlPlayingListResponseBodyResult) GetSinger() *string {
	return s.Singer
}

func (s *IndexControlPlayingListResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *IndexControlPlayingListResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *IndexControlPlayingListResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *IndexControlPlayingListResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *IndexControlPlayingListResponseBodyResult) SetAlbumName(v string) *IndexControlPlayingListResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetAlbumRawId(v string) *IndexControlPlayingListResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetAudioLength(v int32) *IndexControlPlayingListResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetCopyright(v int32) *IndexControlPlayingListResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetCover(v *IndexControlPlayingListResponseBodyResultCover) *IndexControlPlayingListResponseBodyResult {
	s.Cover = v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetDefaultPlayOrder(v int32) *IndexControlPlayingListResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetItemUrl(v string) *IndexControlPlayingListResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetLiked(v bool) *IndexControlPlayingListResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetLyricUrl(v string) *IndexControlPlayingListResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetPlayMode(v string) *IndexControlPlayingListResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetPos(v int32) *IndexControlPlayingListResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetProgress(v int32) *IndexControlPlayingListResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetRawId(v string) *IndexControlPlayingListResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetSinger(v string) *IndexControlPlayingListResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetSource(v string) *IndexControlPlayingListResponseBodyResult {
	s.Source = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetTitle(v string) *IndexControlPlayingListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetType(v string) *IndexControlPlayingListResponseBodyResult {
	s.Type = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) SetValid(v string) *IndexControlPlayingListResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResult) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IndexControlPlayingListResponseBodyResultCover struct {
	// Indicates whether OSS rules can be used for cropping.
	//
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// Default image
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
	// Medium image
	//
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// Small image
	//
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s IndexControlPlayingListResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *IndexControlPlayingListResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *IndexControlPlayingListResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *IndexControlPlayingListResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *IndexControlPlayingListResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *IndexControlPlayingListResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetCanResize(v bool) *IndexControlPlayingListResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetImg(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetLarge(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetMediam(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetMedium(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) SetSmall(v string) *IndexControlPlayingListResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *IndexControlPlayingListResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
