// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlayHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListPlayHistoryResponseBody
	GetCode() *int32
	SetMessage(v string) *ListPlayHistoryResponseBody
	GetMessage() *string
	SetResult(v []*ListPlayHistoryResponseBodyResult) *ListPlayHistoryResponseBody
	GetResult() []*ListPlayHistoryResponseBodyResult
	SetRequestId(v string) *ListPlayHistoryResponseBody
	GetRequestId() *string
}

type ListPlayHistoryResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Result  []*ListPlayHistoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPlayHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPlayHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPlayHistoryResponseBody) GetResult() []*ListPlayHistoryResponseBodyResult {
	return s.Result
}

func (s *ListPlayHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPlayHistoryResponseBody) SetCode(v int32) *ListPlayHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListPlayHistoryResponseBody) SetMessage(v string) *ListPlayHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListPlayHistoryResponseBody) SetResult(v []*ListPlayHistoryResponseBodyResult) *ListPlayHistoryResponseBody {
	s.Result = v
	return s
}

func (s *ListPlayHistoryResponseBody) SetRequestId(v string) *ListPlayHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPlayHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPlayHistoryResponseBodyResult struct {
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Audition *bool                                       `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors  []*ListPlayHistoryResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
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
	CommCateId  *int64                                  `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *ListPlayHistoryResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s ListPlayHistoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResult) GetAlias() []*string {
	return s.Alias
}

func (s *ListPlayHistoryResponseBodyResult) GetAudition() *bool {
	return s.Audition
}

func (s *ListPlayHistoryResponseBodyResult) GetAuthors() []*ListPlayHistoryResponseBodyResultAuthors {
	return s.Authors
}

func (s *ListPlayHistoryResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *ListPlayHistoryResponseBodyResult) GetCharge() *bool {
	return s.Charge
}

func (s *ListPlayHistoryResponseBodyResult) GetCommCateId() *int64 {
	return s.CommCateId
}

func (s *ListPlayHistoryResponseBodyResult) GetCover() *ListPlayHistoryResponseBodyResultCover {
	return s.Cover
}

func (s *ListPlayHistoryResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListPlayHistoryResponseBodyResult) GetHotScore() *float64 {
	return s.HotScore
}

func (s *ListPlayHistoryResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListPlayHistoryResponseBodyResult) GetItemType() *string {
	return s.ItemType
}

func (s *ListPlayHistoryResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListPlayHistoryResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *ListPlayHistoryResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListPlayHistoryResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *ListPlayHistoryResponseBodyResult) SetAlias(v []*string) *ListPlayHistoryResponseBodyResult {
	s.Alias = v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetAudition(v bool) *ListPlayHistoryResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetAuthors(v []*ListPlayHistoryResponseBodyResultAuthors) *ListPlayHistoryResponseBodyResult {
	s.Authors = v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCategory(v string) *ListPlayHistoryResponseBodyResult {
	s.Category = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCharge(v bool) *ListPlayHistoryResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCommCateId(v int64) *ListPlayHistoryResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetCover(v *ListPlayHistoryResponseBodyResultCover) *ListPlayHistoryResponseBodyResult {
	s.Cover = v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetDescription(v string) *ListPlayHistoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetHotScore(v float64) *ListPlayHistoryResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetId(v int64) *ListPlayHistoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetItemType(v string) *ListPlayHistoryResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetSource(v string) *ListPlayHistoryResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetTitle(v string) *ListPlayHistoryResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetType(v string) *ListPlayHistoryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) SetValid(v string) *ListPlayHistoryResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListPlayHistoryResponseBodyResultAuthors struct {
	AuthorTypes []*string                                      `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	Cover       *ListPlayHistoryResponseBodyResultAuthorsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 123123
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// example:
	//
	// qingting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListPlayHistoryResponseBodyResultAuthors) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetAuthorTypes() []*string {
	return s.AuthorTypes
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetCover() *ListPlayHistoryResponseBodyResultAuthorsCover {
	return s.Cover
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetDescription() *string {
	return s.Description
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetGender() *string {
	return s.Gender
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetId() *int64 {
	return s.Id
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetOnline() *bool {
	return s.Online
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetRawId() *string {
	return s.RawId
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetSource() *string {
	return s.Source
}

func (s *ListPlayHistoryResponseBodyResultAuthors) GetTitle() *string {
	return s.Title
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetAuthorTypes(v []*string) *ListPlayHistoryResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetCover(v *ListPlayHistoryResponseBodyResultAuthorsCover) *ListPlayHistoryResponseBodyResultAuthors {
	s.Cover = v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetDescription(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Description = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetGender(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetId(v int64) *ListPlayHistoryResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetOnline(v bool) *ListPlayHistoryResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetRawId(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.RawId = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetSource(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) SetTitle(v string) *ListPlayHistoryResponseBodyResultAuthors {
	s.Title = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthors) Validate() error {
	return dara.Validate(s)
}

type ListPlayHistoryResponseBodyResultAuthorsCover struct {
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

func (s ListPlayHistoryResponseBodyResultAuthorsCover) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResultAuthorsCover) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) GetImg() *string {
	return s.Img
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) GetLarge() *string {
	return s.Large
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) GetMedium() *string {
	return s.Medium
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) GetSmall() *string {
	return s.Small
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetCanResize(v bool) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.CanResize = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetImg(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Img = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetLarge(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Large = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetMedium(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Medium = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) SetSmall(v string) *ListPlayHistoryResponseBodyResultAuthorsCover {
	s.Small = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultAuthorsCover) Validate() error {
	return dara.Validate(s)
}

type ListPlayHistoryResponseBodyResultCover struct {
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

func (s ListPlayHistoryResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ListPlayHistoryResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *ListPlayHistoryResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *ListPlayHistoryResponseBodyResultCover) GetMediam() *string {
	return s.Mediam
}

func (s *ListPlayHistoryResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *ListPlayHistoryResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *ListPlayHistoryResponseBodyResultCover) SetCanResize(v bool) *ListPlayHistoryResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetImg(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetLarge(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetMediam(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Mediam = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetMedium(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) SetSmall(v string) *ListPlayHistoryResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *ListPlayHistoryResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
