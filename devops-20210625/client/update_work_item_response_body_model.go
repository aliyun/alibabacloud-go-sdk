// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateWorkItemResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateWorkItemResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateWorkItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkItemResponseBody
	GetSuccess() *bool
	SetWorkitem(v *UpdateWorkItemResponseBodyWorkitem) *UpdateWorkItemResponseBody
	GetWorkitem() *UpdateWorkItemResponseBodyWorkitem
}

type UpdateWorkItemResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success  *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	Workitem *UpdateWorkItemResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
}

func (s UpdateWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateWorkItemResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkItemResponseBody) GetWorkitem() *UpdateWorkItemResponseBodyWorkitem {
	return s.Workitem
}

func (s *UpdateWorkItemResponseBody) SetErrorCode(v string) *UpdateWorkItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetErrorMessage(v string) *UpdateWorkItemResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetRequestId(v string) *UpdateWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetSuccess(v bool) *UpdateWorkItemResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetWorkitem(v *UpdateWorkItemResponseBodyWorkitem) *UpdateWorkItemResponseBody {
	s.Workitem = v
	return s
}

func (s *UpdateWorkItemResponseBody) Validate() error {
	if s.Workitem != nil {
		if err := s.Workitem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkItemResponseBodyWorkitem struct {
	// example:
	//
	// 19xx7043xxxxxxx914
	AssignedTo *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	// example:
	//
	// Req
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// html格式
	Document       *string `json:"document,omitempty" xml:"document,omitempty"`
	DocumentFormat *string `json:"documentFormat,omitempty" xml:"documentFormat,omitempty"`
	// example:
	//
	// 1640850318000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1640850318000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// NORMAL
	LogicalStatus *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx24
	ParentIdentifier *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	// example:
	//
	// ACFS-1
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// example:
	//
	// e8b26xxxxx6e76aa20xxxxx23
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// 需求项目
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// example:
	//
	// Project
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// example:
	//
	// 75528f17703e92e5a568......
	SprintIdentifier *string `json:"sprintIdentifier,omitempty" xml:"sprintIdentifier,omitempty"`
	// example:
	//
	// 待处理
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 111000
	StatusIdentifier *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	// example:
	//
	// 1
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// example:
	//
	// 测试工作项
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// example:
	//
	// 1640850328000
	UpdateStatusAt *int64 `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	// example:
	//
	// 9uxxxxxxre573f5xxxxxx0
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s UpdateWorkItemResponseBodyWorkitem) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemResponseBodyWorkitem) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetCreator() *string {
	return s.Creator
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetDocument() *string {
	return s.Document
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetDocumentFormat() *string {
	return s.DocumentFormat
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetModifier() *string {
	return s.Modifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetSpaceName() *string {
	return s.SpaceName
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetSpaceType() *string {
	return s.SpaceType
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetSprintIdentifier() *string {
	return s.SprintIdentifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetStatus() *string {
	return s.Status
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetSubject() *string {
	return s.Subject
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetUpdateStatusAt() *int64 {
	return s.UpdateStatusAt
}

func (s *UpdateWorkItemResponseBodyWorkitem) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetAssignedTo(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.AssignedTo = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetCategoryIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.CategoryIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetCreator(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Creator = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetDocument(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Document = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetDocumentFormat(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.DocumentFormat = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetGmtCreate(v int64) *UpdateWorkItemResponseBodyWorkitem {
	s.GmtCreate = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetGmtModified(v int64) *UpdateWorkItemResponseBodyWorkitem {
	s.GmtModified = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Identifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetLogicalStatus(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.LogicalStatus = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetModifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Modifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetParentIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.ParentIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSerialNumber(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SerialNumber = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSpaceIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SpaceIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSpaceName(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SpaceName = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSpaceType(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SpaceType = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSprintIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SprintIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetStatus(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Status = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetStatusIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.StatusIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetStatusStageIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.StatusStageIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSubject(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Subject = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetUpdateStatusAt(v int64) *UpdateWorkItemResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) Validate() error {
	return dara.Validate(s)
}
