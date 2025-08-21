// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetCurrentPlayingItemResponseBody
	GetCode() *int32
	SetMessage(v string) *GetCurrentPlayingItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCurrentPlayingItemResponseBody
	GetRequestId() *string
	SetResult(v *GetCurrentPlayingItemResponseBodyResult) *GetCurrentPlayingItemResponseBody
	GetResult() *GetCurrentPlayingItemResponseBodyResult
	SetSuccess(v string) *GetCurrentPlayingItemResponseBody
	GetSuccess() *string
}

type GetCurrentPlayingItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10002398812
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetCurrentPlayingItemResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCurrentPlayingItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetCurrentPlayingItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCurrentPlayingItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCurrentPlayingItemResponseBody) GetResult() *GetCurrentPlayingItemResponseBodyResult {
	return s.Result
}

func (s *GetCurrentPlayingItemResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetCurrentPlayingItemResponseBody) SetCode(v int32) *GetCurrentPlayingItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetMessage(v string) *GetCurrentPlayingItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetRequestId(v string) *GetCurrentPlayingItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetResult(v *GetCurrentPlayingItemResponseBodyResult) *GetCurrentPlayingItemResponseBody {
	s.Result = v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) SetSuccess(v string) *GetCurrentPlayingItemResponseBody {
	s.Success = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingItemResponseBodyResult struct {
	AlbumName *string `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	// example:
	//
	// 260744
	AlbumRawId *string `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	// example:
	//
	// 190
	AudioLength *int32 `json:"AudioLength,omitempty" xml:"AudioLength,omitempty"`
	// example:
	//
	// 0
	Copyright *int32                                        `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover     *GetCurrentPlayingItemResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// example:
	//
	// 1
	DefaultPlayOrder *int32 `json:"DefaultPlayOrder,omitempty" xml:"DefaultPlayOrder,omitempty"`
	// example:
	//
	// https://openaudio.cos.tx.xmcdn.com/storages/587f-audiofreehighqps/15/CE/GKwRIJIGnb11ABc6SwF59DNb.mp3
	ItemUrl *string `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	// example:
	//
	// false
	Liked *string `json:"Liked,omitempty" xml:"Liked,omitempty"`
	// example:
	//
	// https://aicontent.alibabausercontent.com/lyric/thirdsource/6f4c8408073db134b0d097c122b5a1a1.lrc
	LyricUrl *string `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	// example:
	//
	// Repeat
	PlayMode *string `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	// example:
	//
	// 0
	Pos *int32 `json:"Pos,omitempty" xml:"Pos,omitempty"`
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 550144364
	RawId  *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singer *string `json:"Singer,omitempty" xml:"Singer,omitempty"`
	// example:
	//
	// ximalayav2
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// story
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetCurrentPlayingItemResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetAlbumName() *string {
	return s.AlbumName
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetAlbumRawId() *string {
	return s.AlbumRawId
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetAudioLength() *int32 {
	return s.AudioLength
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetCopyright() *int32 {
	return s.Copyright
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetCover() *GetCurrentPlayingItemResponseBodyResultCover {
	return s.Cover
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetDefaultPlayOrder() *int32 {
	return s.DefaultPlayOrder
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetItemUrl() *string {
	return s.ItemUrl
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetLiked() *string {
	return s.Liked
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetLyricUrl() *string {
	return s.LyricUrl
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetPlayMode() *string {
	return s.PlayMode
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetPos() *int32 {
	return s.Pos
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetSinger() *string {
	return s.Singer
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetCurrentPlayingItemResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetAlbumName(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetAlbumRawId(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetAudioLength(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetCopyright(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetCover(v *GetCurrentPlayingItemResponseBodyResultCover) *GetCurrentPlayingItemResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetDefaultPlayOrder(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetItemUrl(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetLiked(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetLyricUrl(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetPlayMode(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetPos(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetProgress(v int32) *GetCurrentPlayingItemResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetRawId(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetSinger(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetSource(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetTitle(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetType(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) SetValid(v string) *GetCurrentPlayingItemResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingItemResponseBodyResultCover struct {
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Img *string `json:"Img,omitempty" xml:"Img,omitempty"`
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Mediam *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// example:
	//
	// http://imgopen.xmcdn.com/group58/M06/08/9B/wKgLglzTyTjiOy0oAAcOTv16ohg815.jpg!op_type=3&columns=640&rows=640
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s GetCurrentPlayingItemResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetCanResize(v bool) *GetCurrentPlayingItemResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetImg(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetLarge(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetMediam(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetMedium(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) SetSmall(v string) *GetCurrentPlayingItemResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *GetCurrentPlayingItemResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
