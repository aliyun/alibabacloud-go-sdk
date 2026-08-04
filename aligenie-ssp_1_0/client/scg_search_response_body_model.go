// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScgSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ScgSearchResponseBody
	GetCode() *int32
	SetMessage(v string) *ScgSearchResponseBody
	GetMessage() *string
	SetPageNum(v int32) *ScgSearchResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *ScgSearchResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ScgSearchResponseBody
	GetRequestId() *string
	SetResult(v []*ScgSearchResponseBodyResult) *ScgSearchResponseBody
	GetResult() []*ScgSearchResponseBodyResult
}

type ScgSearchResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response message
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of records per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Detailed returned information.
	//
	// example:
	//
	// {"sourceId":0,"copyright":0,"releaseTime":1196438400000,"hotScore":0.9546929544543479,"tppExtendInfo":"{\\"userEvent\\":\\"前台投放\\",\\"scgTopicName\\":\\"inside民族中国风曲风音乐库\\",\\"scgTopicId\\":\\"MC201132\\"}","source":"xiami","title":"题帕三绝","type":"music","x1Pv30d":10307,"valid":10,"cover":{"img":"http://img.xiami.net/qianxun/07d8ec1a38a5462c3afbfac41413b8af/47244c25fcf3a8f67442d02e3127d023-500x432.jpg","canResize":false},"duration":0,"rawId":"1771626071","albumType":0,"alias":["题帕三绝"],"id":268033175,"supportAudition":false,"contentType":"MUSIC_CONTENT","authorIds":[],"singers":"陈力","albumRawId":"1771626071","album":false,"x1PvTotal":14676,"commCateId":80021001,"finished":0,"isAudition":false,"appIds":[],"authorNames":["陈力","王立平","曹雪芹"],"needCharge":false,"isCharge":false,"category":"audio"}
	Result []*ScgSearchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ScgSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ScgSearchResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ScgSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScgSearchResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ScgSearchResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ScgSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScgSearchResponseBody) GetResult() []*ScgSearchResponseBodyResult {
	return s.Result
}

func (s *ScgSearchResponseBody) SetCode(v int32) *ScgSearchResponseBody {
	s.Code = &v
	return s
}

func (s *ScgSearchResponseBody) SetMessage(v string) *ScgSearchResponseBody {
	s.Message = &v
	return s
}

func (s *ScgSearchResponseBody) SetPageNum(v int32) *ScgSearchResponseBody {
	s.PageNum = &v
	return s
}

func (s *ScgSearchResponseBody) SetPageSize(v int32) *ScgSearchResponseBody {
	s.PageSize = &v
	return s
}

func (s *ScgSearchResponseBody) SetRequestId(v string) *ScgSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScgSearchResponseBody) SetResult(v []*ScgSearchResponseBodyResult) *ScgSearchResponseBody {
	s.Result = v
	return s
}

