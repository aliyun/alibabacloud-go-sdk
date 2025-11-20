// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *QueryDentryResponseBody
	GetContentType() *string
	SetCreatedTime(v int64) *QueryDentryResponseBody
	GetCreatedTime() *int64
	SetCreator(v *QueryDentryResponseBodyCreator) *QueryDentryResponseBody
	GetCreator() *QueryDentryResponseBodyCreator
	SetDentryId(v string) *QueryDentryResponseBody
	GetDentryId() *string
	SetDentryType(v string) *QueryDentryResponseBody
	GetDentryType() *string
	SetDentryUuid(v string) *QueryDentryResponseBody
	GetDentryUuid() *string
	SetDocKey(v string) *QueryDentryResponseBody
	GetDocKey() *string
	SetExtension(v string) *QueryDentryResponseBody
	GetExtension() *string
	SetHasChildren(v bool) *QueryDentryResponseBody
	GetHasChildren() *bool
	SetLinkSourceInfo(v *QueryDentryResponseBodyLinkSourceInfo) *QueryDentryResponseBody
	GetLinkSourceInfo() *QueryDentryResponseBodyLinkSourceInfo
	SetName(v string) *QueryDentryResponseBody
	GetName() *string
	SetPath(v string) *QueryDentryResponseBody
	GetPath() *string
	SetRequestId(v string) *QueryDentryResponseBody
	GetRequestId() *string
	SetSpace(v *QueryDentryResponseBodySpace) *QueryDentryResponseBody
	GetSpace() *QueryDentryResponseBodySpace
	SetSpaceId(v string) *QueryDentryResponseBody
	GetSpaceId() *string
	SetUpdatedTime(v int64) *QueryDentryResponseBody
	GetUpdatedTime() *int64
	SetUpdater(v *QueryDentryResponseBodyUpdater) *QueryDentryResponseBody
	GetUpdater() *QueryDentryResponseBodyUpdater
	SetUrl(v string) *QueryDentryResponseBody
	GetUrl() *string
	SetVisitorInfo(v *QueryDentryResponseBodyVisitorInfo) *QueryDentryResponseBody
	GetVisitorInfo() *QueryDentryResponseBodyVisitorInfo
}

type QueryDentryResponseBody struct {
	// example:
	//
	// alidoc
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// example:
	//
	// 12345678
	CreatedTime *int64                          `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator     *QueryDentryResponseBodyCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// abc
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// file
	DentryType *string `json:"dentryType,omitempty" xml:"dentryType,omitempty"`
	// example:
	//
	// cdefg
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// aabbcc
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// example:
	//
	// alidoc
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// false
	HasChildren    *bool                                  `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	LinkSourceInfo *QueryDentryResponseBodyLinkSourceInfo `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// hello
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 测试组织/测试知识库/abc
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Space     *QueryDentryResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
	// example:
	//
	// bcd
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// 12345678
	UpdatedTime *int64                          `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	Updater     *QueryDentryResponseBodyUpdater `json:"updater,omitempty" xml:"updater,omitempty" type:"Struct"`
	// example:
	//
	// https://xxx.yy
	Url         *string                             `json:"url,omitempty" xml:"url,omitempty"`
	VisitorInfo *QueryDentryResponseBodyVisitorInfo `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty" type:"Struct"`
}

func (s QueryDentryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *QueryDentryResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *QueryDentryResponseBody) GetCreator() *QueryDentryResponseBodyCreator {
	return s.Creator
}

func (s *QueryDentryResponseBody) GetDentryId() *string {
	return s.DentryId
}

func (s *QueryDentryResponseBody) GetDentryType() *string {
	return s.DentryType
}

func (s *QueryDentryResponseBody) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *QueryDentryResponseBody) GetDocKey() *string {
	return s.DocKey
}

func (s *QueryDentryResponseBody) GetExtension() *string {
	return s.Extension
}

func (s *QueryDentryResponseBody) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *QueryDentryResponseBody) GetLinkSourceInfo() *QueryDentryResponseBodyLinkSourceInfo {
	return s.LinkSourceInfo
}

