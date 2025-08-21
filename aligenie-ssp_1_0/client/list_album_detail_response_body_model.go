// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAlbumDetailResponseBody
	GetCode() *int32
	SetMessage(v string) *ListAlbumDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlbumDetailResponseBody
	GetRequestId() *string
	SetResult(v *ListAlbumDetailResponseBodyResult) *ListAlbumDetailResponseBody
	GetResult() *ListAlbumDetailResponseBodyResult
}

type ListAlbumDetailResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListAlbumDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAlbumDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAlbumDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlbumDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlbumDetailResponseBody) GetResult() *ListAlbumDetailResponseBodyResult {
	return s.Result
}

func (s *ListAlbumDetailResponseBody) SetCode(v int32) *ListAlbumDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlbumDetailResponseBody) SetMessage(v string) *ListAlbumDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlbumDetailResponseBody) SetRequestId(v string) *ListAlbumDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlbumDetailResponseBody) SetResult(v *ListAlbumDetailResponseBodyResult) *ListAlbumDetailResponseBody {
	s.Result = v
	return s
}

func (s *ListAlbumDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlbumDetailResponseBodyResult struct {
	// example:
	//
	// 1
	CurrentPageNum   *int64                                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	OpenDataItemList []*ListAlbumDetailResponseBodyResultOpenDataItemList `json:"OpenDataItemList,omitempty" xml:"OpenDataItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 21421
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAlbumDetailResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResult) GetCurrentPageNum() *int64 {
	return s.CurrentPageNum
}

func (s *ListAlbumDetailResponseBodyResult) GetOpenDataItemList() []*ListAlbumDetailResponseBodyResultOpenDataItemList {
	return s.OpenDataItemList
}

func (s *ListAlbumDetailResponseBodyResult) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlbumDetailResponseBodyResult) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListAlbumDetailResponseBodyResult) SetCurrentPageNum(v int64) *ListAlbumDetailResponseBodyResult {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResult) SetOpenDataItemList(v []*ListAlbumDetailResponseBodyResultOpenDataItemList) *ListAlbumDetailResponseBodyResult {
	s.OpenDataItemList = v
	return s
}

func (s *ListAlbumDetailResponseBodyResult) SetPageSize(v int64) *ListAlbumDetailResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResult) SetTotalSize(v int64) *ListAlbumDetailResponseBodyResult {
	s.TotalSize = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListAlbumDetailResponseBodyResultOpenDataItemList struct {
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Audition *bool                                                       `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors  []*ListAlbumDetailResponseBodyResultOpenDataItemListAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
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
	CommCateId  *int64                                                  `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *ListAlbumDetailResponseBodyResultOpenDataItemListCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 180
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 10
	HotScore *float64 `json:"HotScore,omitempty" xml:"HotScore,omitempty"`
	// example:
	//
	// 1231231
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ALBUM
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	// example:
	//
	// 1
	OrderIndex *int64 `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// example:
	//
	// 12323423
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

func (s ListAlbumDetailResponseBodyResultOpenDataItemList) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemList) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetAlias() []*string {
	return s.Alias
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetAudition() *bool {
	return s.Audition
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetAuthors() []*ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	return s.Authors
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetCategory() *string {
	return s.Category
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetCharge() *bool {
	return s.Charge
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetCommCateId() *int64 {
	return s.CommCateId
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetCover() *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	return s.Cover
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetDescription() *string {
	return s.Description
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetHotScore() *float64 {
	return s.HotScore
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetId() *int64 {
	return s.Id
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetItemType() *string {
	return s.ItemType
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetRawId() *string {
	return s.RawId
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetSource() *string {
	return s.Source
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetStyles() []*string {
	return s.Styles
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetTitle() *string {
	return s.Title
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetType() *string {
	return s.Type
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) GetValid() *string {
	return s.Valid
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetAlias(v []*string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Alias = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetAudition(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Audition = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetAuthors(v []*ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Authors = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCategory(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Category = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCharge(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Charge = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCommCateId(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.CommCateId = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetCover(v *ListAlbumDetailResponseBodyResultOpenDataItemListCover) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Cover = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetDescription(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Description = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetDuration(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Duration = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetHotScore(v float64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.HotScore = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetId(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Id = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetItemType(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.ItemType = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetOrderIndex(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.OrderIndex = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetRawId(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.RawId = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetSource(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Source = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetStyles(v []*string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Styles = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetTitle(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Title = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetType(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Type = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) SetValid(v string) *ListAlbumDetailResponseBodyResultOpenDataItemList {
	s.Valid = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemList) Validate() error {
	return dara.Validate(s)
}

type ListAlbumDetailResponseBodyResultOpenDataItemListAuthors struct {
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	// example:
	//
	// MALE
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// 12314
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

func (s ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GetAuthorTypes() []*string {
	return s.AuthorTypes
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GetGender() *string {
	return s.Gender
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GetId() *int64 {
	return s.Id
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GetOnline() *bool {
	return s.Online
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GetSource() *string {
	return s.Source
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) GetTitle() *string {
	return s.Title
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetAuthorTypes(v []*string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetGender(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Gender = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetId(v int64) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Id = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetOnline(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Online = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetSource(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Source = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) SetTitle(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors {
	s.Title = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListAuthors) Validate() error {
	return dara.Validate(s)
}

type ListAlbumDetailResponseBodyResultOpenDataItemListCover struct {
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

func (s ListAlbumDetailResponseBodyResultOpenDataItemListCover) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailResponseBodyResultOpenDataItemListCover) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) GetImg() *string {
	return s.Img
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) GetLarge() *string {
	return s.Large
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) GetMedium() *string {
	return s.Medium
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) GetSmall() *string {
	return s.Small
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetCanResize(v bool) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.CanResize = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetImg(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Img = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetLarge(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Large = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetMedium(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Medium = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) SetSmall(v string) *ListAlbumDetailResponseBodyResultOpenDataItemListCover {
	s.Small = &v
	return s
}

func (s *ListAlbumDetailResponseBodyResultOpenDataItemListCover) Validate() error {
	return dara.Validate(s)
}
