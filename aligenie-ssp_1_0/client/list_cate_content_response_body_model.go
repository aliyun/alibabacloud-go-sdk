// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCateContentResponseBody
	GetCode() *int32
	SetMessage(v string) *ListCateContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCateContentResponseBody
	GetRequestId() *string
	SetResult(v *ListCateContentResponseBodyResult) *ListCateContentResponseBody
	GetResult() *ListCateContentResponseBodyResult
}

type ListCateContentResponseBody struct {
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
	// sucess
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return Result
	Result *ListCateContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListCateContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentResponseBody) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCateContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCateContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCateContentResponseBody) GetResult() *ListCateContentResponseBodyResult {
	return s.Result
}

func (s *ListCateContentResponseBody) SetCode(v int32) *ListCateContentResponseBody {
	s.Code = &v
	return s
}

func (s *ListCateContentResponseBody) SetMessage(v string) *ListCateContentResponseBody {
	s.Message = &v
	return s
}

func (s *ListCateContentResponseBody) SetRequestId(v string) *ListCateContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCateContentResponseBody) SetResult(v *ListCateContentResponseBodyResult) *ListCateContentResponseBody {
	s.Result = v
	return s
}

func (s *ListCateContentResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCateContentResponseBodyResult struct {
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// Data information
	OpenDataItemList []*ListCateContentResponseBodyResultOpenDataItemList `json:"OpenDataItemList,omitempty" xml:"OpenDataItemList,omitempty" type:"Repeated"`
	// Number of records per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 12002
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListCateContentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResult) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListCateContentResponseBodyResult) GetOpenDataItemList() []*ListCateContentResponseBodyResultOpenDataItemList {
	return s.OpenDataItemList
}

func (s *ListCateContentResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCateContentResponseBodyResult) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListCateContentResponseBodyResult) SetCurrentPageNum(v int32) *ListCateContentResponseBodyResult {
	s.CurrentPageNum = &v
	return s
}

func (s *ListCateContentResponseBodyResult) SetOpenDataItemList(v []*ListCateContentResponseBodyResultOpenDataItemList) *ListCateContentResponseBodyResult {
	s.OpenDataItemList = v
	return s
}

func (s *ListCateContentResponseBodyResult) SetPageSize(v int32) *ListCateContentResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListCateContentResponseBodyResult) SetTotalSize(v int64) *ListCateContentResponseBodyResult {
	s.TotalSize = &v
	return s
}

