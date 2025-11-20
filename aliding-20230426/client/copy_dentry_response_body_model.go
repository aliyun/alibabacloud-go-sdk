// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *CopyDentryResponseBody
	GetContentType() *string
	SetCreatedTime(v int64) *CopyDentryResponseBody
	GetCreatedTime() *int64
	SetCreator(v map[string]interface{}) *CopyDentryResponseBody
	GetCreator() map[string]interface{}
	SetDentryId(v string) *CopyDentryResponseBody
	GetDentryId() *string
	SetDentryType(v string) *CopyDentryResponseBody
	GetDentryType() *string
	SetDentryUuid(v string) *CopyDentryResponseBody
	GetDentryUuid() *string
	SetDocKey(v string) *CopyDentryResponseBody
	GetDocKey() *string
	SetExtension(v string) *CopyDentryResponseBody
	GetExtension() *string
	SetHasChildren(v bool) *CopyDentryResponseBody
	GetHasChildren() *bool
	SetLinkSourceInfo(v *CopyDentryResponseBodyLinkSourceInfo) *CopyDentryResponseBody
	GetLinkSourceInfo() *CopyDentryResponseBodyLinkSourceInfo
	SetName(v string) *CopyDentryResponseBody
	GetName() *string
	SetPath(v string) *CopyDentryResponseBody
	GetPath() *string
	SetRequestId(v string) *CopyDentryResponseBody
	GetRequestId() *string
	SetSpace(v *CopyDentryResponseBodySpace) *CopyDentryResponseBody
	GetSpace() *CopyDentryResponseBodySpace
	SetSpaceId(v string) *CopyDentryResponseBody
	GetSpaceId() *string
	SetUpdatedTime(v int64) *CopyDentryResponseBody
	GetUpdatedTime() *int64
	SetUpdater(v map[string]interface{}) *CopyDentryResponseBody
	GetUpdater() map[string]interface{}
	SetUrl(v string) *CopyDentryResponseBody
	GetUrl() *string
	SetVendorRequestId(v string) *CopyDentryResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CopyDentryResponseBody
	GetVendorType() *string
	SetVisitorInfo(v map[string]interface{}) *CopyDentryResponseBody
	GetVisitorInfo() map[string]interface{}
}

type CopyDentryResponseBody struct {
	// example:
	//
	// alidoc
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// example:
	//
	// 12345678
	CreatedTime *int64                 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Creator     map[string]interface{} `json:"creator,omitempty" xml:"creator,omitempty"`
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
	HasChildren    *bool                                 `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	LinkSourceInfo *CopyDentryResponseBodyLinkSourceInfo `json:"linkSourceInfo,omitempty" xml:"linkSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// hello
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 测试组织/测试知识库/abc
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Space     *CopyDentryResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
	// example:
	//
	// bcd
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// 12345678
	UpdatedTime *int64 `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// example:
	//
	// {\\"UserId\\": \\"353851\\", \\"Name\\": u\\"\\u848b\\u7fbd\\u4e2d\\"}
	Updater map[string]interface{} `json:"updater,omitempty" xml:"updater,omitempty"`
	// example:
	//
	// https://xxx.yy
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType  *string                `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	VisitorInfo map[string]interface{} `json:"visitorInfo,omitempty" xml:"visitorInfo,omitempty"`
}

func (s CopyDentryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *CopyDentryResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *CopyDentryResponseBody) GetCreator() map[string]interface{} {
	return s.Creator
}

func (s *CopyDentryResponseBody) GetDentryId() *string {
	return s.DentryId
}

func (s *CopyDentryResponseBody) GetDentryType() *string {
	return s.DentryType
}

func (s *CopyDentryResponseBody) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *CopyDentryResponseBody) GetDocKey() *string {
	return s.DocKey
}

func (s *CopyDentryResponseBody) GetExtension() *string {
	return s.Extension
}

func (s *CopyDentryResponseBody) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *CopyDentryResponseBody) GetLinkSourceInfo() *CopyDentryResponseBodyLinkSourceInfo {
	return s.LinkSourceInfo
}

func (s *CopyDentryResponseBody) GetName() *string {
	return s.Name
}