func (s *QueryDentryResponseBody) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBody) GetPath() *string {
	return s.Path
}

func (s *QueryDentryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDentryResponseBody) GetSpace() *QueryDentryResponseBodySpace {
	return s.Space
}

func (s *QueryDentryResponseBody) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentryResponseBody) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *QueryDentryResponseBody) GetUpdater() *QueryDentryResponseBodyUpdater {
	return s.Updater
}

func (s *QueryDentryResponseBody) GetUrl() *string {
	return s.Url
}

func (s *QueryDentryResponseBody) GetVisitorInfo() *QueryDentryResponseBodyVisitorInfo {
	return s.VisitorInfo
}

func (s *QueryDentryResponseBody) SetContentType(v string) *QueryDentryResponseBody {
	s.ContentType = &v
	return s
}

func (s *QueryDentryResponseBody) SetCreatedTime(v int64) *QueryDentryResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *QueryDentryResponseBody) SetCreator(v *QueryDentryResponseBodyCreator) *QueryDentryResponseBody {
	s.Creator = v
	return s
}

func (s *QueryDentryResponseBody) SetDentryId(v string) *QueryDentryResponseBody {
	s.DentryId = &v
	return s
}

func (s *QueryDentryResponseBody) SetDentryType(v string) *QueryDentryResponseBody {
	s.DentryType = &v
	return s
}

func (s *QueryDentryResponseBody) SetDentryUuid(v string) *QueryDentryResponseBody {
	s.DentryUuid = &v
	return s
}

func (s *QueryDentryResponseBody) SetDocKey(v string) *QueryDentryResponseBody {
	s.DocKey = &v
	return s
}

func (s *QueryDentryResponseBody) SetExtension(v string) *QueryDentryResponseBody {
	s.Extension = &v
	return s
}

func (s *QueryDentryResponseBody) SetHasChildren(v bool) *QueryDentryResponseBody {
	s.HasChildren = &v
	return s
}

func (s *QueryDentryResponseBody) SetLinkSourceInfo(v *QueryDentryResponseBodyLinkSourceInfo) *QueryDentryResponseBody {
	s.LinkSourceInfo = v
	return s
}

func (s *QueryDentryResponseBody) SetName(v string) *QueryDentryResponseBody {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBody) SetPath(v string) *QueryDentryResponseBody {
	s.Path = &v
	return s
}

func (s *QueryDentryResponseBody) SetRequestId(v string) *QueryDentryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDentryResponseBody) SetSpace(v *QueryDentryResponseBodySpace) *QueryDentryResponseBody {
	s.Space = v
	return s
}

func (s *QueryDentryResponseBody) SetSpaceId(v string) *QueryDentryResponseBody {
	s.SpaceId = &v
	return s
}

func (s *QueryDentryResponseBody) SetUpdatedTime(v int64) *QueryDentryResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *QueryDentryResponseBody) SetUpdater(v *QueryDentryResponseBodyUpdater) *QueryDentryResponseBody {
	s.Updater = v
	return s
}

func (s *QueryDentryResponseBody) SetUrl(v string) *QueryDentryResponseBody {
	s.Url = &v
	return s
}

func (s *QueryDentryResponseBody) SetVisitorInfo(v *QueryDentryResponseBodyVisitorInfo) *QueryDentryResponseBody {
	s.VisitorInfo = v
	return s
}

func (s *QueryDentryResponseBody) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.LinkSourceInfo != nil {
		if err := s.LinkSourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Space != nil {
		if err := s.Space.Validate(); err != nil {
			return err
		}
	}
	if s.Updater != nil {
		if err := s.Updater.Validate(); err != nil {
			return err
		}
	}
	if s.VisitorInfo != nil {
		if err := s.VisitorInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDentryResponseBodyCreator struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDentryResponseBodyCreator) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodyCreator) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodyCreator) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBodyCreator) GetUserId() *string {
	return s.UserId
}

