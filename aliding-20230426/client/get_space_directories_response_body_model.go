// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpaceDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*GetSpaceDirectoriesResponseBodyChildren) *GetSpaceDirectoriesResponseBody
	GetChildren() []*GetSpaceDirectoriesResponseBodyChildren
	SetHasMore(v bool) *GetSpaceDirectoriesResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *GetSpaceDirectoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetSpaceDirectoriesResponseBody
	GetRequestId() *string
}

type GetSpaceDirectoriesResponseBody struct {
	Children []*GetSpaceDirectoriesResponseBodyChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 1296
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSpaceDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBody) GetChildren() []*GetSpaceDirectoriesResponseBodyChildren {
	return s.Children
}

func (s *GetSpaceDirectoriesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetSpaceDirectoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetSpaceDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSpaceDirectoriesResponseBody) SetChildren(v []*GetSpaceDirectoriesResponseBodyChildren) *GetSpaceDirectoriesResponseBody {
	s.Children = v
	return s
}

func (s *GetSpaceDirectoriesResponseBody) SetHasMore(v bool) *GetSpaceDirectoriesResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBody) SetNextToken(v string) *GetSpaceDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBody) SetRequestId(v string) *GetSpaceDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildren struct {
	// example:
	//
	// alidoc
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 12345678
	CreatedTime *int64                                          `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Creator     *GetSpaceDirectoriesResponseBodyChildrenCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
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
	HasChildren    *bool                                                  `json:"HasChildren,omitempty" xml:"HasChildren,omitempty"`
	LinkSourceInfo *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo `json:"LinkSourceInfo,omitempty" xml:"LinkSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 测试组织/测试知识库/abc
	Path  *string                                       `json:"Path,omitempty" xml:"Path,omitempty"`
	Space *GetSpaceDirectoriesResponseBodyChildrenSpace `json:"Space,omitempty" xml:"Space,omitempty" type:"Struct"`
	// example:
	//
	// bcd
	SpaceId         *string                                                 `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	StatisticalInfo *GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo `json:"StatisticalInfo,omitempty" xml:"StatisticalInfo,omitempty" type:"Struct"`
	// example:
	//
	// 12345678
	UpdatedTime *int64                                          `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	Updater     *GetSpaceDirectoriesResponseBodyChildrenUpdater `json:"Updater,omitempty" xml:"Updater,omitempty" type:"Struct"`
	// example:
	//
	// https://xxx.yy
	Url         *string                                             `json:"Url,omitempty" xml:"Url,omitempty"`
	VisitorInfo *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo `json:"VisitorInfo,omitempty" xml:"VisitorInfo,omitempty" type:"Struct"`
}