func (s *CopyDentryResponseBody) GetPath() *string {
	return s.Path
}

func (s *CopyDentryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyDentryResponseBody) GetSpace() *CopyDentryResponseBodySpace {
	return s.Space
}

func (s *CopyDentryResponseBody) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CopyDentryResponseBody) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *CopyDentryResponseBody) GetUpdater() map[string]interface{} {
	return s.Updater
}

func (s *CopyDentryResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CopyDentryResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CopyDentryResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CopyDentryResponseBody) GetVisitorInfo() map[string]interface{} {
	return s.VisitorInfo
}

func (s *CopyDentryResponseBody) SetContentType(v string) *CopyDentryResponseBody {
	s.ContentType = &v
	return s
}

func (s *CopyDentryResponseBody) SetCreatedTime(v int64) *CopyDentryResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CopyDentryResponseBody) SetCreator(v map[string]interface{}) *CopyDentryResponseBody {
	s.Creator = v
	return s
}

func (s *CopyDentryResponseBody) SetDentryId(v string) *CopyDentryResponseBody {
	s.DentryId = &v
	return s
}

func (s *CopyDentryResponseBody) SetDentryType(v string) *CopyDentryResponseBody {
	s.DentryType = &v
	return s
}

func (s *CopyDentryResponseBody) SetDentryUuid(v string) *CopyDentryResponseBody {
	s.DentryUuid = &v
	return s
}

func (s *CopyDentryResponseBody) SetDocKey(v string) *CopyDentryResponseBody {
	s.DocKey = &v
	return s
}

func (s *CopyDentryResponseBody) SetExtension(v string) *CopyDentryResponseBody {
	s.Extension = &v
	return s
}

func (s *CopyDentryResponseBody) SetHasChildren(v bool) *CopyDentryResponseBody {
	s.HasChildren = &v
	return s
}

func (s *CopyDentryResponseBody) SetLinkSourceInfo(v *CopyDentryResponseBodyLinkSourceInfo) *CopyDentryResponseBody {
	s.LinkSourceInfo = v
	return s
}

func (s *CopyDentryResponseBody) SetName(v string) *CopyDentryResponseBody {
	s.Name = &v
	return s
}

func (s *CopyDentryResponseBody) SetPath(v string) *CopyDentryResponseBody {
	s.Path = &v
	return s
}

func (s *CopyDentryResponseBody) SetRequestId(v string) *CopyDentryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyDentryResponseBody) SetSpace(v *CopyDentryResponseBodySpace) *CopyDentryResponseBody {
	s.Space = v
	return s
}

func (s *CopyDentryResponseBody) SetSpaceId(v string) *CopyDentryResponseBody {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryResponseBody) SetUpdatedTime(v int64) *CopyDentryResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *CopyDentryResponseBody) SetUpdater(v map[string]interface{}) *CopyDentryResponseBody {
	s.Updater = v
	return s
}

func (s *CopyDentryResponseBody) SetUrl(v string) *CopyDentryResponseBody {
	s.Url = &v
	return s
}

func (s *CopyDentryResponseBody) SetVendorRequestId(v string) *CopyDentryResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CopyDentryResponseBody) SetVendorType(v string) *CopyDentryResponseBody {
	s.VendorType = &v
	return s
}

func (s *CopyDentryResponseBody) SetVisitorInfo(v map[string]interface{}) *CopyDentryResponseBody {
	s.VisitorInfo = v
	return s
}

func (s *CopyDentryResponseBody) Validate() error {
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
	return nil
}

type CopyDentryResponseBodyLinkSourceInfo struct {
	// example:
	//
	// docx
	Extension *string                                      `json:"Extension,omitempty" xml:"Extension,omitempty"`
	IconUrl   *CopyDentryResponseBodyLinkSourceInfoIconUrl `json:"IconUrl,omitempty" xml:"IconUrl,omitempty" type:"Struct"`
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

func (s CopyDentryResponseBodyLinkSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodyLinkSourceInfo) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodyLinkSourceInfo) GetExtension() *string {
	return s.Extension
}

