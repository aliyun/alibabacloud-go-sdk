// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListOAuth2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePlayingListOAuth2ResponseBody
	GetCode() *int32
	SetMessage(v string) *CreatePlayingListOAuth2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePlayingListOAuth2ResponseBody
	GetRequestId() *string
	SetResult(v *CreatePlayingListOAuth2ResponseBodyResult) *CreatePlayingListOAuth2ResponseBody
	GetResult() *CreatePlayingListOAuth2ResponseBodyResult
	SetSuccess(v string) *CreatePlayingListOAuth2ResponseBody
	GetSuccess() *string
}

type CreatePlayingListOAuth2ResponseBody struct {
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
	Result    *CreatePlayingListOAuth2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePlayingListOAuth2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2ResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePlayingListOAuth2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePlayingListOAuth2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePlayingListOAuth2ResponseBody) GetResult() *CreatePlayingListOAuth2ResponseBodyResult {
	return s.Result
}

func (s *CreatePlayingListOAuth2ResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreatePlayingListOAuth2ResponseBody) SetCode(v int32) *CreatePlayingListOAuth2ResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBody) SetMessage(v string) *CreatePlayingListOAuth2ResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBody) SetRequestId(v string) *CreatePlayingListOAuth2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBody) SetResult(v *CreatePlayingListOAuth2ResponseBodyResult) *CreatePlayingListOAuth2ResponseBody {
	s.Result = v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBody) SetSuccess(v string) *CreatePlayingListOAuth2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListOAuth2ResponseBodyResult struct {
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
	Copyright *int32                                          `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	Cover     *CreatePlayingListOAuth2ResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
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
	// 1
	Pos *int32 `json:"Pos,omitempty" xml:"Pos,omitempty"`
	// example:
	//
	// 96.0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 123123
	RawId  *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	Singer *string `json:"Singer,omitempty" xml:"Singer,omitempty"`
	// example:
	//
	// qignting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// program
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s CreatePlayingListOAuth2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetAlbumName() *string {
	return s.AlbumName
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetAlbumRawId() *string {
	return s.AlbumRawId
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetAudioLength() *int32 {
	return s.AudioLength
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetCopyright() *int32 {
	return s.Copyright
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetCover() *CreatePlayingListOAuth2ResponseBodyResultCover {
	return s.Cover
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetDefaultPlayOrder() *int32 {
	return s.DefaultPlayOrder
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetItemUrl() *string {
	return s.ItemUrl
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetLiked() *bool {
	return s.Liked
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetLyricUrl() *string {
	return s.LyricUrl
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetPlayMode() *string {
	return s.PlayMode
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetPos() *int32 {
	return s.Pos
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetSinger() *string {
	return s.Singer
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetAlbumName(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.AlbumName = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetAlbumRawId(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetAudioLength(v int32) *CreatePlayingListOAuth2ResponseBodyResult {
	s.AudioLength = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetCopyright(v int32) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Copyright = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetCover(v *CreatePlayingListOAuth2ResponseBodyResultCover) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Cover = v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetDefaultPlayOrder(v int32) *CreatePlayingListOAuth2ResponseBodyResult {
	s.DefaultPlayOrder = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetItemUrl(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetLiked(v bool) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Liked = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetLyricUrl(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.LyricUrl = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetPlayMode(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetPos(v int32) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Pos = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetProgress(v int32) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetRawId(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetSinger(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Singer = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetSource(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Source = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetTitle(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetType(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) SetValid(v string) *CreatePlayingListOAuth2ResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListOAuth2ResponseBodyResultCover struct {
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Img *string `json:"Img,omitempty" xml:"Img,omitempty"`
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Mediam *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s CreatePlayingListOAuth2ResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2ResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) SetCanResize(v bool) *CreatePlayingListOAuth2ResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) SetImg(v string) *CreatePlayingListOAuth2ResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) SetLarge(v string) *CreatePlayingListOAuth2ResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) SetMediam(v string) *CreatePlayingListOAuth2ResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) SetMedium(v string) *CreatePlayingListOAuth2ResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) SetSmall(v string) *CreatePlayingListOAuth2ResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *CreatePlayingListOAuth2ResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