func (s *ListCateContentResponseBodyResult) Validate() error {
	if s.OpenDataItemList != nil {
		for _, item := range s.OpenDataItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCateContentResponseBodyResultOpenDataItemList struct {
	// Alias
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// Is audition available
	//
	// example:
	//
	// true
	Audition *bool `json:"Audition,omitempty" xml:"Audition,omitempty"`
	// author information
	Authors []*ListCateContentResponseBodyResultOpenDataItemListAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	// Transform controlType based on the assigned public category
	//
	// example:
	//
	// audio
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Whether it is charged
	//
	// example:
	//
	// false
	Charge *bool `json:"Charge,omitempty" xml:"Charge,omitempty"`
	// Category ID
	//
	// example:
	//
	// 80012017
	CommCateId *string `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	// Album thumbnail image
	Cover *ListCateContentResponseBodyResultOpenDataItemListCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// Content description
	//
	// example:
	//
	// 作者：月不醉 简介：啥，假冒你男友？退役兵王回归都市后做了一名小保安，机缘巧合下却被冷艳女总拉去客串男友，从此之后，他与冷艳老总开始纠缠不清···拳打恶一代，脚踩高富帅，哥哥我专治各种不服。面对各类的阴谋阳谋，面对各种黑恶势力，且看哥如何称霸都市，踏上巅峰！
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Popularity score
	//
	// example:
	//
	// 0
	HotScore *float64 `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	// Type of content, such as music, audio, radio, jokes, etc.
	//
	// example:
	//
	// ALBUM
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	// Third-party ID of the album
	//
	// example:
	//
	// 206775
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// Source
	//
	// example:
	//
	// qingting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Title
	//
	// example:
	//
	// 超能狂少在都市
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Transform favoriteType based on the associated public category.
	//
	// example:
	//
	// program
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the content is playable.
	//
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 26152778
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListCateContentResponseBodyResultOpenDataItemList) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemList) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetAlias() []*string {
	return s.Alias
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetAudition() *bool {
	return s.Audition
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetAuthors() []*ListCateContentResponseBodyResultOpenDataItemListAuthors {
	return s.Authors
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetCategory() *string {
	return s.Category
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetCharge() *bool {
	return s.Charge
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetCommCateId() *string {
	return s.CommCateId
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetCover() *ListCateContentResponseBodyResultOpenDataItemListCover {
	return s.Cover
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetDescription() *string {
	return s.Description
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetHotScore() *float64 {
	return s.HotScore
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetItemType() *string {
	return s.ItemType
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetRawId() *string {
	return s.RawId
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetSource() *string {
	return s.Source
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetTitle() *string {
	return s.Title
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetType() *string {
	return s.Type
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetValid() *string {
	return s.Valid
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) GetId() *int64 {
	return s.Id
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetAlias(v []*string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Alias = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetAudition(v bool) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Audition = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetAuthors(v []*ListCateContentResponseBodyResultOpenDataItemListAuthors) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Authors = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCategory(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Category = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCharge(v bool) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Charge = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCommCateId(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.CommCateId = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetCover(v *ListCateContentResponseBodyResultOpenDataItemListCover) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Cover = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetDescription(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Description = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetHotScore(v float64) *ListCateContentResponseBodyResultOpenDataItemList {
	s.HotScore = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetItemType(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.ItemType = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetRawId(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.RawId = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetSource(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Source = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetTitle(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Title = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetType(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Type = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetValid(v string) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Valid = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) SetId(v int64) *ListCateContentResponseBodyResultOpenDataItemList {
	s.Id = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemList) Validate() error {
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

type ListCateContentResponseBodyResultOpenDataItemListAuthors struct {
	// Author type
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	// profile picture
	Cover *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
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
	// Primary key ID of the author information
	//
	// example:
	//
	// 13597709
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Whether it is online
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// third-party author ID
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

func (s ListCateContentResponseBodyResultOpenDataItemListAuthors) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemListAuthors) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetAuthorTypes() []*string {
	return s.AuthorTypes
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetCover() *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	return s.Cover
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetDescription() *string {
	return s.Description
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetGender() *string {
	return s.Gender
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetId() *int64 {
	return s.Id
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetOnline() *bool {
	return s.Online
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetRawId() *string {
	return s.RawId
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetSource() *string {
	return s.Source
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) GetTitle() *string {
	return s.Title
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetAuthorTypes(v []*string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetCover(v *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Cover = v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetDescription(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Description = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetGender(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Gender = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetId(v int64) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Id = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetOnline(v bool) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Online = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetRawId(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.RawId = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetSource(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Source = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) SetTitle(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthors {
	s.Title = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthors) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCateContentResponseBodyResultOpenDataItemListAuthorsCover struct {
	// Indicates whether the image can be cropped using OSS rules.
	//
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// default image
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
	// Medium image (Deprecated)
	//
	// example:
	//
	// https://a.jpg
	Mediam *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	// medium image
	//
	// example:
	//
	// https://a.jpg
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// small image
	//
	// example:
	//
	// https://a.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GetImg() *string {
	return s.Img
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GetLarge() *string {
	return s.Large
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GetMediam() *string {
	return s.Mediam
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GetMedium() *string {
	return s.Medium
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) GetSmall() *string {
	return s.Small
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetCanResize(v bool) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetImg(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Img = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetLarge(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Large = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetMediam(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Mediam = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetMedium(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Medium = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) SetSmall(v string) *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover {
	s.Small = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListAuthorsCover) Validate() error {
	return dara.Validate(s)
}

type ListCateContentResponseBodyResultOpenDataItemListCover struct {
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
	// http://pic.qtfm.cn/2017/0207/2017020718275.jpg
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// Medium image (Deprecated)
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020718275.jpg
	Mediam *string `json:"Mediam,omitempty" xml:"Mediam,omitempty"`
	// Medium image
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/20170207175.jpg
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// Small image
	//
	// example:
	//
	// http://pic.qtfm.cn/2017/0207/2017020675.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
	// Indicates whether cropping can be performed using OSS rules.
	//
	// example:
	//
	// false
	CanResize *bool `json:"canResize,omitempty" xml:"canResize,omitempty"`
}

func (s ListCateContentResponseBodyResultOpenDataItemListCover) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentResponseBodyResultOpenDataItemListCover) GoString() string {
	return s.String()
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) GetImg() *string {
	return s.Img
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) GetLarge() *string {
	return s.Large
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) GetMediam() *string {
	return s.Mediam
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) GetMedium() *string {
	return s.Medium
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) GetSmall() *string {
	return s.Small
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetImg(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Img = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetLarge(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Large = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetMediam(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Mediam = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetMedium(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Medium = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetSmall(v string) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.Small = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) SetCanResize(v bool) *ListCateContentResponseBodyResultOpenDataItemListCover {
	s.CanResize = &v
	return s
}

func (s *ListCateContentResponseBodyResultOpenDataItemListCover) Validate() error {
	return dara.Validate(s)
}
