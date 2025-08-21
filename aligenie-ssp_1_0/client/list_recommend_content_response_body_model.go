// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListRecommendContentResponseBody
	GetCode() *int32
	SetMessage(v string) *ListRecommendContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRecommendContentResponseBody
	GetRequestId() *string
	SetResult(v []*ListRecommendContentResponseBodyResult) *ListRecommendContentResponseBody
	GetResult() []*ListRecommendContentResponseBodyResult
}

type ListRecommendContentResponseBody struct {
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
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListRecommendContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRecommendContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListRecommendContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRecommendContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecommendContentResponseBody) GetResult() []*ListRecommendContentResponseBodyResult {
	return s.Result
}

func (s *ListRecommendContentResponseBody) SetCode(v int32) *ListRecommendContentResponseBody {
	s.Code = &v
	return s
}

func (s *ListRecommendContentResponseBody) SetMessage(v string) *ListRecommendContentResponseBody {
	s.Message = &v
	return s
}

func (s *ListRecommendContentResponseBody) SetRequestId(v string) *ListRecommendContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecommendContentResponseBody) SetResult(v []*ListRecommendContentResponseBodyResult) *ListRecommendContentResponseBody {
	s.Result = v
	return s
}

func (s *ListRecommendContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentResponseBodyResult struct {
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Audition *bool                                            `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors  []*ListRecommendContentResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
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
	CommCateId  *int64                                       `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *ListRecommendContentResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10
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
	// 123123
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// example:
	//
	// qingting
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

func (s ListRecommendContentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResult) GetAlias() []*string {
	return s.Alias
}

func (s *ListRecommendContentResponseBodyResult) GetAudition() *bool {
	return s.Audition
}

func (s *ListRecommendContentResponseBodyResult) GetAuthors() []*ListRecommendContentResponseBodyResultAuthors {
	return s.Authors
}

func (s *ListRecommendContentResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *ListRecommendContentResponseBodyResult) GetCharge() *bool {
	return s.Charge
}

func (s *ListRecommendContentResponseBodyResult) GetCommCateId() *int64 {
	return s.CommCateId
}

func (s *ListRecommendContentResponseBodyResult) GetCover() *ListRecommendContentResponseBodyResultCover {
	return s.Cover
}

func (s *ListRecommendContentResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListRecommendContentResponseBodyResult) GetHotScore() *float64 {
	return s.HotScore
}

func (s *ListRecommendContentResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListRecommendContentResponseBodyResult) GetItemType() *string {
	return s.ItemType
}

func (s *ListRecommendContentResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *ListRecommendContentResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListRecommendContentResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *ListRecommendContentResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListRecommendContentResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *ListRecommendContentResponseBodyResult) SetAlias(v []*string) *ListRecommendContentResponseBodyResult {
	s.Alias = v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetAudition(v bool) *ListRecommendContentResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetAuthors(v []*ListRecommendContentResponseBodyResultAuthors) *ListRecommendContentResponseBodyResult {
	s.Authors = v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCategory(v string) *ListRecommendContentResponseBodyResult {
	s.Category = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCharge(v bool) *ListRecommendContentResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCommCateId(v int64) *ListRecommendContentResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetCover(v *ListRecommendContentResponseBodyResultCover) *ListRecommendContentResponseBodyResult {
	s.Cover = v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetDescription(v string) *ListRecommendContentResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetHotScore(v float64) *ListRecommendContentResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetId(v int64) *ListRecommendContentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetItemType(v string) *ListRecommendContentResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetRawId(v string) *ListRecommendContentResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetSource(v string) *ListRecommendContentResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetTitle(v string) *ListRecommendContentResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetType(v string) *ListRecommendContentResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) SetValid(v string) *ListRecommendContentResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *ListRecommendContentResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentResponseBodyResultAuthors struct {
	AuthorTypes []*string                                           `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Cover       *ListRecommendContentResponseBodyResultAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// MALE
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// 13597709
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// 12311
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// example:
	//
	// qingting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListRecommendContentResponseBodyResultAuthors) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetAuthorTypes() []*string {
	return s.AuthorTypes
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetCover() *ListRecommendContentResponseBodyResultAuthorsCover {
	return s.Cover
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetDescription() *string {
	return s.Description
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetGender() *string {
	return s.Gender
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetId() *int64 {
	return s.Id
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetOnline() *bool {
	return s.Online
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetRawId() *string {
	return s.RawId
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetSource() *string {
	return s.Source
}

func (s *ListRecommendContentResponseBodyResultAuthors) GetTitle() *string {
	return s.Title
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetAuthorTypes(v []*string) *ListRecommendContentResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetCover(v *ListRecommendContentResponseBodyResultAuthorsCover) *ListRecommendContentResponseBodyResultAuthors {
	s.Cover = v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetDescription(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Description = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetGender(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetId(v int64) *ListRecommendContentResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetOnline(v bool) *ListRecommendContentResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetRawId(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.RawId = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetSource(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) SetTitle(v string) *ListRecommendContentResponseBodyResultAuthors {
	s.Title = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthors) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentResponseBodyResultAuthorsCover struct {
	// example:
	//
	// false
	CanResize *bool `json:"CanResize,omitempty" xml:"CanResize,omitempty"`
	// example:
	//
	// https://a.jpg
	Img *string `json:"Img,omitempty" xml:"Img,omitempty"`
	// example:
	//
	// https://a.jpg
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// example:
	//
	// https://a.jpg
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// example:
	//
	// https://a.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s ListRecommendContentResponseBodyResultAuthorsCover) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentResponseBodyResultAuthorsCover) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) GetImg() *string {
	return s.Img
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) GetLarge() *string {
	return s.Large
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) GetMedium() *string {
	return s.Medium
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) GetSmall() *string {
	return s.Small
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetCanResize(v bool) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetImg(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Img = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetLarge(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Large = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetMedium(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Medium = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) SetSmall(v string) *ListRecommendContentResponseBodyResultAuthorsCover {
	s.Small = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultAuthorsCover) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentResponseBodyResultCover struct {
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

func (s ListRecommendContentResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ListRecommendContentResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *ListRecommendContentResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *ListRecommendContentResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *ListRecommendContentResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *ListRecommendContentResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *ListRecommendContentResponseBodyResultCover) SetCanResize(v bool) *ListRecommendContentResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetImg(v string) *ListRecommendContentResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetLarge(v string) *ListRecommendContentResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetMediam(v string) *ListRecommendContentResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetMedium(v string) *ListRecommendContentResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) SetSmall(v string) *ListRecommendContentResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *ListRecommendContentResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
