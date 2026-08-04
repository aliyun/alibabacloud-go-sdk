// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SearchContentResponseBody
	GetCode() *int32
	SetMessage(v string) *SearchContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchContentResponseBody
	GetRequestId() *string
	SetResult(v []*SearchContentResponseBodyResult) *SearchContentResponseBody
	GetResult() []*SearchContentResponseBodyResult
}

type SearchContentResponseBody struct {
	// code encoding
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// message information
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// return information
	Result []*SearchContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s SearchContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchContentResponseBody) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SearchContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchContentResponseBody) GetResult() []*SearchContentResponseBodyResult {
	return s.Result
}

func (s *SearchContentResponseBody) SetCode(v int32) *SearchContentResponseBody {
	s.Code = &v
	return s
}

func (s *SearchContentResponseBody) SetMessage(v string) *SearchContentResponseBody {
	s.Message = &v
	return s
}

func (s *SearchContentResponseBody) SetRequestId(v string) *SearchContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchContentResponseBody) SetResult(v []*SearchContentResponseBodyResult) *SearchContentResponseBody {
	s.Result = v
	return s
}

func (s *SearchContentResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchContentResponseBodyResult struct {
	// The corresponding album ID
	//
	// example:
	//
	// 13123
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// Alias
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// Is audition available
	//
	// example:
	//
	// false
	Audition *bool `json:"Audition,omitempty" xml:"Audition,omitempty"`
	// Authors of the content
	Authors []*SearchContentResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	// Transform controlType based on the associated public category
	//
	// example:
	//
	// audio
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Is charged
	//
	// example:
	//
	// fasle
	Charge *bool `json:"Charge,omitempty" xml:"Charge,omitempty"`
	// The corresponding category ID
	//
	// example:
	//
	// 80012017
	CommCateId *int64 `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	// Album thumbnail image
	Cover *SearchContentResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// Content description
	//
	// example:
	//
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Duration information
	//
	// example:
	//
	// 180
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Popularity score
	//
	// example:
	//
	// 10
	HotScore *float64 `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	// Content ID
	//
	// example:
	//
	// 13597709
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of content, such as music, audio, radio, jokes, etc.
	//
	// example:
	//
	// ALBUM
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	// Lyric information
	//
	// example:
	//
	// http://a1231.lrc
	Lyric *string `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
	// Content sorting value
	//
	// example:
	//
	// 1
	OrderIndex *string `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// Source
	//
	// example:
	//
	// qingting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Music styles
	Styles []*string `json:"Styles,omitempty" xml:"Styles,omitempty" type:"Repeated"`
	// Title
	//
	// example:
	//
	// 超能狂少在都市
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Transform favoriteType based on the associated public category
	//
	// example:
	//
	// program
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the item is playable
	//
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s SearchContentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SearchContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResult) GetAlbumId() *string {
	return s.AlbumId
}

func (s *SearchContentResponseBodyResult) GetAlias() []*string {
	return s.Alias
}

func (s *SearchContentResponseBodyResult) GetAudition() *bool {
	return s.Audition
}

func (s *SearchContentResponseBodyResult) GetAuthors() []*SearchContentResponseBodyResultAuthors {
	return s.Authors
}

func (s *SearchContentResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *SearchContentResponseBodyResult) GetCharge() *bool {
	return s.Charge
}

func (s *SearchContentResponseBodyResult) GetCommCateId() *int64 {
	return s.CommCateId
}

func (s *SearchContentResponseBodyResult) GetCover() *SearchContentResponseBodyResultCover {
	return s.Cover
}

func (s *SearchContentResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *SearchContentResponseBodyResult) GetDuration() *int64 {
	return s.Duration
}

func (s *SearchContentResponseBodyResult) GetHotScore() *float64 {
	return s.HotScore
}

func (s *SearchContentResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *SearchContentResponseBodyResult) GetItemType() *string {
	return s.ItemType
}

func (s *SearchContentResponseBodyResult) GetLyric() *string {
	return s.Lyric
}

func (s *SearchContentResponseBodyResult) GetOrderIndex() *string {
	return s.OrderIndex
}

func (s *SearchContentResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *SearchContentResponseBodyResult) GetStyles() []*string {
	return s.Styles
}

func (s *SearchContentResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *SearchContentResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *SearchContentResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *SearchContentResponseBodyResult) SetAlbumId(v string) *SearchContentResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetAlias(v []*string) *SearchContentResponseBodyResult {
	s.Alias = v
	return s
}

func (s *SearchContentResponseBodyResult) SetAudition(v bool) *SearchContentResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetAuthors(v []*SearchContentResponseBodyResultAuthors) *SearchContentResponseBodyResult {
	s.Authors = v
	return s
}

func (s *SearchContentResponseBodyResult) SetCategory(v string) *SearchContentResponseBodyResult {
	s.Category = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetCharge(v bool) *SearchContentResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetCommCateId(v int64) *SearchContentResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetCover(v *SearchContentResponseBodyResultCover) *SearchContentResponseBodyResult {
	s.Cover = v
	return s
}

func (s *SearchContentResponseBodyResult) SetDescription(v string) *SearchContentResponseBodyResult {
	s.Description = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetDuration(v int64) *SearchContentResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetHotScore(v float64) *SearchContentResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetId(v int64) *SearchContentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetItemType(v string) *SearchContentResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetLyric(v string) *SearchContentResponseBodyResult {
	s.Lyric = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetOrderIndex(v string) *SearchContentResponseBodyResult {
	s.OrderIndex = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetSource(v string) *SearchContentResponseBodyResult {
	s.Source = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetStyles(v []*string) *SearchContentResponseBodyResult {
	s.Styles = v
	return s
}

func (s *SearchContentResponseBodyResult) SetTitle(v string) *SearchContentResponseBodyResult {
	s.Title = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetType(v string) *SearchContentResponseBodyResult {
	s.Type = &v
	return s
}

func (s *SearchContentResponseBodyResult) SetValid(v string) *SearchContentResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *SearchContentResponseBodyResult) Validate() error {
	if s.Authors != nil {
		for _, item := range s.Authors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchContentResponseBodyResultAuthors struct {
	// Author type
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	// Profile picture
	Cover *SearchContentResponseBodyResultAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// Author description
	//
	// example:
	//
	// 播音呆瓜小贼
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Gender
	//
	// example:
	//
	// MALE
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// Author primary key ID
	//
	// example:
	//
	// 13597709
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Is online
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// Third-party author ID
	//
	// example:
	//
	// 123123
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// Source
	//
	// example:
	//
	// qingting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Author title
	//
	// example:
	//
	// 播音呆瓜小贼
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchContentResponseBodyResultAuthors) String() string {
	return dara.Prettify(s)
}

func (s SearchContentResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResultAuthors) GetAuthorTypes() []*string {
	return s.AuthorTypes
}

func (s *SearchContentResponseBodyResultAuthors) GetCover() *SearchContentResponseBodyResultAuthorsCover {
	return s.Cover
}

func (s *SearchContentResponseBodyResultAuthors) GetDescription() *string {
	return s.Description
}

func (s *SearchContentResponseBodyResultAuthors) GetGender() *string {
	return s.Gender
}

func (s *SearchContentResponseBodyResultAuthors) GetId() *int64 {
	return s.Id
}

func (s *SearchContentResponseBodyResultAuthors) GetOnline() *bool {
	return s.Online
}

func (s *SearchContentResponseBodyResultAuthors) GetRawId() *string {
	return s.RawId
}

func (s *SearchContentResponseBodyResultAuthors) GetSource() *string {
	return s.Source
}

func (s *SearchContentResponseBodyResultAuthors) GetTitle() *string {
	return s.Title
}

func (s *SearchContentResponseBodyResultAuthors) SetAuthorTypes(v []*string) *SearchContentResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetCover(v *SearchContentResponseBodyResultAuthorsCover) *SearchContentResponseBodyResultAuthors {
	s.Cover = v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetDescription(v string) *SearchContentResponseBodyResultAuthors {
	s.Description = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetGender(v string) *SearchContentResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetId(v int64) *SearchContentResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetOnline(v bool) *SearchContentResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetRawId(v string) *SearchContentResponseBodyResultAuthors {
	s.RawId = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetSource(v string) *SearchContentResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) SetTitle(v string) *SearchContentResponseBodyResultAuthors {
	s.Title = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthors) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchContentResponseBodyResultAuthorsCover struct {
	// Indicates whether OSS rules can be used for cropping
	//
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// Default image
	//
	// example:
	//
	// https://a.jpg
	Img *string `json:"Img,omitempty" xml:"Img,omitempty"`
	// Large image
	//
	// example:
	//
	// https://a.jpg
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// Medium image
	//
	// example:
	//
	// https://a.jpg
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// Small image
	//
	// example:
	//
	// https://a.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s SearchContentResponseBodyResultAuthorsCover) String() string {
	return dara.Prettify(s)
}

func (s SearchContentResponseBodyResultAuthorsCover) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResultAuthorsCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *SearchContentResponseBodyResultAuthorsCover) GetImg() *string {
	return s.Img
}

func (s *SearchContentResponseBodyResultAuthorsCover) GetLarge() *string {
	return s.Large
}

func (s *SearchContentResponseBodyResultAuthorsCover) GetMedium() *string {
	return s.Medium
}

func (s *SearchContentResponseBodyResultAuthorsCover) GetSmall() *string {
	return s.Small
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetCanResize(v bool) *SearchContentResponseBodyResultAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetImg(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Img = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetLarge(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Large = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetMedium(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Medium = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) SetSmall(v string) *SearchContentResponseBodyResultAuthorsCover {
	s.Small = &v
	return s
}

func (s *SearchContentResponseBodyResultAuthorsCover) Validate() error {
	return dara.Validate(s)
}

type SearchContentResponseBodyResultCover struct {
	// Indicates whether OSS rules can be used to crop the image
	//
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// Default image
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Img *string `json:"Img,omitempty" xml:"Img,omitempty"`
	// Large image
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// Medium image (Deprecated)
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Mediam *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	// Medium image
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// Small image
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718285.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s SearchContentResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s SearchContentResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *SearchContentResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *SearchContentResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *SearchContentResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *SearchContentResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *SearchContentResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *SearchContentResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *SearchContentResponseBodyResultCover) SetCanResize(v bool) *SearchContentResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetImg(v string) *SearchContentResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetLarge(v string) *SearchContentResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetMediam(v string) *SearchContentResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetMedium(v string) *SearchContentResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) SetSmall(v string) *SearchContentResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *SearchContentResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
