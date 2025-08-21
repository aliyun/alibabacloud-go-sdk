// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetCurrentPlayingListResponseBody
	GetCode() *int32
	SetMessage(v string) *GetCurrentPlayingListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCurrentPlayingListResponseBody
	GetRequestId() *string
	SetResult(v []*GetCurrentPlayingListResponseBodyResult) *GetCurrentPlayingListResponseBody
	GetResult() []*GetCurrentPlayingListResponseBodyResult
	SetSuccess(v string) *GetCurrentPlayingListResponseBody
	GetSuccess() *string
}

type GetCurrentPlayingListResponseBody struct {
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
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetCurrentPlayingListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCurrentPlayingListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetCurrentPlayingListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCurrentPlayingListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCurrentPlayingListResponseBody) GetResult() []*GetCurrentPlayingListResponseBodyResult {
	return s.Result
}

func (s *GetCurrentPlayingListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetCurrentPlayingListResponseBody) SetCode(v int32) *GetCurrentPlayingListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetMessage(v string) *GetCurrentPlayingListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetRequestId(v string) *GetCurrentPlayingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetResult(v []*GetCurrentPlayingListResponseBodyResult) *GetCurrentPlayingListResponseBody {
	s.Result = v
	return s
}

func (s *GetCurrentPlayingListResponseBody) SetSuccess(v string) *GetCurrentPlayingListResponseBody {
	s.Success = &v
	return s
}

func (s *GetCurrentPlayingListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingListResponseBodyResult struct {
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
	Cover     *GetCurrentPlayingListResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
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
	Liked *bool `json:"Liked,omitempty" xml:"Liked,omitempty"`
	// example:
	//
	// https://aicontent.alibabausercontent.com/lyric/thirdsource/6f4c8408073db134b0d097c122b5a1a1.lrc
	LyricUrl *string `json:"LyricUrl,omitempty" xml:"LyricUrl,omitempty"`
	// example:
	//
	// Normal
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

func (s GetCurrentPlayingListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponseBodyResult) GetAlbumName() *string {
	return s.AlbumName
}

func (s *GetCurrentPlayingListResponseBodyResult) GetAlbumRawId() *string {
	return s.AlbumRawId
}

func (s *GetCurrentPlayingListResponseBodyResult) GetAudioLength() *int32 {
	return s.AudioLength
}

func (s *GetCurrentPlayingListResponseBodyResult) GetCopyright() *int32 {
	return s.Copyright
}

func (s *GetCurrentPlayingListResponseBodyResult) GetCover() *GetCurrentPlayingListResponseBodyResultCover {
	return s.Cover
}

func (s *GetCurrentPlayingListResponseBodyResult) GetDefaultPlayOrder() *int32 {
	return s.DefaultPlayOrder
}

func (s *GetCurrentPlayingListResponseBodyResult) GetItemUrl() *string {
	return s.ItemUrl
}

func (s *GetCurrentPlayingListResponseBodyResult) GetLiked() *bool {
	return s.Liked
}

func (s *GetCurrentPlayingListResponseBodyResult) GetLyricUrl() *string {
	return s.LyricUrl
}

func (s *GetCurrentPlayingListResponseBodyResult) GetPlayMode() *string {
	return s.PlayMode
}

func (s *GetCurrentPlayingListResponseBodyResult) GetPos() *int32 {
	return s.Pos
}

func (s *GetCurrentPlayingListResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *GetCurrentPlayingListResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *GetCurrentPlayingListResponseBodyResult) GetSinger() *string {
	return s.Singer
}

func (s *GetCurrentPlayingListResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *GetCurrentPlayingListResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetCurrentPlayingListResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetCurrentPlayingListResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *GetCurrentPlayingListResponseBodyResult) SetAlbumName(v string) *GetCurrentPlayingListResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetAlbumRawId(v string) *GetCurrentPlayingListResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetAudioLength(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetCopyright(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetCover(v *GetCurrentPlayingListResponseBodyResultCover) *GetCurrentPlayingListResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetDefaultPlayOrder(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetItemUrl(v string) *GetCurrentPlayingListResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetLiked(v bool) *GetCurrentPlayingListResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetLyricUrl(v string) *GetCurrentPlayingListResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetPlayMode(v string) *GetCurrentPlayingListResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetPos(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetProgress(v int32) *GetCurrentPlayingListResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetRawId(v string) *GetCurrentPlayingListResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetSinger(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetSource(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetTitle(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetType(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) SetValid(v string) *GetCurrentPlayingListResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingListResponseBodyResultCover struct {
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

func (s GetCurrentPlayingListResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *GetCurrentPlayingListResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *GetCurrentPlayingListResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *GetCurrentPlayingListResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *GetCurrentPlayingListResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *GetCurrentPlayingListResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetCanResize(v bool) *GetCurrentPlayingListResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetImg(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetLarge(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetMediam(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetMedium(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) SetSmall(v string) *GetCurrentPlayingListResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *GetCurrentPlayingListResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
