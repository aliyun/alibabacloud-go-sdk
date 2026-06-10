// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDocHits(v []*SearchDocResponseBodyDocHits) *SearchDocResponseBody
	GetDocHits() []*SearchDocResponseBodyDocHits
	SetPageNumber(v int32) *SearchDocResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchDocResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchDocResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *SearchDocResponseBody
	GetTotalCount() *int32
}

type SearchDocResponseBody struct {
	// A list of matching documents.
	DocHits []*SearchDocResponseBodyDocHits `json:"DocHits,omitempty" xml:"DocHits,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E3E5C779-A630-45AC-B0F2-A4506A4212F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of matching entries.
	//
	// example:
	//
	// 141
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchDocResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDocResponseBody) GetDocHits() []*SearchDocResponseBodyDocHits {
	return s.DocHits
}

func (s *SearchDocResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchDocResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchDocResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchDocResponseBody) SetDocHits(v []*SearchDocResponseBodyDocHits) *SearchDocResponseBody {
	s.DocHits = v
	return s
}

func (s *SearchDocResponseBody) SetPageNumber(v int32) *SearchDocResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchDocResponseBody) SetPageSize(v int32) *SearchDocResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchDocResponseBody) SetRequestId(v string) *SearchDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchDocResponseBody) SetTotalCount(v int32) *SearchDocResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchDocResponseBody) Validate() error {
	if s.DocHits != nil {
		for _, item := range s.DocHits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchDocResponseBodyDocHits struct {
	// The business code.
	//
	// example:
	//
	// cn_dytns
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 30000135654
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The splitter for the document. Key: `Splitter`. Valid values:<br>• `paragraphSplitter` (default): Splits the document by paragraph.<br>• `treeSplitter`: Splits the document based on a rule-based hierarchy.<br><br>
	//
	// The size of each document chunk. Key: `ChunkSize`. Default value: 500. Value range: [200, 800].
	//
	// The patterns for the rule-based hierarchy. Key: `TreePatterns`. Default value: `[]`.
	//
	// The source of the document title. Key: `TitleSource`. Valid values:<br>• `ocrTitle` (default): Uses the OCR-identified title.<br>• `docName`: Uses the document name as the title.<br><br>
	//
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":"docName"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The time the document was created (UTC).
	//
	// example:
	//
	// 2023-06-22T03:53:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 111111111
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// 张三
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// The document name.
	//
	// example:
	//
	// 文档名称
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// A list of the document\\"s tags.
	DocTags []*SearchDocResponseBodyDocHitsDocTags `json:"DocTags,omitempty" xml:"DocTags,omitempty" type:"Repeated"`
	// The knowledge\\"s effective status, calculated based on `StartDate` and `EndDate`. Valid values:<br>• 20: Active<br>• 21: Expired<br>• 22: Pending<br><br><br>
	//
	// example:
	//
	// 20
	EffectStatus *int32 `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// The time the knowledge expires (UTC).
	//
	// example:
	//
	// 2099-12-31T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The knowledge ID.
	//
	// example:
	//
	// 30002692007
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The document metadata.
	//
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The time the document was last modified (UTC).
	//
	// example:
	//
	// 2023-06-25T02:27:42Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The modifier ID.
	//
	// example:
	//
	// 222222222
	ModifyUserId *int64 `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// The name of the last modifier.
	//
	// example:
	//
	// 李四
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// Indicates whether the task can be retried.<br>• `true`: The task can be retried.<br>• `false`: The task cannot be retried.<br><br>
	//
	// example:
	//
	// true
	ProcessCanRetry *bool `json:"ProcessCanRetry,omitempty" xml:"ProcessCanRetry,omitempty"`
	// The processing message for the task.
	//
	// example:
	//
	// 任务处理成功
	ProcessMessage *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty"`
	// The task\\"s processing status. Valid values:<br>• -1: Queued<br>• 0: Succeeded<br>• 1: Parsing<br>• 2: Processing<br>• 3: Failed<br><br><br><br><br>
	//
	// example:
	//
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The time the knowledge takes effect (UTC).
	//
	// example:
	//
	// 2023-02-28T11:40:18Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The editing status of the document. Valid values:<br>• 1: Unpublished<br>• 2: Published<br>• 3: Updated but not published<br><br><br>
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The OSS URL of the object.
	//
	// example:
	//
	// https://doc2bot-bucket-cloud.oss-cn-shanghai.aliyuncs.com/doc2bot/input/文档问答说明书_V1.pdf
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchDocResponseBodyDocHits) String() string {
	return dara.Prettify(s)
}

func (s SearchDocResponseBodyDocHits) GoString() string {
	return s.String()
}

func (s *SearchDocResponseBodyDocHits) GetBizCode() *string {
	return s.BizCode
}

func (s *SearchDocResponseBodyDocHits) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *SearchDocResponseBodyDocHits) GetConfig() *string {
	return s.Config
}

func (s *SearchDocResponseBodyDocHits) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchDocResponseBodyDocHits) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *SearchDocResponseBodyDocHits) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *SearchDocResponseBodyDocHits) GetDocName() *string {
	return s.DocName
}

func (s *SearchDocResponseBodyDocHits) GetDocTags() []*SearchDocResponseBodyDocHitsDocTags {
	return s.DocTags
}