func (s *QueryDentryResponseBodyCreator) SetName(v string) *QueryDentryResponseBodyCreator {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBodyCreator) SetUserId(v string) *QueryDentryResponseBodyCreator {
	s.UserId = &v
	return s
}

func (s *QueryDentryResponseBodyCreator) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodyLinkSourceInfo struct {
	// example:
	//
	// docx
	Extension *string                                       `json:"Extension,omitempty" xml:"Extension,omitempty"`
	IconUrl   *QueryDentryResponseBodyLinkSourceInfoIconUrl `json:"IconUrl,omitempty" xml:"IconUrl,omitempty" type:"Struct"`
	// example:
	//
	// def
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	LinkType *int64 `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	// example:
	//
	// def
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s QueryDentryResponseBodyLinkSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodyLinkSourceInfo) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodyLinkSourceInfo) GetExtension() *string {
	return s.Extension
}

func (s *QueryDentryResponseBodyLinkSourceInfo) GetIconUrl() *QueryDentryResponseBodyLinkSourceInfoIconUrl {
	return s.IconUrl
}

func (s *QueryDentryResponseBodyLinkSourceInfo) GetId() *string {
	return s.Id
}

func (s *QueryDentryResponseBodyLinkSourceInfo) GetLinkType() *int64 {
	return s.LinkType
}

func (s *QueryDentryResponseBodyLinkSourceInfo) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentryResponseBodyLinkSourceInfo) SetExtension(v string) *QueryDentryResponseBodyLinkSourceInfo {
	s.Extension = &v
	return s
}

func (s *QueryDentryResponseBodyLinkSourceInfo) SetIconUrl(v *QueryDentryResponseBodyLinkSourceInfoIconUrl) *QueryDentryResponseBodyLinkSourceInfo {
	s.IconUrl = v
	return s
}

func (s *QueryDentryResponseBodyLinkSourceInfo) SetId(v string) *QueryDentryResponseBodyLinkSourceInfo {
	s.Id = &v
	return s
}

func (s *QueryDentryResponseBodyLinkSourceInfo) SetLinkType(v int64) *QueryDentryResponseBodyLinkSourceInfo {
	s.LinkType = &v
	return s
}

func (s *QueryDentryResponseBodyLinkSourceInfo) SetSpaceId(v string) *QueryDentryResponseBodyLinkSourceInfo {
	s.SpaceId = &v
	return s
}

func (s *QueryDentryResponseBodyLinkSourceInfo) Validate() error {
	if s.IconUrl != nil {
		if err := s.IconUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDentryResponseBodyLinkSourceInfoIconUrl struct {
	// example:
	//
	// gh
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// def
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s QueryDentryResponseBodyLinkSourceInfoIconUrl) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodyLinkSourceInfoIconUrl) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodyLinkSourceInfoIconUrl) GetLine() *string {
	return s.Line
}

func (s *QueryDentryResponseBodyLinkSourceInfoIconUrl) GetSmall() *string {
	return s.Small
}

func (s *QueryDentryResponseBodyLinkSourceInfoIconUrl) SetLine(v string) *QueryDentryResponseBodyLinkSourceInfoIconUrl {
	s.Line = &v
	return s
}

func (s *QueryDentryResponseBodyLinkSourceInfoIconUrl) SetSmall(v string) *QueryDentryResponseBodyLinkSourceInfoIconUrl {
	s.Small = &v
	return s
}

func (s *QueryDentryResponseBodyLinkSourceInfoIconUrl) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpace struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1xxxxx.png
	Cover *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// example:
	//
	// 这是简介
	Description *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	HdIconVO    *QueryDentryResponseBodySpaceHdIconVO `json:"HdIconVO,omitempty" xml:"HdIconVO,omitempty" type:"Struct"`
	IconVO      *QueryDentryResponseBodySpaceIconVO   `json:"IconVO,omitempty" xml:"IconVO,omitempty" type:"Struct"`
	// example:
	//
	// n9XJxxxxx
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 测试知识库
	Owner      *QueryDentryResponseBodySpaceOwner        `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	RecentList []*QueryDentryResponseBodySpaceRecentList `json:"RecentList,omitempty" xml:"RecentList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://alidocs.dingtalk.com/i/spaces/n9XJ*******Xy/overview
	Url         *string                                  `json:"Url,omitempty" xml:"Url,omitempty"`
	VisitorInfo *QueryDentryResponseBodySpaceVisitorInfo `json:"VisitorInfo,omitempty" xml:"VisitorInfo,omitempty" type:"Struct"`
}

func (s QueryDentryResponseBodySpace) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpace) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpace) GetCover() *string {
	return s.Cover
}

func (s *QueryDentryResponseBodySpace) GetDescription() *string {
	return s.Description
}

func (s *QueryDentryResponseBodySpace) GetHdIconVO() *QueryDentryResponseBodySpaceHdIconVO {
	return s.HdIconVO
}

func (s *QueryDentryResponseBodySpace) GetIconVO() *QueryDentryResponseBodySpaceIconVO {
	return s.IconVO
}

func (s *QueryDentryResponseBodySpace) GetId() *string {
	return s.Id
}

func (s *QueryDentryResponseBodySpace) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBodySpace) GetOwner() *QueryDentryResponseBodySpaceOwner {
	return s.Owner
}

func (s *QueryDentryResponseBodySpace) GetRecentList() []*QueryDentryResponseBodySpaceRecentList {
	return s.RecentList
}

func (s *QueryDentryResponseBodySpace) GetType() *int32 {
	return s.Type
}

func (s *QueryDentryResponseBodySpace) GetUrl() *string {
	return s.Url
}

func (s *QueryDentryResponseBodySpace) GetVisitorInfo() *QueryDentryResponseBodySpaceVisitorInfo {
	return s.VisitorInfo
}

func (s *QueryDentryResponseBodySpace) SetCover(v string) *QueryDentryResponseBodySpace {
	s.Cover = &v
	return s
}

func (s *QueryDentryResponseBodySpace) SetDescription(v string) *QueryDentryResponseBodySpace {
	s.Description = &v
	return s
}

func (s *QueryDentryResponseBodySpace) SetHdIconVO(v *QueryDentryResponseBodySpaceHdIconVO) *QueryDentryResponseBodySpace {
	s.HdIconVO = v
	return s
}

func (s *QueryDentryResponseBodySpace) SetIconVO(v *QueryDentryResponseBodySpaceIconVO) *QueryDentryResponseBodySpace {
	s.IconVO = v
	return s
}

func (s *QueryDentryResponseBodySpace) SetId(v string) *QueryDentryResponseBodySpace {
	s.Id = &v
	return s
}

func (s *QueryDentryResponseBodySpace) SetName(v string) *QueryDentryResponseBodySpace {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBodySpace) SetOwner(v *QueryDentryResponseBodySpaceOwner) *QueryDentryResponseBodySpace {
	s.Owner = v
	return s
}

func (s *QueryDentryResponseBodySpace) SetRecentList(v []*QueryDentryResponseBodySpaceRecentList) *QueryDentryResponseBodySpace {
	s.RecentList = v
	return s
}

func (s *QueryDentryResponseBodySpace) SetType(v int32) *QueryDentryResponseBodySpace {
	s.Type = &v
	return s
}

func (s *QueryDentryResponseBodySpace) SetUrl(v string) *QueryDentryResponseBodySpace {
	s.Url = &v
	return s
}

func (s *QueryDentryResponseBodySpace) SetVisitorInfo(v *QueryDentryResponseBodySpaceVisitorInfo) *QueryDentryResponseBodySpace {
	s.VisitorInfo = v
	return s
}

func (s *QueryDentryResponseBodySpace) Validate() error {
	if s.HdIconVO != nil {
		if err := s.HdIconVO.Validate(); err != nil {
			return err
		}
	}
	if s.IconVO != nil {
		if err := s.IconVO.Validate(); err != nil {
			return err
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	if s.RecentList != nil {
		for _, item := range s.RecentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VisitorInfo != nil {
		if err := s.VisitorInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDentryResponseBodySpaceHdIconVO struct {
	// example:
	//
	// http://
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDentryResponseBodySpaceHdIconVO) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceHdIconVO) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceHdIconVO) GetIcon() *string {
	return s.Icon
}

func (s *QueryDentryResponseBodySpaceHdIconVO) GetType() *string {
	return s.Type
}

func (s *QueryDentryResponseBodySpaceHdIconVO) SetIcon(v string) *QueryDentryResponseBodySpaceHdIconVO {
	s.Icon = &v
	return s
}

func (s *QueryDentryResponseBodySpaceHdIconVO) SetType(v string) *QueryDentryResponseBodySpaceHdIconVO {
	s.Type = &v
	return s
}

func (s *QueryDentryResponseBodySpaceHdIconVO) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceIconVO struct {
	// example:
	//
	// http://
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDentryResponseBodySpaceIconVO) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceIconVO) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceIconVO) GetIcon() *string {
	return s.Icon
}

func (s *QueryDentryResponseBodySpaceIconVO) GetType() *string {
	return s.Type
}

func (s *QueryDentryResponseBodySpaceIconVO) SetIcon(v string) *QueryDentryResponseBodySpaceIconVO {
	s.Icon = &v
	return s
}

func (s *QueryDentryResponseBodySpaceIconVO) SetType(v string) *QueryDentryResponseBodySpaceIconVO {
	s.Type = &v
	return s
}

func (s *QueryDentryResponseBodySpaceIconVO) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceOwner struct {
	// example:
	//
	// 小钉
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDentryResponseBodySpaceOwner) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceOwner) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceOwner) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBodySpaceOwner) GetUserId() *string {
	return s.UserId
}

func (s *QueryDentryResponseBodySpaceOwner) SetName(v string) *QueryDentryResponseBodySpaceOwner {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBodySpaceOwner) SetUserId(v string) *QueryDentryResponseBodySpaceOwner {
	s.UserId = &v
	return s
}

func (s *QueryDentryResponseBodySpaceOwner) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceRecentList struct {
	// example:
	//
	// alidoc
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 12345678
	CreatedTime *int64                                         `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Creator     *QueryDentryResponseBodySpaceRecentListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// abc
	DentryId *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// example:
	//
	// file
	DentryType *string `json:"DentryType,omitempty" xml:"DentryType,omitempty"`
	// example:
	//
	// cdefg
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// example:
	//
	// aabbcc
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	// example:
	//
	// alidoc
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// false
	HasChildren    *bool                                                 `json:"HasChildren,omitempty" xml:"HasChildren,omitempty"`
	LinkSourceInfo *QueryDentryResponseBodySpaceRecentListLinkSourceInfo `json:"LinkSourceInfo,omitempty" xml:"LinkSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 测试组织/测试知识库/abc
	Path  *string     `json:"Path,omitempty" xml:"Path,omitempty"`
	Space interface{} `json:"Space,omitempty" xml:"Space,omitempty"`
	// example:
	//
	// bcd
	SpaceId         *string                                                `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	StatisticalInfo *QueryDentryResponseBodySpaceRecentListStatisticalInfo `json:"StatisticalInfo,omitempty" xml:"StatisticalInfo,omitempty" type:"Struct"`
	// example:
	//
	// 12345678
	UpdatedTime *int64                                         `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	Updater     *QueryDentryResponseBodySpaceRecentListUpdater `json:"Updater,omitempty" xml:"Updater,omitempty" type:"Struct"`
	// example:
	//
	// https://xxx.yy
	Url         *string                                            `json:"Url,omitempty" xml:"Url,omitempty"`
	VisitorInfo *QueryDentryResponseBodySpaceRecentListVisitorInfo `json:"VisitorInfo,omitempty" xml:"VisitorInfo,omitempty" type:"Struct"`
}

func (s QueryDentryResponseBodySpaceRecentList) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceRecentList) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceRecentList) GetContentType() *string {
	return s.ContentType
}

func (s *QueryDentryResponseBodySpaceRecentList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *QueryDentryResponseBodySpaceRecentList) GetCreator() *QueryDentryResponseBodySpaceRecentListCreator {
	return s.Creator
}

func (s *QueryDentryResponseBodySpaceRecentList) GetDentryId() *string {
	return s.DentryId
}

func (s *QueryDentryResponseBodySpaceRecentList) GetDentryType() *string {
	return s.DentryType
}

func (s *QueryDentryResponseBodySpaceRecentList) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *QueryDentryResponseBodySpaceRecentList) GetDocKey() *string {
	return s.DocKey
}

func (s *QueryDentryResponseBodySpaceRecentList) GetExtension() *string {
	return s.Extension
}

func (s *QueryDentryResponseBodySpaceRecentList) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *QueryDentryResponseBodySpaceRecentList) GetLinkSourceInfo() *QueryDentryResponseBodySpaceRecentListLinkSourceInfo {
	return s.LinkSourceInfo
}

func (s *QueryDentryResponseBodySpaceRecentList) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBodySpaceRecentList) GetPath() *string {
	return s.Path
}

func (s *QueryDentryResponseBodySpaceRecentList) GetSpace() interface{} {
	return s.Space
}

func (s *QueryDentryResponseBodySpaceRecentList) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentryResponseBodySpaceRecentList) GetStatisticalInfo() *QueryDentryResponseBodySpaceRecentListStatisticalInfo {
	return s.StatisticalInfo
}

func (s *QueryDentryResponseBodySpaceRecentList) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *QueryDentryResponseBodySpaceRecentList) GetUpdater() *QueryDentryResponseBodySpaceRecentListUpdater {
	return s.Updater
}

func (s *QueryDentryResponseBodySpaceRecentList) GetUrl() *string {
	return s.Url
}

func (s *QueryDentryResponseBodySpaceRecentList) GetVisitorInfo() *QueryDentryResponseBodySpaceRecentListVisitorInfo {
	return s.VisitorInfo
}

func (s *QueryDentryResponseBodySpaceRecentList) SetContentType(v string) *QueryDentryResponseBodySpaceRecentList {
	s.ContentType = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetCreatedTime(v int64) *QueryDentryResponseBodySpaceRecentList {
	s.CreatedTime = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetCreator(v *QueryDentryResponseBodySpaceRecentListCreator) *QueryDentryResponseBodySpaceRecentList {
	s.Creator = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetDentryId(v string) *QueryDentryResponseBodySpaceRecentList {
	s.DentryId = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetDentryType(v string) *QueryDentryResponseBodySpaceRecentList {
	s.DentryType = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetDentryUuid(v string) *QueryDentryResponseBodySpaceRecentList {
	s.DentryUuid = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetDocKey(v string) *QueryDentryResponseBodySpaceRecentList {
	s.DocKey = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetExtension(v string) *QueryDentryResponseBodySpaceRecentList {
	s.Extension = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetHasChildren(v bool) *QueryDentryResponseBodySpaceRecentList {
	s.HasChildren = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetLinkSourceInfo(v *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) *QueryDentryResponseBodySpaceRecentList {
	s.LinkSourceInfo = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetName(v string) *QueryDentryResponseBodySpaceRecentList {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetPath(v string) *QueryDentryResponseBodySpaceRecentList {
	s.Path = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetSpace(v interface{}) *QueryDentryResponseBodySpaceRecentList {
	s.Space = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetSpaceId(v string) *QueryDentryResponseBodySpaceRecentList {
	s.SpaceId = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetStatisticalInfo(v *QueryDentryResponseBodySpaceRecentListStatisticalInfo) *QueryDentryResponseBodySpaceRecentList {
	s.StatisticalInfo = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetUpdatedTime(v int64) *QueryDentryResponseBodySpaceRecentList {
	s.UpdatedTime = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetUpdater(v *QueryDentryResponseBodySpaceRecentListUpdater) *QueryDentryResponseBodySpaceRecentList {
	s.Updater = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetUrl(v string) *QueryDentryResponseBodySpaceRecentList {
	s.Url = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) SetVisitorInfo(v *QueryDentryResponseBodySpaceRecentListVisitorInfo) *QueryDentryResponseBodySpaceRecentList {
	s.VisitorInfo = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentList) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.LinkSourceInfo != nil {
		if err := s.LinkSourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.StatisticalInfo != nil {
		if err := s.StatisticalInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Updater != nil {
		if err := s.Updater.Validate(); err != nil {
			return err
		}
	}
	if s.VisitorInfo != nil {
		if err := s.VisitorInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDentryResponseBodySpaceRecentListCreator struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDentryResponseBodySpaceRecentListCreator) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceRecentListCreator) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceRecentListCreator) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBodySpaceRecentListCreator) GetUserId() *string {
	return s.UserId
}

func (s *QueryDentryResponseBodySpaceRecentListCreator) SetName(v string) *QueryDentryResponseBodySpaceRecentListCreator {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListCreator) SetUserId(v string) *QueryDentryResponseBodySpaceRecentListCreator {
	s.UserId = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListCreator) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceRecentListLinkSourceInfo struct {
	// example:
	//
	// docx
	Extension *string                                                      `json:"Extension,omitempty" xml:"Extension,omitempty"`
	IconUrl   *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl `json:"IconUrl,omitempty" xml:"IconUrl,omitempty" type:"Struct"`
	// example:
	//
	// def
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	LinkType *int64 `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	// example:
	//
	// def
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s QueryDentryResponseBodySpaceRecentListLinkSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceRecentListLinkSourceInfo) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) GetExtension() *string {
	return s.Extension
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) GetIconUrl() *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl {
	return s.IconUrl
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) GetId() *string {
	return s.Id
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) GetLinkType() *int64 {
	return s.LinkType
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) SetExtension(v string) *QueryDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.Extension = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) SetIconUrl(v *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) *QueryDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.IconUrl = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) SetId(v string) *QueryDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.Id = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) SetLinkType(v int64) *QueryDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.LinkType = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) SetSpaceId(v string) *QueryDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.SpaceId = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfo) Validate() error {
	if s.IconUrl != nil {
		if err := s.IconUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl struct {
	// example:
	//
	// gh
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// def
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) GetLine() *string {
	return s.Line
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) GetSmall() *string {
	return s.Small
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) SetLine(v string) *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl {
	s.Line = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) SetSmall(v string) *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl {
	s.Small = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceRecentListStatisticalInfo struct {
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s QueryDentryResponseBodySpaceRecentListStatisticalInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceRecentListStatisticalInfo) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceRecentListStatisticalInfo) GetWordCount() *int64 {
	return s.WordCount
}

func (s *QueryDentryResponseBodySpaceRecentListStatisticalInfo) SetWordCount(v int64) *QueryDentryResponseBodySpaceRecentListStatisticalInfo {
	s.WordCount = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListStatisticalInfo) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceRecentListUpdater struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDentryResponseBodySpaceRecentListUpdater) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceRecentListUpdater) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceRecentListUpdater) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBodySpaceRecentListUpdater) GetUserId() *string {
	return s.UserId
}

