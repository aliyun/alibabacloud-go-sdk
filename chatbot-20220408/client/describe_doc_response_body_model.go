// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *DescribeDocResponseBody
	GetBizCode() *string
	SetCategoryId(v int64) *DescribeDocResponseBody
	GetCategoryId() *int64
	SetConfig(v string) *DescribeDocResponseBody
	GetConfig() *string
	SetCreateTime(v string) *DescribeDocResponseBody
	GetCreateTime() *string
	SetCreateUserId(v int64) *DescribeDocResponseBody
	GetCreateUserId() *int64
	SetCreateUserName(v string) *DescribeDocResponseBody
	GetCreateUserName() *string
	SetDocInfo(v *DescribeDocResponseBodyDocInfo) *DescribeDocResponseBody
	GetDocInfo() *DescribeDocResponseBodyDocInfo
	SetDocMetadata(v []*DescribeDocResponseBodyDocMetadata) *DescribeDocResponseBody
	GetDocMetadata() []*DescribeDocResponseBodyDocMetadata
	SetDocName(v string) *DescribeDocResponseBody
	GetDocName() *string
	SetDocTags(v []*DescribeDocResponseBodyDocTags) *DescribeDocResponseBody
	GetDocTags() []*DescribeDocResponseBodyDocTags
	SetEffectStatus(v int32) *DescribeDocResponseBody
	GetEffectStatus() *int32
	SetEndDate(v string) *DescribeDocResponseBody
	GetEndDate() *string
	SetKnowledgeId(v int64) *DescribeDocResponseBody
	GetKnowledgeId() *int64
	SetMeta(v string) *DescribeDocResponseBody
	GetMeta() *string
	SetModifyTime(v string) *DescribeDocResponseBody
	GetModifyTime() *string
	SetModifyUserId(v int64) *DescribeDocResponseBody
	GetModifyUserId() *int64
	SetModifyUserName(v string) *DescribeDocResponseBody
	GetModifyUserName() *string
	SetProcessCanRetry(v bool) *DescribeDocResponseBody
	GetProcessCanRetry() *bool
	SetProcessMessage(v string) *DescribeDocResponseBody
	GetProcessMessage() *string
	SetProcessStatus(v int32) *DescribeDocResponseBody
	GetProcessStatus() *int32
	SetRequestId(v string) *DescribeDocResponseBody
	GetRequestId() *string
	SetStartDate(v string) *DescribeDocResponseBody
	GetStartDate() *string
	SetStatus(v int32) *DescribeDocResponseBody
	GetStatus() *int32
	SetTitle(v string) *DescribeDocResponseBody
	GetTitle() *string
	SetUrl(v string) *DescribeDocResponseBody
	GetUrl() *string
}

type DescribeDocResponseBody struct {
	// example:
	//
	// bizcode123
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":"docName"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1111111111
	CreateUserId   *int64                                `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	CreateUserName *string                               `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DocInfo        *DescribeDocResponseBodyDocInfo       `json:"DocInfo,omitempty" xml:"DocInfo,omitempty" type:"Struct"`
	DocMetadata    []*DescribeDocResponseBodyDocMetadata `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty" type:"Repeated"`
	DocName        *string                               `json:"DocName,omitempty" xml:"DocName,omitempty"`
	DocTags        []*DescribeDocResponseBodyDocTags     `json:"DocTags,omitempty" xml:"DocTags,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	EffectStatus *int32 `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// example:
	//
	// 2023-04-27T06:08:54Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2020-11-25T08:56:55Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 2222222222
	ModifyUserId   *int64  `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// true
	ProcessCanRetry *bool   `json:"ProcessCanRetry,omitempty" xml:"ProcessCanRetry,omitempty"`
	ProcessMessage  *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty"`
	// example:
	//
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F132693-212A-40A9-8A81-11E7694E478B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1979-12-31T16:00:00Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 1
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBody) GetBizCode() *string {
	return s.BizCode
}

func (s *DescribeDocResponseBody) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DescribeDocResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeDocResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDocResponseBody) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *DescribeDocResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeDocResponseBody) GetDocInfo() *DescribeDocResponseBodyDocInfo {
	return s.DocInfo
}

func (s *DescribeDocResponseBody) GetDocMetadata() []*DescribeDocResponseBodyDocMetadata {
	return s.DocMetadata
}

func (s *DescribeDocResponseBody) GetDocName() *string {
	return s.DocName
}

func (s *DescribeDocResponseBody) GetDocTags() []*DescribeDocResponseBodyDocTags {
	return s.DocTags
}

func (s *DescribeDocResponseBody) GetEffectStatus() *int32 {
	return s.EffectStatus
}

