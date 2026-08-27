// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBaseContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetCode() *string
	SetItems(v []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetItems() []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems
	SetMessage(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody
	GetTotalCount() *int64
}

type ListUserVisibleKnowledgeBaseContentsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of skill cards.
	Items []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetItems() []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	return s.Items
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetCode(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetItems(v []*ListUserVisibleKnowledgeBaseContentsResponseBodyItems) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Items = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetMessage(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.Page = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetRequestId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) SetTotalCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserVisibleKnowledgeBaseContentsResponseBodyItems struct {
	// The name of the creator.
	//
	// example:
	//
	// John
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Product materials and user guide
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory type.
	//
	// example:
	//
	// normal
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// The directory type.
	//
	// example:
	//
	// TENANT
	DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The item ID.
	//
	// example:
	//
	// source_example
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The item type.
	//
	// example:
	//
	// resource
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// Indicates whether the resource has a pending knowledge base submission record.
	//
	// example:
	//
	// false
	KbSubmissionPending *bool `json:"kbSubmissionPending,omitempty" xml:"kbSubmissionPending,omitempty"`
	// The name of the modifier.
	//
	// example:
	//
	// Jane
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// The name.
	//
	// example:
	//
	// Product description.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object bindings.
	ObjectBindings []map[string]interface{} `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The visibility mode of the knowledge base to digital employees.
	//
	// example:
	//
	// PUBLIC
	OoVisibilityMode *string `json:"ooVisibilityMode,omitempty" xml:"ooVisibilityMode,omitempty"`
	// Indicates whether the item is read-only.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The sharing information.
	ShareInfos []*ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos `json:"shareInfos,omitempty" xml:"shareInfos,omitempty" type:"Repeated"`
	// Indicates whether shared access is allowed.
	//
	// example:
	//
	// false
	Shared *bool `json:"shared,omitempty" xml:"shared,omitempty"`
	// The number of resources in FAILED status. Returned only when listing top-level KB directories.
	//
	// example:
	//
	// 0
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// The knowledge base affiliation type. Valid values: aliding_kb_doc (DingTalk knowledge base document), normal (common knowledge).
	//
	// example:
	//
	// normal
	SourceKind *string `json:"sourceKind,omitempty" xml:"sourceKind,omitempty"`
	// The number of resources in READY status. Returned only when listing top-level KB directories.
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// The resource status. This field has a value only when itemType is resource.
	//
	// example:
	//
	// READY
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// The total number of resources under the directory and its subdirectories. Returned only when listing top-level KB directories.
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
	// The data source type.
	//
	// example:
	//
	// FILE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetItemId() *string {
	return s.ItemId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetKbSubmissionPending() *bool {
	return s.KbSubmissionPending
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetObjectBindings() []map[string]interface{} {
	return s.ObjectBindings
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetOoVisibilityMode() *string {
	return s.OoVisibilityMode
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetShareInfos() []*ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	return s.ShareInfos
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetShared() *bool {
	return s.Shared
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceKind() *string {
	return s.SourceKind
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetCreatorName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetDescription(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetDirectoryKind(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.DirectoryKind = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetDirectoryType(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.DirectoryType = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetGmtCreate(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetGmtModified(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetItemId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ItemId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetItemType(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetKbSubmissionPending(v bool) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.KbSubmissionPending = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetModifierName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ModifierName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetObjectBindings(v []map[string]interface{}) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ObjectBindings = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetOoVisibilityMode(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.OoVisibilityMode = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetReadOnly(v bool) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ReadOnly = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetShareInfos(v []*ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.ShareInfos = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetShared(v bool) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.Shared = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceFailedCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceKind(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceKind = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceReadyCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceStatus(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceStatus = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceTotalCount(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) SetSourceType(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItems) Validate() error {
	if s.ShareInfos != nil {
		for _, item := range s.ShareInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos struct {
	// The Ray Job ID.
	//
	// example:
	//
	// submission_example
	SubmissionId *string `json:"submissionId,omitempty" xml:"submissionId,omitempty"`
	// The user ID of the submitter.
	//
	// example:
	//
	// 1
	SubmitterId *int64 `json:"submitterId,omitempty" xml:"submitterId,omitempty"`
	// The submitter name.
	//
	// example:
	//
	// John
	SubmitterName *string `json:"submitterName,omitempty" xml:"submitterName,omitempty"`
	// The target directory ID.
	//
	// example:
	//
	// dir_target
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// The target directory name.
	//
	// example:
	//
	// Shared directory
	TargetDirectoryName *string `json:"targetDirectoryName,omitempty" xml:"targetDirectoryName,omitempty"`
	// The root directory ID of the target enterprise knowledge base.
	//
	// example:
	//
	// dir_kb_root
	TargetKbRootDirectoryId *string `json:"targetKbRootDirectoryId,omitempty" xml:"targetKbRootDirectoryId,omitempty"`
	// The name of the target enterprise knowledge base.
	//
	// example:
	//
	// Product knowledge base
	TargetKbRootDirectoryName *string `json:"targetKbRootDirectoryName,omitempty" xml:"targetKbRootDirectoryName,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetSubmissionId() *string {
	return s.SubmissionId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetSubmitterId() *int64 {
	return s.SubmitterId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetSubmitterName() *string {
	return s.SubmitterName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetDirectoryName() *string {
	return s.TargetDirectoryName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetKbRootDirectoryId() *string {
	return s.TargetKbRootDirectoryId
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) GetTargetKbRootDirectoryName() *string {
	return s.TargetKbRootDirectoryName
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetSubmissionId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.SubmissionId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetSubmitterId(v int64) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.SubmitterId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetSubmitterName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.SubmitterName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetDirectoryId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetDirectoryName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetDirectoryName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetKbRootDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetKbRootDirectoryId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) SetTargetKbRootDirectoryName(v string) *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos {
	s.TargetKbRootDirectoryName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponseBodyItemsShareInfos) Validate() error {
	return dara.Validate(s)
}
