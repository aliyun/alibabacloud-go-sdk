// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetContentResponseBody
	GetCode() *int32
	SetMessage(v string) *GetContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContentResponseBody
	GetRequestId() *string
	SetResult(v *GetContentResponseBodyResult) *GetContentResponseBody
	GetResult() *GetContentResponseBodyResult
}

type GetContentResponseBody struct {
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
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetContentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContentResponseBody) GetResult() *GetContentResponseBodyResult {
	return s.Result
}

func (s *GetContentResponseBody) SetCode(v int32) *GetContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetContentResponseBody) SetMessage(v string) *GetContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetContentResponseBody) SetRequestId(v string) *GetContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContentResponseBody) SetResult(v *GetContentResponseBodyResult) *GetContentResponseBody {
	s.Result = v
	return s
}

func (s *GetContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetContentResponseBodyResult struct {
	// example:
	//
	// 1231
	AlbumId *string   `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	Alias   []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Audition *bool                                  `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors  []*GetContentResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	// example:
	//
	// audio
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// false
	Charge *bool `json:"Charge,omitempty" xml:"Charge,omitempty"`
	// example:
	//
	// 80012017
	CommCateId  *int64                             `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *GetContentResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 180
	Duration *int64   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HotScore *float64 `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	// example:
	//
	// 13597709
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ALBUM
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	// example:
	//
	// http://1231.lrc
	Lyric *string `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
	// example:
	//
	// 1231231
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// example:
	//
	// qingting
	Source *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Styles []*string `json:"Styles,omitempty" xml:"Styles,omitempty" type:"Repeated"`
	Title  *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// program
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetContentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetContentResponseBodyResult) GetAlbumId() *string {
	return s.AlbumId
}

func (s *GetContentResponseBodyResult) GetAlias() []*string {
	return s.Alias
}

func (s *GetContentResponseBodyResult) GetAudition() *bool {
	return s.Audition
}

func (s *GetContentResponseBodyResult) GetAuthors() []*GetContentResponseBodyResultAuthors {
	return s.Authors
}

func (s *GetContentResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *GetContentResponseBodyResult) GetCharge() *bool {
	return s.Charge
}

func (s *GetContentResponseBodyResult) GetCommCateId() *int64 {
	return s.CommCateId
}

func (s *GetContentResponseBodyResult) GetCover() *GetContentResponseBodyResultCover {
	return s.Cover
}

func (s *GetContentResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetContentResponseBodyResult) GetDuration() *int64 {
	return s.Duration
}

func (s *GetContentResponseBodyResult) GetHotScore() *float64 {
	return s.HotScore
}

func (s *GetContentResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetContentResponseBodyResult) GetItemType() *string {
	return s.ItemType
}

func (s *GetContentResponseBodyResult) GetLyric() *string {
	return s.Lyric
}

func (s *GetContentResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *GetContentResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *GetContentResponseBodyResult) GetStyles() []*string {
	return s.Styles
}

func (s *GetContentResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetContentResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetContentResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *GetContentResponseBodyResult) SetAlbumId(v string) *GetContentResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *GetContentResponseBodyResult) SetAlias(v []*string) *GetContentResponseBodyResult {
	s.Alias = v
	return s
}

func (s *GetContentResponseBodyResult) SetAudition(v bool) *GetContentResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *GetContentResponseBodyResult) SetAuthors(v []*GetContentResponseBodyResultAuthors) *GetContentResponseBodyResult {
	s.Authors = v
	return s
}

func (s *GetContentResponseBodyResult) SetCategory(v string) *GetContentResponseBodyResult {
	s.Category = &v
	return s
}

func (s *GetContentResponseBodyResult) SetCharge(v bool) *GetContentResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *GetContentResponseBodyResult) SetCommCateId(v int64) *GetContentResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *GetContentResponseBodyResult) SetCover(v *GetContentResponseBodyResultCover) *GetContentResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetContentResponseBodyResult) SetDescription(v string) *GetContentResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetContentResponseBodyResult) SetDuration(v int64) *GetContentResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *GetContentResponseBodyResult) SetHotScore(v float64) *GetContentResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *GetContentResponseBodyResult) SetId(v int64) *GetContentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetContentResponseBodyResult) SetItemType(v string) *GetContentResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *GetContentResponseBodyResult) SetLyric(v string) *GetContentResponseBodyResult {
	s.Lyric = &v
	return s
}

func (s *GetContentResponseBodyResult) SetRawId(v string) *GetContentResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *GetContentResponseBodyResult) SetSource(v string) *GetContentResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetContentResponseBodyResult) SetStyles(v []*string) *GetContentResponseBodyResult {
	s.Styles = v
	return s
}

func (s *GetContentResponseBodyResult) SetTitle(v string) *GetContentResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetContentResponseBodyResult) SetType(v string) *GetContentResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetContentResponseBodyResult) SetValid(v string) *GetContentResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *GetContentResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetContentResponseBodyResultAuthors struct {
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	// example:
	//
	// MALE
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// 123123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// qingting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetContentResponseBodyResultAuthors) String() string {
	return dara.Prettify(s)
}

func (s GetContentResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *GetContentResponseBodyResultAuthors) GetAuthorTypes() []*string {
	return s.AuthorTypes
}

func (s *GetContentResponseBodyResultAuthors) GetGender() *string {
	return s.Gender
}

func (s *GetContentResponseBodyResultAuthors) GetId() *int64 {
	return s.Id
}

func (s *GetContentResponseBodyResultAuthors) GetOnline() *bool {
	return s.Online
}

func (s *GetContentResponseBodyResultAuthors) GetSource() *string {
	return s.Source
}

func (s *GetContentResponseBodyResultAuthors) GetTitle() *string {
	return s.Title
}

func (s *GetContentResponseBodyResultAuthors) SetAuthorTypes(v []*string) *GetContentResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetGender(v string) *GetContentResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetId(v int64) *GetContentResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetOnline(v bool) *GetContentResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetSource(v string) *GetContentResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) SetTitle(v string) *GetContentResponseBodyResultAuthors {
	s.Title = &v
	return s
}

func (s *GetContentResponseBodyResultAuthors) Validate() error {
	return dara.Validate(s)
}

type GetContentResponseBodyResultCover struct {
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
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s GetContentResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s GetContentResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetContentResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *GetContentResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *GetContentResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *GetContentResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *GetContentResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *GetContentResponseBodyResultCover) SetCanResize(v bool) *GetContentResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetImg(v string) *GetContentResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetLarge(v string) *GetContentResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetMedium(v string) *GetContentResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetContentResponseBodyResultCover) SetSmall(v string) *GetContentResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *GetContentResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