func (s *CopyDentryResponseBodyLinkSourceInfo) GetIconUrl() *CopyDentryResponseBodyLinkSourceInfoIconUrl {
	return s.IconUrl
}

func (s *CopyDentryResponseBodyLinkSourceInfo) GetId() *string {
	return s.Id
}

func (s *CopyDentryResponseBodyLinkSourceInfo) GetLinkType() *int64 {
	return s.LinkType
}

func (s *CopyDentryResponseBodyLinkSourceInfo) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CopyDentryResponseBodyLinkSourceInfo) SetExtension(v string) *CopyDentryResponseBodyLinkSourceInfo {
	s.Extension = &v
	return s
}

func (s *CopyDentryResponseBodyLinkSourceInfo) SetIconUrl(v *CopyDentryResponseBodyLinkSourceInfoIconUrl) *CopyDentryResponseBodyLinkSourceInfo {
	s.IconUrl = v
	return s
}

func (s *CopyDentryResponseBodyLinkSourceInfo) SetId(v string) *CopyDentryResponseBodyLinkSourceInfo {
	s.Id = &v
	return s
}

func (s *CopyDentryResponseBodyLinkSourceInfo) SetLinkType(v int64) *CopyDentryResponseBodyLinkSourceInfo {
	s.LinkType = &v
	return s
}

func (s *CopyDentryResponseBodyLinkSourceInfo) SetSpaceId(v string) *CopyDentryResponseBodyLinkSourceInfo {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryResponseBodyLinkSourceInfo) Validate() error {
	if s.IconUrl != nil {
		if err := s.IconUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyDentryResponseBodyLinkSourceInfoIconUrl struct {
	// example:
	//
	// gh
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// def
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s CopyDentryResponseBodyLinkSourceInfoIconUrl) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodyLinkSourceInfoIconUrl) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodyLinkSourceInfoIconUrl) GetLine() *string {
	return s.Line
}

func (s *CopyDentryResponseBodyLinkSourceInfoIconUrl) GetSmall() *string {
	return s.Small
}

func (s *CopyDentryResponseBodyLinkSourceInfoIconUrl) SetLine(v string) *CopyDentryResponseBodyLinkSourceInfoIconUrl {
	s.Line = &v
	return s
}

func (s *CopyDentryResponseBodyLinkSourceInfoIconUrl) SetSmall(v string) *CopyDentryResponseBodyLinkSourceInfoIconUrl {
	s.Small = &v
	return s
}

func (s *CopyDentryResponseBodyLinkSourceInfoIconUrl) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpace struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1xxxxx.png
	Cover *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// example:
	//
	// 这是简介
	Description *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	HdIconVO    *CopyDentryResponseBodySpaceHdIconVO `json:"HdIconVO,omitempty" xml:"HdIconVO,omitempty" type:"Struct"`
	IconVO      *CopyDentryResponseBodySpaceIconVO   `json:"IconVO,omitempty" xml:"IconVO,omitempty" type:"Struct"`
	// example:
	//
	// n9XJxxxxx
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 测试知识库
	Owner      *CopyDentryResponseBodySpaceOwner        `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	RecentList []*CopyDentryResponseBodySpaceRecentList `json:"RecentList,omitempty" xml:"RecentList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://alidocs.dingtalk.com/i/spaces/n9XJ*******Xy/overview
	Url         *string                                 `json:"Url,omitempty" xml:"Url,omitempty"`
	VisitorInfo *CopyDentryResponseBodySpaceVisitorInfo `json:"VisitorInfo,omitempty" xml:"VisitorInfo,omitempty" type:"Struct"`
}

func (s CopyDentryResponseBodySpace) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpace) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpace) GetCover() *string {
	return s.Cover
}

func (s *CopyDentryResponseBodySpace) GetDescription() *string {
	return s.Description
}

func (s *CopyDentryResponseBodySpace) GetHdIconVO() *CopyDentryResponseBodySpaceHdIconVO {
	return s.HdIconVO
}

func (s *CopyDentryResponseBodySpace) GetIconVO() *CopyDentryResponseBodySpaceIconVO {
	return s.IconVO
}

func (s *CopyDentryResponseBodySpace) GetId() *string {
	return s.Id
}

