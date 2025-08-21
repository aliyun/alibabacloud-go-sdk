// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAlbumResponseBody
	GetCode() *int32
	SetRequestId(v string) *GetAlbumResponseBody
	GetRequestId() *string
	SetResult(v *GetAlbumResponseBodyResult) *GetAlbumResponseBody
	GetResult() *GetAlbumResponseBodyResult
}

type GetAlbumResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAlbumResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAlbumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAlbumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlbumResponseBody) GetResult() *GetAlbumResponseBodyResult {
	return s.Result
}

func (s *GetAlbumResponseBody) SetCode(v int32) *GetAlbumResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlbumResponseBody) SetRequestId(v string) *GetAlbumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlbumResponseBody) SetResult(v *GetAlbumResponseBodyResult) *GetAlbumResponseBody {
	s.Result = v
	return s
}

func (s *GetAlbumResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAlbumResponseBodyResult struct {
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Audition *bool                                `json:"Audition,omitempty" xml:"Audition,omitempty"`
	Authors  []*GetAlbumResponseBodyResultAuthors `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
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
	CommCateId  *int64                           `json:"CommCateId,omitempty" xml:"CommCateId,omitempty"`
	Cover       *GetAlbumResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Description *string                          `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// FINISHED
	Finished *string `json:"Finished,omitempty" xml:"Finished,omitempty"`
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
	// 23242
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// example:
	//
	// qignting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 12
	TotalEpisode *int32 `json:"TotalEpisode,omitempty" xml:"TotalEpisode,omitempty"`
	// example:
	//
	// program
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// VALID
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetAlbumResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBodyResult) GetAlias() []*string {
	return s.Alias
}

func (s *GetAlbumResponseBodyResult) GetAudition() *bool {
	return s.Audition
}

func (s *GetAlbumResponseBodyResult) GetAuthors() []*GetAlbumResponseBodyResultAuthors {
	return s.Authors
}

func (s *GetAlbumResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *GetAlbumResponseBodyResult) GetCharge() *bool {
	return s.Charge
}

func (s *GetAlbumResponseBodyResult) GetCommCateId() *int64 {
	return s.CommCateId
}

func (s *GetAlbumResponseBodyResult) GetCover() *GetAlbumResponseBodyResultCover {
	return s.Cover
}

func (s *GetAlbumResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetAlbumResponseBodyResult) GetFinished() *string {
	return s.Finished
}

func (s *GetAlbumResponseBodyResult) GetHotScore() *float64 {
	return s.HotScore
}

func (s *GetAlbumResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetAlbumResponseBodyResult) GetItemType() *string {
	return s.ItemType
}

func (s *GetAlbumResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *GetAlbumResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *GetAlbumResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetAlbumResponseBodyResult) GetTotalEpisode() *int32 {
	return s.TotalEpisode
}

func (s *GetAlbumResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetAlbumResponseBodyResult) GetValid() *string {
	return s.Valid
}

func (s *GetAlbumResponseBodyResult) SetAlias(v []*string) *GetAlbumResponseBodyResult {
	s.Alias = v
	return s
}

func (s *GetAlbumResponseBodyResult) SetAudition(v bool) *GetAlbumResponseBodyResult {
	s.Audition = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetAuthors(v []*GetAlbumResponseBodyResultAuthors) *GetAlbumResponseBodyResult {
	s.Authors = v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCategory(v string) *GetAlbumResponseBodyResult {
	s.Category = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCharge(v bool) *GetAlbumResponseBodyResult {
	s.Charge = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCommCateId(v int64) *GetAlbumResponseBodyResult {
	s.CommCateId = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetCover(v *GetAlbumResponseBodyResultCover) *GetAlbumResponseBodyResult {
	s.Cover = v
	return s
}

func (s *GetAlbumResponseBodyResult) SetDescription(v string) *GetAlbumResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetFinished(v string) *GetAlbumResponseBodyResult {
	s.Finished = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetHotScore(v float64) *GetAlbumResponseBodyResult {
	s.HotScore = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetId(v int64) *GetAlbumResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetItemType(v string) *GetAlbumResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetRawId(v string) *GetAlbumResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetSource(v string) *GetAlbumResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetTitle(v string) *GetAlbumResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetTotalEpisode(v int32) *GetAlbumResponseBodyResult {
	s.TotalEpisode = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetType(v string) *GetAlbumResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetAlbumResponseBodyResult) SetValid(v string) *GetAlbumResponseBodyResult {
	s.Valid = &v
	return s
}

func (s *GetAlbumResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetAlbumResponseBodyResultAuthors struct {
	AuthorTypes []*string `json:"AuthorTypes,omitempty" xml:"AuthorTypes,omitempty" type:"Repeated"`
	// example:
	//
	// MALE
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// 13123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// qignting
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetAlbumResponseBodyResultAuthors) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumResponseBodyResultAuthors) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBodyResultAuthors) GetAuthorTypes() []*string {
	return s.AuthorTypes
}

func (s *GetAlbumResponseBodyResultAuthors) GetGender() *string {
	return s.Gender
}

func (s *GetAlbumResponseBodyResultAuthors) GetId() *int64 {
	return s.Id
}

func (s *GetAlbumResponseBodyResultAuthors) GetOnline() *bool {
	return s.Online
}

func (s *GetAlbumResponseBodyResultAuthors) GetSource() *string {
	return s.Source
}

func (s *GetAlbumResponseBodyResultAuthors) GetTitle() *string {
	return s.Title
}

func (s *GetAlbumResponseBodyResultAuthors) SetAuthorTypes(v []*string) *GetAlbumResponseBodyResultAuthors {
	s.AuthorTypes = v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetGender(v string) *GetAlbumResponseBodyResultAuthors {
	s.Gender = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetId(v int64) *GetAlbumResponseBodyResultAuthors {
	s.Id = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetOnline(v bool) *GetAlbumResponseBodyResultAuthors {
	s.Online = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetSource(v string) *GetAlbumResponseBodyResultAuthors {
	s.Source = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) SetTitle(v string) *GetAlbumResponseBodyResultAuthors {
	s.Title = &v
	return s
}

func (s *GetAlbumResponseBodyResultAuthors) Validate() error {
	return dara.Validate(s)
}

type GetAlbumResponseBodyResultCover struct {
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

func (s GetAlbumResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *GetAlbumResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *GetAlbumResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *GetAlbumResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *GetAlbumResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *GetAlbumResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *GetAlbumResponseBodyResultCover) SetCanResize(v bool) *GetAlbumResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetImg(v string) *GetAlbumResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetLarge(v string) *GetAlbumResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetMedium(v string) *GetAlbumResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) SetSmall(v string) *GetAlbumResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *GetAlbumResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