func (s GetSpaceDirectoriesResponseBodyChildren) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildren) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetContentType() *string {
	return s.ContentType
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetCreator() *GetSpaceDirectoriesResponseBodyChildrenCreator {
	return s.Creator
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetDentryId() *string {
	return s.DentryId
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetDentryType() *string {
	return s.DentryType
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetDocKey() *string {
	return s.DocKey
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetExtension() *string {
	return s.Extension
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetLinkSourceInfo() *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo {
	return s.LinkSourceInfo
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetName() *string {
	return s.Name
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetPath() *string {
	return s.Path
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetSpace() *GetSpaceDirectoriesResponseBodyChildrenSpace {
	return s.Space
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetSpaceId() *string {
	return s.SpaceId
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetStatisticalInfo() *GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo {
	return s.StatisticalInfo
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetUpdater() *GetSpaceDirectoriesResponseBodyChildrenUpdater {
	return s.Updater
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetUrl() *string {
	return s.Url
}

func (s *GetSpaceDirectoriesResponseBodyChildren) GetVisitorInfo() *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo {
	return s.VisitorInfo
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetContentType(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.ContentType = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetCreatedTime(v int64) *GetSpaceDirectoriesResponseBodyChildren {
	s.CreatedTime = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetCreator(v *GetSpaceDirectoriesResponseBodyChildrenCreator) *GetSpaceDirectoriesResponseBodyChildren {
	s.Creator = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetDentryId(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.DentryId = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetDentryType(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.DentryType = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetDentryUuid(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.DentryUuid = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetDocKey(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.DocKey = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetExtension(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.Extension = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetHasChildren(v bool) *GetSpaceDirectoriesResponseBodyChildren {
	s.HasChildren = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetLinkSourceInfo(v *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) *GetSpaceDirectoriesResponseBodyChildren {
	s.LinkSourceInfo = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetName(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.Name = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetPath(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.Path = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetSpace(v *GetSpaceDirectoriesResponseBodyChildrenSpace) *GetSpaceDirectoriesResponseBodyChildren {
	s.Space = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetSpaceId(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.SpaceId = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetStatisticalInfo(v *GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo) *GetSpaceDirectoriesResponseBodyChildren {
	s.StatisticalInfo = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetUpdatedTime(v int64) *GetSpaceDirectoriesResponseBodyChildren {
	s.UpdatedTime = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetUpdater(v *GetSpaceDirectoriesResponseBodyChildrenUpdater) *GetSpaceDirectoriesResponseBodyChildren {
	s.Updater = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetUrl(v string) *GetSpaceDirectoriesResponseBodyChildren {
	s.Url = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) SetVisitorInfo(v *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) *GetSpaceDirectoriesResponseBodyChildren {
	s.VisitorInfo = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildren) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenCreator struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenCreator) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenCreator) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenCreator) GetName() *string {
	return s.Name
}

func (s *GetSpaceDirectoriesResponseBodyChildrenCreator) GetUserId() *string {
	return s.UserId
}

func (s *GetSpaceDirectoriesResponseBodyChildrenCreator) SetName(v string) *GetSpaceDirectoriesResponseBodyChildrenCreator {
	s.Name = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenCreator) SetUserId(v string) *GetSpaceDirectoriesResponseBodyChildrenCreator {
	s.UserId = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenCreator) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo struct {
	// example:
	//
	// docx
	Extension *string                                                       `json:"Extension,omitempty" xml:"Extension,omitempty"`
	IconUrl   *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl `json:"IconUrl,omitempty" xml:"IconUrl,omitempty" type:"Struct"`
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

func (s GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) GetExtension() *string {
	return s.Extension
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) GetIconUrl() *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl {
	return s.IconUrl
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) GetId() *string {
	return s.Id
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) GetLinkType() *int64 {
	return s.LinkType
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) GetSpaceId() *string {
	return s.SpaceId
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) SetExtension(v string) *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo {
	s.Extension = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) SetIconUrl(v *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo {
	s.IconUrl = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) SetId(v string) *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo {
	s.Id = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) SetLinkType(v int64) *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo {
	s.LinkType = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) SetSpaceId(v string) *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo {
	s.SpaceId = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfo) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl struct {
	// example:
	//
	// gh
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// def
	Small *string `json:"Small,omitempty" xml:"Small,omitempty"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) GetLine() *string {
	return s.Line
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) GetSmall() *string {
	return s.Small
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) SetLine(v string) *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl {
	s.Line = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) SetSmall(v string) *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl {
	s.Small = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenLinkSourceInfoIconUrl) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenSpace struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1xxxxx.png
	Cover *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// example:
	//
	// 这是简介
	Description *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	HdIconVO    *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO `json:"HdIconVO,omitempty" xml:"HdIconVO,omitempty" type:"Struct"`
	IconVO      *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO   `json:"IconVO,omitempty" xml:"IconVO,omitempty" type:"Struct"`
	// example:
	//
	// n9XJxxxxx
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 测试知识库
	Owner      *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	RecentList []interface{}                                      `json:"RecentList,omitempty" xml:"RecentList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://alidocs.dingtalk.com/i/spaces/n9XJ*******Xy/overview
	Url         *string                                                  `json:"Url,omitempty" xml:"Url,omitempty"`
	VisitorInfo *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo `json:"VisitorInfo,omitempty" xml:"VisitorInfo,omitempty" type:"Struct"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpace) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpace) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetCover() *string {
	return s.Cover
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetDescription() *string {
	return s.Description
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetHdIconVO() *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO {
	return s.HdIconVO
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetIconVO() *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO {
	return s.IconVO
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetId() *string {
	return s.Id
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetName() *string {
	return s.Name
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetOwner() *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner {
	return s.Owner
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetRecentList() []interface{} {
	return s.RecentList
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetType() *int32 {
	return s.Type
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetUrl() *string {
	return s.Url
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) GetVisitorInfo() *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo {
	return s.VisitorInfo
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetCover(v string) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.Cover = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetDescription(v string) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.Description = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetHdIconVO(v *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.HdIconVO = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetIconVO(v *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.IconVO = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetId(v string) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.Id = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetName(v string) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.Name = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetOwner(v *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.Owner = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetRecentList(v []interface{}) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.RecentList = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetType(v int32) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.Type = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetUrl(v string) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.Url = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) SetVisitorInfo(v *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) *GetSpaceDirectoriesResponseBodyChildrenSpace {
	s.VisitorInfo = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpace) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/xxxxx.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) GetIcon() *string {
	return s.Icon
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) GetType() *string {
	return s.Type
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) SetIcon(v string) *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO {
	s.Icon = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) SetType(v string) *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO {
	s.Type = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceHdIconVO) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/xxxxx.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) GetIcon() *string {
	return s.Icon
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) GetType() *string {
	return s.Type
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) SetIcon(v string) *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO {
	s.Icon = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) SetType(v string) *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO {
	s.Type = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceIconVO) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenSpaceOwner struct {
	// example:
	//
	// 小钉
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) GetName() *string {
	return s.Name
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) GetUserId() *string {
	return s.UserId
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) SetName(v string) *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner {
	s.Name = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) SetUserId(v string) *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner {
	s.UserId = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceOwner) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo struct {
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

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) GetDentryActions() []*string {
	return s.DentryActions
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) GetRoleCode() *string {
	return s.RoleCode
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) GetSpaceActions() []*string {
	return s.SpaceActions
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) SetDentryActions(v []*string) *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) SetRoleCode(v string) *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) SetSpaceActions(v []*string) *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo {
	s.SpaceActions = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenSpaceVisitorInfo) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo struct {
	// WordCount
	//
	// example:
	//
	// 10
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo) GetWordCount() *int64 {
	return s.WordCount
}

func (s *GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo) SetWordCount(v int64) *GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo {
	s.WordCount = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenStatisticalInfo) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenUpdater struct {
	// example:
	//
	// hello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenUpdater) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenUpdater) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenUpdater) GetName() *string {
	return s.Name
}

func (s *GetSpaceDirectoriesResponseBodyChildrenUpdater) GetUserId() *string {
	return s.UserId
}

func (s *GetSpaceDirectoriesResponseBodyChildrenUpdater) SetName(v string) *GetSpaceDirectoriesResponseBodyChildrenUpdater {
	s.Name = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenUpdater) SetUserId(v string) *GetSpaceDirectoriesResponseBodyChildrenUpdater {
	s.UserId = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenUpdater) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesResponseBodyChildrenVisitorInfo struct {
	DentryActions []*string `json:"DentryActions,omitempty" xml:"DentryActions,omitempty" type:"Repeated"`
	// example:
	//
	// OWNER
	RoleCode     *string   `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	SpaceActions []*string `json:"SpaceActions,omitempty" xml:"SpaceActions,omitempty" type:"Repeated"`
}

func (s GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) GetDentryActions() []*string {
	return s.DentryActions
}

func (s *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) GetRoleCode() *string {
	return s.RoleCode
}

func (s *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) GetSpaceActions() []*string {
	return s.SpaceActions
}

func (s *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) SetDentryActions(v []*string) *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo {
	s.DentryActions = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) SetRoleCode(v string) *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo {
	s.RoleCode = &v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) SetSpaceActions(v []*string) *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo {
	s.SpaceActions = v
	return s
}

func (s *GetSpaceDirectoriesResponseBodyChildrenVisitorInfo) Validate() error {
	return dara.Validate(s)
}