func (s *CopyDentryResponseBodySpace) GetName() *string {
	return s.Name
}

func (s *CopyDentryResponseBodySpace) GetOwner() *CopyDentryResponseBodySpaceOwner {
	return s.Owner
}

func (s *CopyDentryResponseBodySpace) GetRecentList() []*CopyDentryResponseBodySpaceRecentList {
	return s.RecentList
}

func (s *CopyDentryResponseBodySpace) GetType() *int32 {
	return s.Type
}

func (s *CopyDentryResponseBodySpace) GetUrl() *string {
	return s.Url
}

func (s *CopyDentryResponseBodySpace) GetVisitorInfo() *CopyDentryResponseBodySpaceVisitorInfo {
	return s.VisitorInfo
}

func (s *CopyDentryResponseBodySpace) SetCover(v string) *CopyDentryResponseBodySpace {
	s.Cover = &v
	return s
}

func (s *CopyDentryResponseBodySpace) SetDescription(v string) *CopyDentryResponseBodySpace {
	s.Description = &v
	return s
}

func (s *CopyDentryResponseBodySpace) SetHdIconVO(v *CopyDentryResponseBodySpaceHdIconVO) *CopyDentryResponseBodySpace {
	s.HdIconVO = v
	return s
}

func (s *CopyDentryResponseBodySpace) SetIconVO(v *CopyDentryResponseBodySpaceIconVO) *CopyDentryResponseBodySpace {
	s.IconVO = v
	return s
}

func (s *CopyDentryResponseBodySpace) SetId(v string) *CopyDentryResponseBodySpace {
	s.Id = &v
	return s
}

func (s *CopyDentryResponseBodySpace) SetName(v string) *CopyDentryResponseBodySpace {
	s.Name = &v
	return s
}

func (s *CopyDentryResponseBodySpace) SetOwner(v *CopyDentryResponseBodySpaceOwner) *CopyDentryResponseBodySpace {
	s.Owner = v
	return s
}

func (s *CopyDentryResponseBodySpace) SetRecentList(v []*CopyDentryResponseBodySpaceRecentList) *CopyDentryResponseBodySpace {
	s.RecentList = v
	return s
}

func (s *CopyDentryResponseBodySpace) SetType(v int32) *CopyDentryResponseBodySpace {
	s.Type = &v
	return s
}

func (s *CopyDentryResponseBodySpace) SetUrl(v string) *CopyDentryResponseBodySpace {
	s.Url = &v
	return s
}

func (s *CopyDentryResponseBodySpace) SetVisitorInfo(v *CopyDentryResponseBodySpaceVisitorInfo) *CopyDentryResponseBodySpace {
	s.VisitorInfo = v
	return s
}

func (s *CopyDentryResponseBodySpace) Validate() error {
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

type CopyDentryResponseBodySpaceHdIconVO struct {
	// example:
	//
	// http://
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CopyDentryResponseBodySpaceHdIconVO) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceHdIconVO) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceHdIconVO) GetIcon() *string {
	return s.Icon
}

func (s *CopyDentryResponseBodySpaceHdIconVO) GetType() *string {
	return s.Type
}

func (s *CopyDentryResponseBodySpaceHdIconVO) SetIcon(v string) *CopyDentryResponseBodySpaceHdIconVO {
	s.Icon = &v
	return s
}

func (s *CopyDentryResponseBodySpaceHdIconVO) SetType(v string) *CopyDentryResponseBodySpaceHdIconVO {
	s.Type = &v
	return s
}

func (s *CopyDentryResponseBodySpaceHdIconVO) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceIconVO struct {
	// example:
	//
	// http://
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CopyDentryResponseBodySpaceIconVO) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceIconVO) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceIconVO) GetIcon() *string {
	return s.Icon
}

func (s *CopyDentryResponseBodySpaceIconVO) GetType() *string {
	return s.Type
}

func (s *CopyDentryResponseBodySpaceIconVO) SetIcon(v string) *CopyDentryResponseBodySpaceIconVO {
	s.Icon = &v
	return s
}

func (s *CopyDentryResponseBodySpaceIconVO) SetType(v string) *CopyDentryResponseBodySpaceIconVO {
	s.Type = &v
	return s
}