func (s *QueryDentryResponseBodySpaceRecentListUpdater) SetName(v string) *QueryDentryResponseBodySpaceRecentListUpdater {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListUpdater) SetUserId(v string) *QueryDentryResponseBodySpaceRecentListUpdater {
	s.UserId = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListUpdater) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceRecentListVisitorInfo struct {
	DentryActions []*string `json:"DentryActions,omitempty" xml:"DentryActions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RoleCode     *string   `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	SpaceActions []*string `json:"SpaceActions,omitempty" xml:"SpaceActions,omitempty" type:"Repeated"`
}

func (s QueryDentryResponseBodySpaceRecentListVisitorInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceRecentListVisitorInfo) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceRecentListVisitorInfo) GetDentryActions() []*string {
	return s.DentryActions
}

func (s *QueryDentryResponseBodySpaceRecentListVisitorInfo) GetRoleCode() *string {
	return s.RoleCode
}

func (s *QueryDentryResponseBodySpaceRecentListVisitorInfo) GetSpaceActions() []*string {
	return s.SpaceActions
}

func (s *QueryDentryResponseBodySpaceRecentListVisitorInfo) SetDentryActions(v []*string) *QueryDentryResponseBodySpaceRecentListVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListVisitorInfo) SetRoleCode(v string) *QueryDentryResponseBodySpaceRecentListVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListVisitorInfo) SetSpaceActions(v []*string) *QueryDentryResponseBodySpaceRecentListVisitorInfo {
	s.SpaceActions = v
	return s
}

