// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCmsV1MigrationTaskView interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *CmsV1MigrationTaskView
	GetBatchId() *string
	SetCanRecover(v bool) *CmsV1MigrationTaskView
	GetCanRecover() *bool
	SetErrorMessage(v string) *CmsV1MigrationTaskView
	GetErrorMessage() *string
	SetGroupName(v string) *CmsV1MigrationTaskView
	GetGroupName() *string
	SetObserveGroupId(v string) *CmsV1MigrationTaskView
	GetObserveGroupId() *string
	SetOriginGroupId(v string) *CmsV1MigrationTaskView
	GetOriginGroupId() *string
	SetSourceReadOnly(v bool) *CmsV1MigrationTaskView
	GetSourceReadOnly() *bool
	SetStatus(v string) *CmsV1MigrationTaskView
	GetStatus() *string
	SetTakeoverStage(v string) *CmsV1MigrationTaskView
	GetTakeoverStage() *string
	SetTargetGroupName(v string) *CmsV1MigrationTaskView
	GetTargetGroupName() *string
	SetTaskId(v int64) *CmsV1MigrationTaskView
	GetTaskId() *int64
}

type CmsV1MigrationTaskView struct {
	// The migration batch ID.
	//
	// example:
	//
	// 00112233-4455-4677-8899-aabbccddeeff
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// Indicates whether to provide a recovery prompt for existing frozen tasks. This field does not grant authorization. Each recovery re-verifies the original bindings. Tasks with the SUCCESS status are not retried.
	//
	// example:
	//
	// false
	CanRecover *bool `json:"canRecover,omitempty" xml:"canRecover,omitempty"`
	// The task failure description.
	//
	// example:
	//
	// UNSUPPORTED_RESOURCE: source members cannot be translated without loss
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The source application group name.
	//
	// example:
	//
	// order-service
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The target observation group ID.
	//
	// example:
	//
	// og-0123456789abcdef
	ObserveGroupId *string `json:"observeGroupId,omitempty" xml:"observeGroupId,omitempty"`
	// The source CloudMonitor 1.0 application group ID.
	//
	// example:
	//
	// 247675311
	OriginGroupId *string `json:"originGroupId,omitempty" xml:"originGroupId,omitempty"`
	// Indicates whether the source 1.0 group is persistently set to read-only.
	//
	// example:
	//
	// false
	SourceReadOnly *bool `json:"sourceReadOnly,omitempty" xml:"sourceReadOnly,omitempty"`
	// The migration task status. Valid values: PENDING, RUNNING, SUCCESS, SKIPPED, and FAILED.
	//
	// example:
	//
	// FAILED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The takeover stage. Valid values: PENDING, SOURCE_LOCKED, TARGET_CREATED, READ_ONLY, and NEEDS_REVIEW.
	//
	// example:
	//
	// SOURCE_LOCKED
	TakeoverStage *string `json:"takeoverStage,omitempty" xml:"takeoverStage,omitempty"`
	// The target observation group name.
	//
	// example:
	//
	// order-service
	TargetGroupName *string `json:"targetGroupName,omitempty" xml:"targetGroupName,omitempty"`
	// The migration task ID.
	//
	// example:
	//
	// 3
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CmsV1MigrationTaskView) String() string {
	return dara.Prettify(s)
}

func (s CmsV1MigrationTaskView) GoString() string {
	return s.String()
}

func (s *CmsV1MigrationTaskView) GetBatchId() *string {
	return s.BatchId
}

func (s *CmsV1MigrationTaskView) GetCanRecover() *bool {
	return s.CanRecover
}

func (s *CmsV1MigrationTaskView) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CmsV1MigrationTaskView) GetGroupName() *string {
	return s.GroupName
}

func (s *CmsV1MigrationTaskView) GetObserveGroupId() *string {
	return s.ObserveGroupId
}

func (s *CmsV1MigrationTaskView) GetOriginGroupId() *string {
	return s.OriginGroupId
}

func (s *CmsV1MigrationTaskView) GetSourceReadOnly() *bool {
	return s.SourceReadOnly
}

func (s *CmsV1MigrationTaskView) GetStatus() *string {
	return s.Status
}

func (s *CmsV1MigrationTaskView) GetTakeoverStage() *string {
	return s.TakeoverStage
}

func (s *CmsV1MigrationTaskView) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *CmsV1MigrationTaskView) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CmsV1MigrationTaskView) SetBatchId(v string) *CmsV1MigrationTaskView {
	s.BatchId = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetCanRecover(v bool) *CmsV1MigrationTaskView {
	s.CanRecover = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetErrorMessage(v string) *CmsV1MigrationTaskView {
	s.ErrorMessage = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetGroupName(v string) *CmsV1MigrationTaskView {
	s.GroupName = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetObserveGroupId(v string) *CmsV1MigrationTaskView {
	s.ObserveGroupId = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetOriginGroupId(v string) *CmsV1MigrationTaskView {
	s.OriginGroupId = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetSourceReadOnly(v bool) *CmsV1MigrationTaskView {
	s.SourceReadOnly = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetStatus(v string) *CmsV1MigrationTaskView {
	s.Status = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetTakeoverStage(v string) *CmsV1MigrationTaskView {
	s.TakeoverStage = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetTargetGroupName(v string) *CmsV1MigrationTaskView {
	s.TargetGroupName = &v
	return s
}

func (s *CmsV1MigrationTaskView) SetTaskId(v int64) *CmsV1MigrationTaskView {
	s.TaskId = &v
	return s
}

func (s *CmsV1MigrationTaskView) Validate() error {
	return dara.Validate(s)
}