func (s *CopyDentryResponseBodySpaceIconVO) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceOwner struct {
	// example:
	//
	// 小钉
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CopyDentryResponseBodySpaceOwner) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceOwner) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceOwner) GetName() *string {
	return s.Name
}

func (s *CopyDentryResponseBodySpaceOwner) GetUserId() *string {
	return s.UserId
}

func (s *CopyDentryResponseBodySpaceOwner) SetName(v string) *CopyDentryResponseBodySpaceOwner {
	s.Name = &v
	return s
}

func (s *CopyDentryResponseBodySpaceOwner) SetUserId(v string) *CopyDentryResponseBodySpaceOwner {
	s.UserId = &v
	return s
}

func (s *CopyDentryResponseBodySpaceOwner) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceRecentList struct {
	// example:
	//
	// alidoc
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 12345678
	CreatedTime *int64                                        `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Creator     *CopyDentryResponseBodySpaceRecentListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
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
	HasChildren    *bool                                                `json:"HasChildren,omitempty" xml:"HasChildren,omitempty"`
	LinkSourceInfo *CopyDentryResponseBodySpaceRecentListLinkSourceInfo `json:"LinkSourceInfo,omitempty" xml:"LinkSourceInfo,omitempty" type:"Struct"`
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
	SpaceId         *string                                               `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	StatisticalInfo *CopyDentryResponseBodySpaceRecentListStatisticalInfo `json:"StatisticalInfo,omitempty" xml:"StatisticalInfo,omitempty" type:"Struct"`
	// example:
	//
	// 12345678
	UpdatedTime *int64                                        `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	Updater     *CopyDentryResponseBodySpaceRecentListUpdater `json:"Updater,omitempty" xml:"Updater,omitempty" type:"Struct"`
	// example:
	//
	// https://xxx.yy
	Url         *string                                           `json:"Url,omitempty" xml:"Url,omitempty"`
	VisitorInfo *CopyDentryResponseBodySpaceRecentListVisitorInfo `json:"VisitorInfo,omitempty" xml:"VisitorInfo,omitempty" type:"Struct"`
}

func (s CopyDentryResponseBodySpaceRecentList) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceRecentList) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceRecentList) GetContentType() *string {
	return s.ContentType
}

func (s *CopyDentryResponseBodySpaceRecentList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *CopyDentryResponseBodySpaceRecentList) GetCreator() *CopyDentryResponseBodySpaceRecentListCreator {
	return s.Creator
}

func (s *CopyDentryResponseBodySpaceRecentList) GetDentryId() *string {
	return s.DentryId
}

func (s *CopyDentryResponseBodySpaceRecentList) GetDentryType() *string {
	return s.DentryType
}

func (s *CopyDentryResponseBodySpaceRecentList) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *CopyDentryResponseBodySpaceRecentList) GetDocKey() *string {
	return s.DocKey
}

func (s *CopyDentryResponseBodySpaceRecentList) GetExtension() *string {
	return s.Extension
}

func (s *CopyDentryResponseBodySpaceRecentList) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *CopyDentryResponseBodySpaceRecentList) GetLinkSourceInfo() *CopyDentryResponseBodySpaceRecentListLinkSourceInfo {
	return s.LinkSourceInfo
}

func (s *CopyDentryResponseBodySpaceRecentList) GetName() *string {
	return s.Name
}

func (s *CopyDentryResponseBodySpaceRecentList) GetPath() *string {
	return s.Path
}

func (s *CopyDentryResponseBodySpaceRecentList) GetSpace() interface{} {
	return s.Space
}

func (s *CopyDentryResponseBodySpaceRecentList) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CopyDentryResponseBodySpaceRecentList) GetStatisticalInfo() *CopyDentryResponseBodySpaceRecentListStatisticalInfo {
	return s.StatisticalInfo
}

func (s *CopyDentryResponseBodySpaceRecentList) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *CopyDentryResponseBodySpaceRecentList) GetUpdater() *CopyDentryResponseBodySpaceRecentListUpdater {
	return s.Updater
}

func (s *CopyDentryResponseBodySpaceRecentList) GetUrl() *string {
	return s.Url
}

func (s *CopyDentryResponseBodySpaceRecentList) GetVisitorInfo() *CopyDentryResponseBodySpaceRecentListVisitorInfo {
	return s.VisitorInfo
}

func (s *CopyDentryResponseBodySpaceRecentList) SetContentType(v string) *CopyDentryResponseBodySpaceRecentList {
	s.ContentType = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetCreatedTime(v int64) *CopyDentryResponseBodySpaceRecentList {
	s.CreatedTime = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetCreator(v *CopyDentryResponseBodySpaceRecentListCreator) *CopyDentryResponseBodySpaceRecentList {
	s.Creator = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetDentryId(v string) *CopyDentryResponseBodySpaceRecentList {
	s.DentryId = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetDentryType(v string) *CopyDentryResponseBodySpaceRecentList {
	s.DentryType = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetDentryUuid(v string) *CopyDentryResponseBodySpaceRecentList {
	s.DentryUuid = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetDocKey(v string) *CopyDentryResponseBodySpaceRecentList {
	s.DocKey = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetExtension(v string) *CopyDentryResponseBodySpaceRecentList {
	s.Extension = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetHasChildren(v bool) *CopyDentryResponseBodySpaceRecentList {
	s.HasChildren = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetLinkSourceInfo(v *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) *CopyDentryResponseBodySpaceRecentList {
	s.LinkSourceInfo = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetName(v string) *CopyDentryResponseBodySpaceRecentList {
	s.Name = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetPath(v string) *CopyDentryResponseBodySpaceRecentList {
	s.Path = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetSpace(v interface{}) *CopyDentryResponseBodySpaceRecentList {
	s.Space = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetSpaceId(v string) *CopyDentryResponseBodySpaceRecentList {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetStatisticalInfo(v *CopyDentryResponseBodySpaceRecentListStatisticalInfo) *CopyDentryResponseBodySpaceRecentList {
	s.StatisticalInfo = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetUpdatedTime(v int64) *CopyDentryResponseBodySpaceRecentList {
	s.UpdatedTime = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetUpdater(v *CopyDentryResponseBodySpaceRecentListUpdater) *CopyDentryResponseBodySpaceRecentList {
	s.Updater = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetUrl(v string) *CopyDentryResponseBodySpaceRecentList {
	s.Url = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) SetVisitorInfo(v *CopyDentryResponseBodySpaceRecentListVisitorInfo) *CopyDentryResponseBodySpaceRecentList {
	s.VisitorInfo = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentList) Validate() error {
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

type CopyDentryResponseBodySpaceRecentListCreator struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CopyDentryResponseBodySpaceRecentListCreator) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceRecentListCreator) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceRecentListCreator) GetName() *string {
	return s.Name
}

func (s *CopyDentryResponseBodySpaceRecentListCreator) GetUserId() *string {
	return s.UserId
}

func (s *CopyDentryResponseBodySpaceRecentListCreator) SetName(v string) *CopyDentryResponseBodySpaceRecentListCreator {
	s.Name = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListCreator) SetUserId(v string) *CopyDentryResponseBodySpaceRecentListCreator {
	s.UserId = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListCreator) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceRecentListLinkSourceInfo struct {
	// example:
	//
	// docx
	Extension *string                                                     `json:"Extension,omitempty" xml:"Extension,omitempty"`
	IconUrl   *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl `json:"IconUrl,omitempty" xml:"IconUrl,omitempty" type:"Struct"`
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

func (s CopyDentryResponseBodySpaceRecentListLinkSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceRecentListLinkSourceInfo) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) GetExtension() *string {
	return s.Extension
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) GetIconUrl() *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl {
	return s.IconUrl
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) GetId() *string {
	return s.Id
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) GetLinkType() *int64 {
	return s.LinkType
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) SetExtension(v string) *CopyDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.Extension = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) SetIconUrl(v *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) *CopyDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.IconUrl = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) SetId(v string) *CopyDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.Id = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) SetLinkType(v int64) *CopyDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.LinkType = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) SetSpaceId(v string) *CopyDentryResponseBodySpaceRecentListLinkSourceInfo {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfo) Validate() error {
	if s.IconUrl != nil {
		if err := s.IconUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl struct {
	// example:
	//
	// gh
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// def
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) GetLine() *string {
	return s.Line
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) GetSmall() *string {
	return s.Small
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) SetLine(v string) *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl {
	s.Line = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) SetSmall(v string) *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl {
	s.Small = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListLinkSourceInfoIconUrl) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceRecentListStatisticalInfo struct {
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s CopyDentryResponseBodySpaceRecentListStatisticalInfo) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceRecentListStatisticalInfo) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceRecentListStatisticalInfo) GetWordCount() *int64 {
	return s.WordCount
}

func (s *CopyDentryResponseBodySpaceRecentListStatisticalInfo) SetWordCount(v int64) *CopyDentryResponseBodySpaceRecentListStatisticalInfo {
	s.WordCount = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListStatisticalInfo) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceRecentListUpdater struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CopyDentryResponseBodySpaceRecentListUpdater) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceRecentListUpdater) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceRecentListUpdater) GetName() *string {
	return s.Name
}

func (s *CopyDentryResponseBodySpaceRecentListUpdater) GetUserId() *string {
	return s.UserId
}

func (s *CopyDentryResponseBodySpaceRecentListUpdater) SetName(v string) *CopyDentryResponseBodySpaceRecentListUpdater {
	s.Name = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListUpdater) SetUserId(v string) *CopyDentryResponseBodySpaceRecentListUpdater {
	s.UserId = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListUpdater) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceRecentListVisitorInfo struct {
	DentryActions []*string `json:"DentryActions,omitempty" xml:"DentryActions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RoleCode     *string   `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	SpaceActions []*string `json:"SpaceActions,omitempty" xml:"SpaceActions,omitempty" type:"Repeated"`
}