func (s *DescribeDocResponseBody) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDocResponseBody) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *DescribeDocResponseBody) GetMeta() *string {
	return s.Meta
}

func (s *DescribeDocResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeDocResponseBody) GetModifyUserId() *int64 {
	return s.ModifyUserId
}

func (s *DescribeDocResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeDocResponseBody) GetProcessCanRetry() *bool {
	return s.ProcessCanRetry
}

func (s *DescribeDocResponseBody) GetProcessMessage() *string {
	return s.ProcessMessage
}

func (s *DescribeDocResponseBody) GetProcessStatus() *int32 {
	return s.ProcessStatus
}

func (s *DescribeDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDocResponseBody) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDocResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDocResponseBody) GetTitle() *string {
	return s.Title
}

func (s *DescribeDocResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeDocResponseBody) SetBizCode(v string) *DescribeDocResponseBody {
	s.BizCode = &v
	return s
}

func (s *DescribeDocResponseBody) SetCategoryId(v int64) *DescribeDocResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeDocResponseBody) SetConfig(v string) *DescribeDocResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDocResponseBody) SetCreateTime(v string) *DescribeDocResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDocResponseBody) SetCreateUserId(v int64) *DescribeDocResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeDocResponseBody) SetCreateUserName(v string) *DescribeDocResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeDocResponseBody) SetDocInfo(v *DescribeDocResponseBodyDocInfo) *DescribeDocResponseBody {
	s.DocInfo = v
	return s
}

func (s *DescribeDocResponseBody) SetDocMetadata(v []*DescribeDocResponseBodyDocMetadata) *DescribeDocResponseBody {
	s.DocMetadata = v
	return s
}

func (s *DescribeDocResponseBody) SetDocName(v string) *DescribeDocResponseBody {
	s.DocName = &v
	return s
}

func (s *DescribeDocResponseBody) SetDocTags(v []*DescribeDocResponseBodyDocTags) *DescribeDocResponseBody {
	s.DocTags = v
	return s
}

func (s *DescribeDocResponseBody) SetEffectStatus(v int32) *DescribeDocResponseBody {
	s.EffectStatus = &v
	return s
}

func (s *DescribeDocResponseBody) SetEndDate(v string) *DescribeDocResponseBody {
	s.EndDate = &v
	return s
}

func (s *DescribeDocResponseBody) SetKnowledgeId(v int64) *DescribeDocResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeDocResponseBody) SetMeta(v string) *DescribeDocResponseBody {
	s.Meta = &v
	return s
}

func (s *DescribeDocResponseBody) SetModifyTime(v string) *DescribeDocResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDocResponseBody) SetModifyUserId(v int64) *DescribeDocResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeDocResponseBody) SetModifyUserName(v string) *DescribeDocResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeDocResponseBody) SetProcessCanRetry(v bool) *DescribeDocResponseBody {
	s.ProcessCanRetry = &v
	return s
}

func (s *DescribeDocResponseBody) SetProcessMessage(v string) *DescribeDocResponseBody {
	s.ProcessMessage = &v
	return s
}

func (s *DescribeDocResponseBody) SetProcessStatus(v int32) *DescribeDocResponseBody {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDocResponseBody) SetRequestId(v string) *DescribeDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocResponseBody) SetStartDate(v string) *DescribeDocResponseBody {
	s.StartDate = &v
	return s
}

func (s *DescribeDocResponseBody) SetStatus(v int32) *DescribeDocResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDocResponseBody) SetTitle(v string) *DescribeDocResponseBody {
	s.Title = &v
	return s
}

func (s *DescribeDocResponseBody) SetUrl(v string) *DescribeDocResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeDocResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDocResponseBodyDocInfo struct {
	DocParas []*DescribeDocResponseBodyDocInfoDocParas `json:"DocParas,omitempty" xml:"DocParas,omitempty" type:"Repeated"`
}

func (s DescribeDocResponseBodyDocInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocResponseBodyDocInfo) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocInfo) GetDocParas() []*DescribeDocResponseBodyDocInfoDocParas {
	return s.DocParas
}

func (s *DescribeDocResponseBodyDocInfo) SetDocParas(v []*DescribeDocResponseBodyDocInfoDocParas) *DescribeDocResponseBodyDocInfo {
	s.DocParas = v
	return s
}

func (s *DescribeDocResponseBodyDocInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDocResponseBodyDocInfoDocParas struct {
	// example:
	//
	// 1
	ParaLevel *int32 `json:"ParaLevel,omitempty" xml:"ParaLevel,omitempty"`
	// example:
	//
	// 1
	ParaNo   *int32  `json:"ParaNo,omitempty" xml:"ParaNo,omitempty"`
	ParaText *string `json:"ParaText,omitempty" xml:"ParaText,omitempty"`
	// example:
	//
	// text
	ParaType *string `json:"ParaType,omitempty" xml:"ParaType,omitempty"`
}

func (s DescribeDocResponseBodyDocInfoDocParas) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocResponseBodyDocInfoDocParas) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocInfoDocParas) GetParaLevel() *int32 {
	return s.ParaLevel
}

func (s *DescribeDocResponseBodyDocInfoDocParas) GetParaNo() *int32 {
	return s.ParaNo
}

func (s *DescribeDocResponseBodyDocInfoDocParas) GetParaText() *string {
	return s.ParaText
}

func (s *DescribeDocResponseBodyDocInfoDocParas) GetParaType() *string {
	return s.ParaType
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaLevel(v int32) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaLevel = &v
	return s
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaNo(v int32) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaNo = &v
	return s
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaText(v string) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaText = &v
	return s
}

func (s *DescribeDocResponseBodyDocInfoDocParas) SetParaType(v string) *DescribeDocResponseBodyDocInfoDocParas {
	s.ParaType = &v
	return s
}

func (s *DescribeDocResponseBodyDocInfoDocParas) Validate() error {
	return dara.Validate(s)
}

type DescribeDocResponseBodyDocMetadata struct {
	BusinessViewId      *string                                                  `json:"BusinessViewId,omitempty" xml:"BusinessViewId,omitempty"`
	BusinessViewName    *string                                                  `json:"BusinessViewName,omitempty" xml:"BusinessViewName,omitempty"`
	MetaCellInfoDTOList []*DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList `json:"MetaCellInfoDTOList,omitempty" xml:"MetaCellInfoDTOList,omitempty" type:"Repeated"`
}

func (s DescribeDocResponseBodyDocMetadata) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocResponseBodyDocMetadata) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocMetadata) GetBusinessViewId() *string {
	return s.BusinessViewId
}

func (s *DescribeDocResponseBodyDocMetadata) GetBusinessViewName() *string {
	return s.BusinessViewName
}

func (s *DescribeDocResponseBodyDocMetadata) GetMetaCellInfoDTOList() []*DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList {
	return s.MetaCellInfoDTOList
}

func (s *DescribeDocResponseBodyDocMetadata) SetBusinessViewId(v string) *DescribeDocResponseBodyDocMetadata {
	s.BusinessViewId = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadata) SetBusinessViewName(v string) *DescribeDocResponseBodyDocMetadata {
	s.BusinessViewName = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadata) SetMetaCellInfoDTOList(v []*DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) *DescribeDocResponseBodyDocMetadata {
	s.MetaCellInfoDTOList = v
	return s
}

func (s *DescribeDocResponseBodyDocMetadata) Validate() error {
	return dara.Validate(s)
}

type DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList struct {
	FieldCode *string `json:"FieldCode,omitempty" xml:"FieldCode,omitempty"`
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) GetFieldCode() *string {
	return s.FieldCode
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) GetValue() *string {
	return s.Value
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) SetFieldCode(v string) *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList {
	s.FieldCode = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) SetFieldName(v string) *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList {
	s.FieldName = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) SetValue(v string) *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList {
	s.Value = &v
	return s
}

func (s *DescribeDocResponseBodyDocMetadataMetaCellInfoDTOList) Validate() error {
	return dara.Validate(s)
}

type DescribeDocResponseBodyDocTags struct {
	DefaultTag *bool   `json:"DefaultTag,omitempty" xml:"DefaultTag,omitempty"`
	GroupId    *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	TagId      *int64  `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeDocResponseBodyDocTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocResponseBodyDocTags) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyDocTags) GetDefaultTag() *bool {
	return s.DefaultTag
}

func (s *DescribeDocResponseBodyDocTags) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeDocResponseBodyDocTags) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDocResponseBodyDocTags) GetTagId() *int64 {
	return s.TagId
}

func (s *DescribeDocResponseBodyDocTags) GetTagName() *string {
	return s.TagName
}

func (s *DescribeDocResponseBodyDocTags) SetDefaultTag(v bool) *DescribeDocResponseBodyDocTags {
	s.DefaultTag = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetGroupId(v int64) *DescribeDocResponseBodyDocTags {
	s.GroupId = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetGroupName(v string) *DescribeDocResponseBodyDocTags {
	s.GroupName = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetTagId(v int64) *DescribeDocResponseBodyDocTags {
	s.TagId = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) SetTagName(v string) *DescribeDocResponseBodyDocTags {
	s.TagName = &v
	return s
}

func (s *DescribeDocResponseBodyDocTags) Validate() error {
	return dara.Validate(s)
}
