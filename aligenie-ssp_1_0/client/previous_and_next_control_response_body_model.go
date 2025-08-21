// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviousAndNextControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PreviousAndNextControlResponseBody
	GetCode() *int32
	SetMessage(v string) *PreviousAndNextControlResponseBody
	GetMessage() *string
	SetRequestId(v string) *PreviousAndNextControlResponseBody
	GetRequestId() *string
	SetResult(v *PreviousAndNextControlResponseBodyResult) *PreviousAndNextControlResponseBody
	GetResult() *PreviousAndNextControlResponseBodyResult
	SetSuccess(v string) *PreviousAndNextControlResponseBody
	GetSuccess() *string
}

type PreviousAndNextControlResponseBody struct {
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
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *PreviousAndNextControlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PreviousAndNextControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlResponseBody) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PreviousAndNextControlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreviousAndNextControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviousAndNextControlResponseBody) GetResult() *PreviousAndNextControlResponseBodyResult {
	return s.Result
}

func (s *PreviousAndNextControlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PreviousAndNextControlResponseBody) SetCode(v int32) *PreviousAndNextControlResponseBody {
	s.Code = &v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetMessage(v string) *PreviousAndNextControlResponseBody {
	s.Message = &v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetRequestId(v string) *PreviousAndNextControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetResult(v *PreviousAndNextControlResponseBodyResult) *PreviousAndNextControlResponseBody {
	s.Result = v
	return s
}

func (s *PreviousAndNextControlResponseBody) SetSuccess(v string) *PreviousAndNextControlResponseBody {
	s.Success = &v
	return s
}

func (s *PreviousAndNextControlResponseBody) Validate() error {
	return dara.Validate(s)
}

type PreviousAndNextControlResponseBodyResult struct {
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
	Copyright *int32                                         `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover     *PreviousAndNextControlResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
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

func (s PreviousAndNextControlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponseBodyResult) GetAlbumName() *string {
	return s.AlbumName
}

func (s *PreviousAndNextControlResponseBodyResult) GetAlbumRawId() *string {
	return s.AlbumRawId
}

func (s *PreviousAndNextControlResponseBodyResult) GetAudioLength() *int32 {
	return s.AudioLength
}

func (s *PreviousAndNextControlResponseBodyResult) GetCopyright() *int32 {
	return s.Copyright
}

func (s *PreviousAndNextControlResponseBodyResult) GetCover() *PreviousAndNextControlResponseBodyResultCover {
	return s.Cover
}

func (s *PreviousAndNextControlResponseBodyResult) GetDefaultPlayOrder() *int32 {
	return s.DefaultPlayOrder
}

func (s *PreviousAndNextControlResponseBodyResult) GetItemUrl() *string {
	return s.ItemUrl
}

func (s *PreviousAndNextControlResponseBodyResult) GetLiked() *bool {
	return s.Liked
}

func (s *PreviousAndNextControlResponseBodyResult) GetLyricUrl() *string {
	return s.LyricUrl
}

func (s *PreviousAndNextControlResponseBodyResult) GetPlayMode() *string {
	return s.PlayMode
}

func (s *PreviousAndNextControlResponseBodyResult) GetPos() *int32 {
	return s.Pos
}

func (s *PreviousAndNextControlResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *PreviousAndNextControlResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *PreviousAndNextControlResponseBodyResult) GetSinger() *string {
	return s.Singer
}

func (s *PreviousAndNextControlResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *PreviousAndNextControlResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *PreviousAndNextControlResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *PreviousAndNextControlResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *PreviousAndNextControlResponseBodyResult) SetAlbumName(v string) *PreviousAndNextControlResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetAlbumRawId(v string) *PreviousAndNextControlResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetAudioLength(v int32) *PreviousAndNextControlResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetCopyright(v int32) *PreviousAndNextControlResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetCover(v *PreviousAndNextControlResponseBodyResultCover) *PreviousAndNextControlResponseBodyResult {
	s.Cover = v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetDefaultPlayOrder(v int32) *PreviousAndNextControlResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetItemUrl(v string) *PreviousAndNextControlResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetLiked(v bool) *PreviousAndNextControlResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetLyricUrl(v string) *PreviousAndNextControlResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetPlayMode(v string) *PreviousAndNextControlResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetPos(v int32) *PreviousAndNextControlResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetProgress(v int32) *PreviousAndNextControlResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetRawId(v string) *PreviousAndNextControlResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetSinger(v string) *PreviousAndNextControlResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetSource(v string) *PreviousAndNextControlResponseBodyResult {
	s.Source = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetTitle(v string) *PreviousAndNextControlResponseBodyResult {
	s.Title = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetType(v string) *PreviousAndNextControlResponseBodyResult {
	s.Type = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) SetValid(v string) *PreviousAndNextControlResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type PreviousAndNextControlResponseBodyResultCover struct {
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

func (s PreviousAndNextControlResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *PreviousAndNextControlResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *PreviousAndNextControlResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *PreviousAndNextControlResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *PreviousAndNextControlResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *PreviousAndNextControlResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetCanResize(v bool) *PreviousAndNextControlResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetImg(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetLarge(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetMediam(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetMedium(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) SetSmall(v string) *PreviousAndNextControlResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *PreviousAndNextControlResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
