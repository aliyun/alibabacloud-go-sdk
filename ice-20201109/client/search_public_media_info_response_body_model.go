// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPublicMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicMediaInfos(v []*SearchPublicMediaInfoResponseBodyPublicMediaInfos) *SearchPublicMediaInfoResponseBody
	GetPublicMediaInfos() []*SearchPublicMediaInfoResponseBodyPublicMediaInfos
	SetRequestId(v string) *SearchPublicMediaInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchPublicMediaInfoResponseBody
	GetTotalCount() *int64
}

type SearchPublicMediaInfoResponseBody struct {
	PublicMediaInfos []*SearchPublicMediaInfoResponseBodyPublicMediaInfos `json:"PublicMediaInfos,omitempty" xml:"PublicMediaInfos,omitempty" type:"Repeated"`
	// example:
	//
	// ****3CFB-2767-54FD-B311-BD15A4C1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchPublicMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchPublicMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPublicMediaInfoResponseBody) GetPublicMediaInfos() []*SearchPublicMediaInfoResponseBodyPublicMediaInfos {
	return s.PublicMediaInfos
}

func (s *SearchPublicMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchPublicMediaInfoResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchPublicMediaInfoResponseBody) SetPublicMediaInfos(v []*SearchPublicMediaInfoResponseBodyPublicMediaInfos) *SearchPublicMediaInfoResponseBody {
	s.PublicMediaInfos = v
	return s
}

func (s *SearchPublicMediaInfoResponseBody) SetRequestId(v string) *SearchPublicMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBody) SetTotalCount(v int64) *SearchPublicMediaInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchPublicMediaInfoResponseBodyPublicMediaInfos struct {
	// example:
	//
	// true
	Authorized *bool `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	// example:
	//
	// true
	Favorite  *bool                                                       `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	MediaInfo *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// example:
	//
	// 100
	RemainingAuthTime *string `json:"RemainingAuthTime,omitempty" xml:"RemainingAuthTime,omitempty"`
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfos) GoString() string {
	return s.String()
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) GetAuthorized() *bool {
	return s.Authorized
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) GetFavorite() *bool {
	return s.Favorite
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) GetMediaInfo() *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo {
	return s.MediaInfo
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) GetRemainingAuthTime() *string {
	return s.RemainingAuthTime
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) SetAuthorized(v bool) *SearchPublicMediaInfoResponseBodyPublicMediaInfos {
	s.Authorized = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) SetFavorite(v bool) *SearchPublicMediaInfoResponseBodyPublicMediaInfos {
	s.Favorite = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) SetMediaInfo(v *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) *SearchPublicMediaInfoResponseBodyPublicMediaInfos {
	s.MediaInfo = v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) SetRemainingAuthTime(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfos {
	s.RemainingAuthTime = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfos) Validate() error {
	return dara.Validate(s)
}

type SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo struct {
	DynamicMetaData *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty" type:"Struct"`
	// BasicInfo
	MediaBasicInfo *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// example:
	//
	// icepublic-****87b921bb4a55908a72a0537e****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) GoString() string {
	return s.String()
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) GetDynamicMetaData() *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData {
	return s.DynamicMetaData
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) GetMediaBasicInfo() *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) SetDynamicMetaData(v *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo {
	s.DynamicMetaData = v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) SetMediaBasicInfo(v *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo {
	s.MediaBasicInfo = v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) SetMediaId(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo {
	s.MediaId = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfo) Validate() error {
	return dara.Validate(s)
}

type SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData struct {
	// example:
	//
	// "{\\"AuditionUrl\\": \\"http://xxx\\", \\"AuditionCount\\": 3...}"
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) String() string {
	return dara.Prettify(s)
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) GoString() string {
	return s.String()
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) GetData() *string {
	return s.Data
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) GetType() *string {
	return s.Type
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) SetData(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData {
	s.Data = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) SetType(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData {
	s.Type = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoDynamicMetaData) Validate() error {
	return dara.Validate(s)
}

type SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo struct {
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// category
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// 2020-12-26T06:04:49Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2020-12-29T06:04:49Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// MediaId
	//
	// example:
	//
	// icepublic-****87b921bb4a55908a72a0537e****
	MediaId   *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// audio
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 2020-12-26T06:04:50Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// userDataTest
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetBusinessType(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetCategory(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetCoverURL(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetCreateTime(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetDeletedTime(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetDescription(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetMediaId(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetMediaTags(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetMediaType(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetModifiedTime(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetSource(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetSpriteImages(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetStatus(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetTitle(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) SetUserData(v string) *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *SearchPublicMediaInfoResponseBodyPublicMediaInfosMediaInfoMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