func (s *QueryDentryResponseBodySpaceRecentListVisitorInfo) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodySpaceVisitorInfo struct {
	// example:
	//
	// GET_DENTRY
	DentryActions []*string `json:"DentryActions,omitempty" xml:"DentryActions,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// example:
	//
	// GET_SPACE
	SpaceActions []*string `json:"SpaceActions,omitempty" xml:"SpaceActions,omitempty" type:"Repeated"`
}

func (s QueryDentryResponseBodySpaceVisitorInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodySpaceVisitorInfo) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodySpaceVisitorInfo) GetDentryActions() []*string {
	return s.DentryActions
}

func (s *QueryDentryResponseBodySpaceVisitorInfo) GetRoleCode() *string {
	return s.RoleCode
}

func (s *QueryDentryResponseBodySpaceVisitorInfo) GetSpaceActions() []*string {
	return s.SpaceActions
}

func (s *QueryDentryResponseBodySpaceVisitorInfo) SetDentryActions(v []*string) *QueryDentryResponseBodySpaceVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *QueryDentryResponseBodySpaceVisitorInfo) SetRoleCode(v string) *QueryDentryResponseBodySpaceVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *QueryDentryResponseBodySpaceVisitorInfo) SetSpaceActions(v []*string) *QueryDentryResponseBodySpaceVisitorInfo {
	s.SpaceActions = v
	return s
}

func (s *QueryDentryResponseBodySpaceVisitorInfo) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodyUpdater struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDentryResponseBodyUpdater) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodyUpdater) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodyUpdater) GetName() *string {
	return s.Name
}

func (s *QueryDentryResponseBodyUpdater) GetUserId() *string {
	return s.UserId
}

func (s *QueryDentryResponseBodyUpdater) SetName(v string) *QueryDentryResponseBodyUpdater {
	s.Name = &v
	return s
}

func (s *QueryDentryResponseBodyUpdater) SetUserId(v string) *QueryDentryResponseBodyUpdater {
	s.UserId = &v
	return s
}

func (s *QueryDentryResponseBodyUpdater) Validate() error {
	return dara.Validate(s)
}

type QueryDentryResponseBodyVisitorInfo struct {
	DentryActions []*string `json:"DentryActions,omitempty" xml:"DentryActions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RoleCode     *string   `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	SpaceActions []*string `json:"SpaceActions,omitempty" xml:"SpaceActions,omitempty" type:"Repeated"`
}

func (s QueryDentryResponseBodyVisitorInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryResponseBodyVisitorInfo) GoString() string {
	return s.String()
}

func (s *QueryDentryResponseBodyVisitorInfo) GetDentryActions() []*string {
	return s.DentryActions
}

func (s *QueryDentryResponseBodyVisitorInfo) GetRoleCode() *string {
	return s.RoleCode
}

func (s *QueryDentryResponseBodyVisitorInfo) GetSpaceActions() []*string {
	return s.SpaceActions
}

func (s *QueryDentryResponseBodyVisitorInfo) SetDentryActions(v []*string) *QueryDentryResponseBodyVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *QueryDentryResponseBodyVisitorInfo) SetRoleCode(v string) *QueryDentryResponseBodyVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *QueryDentryResponseBodyVisitorInfo) SetSpaceActions(v []*string) *QueryDentryResponseBodyVisitorInfo {
	s.SpaceActions = v
	return s
}

func (s *QueryDentryResponseBodyVisitorInfo) Validate() error {
	return dara.Validate(s)
}