func (s *ScgSearchResponseBody) Validate() error {
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

type ScgSearchResponseBodyResult struct {
	// Whether it is an album
	//
	// example:
	//
	// false
	Album *bool `json:"Album,omitempty" xml:"Album,omitempty"`
	// Album ID
	//
	// example:
	//
	// 1795716629
	AlbumRawId *string `json:"AlbumRawId,omitempty" xml:"AlbumRawId,omitempty"`
	// Album type
	//
	// example:
	//
	// 0
	AlbumType *int32 `json:"AlbumType,omitempty" xml:"AlbumType,omitempty"`
	// Alias
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// Author ID
	AuthorIds []*int64 `json:"AuthorIds,omitempty" xml:"AuthorIds,omitempty" type:"Repeated"`
	// Author names
	AuthorNames []*string `json:"AuthorNames,omitempty" xml:"AuthorNames,omitempty" type:"Repeated"`
	// Category
	//
	// example:
	//
	// audio
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Content type
	//
	// example:
	//
	// MUSIC_CONTENT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Thumbnail
	Cover *ScgSearchResponseBodyResultCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// Whether audition is available
	//
	// example:
	//
	// false
	IsAudition *bool `json:"IsAudition,omitempty" xml:"IsAudition,omitempty"`
	// Is charged
	//
	// example:
	//
	// false
	IsCharge *string `json:"IsCharge,omitempty" xml:"IsCharge,omitempty"`
	// Whether charging is required
	//
	// example:
	//
	// false
	NeedCharge *bool `json:"NeedCharge,omitempty" xml:"NeedCharge,omitempty"`
	// Third-party content ID
	//
	// example:
	//
	// 1795716629
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// Singer name
	//
	// example:
	//
	// 黎偌天
	Singers *string `json:"Singers,omitempty" xml:"Singers,omitempty"`
	// Content source
	//
	// example:
	//
	// xiami
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Whether audition is supported
	//
	// example:
	//
	// false
	SupportAudition *bool `json:"SupportAudition,omitempty" xml:"SupportAudition,omitempty"`
	// Content title
	//
	// example:
	//
	// 那个人
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Content type
	//
	// example:
	//
	// music
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ScgSearchResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ScgSearchResponseBodyResult) GetAlbum() *bool {
	return s.Album
}

func (s *ScgSearchResponseBodyResult) GetAlbumRawId() *string {
	return s.AlbumRawId
}

func (s *ScgSearchResponseBodyResult) GetAlbumType() *int32 {
	return s.AlbumType
}

func (s *ScgSearchResponseBodyResult) GetAlias() []*string {
	return s.Alias
}

func (s *ScgSearchResponseBodyResult) GetAuthorIds() []*int64 {
	return s.AuthorIds
}

func (s *ScgSearchResponseBodyResult) GetAuthorNames() []*string {
	return s.AuthorNames
}

func (s *ScgSearchResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *ScgSearchResponseBodyResult) GetContentType() *string {
	return s.ContentType
}

func (s *ScgSearchResponseBodyResult) GetCover() *ScgSearchResponseBodyResultCover {
	return s.Cover
}

func (s *ScgSearchResponseBodyResult) GetIsAudition() *bool {
	return s.IsAudition
}

func (s *ScgSearchResponseBodyResult) GetIsCharge() *string {
	return s.IsCharge
}

func (s *ScgSearchResponseBodyResult) GetNeedCharge() *bool {
	return s.NeedCharge
}

func (s *ScgSearchResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *ScgSearchResponseBodyResult) GetSingers() *string {
	return s.Singers
}

func (s *ScgSearchResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ScgSearchResponseBodyResult) GetSupportAudition() *bool {
	return s.SupportAudition
}

func (s *ScgSearchResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *ScgSearchResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ScgSearchResponseBodyResult) SetAlbum(v bool) *ScgSearchResponseBodyResult {
	s.Album = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAlbumRawId(v string) *ScgSearchResponseBodyResult {
	s.AlbumRawId = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAlbumType(v int32) *ScgSearchResponseBodyResult {
	s.AlbumType = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAlias(v []*string) *ScgSearchResponseBodyResult {
	s.Alias = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAuthorIds(v []*int64) *ScgSearchResponseBodyResult {
	s.AuthorIds = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetAuthorNames(v []*string) *ScgSearchResponseBodyResult {
	s.AuthorNames = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetCategory(v string) *ScgSearchResponseBodyResult {
	s.Category = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetContentType(v string) *ScgSearchResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetCover(v *ScgSearchResponseBodyResultCover) *ScgSearchResponseBodyResult {
	s.Cover = v
	return s
}

func (s *ScgSearchResponseBodyResult) SetIsAudition(v bool) *ScgSearchResponseBodyResult {
	s.IsAudition = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetIsCharge(v string) *ScgSearchResponseBodyResult {
	s.IsCharge = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetNeedCharge(v bool) *ScgSearchResponseBodyResult {
	s.NeedCharge = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetRawId(v string) *ScgSearchResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetSingers(v string) *ScgSearchResponseBodyResult {
	s.Singers = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetSource(v string) *ScgSearchResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetSupportAudition(v bool) *ScgSearchResponseBodyResult {
	s.SupportAudition = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetTitle(v string) *ScgSearchResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ScgSearchResponseBodyResult) SetType(v string) *ScgSearchResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ScgSearchResponseBodyResult) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScgSearchResponseBodyResultCover struct {
	// Thumbnail image (Img, Large, Medium, and Small may not appear simultaneously; only one of them may be present)
	//
	// example:
	//
	// http://img.xiami.net/images/album/img59/56/58da2153e3133_2826959_1490690387.jpg
	Img *string `json:"Img,omitempty" xml:"Img,omitempty"`
	// Large graph
	//
	// example:
	//
	// http://img.xiami.net/images/album/img59/56/58da2153e3133_2826959_1490690387.jpg
	Large *string `json:"Large,omitempty" xml:"Large,omitempty"`
	// Medium image
	//
	// example:
	//
	// http://img.xiami.net/images/album/img59/56/58da2153e3133_2826959_1490690387.jpg
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// Small image
	//
	// example:
	//
	// http://img.xiami.net/images/album/img59/56/58da2153e3133_2826959_1490690387.jpg
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
	// Whether scaling is supported
	//
	// example:
	//
	// false
	CanResize *bool `json:"canResize,omitempty" xml:"canResize,omitempty"`
}

func (s ScgSearchResponseBodyResultCover) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchResponseBodyResultCover) GoString() string {
	return s.String()
}

func (s *ScgSearchResponseBodyResultCover) GetImg() *string {
	return s.Img
}

func (s *ScgSearchResponseBodyResultCover) GetLarge() *string {
	return s.Large
}

func (s *ScgSearchResponseBodyResultCover) GetMedium() *string {
	return s.Medium
}

func (s *ScgSearchResponseBodyResultCover) GetSmall() *string {
	return s.Small
}

func (s *ScgSearchResponseBodyResultCover) GetCanResize() *bool {
	return s.CanResize
}

func (s *ScgSearchResponseBodyResultCover) SetImg(v string) *ScgSearchResponseBodyResultCover {
	s.Img = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetLarge(v string) *ScgSearchResponseBodyResultCover {
	s.Large = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetMedium(v string) *ScgSearchResponseBodyResultCover {
	s.Medium = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetSmall(v string) *ScgSearchResponseBodyResultCover {
	s.Small = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) SetCanResize(v bool) *ScgSearchResponseBodyResultCover {
	s.CanResize = &v
	return s
}

func (s *ScgSearchResponseBodyResultCover) Validate() error {
	return dara.Validate(s)
}