func (s *SearchDocResponseBodyDocHits) GetEffectStatus() *int32 {
	return s.EffectStatus
}

func (s *SearchDocResponseBodyDocHits) GetEndDate() *string {
	return s.EndDate
}

func (s *SearchDocResponseBodyDocHits) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *SearchDocResponseBodyDocHits) GetMeta() *string {
	return s.Meta
}

func (s *SearchDocResponseBodyDocHits) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *SearchDocResponseBodyDocHits) GetModifyUserId() *int64 {
	return s.ModifyUserId
}

func (s *SearchDocResponseBodyDocHits) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *SearchDocResponseBodyDocHits) GetProcessCanRetry() *bool {
	return s.ProcessCanRetry
}

func (s *SearchDocResponseBodyDocHits) GetProcessMessage() *string {
	return s.ProcessMessage
}

func (s *SearchDocResponseBodyDocHits) GetProcessStatus() *int32 {
	return s.ProcessStatus
}

func (s *SearchDocResponseBodyDocHits) GetStartDate() *string {
	return s.StartDate
}

func (s *SearchDocResponseBodyDocHits) GetStatus() *int32 {
	return s.Status
}

func (s *SearchDocResponseBodyDocHits) GetUrl() *string {
	return s.Url
}

func (s *SearchDocResponseBodyDocHits) SetBizCode(v string) *SearchDocResponseBodyDocHits {
	s.BizCode = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCategoryId(v int64) *SearchDocResponseBodyDocHits {
	s.CategoryId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetConfig(v string) *SearchDocResponseBodyDocHits {
	s.Config = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCreateTime(v string) *SearchDocResponseBodyDocHits {
	s.CreateTime = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCreateUserId(v int64) *SearchDocResponseBodyDocHits {
	s.CreateUserId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetCreateUserName(v string) *SearchDocResponseBodyDocHits {
	s.CreateUserName = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetDocName(v string) *SearchDocResponseBodyDocHits {
	s.DocName = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetDocTags(v []*SearchDocResponseBodyDocHitsDocTags) *SearchDocResponseBodyDocHits {
	s.DocTags = v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetEffectStatus(v int32) *SearchDocResponseBodyDocHits {
	s.EffectStatus = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetEndDate(v string) *SearchDocResponseBodyDocHits {
	s.EndDate = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetKnowledgeId(v int64) *SearchDocResponseBodyDocHits {
	s.KnowledgeId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetMeta(v string) *SearchDocResponseBodyDocHits {
	s.Meta = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetModifyTime(v string) *SearchDocResponseBodyDocHits {
	s.ModifyTime = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetModifyUserId(v int64) *SearchDocResponseBodyDocHits {
	s.ModifyUserId = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetModifyUserName(v string) *SearchDocResponseBodyDocHits {
	s.ModifyUserName = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetProcessCanRetry(v bool) *SearchDocResponseBodyDocHits {
	s.ProcessCanRetry = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetProcessMessage(v string) *SearchDocResponseBodyDocHits {
	s.ProcessMessage = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetProcessStatus(v int32) *SearchDocResponseBodyDocHits {
	s.ProcessStatus = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetStartDate(v string) *SearchDocResponseBodyDocHits {
	s.StartDate = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetStatus(v int32) *SearchDocResponseBodyDocHits {
	s.Status = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) SetUrl(v string) *SearchDocResponseBodyDocHits {
	s.Url = &v
	return s
}

func (s *SearchDocResponseBodyDocHits) Validate() error {
	if s.DocTags != nil {
		for _, item := range s.DocTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchDocResponseBodyDocHitsDocTags struct {
	// Indicates whether this is a default tag.
	//
	// example:
	//
	// false
	DefaultTag *bool `json:"DefaultTag,omitempty" xml:"DefaultTag,omitempty"`
	// The tag group ID.
	//
	// example:
	//
	// 3610
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The tag group name.
	//
	// example:
	//
	// 文章
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 1201
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 小说
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s SearchDocResponseBodyDocHitsDocTags) String() string {
	return dara.Prettify(s)
}

func (s SearchDocResponseBodyDocHitsDocTags) GoString() string {
	return s.String()
}

func (s *SearchDocResponseBodyDocHitsDocTags) GetDefaultTag() *bool {
	return s.DefaultTag
}

func (s *SearchDocResponseBodyDocHitsDocTags) GetGroupId() *int64 {
	return s.GroupId
}

func (s *SearchDocResponseBodyDocHitsDocTags) GetGroupName() *string {
	return s.GroupName
}

func (s *SearchDocResponseBodyDocHitsDocTags) GetTagId() *int64 {
	return s.TagId
}

func (s *SearchDocResponseBodyDocHitsDocTags) GetTagName() *string {
	return s.TagName
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetDefaultTag(v bool) *SearchDocResponseBodyDocHitsDocTags {
	s.DefaultTag = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetGroupId(v int64) *SearchDocResponseBodyDocHitsDocTags {
	s.GroupId = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetGroupName(v string) *SearchDocResponseBodyDocHitsDocTags {
	s.GroupName = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetTagId(v int64) *SearchDocResponseBodyDocHitsDocTags {
	s.TagId = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) SetTagName(v string) *SearchDocResponseBodyDocHitsDocTags {
	s.TagName = &v
	return s
}

func (s *SearchDocResponseBodyDocHitsDocTags) Validate() error {
	return dara.Validate(s)
}