func (s CopyDentryResponseBodySpaceRecentListVisitorInfo) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceRecentListVisitorInfo) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceRecentListVisitorInfo) GetDentryActions() []*string {
	return s.DentryActions
}

func (s *CopyDentryResponseBodySpaceRecentListVisitorInfo) GetRoleCode() *string {
	return s.RoleCode
}

func (s *CopyDentryResponseBodySpaceRecentListVisitorInfo) GetSpaceActions() []*string {
	return s.SpaceActions
}

func (s *CopyDentryResponseBodySpaceRecentListVisitorInfo) SetDentryActions(v []*string) *CopyDentryResponseBodySpaceRecentListVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListVisitorInfo) SetRoleCode(v string) *CopyDentryResponseBodySpaceRecentListVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListVisitorInfo) SetSpaceActions(v []*string) *CopyDentryResponseBodySpaceRecentListVisitorInfo {
	s.SpaceActions = v
	return s
}

func (s *CopyDentryResponseBodySpaceRecentListVisitorInfo) Validate() error {
	return dara.Validate(s)
}

type CopyDentryResponseBodySpaceVisitorInfo struct {
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

func (s CopyDentryResponseBodySpaceVisitorInfo) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryResponseBodySpaceVisitorInfo) GoString() string {
	return s.String()
}

func (s *CopyDentryResponseBodySpaceVisitorInfo) GetDentryActions() []*string {
	return s.DentryActions
}

func (s *CopyDentryResponseBodySpaceVisitorInfo) GetRoleCode() *string {
	return s.RoleCode
}

func (s *CopyDentryResponseBodySpaceVisitorInfo) GetSpaceActions() []*string {
	return s.SpaceActions
}

func (s *CopyDentryResponseBodySpaceVisitorInfo) SetDentryActions(v []*string) *CopyDentryResponseBodySpaceVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *CopyDentryResponseBodySpaceVisitorInfo) SetRoleCode(v string) *CopyDentryResponseBodySpaceVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *CopyDentryResponseBodySpaceVisitorInfo) SetSpaceActions(v []*string) *CopyDentryResponseBodySpaceVisitorInfo {
	s.SpaceActions = v
	return s
}

func (s *CopyDentryResponseBodySpaceVisitorInfo) Validate() error {
	return dara.Validate(s)
}
