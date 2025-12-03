// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateWorkitemResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateWorkitemResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateWorkitemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkitemResponseBody
	GetSuccess() *bool
	SetWorkitem(v *CreateWorkitemResponseBodyWorkitem) *CreateWorkitemResponseBody
	GetWorkitem() *CreateWorkitemResponseBodyWorkitem
}

type CreateWorkitemResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success  *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	Workitem *CreateWorkitemResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
}

func (s CreateWorkitemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateWorkitemResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateWorkitemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkitemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkitemResponseBody) GetWorkitem() *CreateWorkitemResponseBodyWorkitem {
	return s.Workitem
}

func (s *CreateWorkitemResponseBody) SetErrorCode(v string) *CreateWorkitemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetErrorMsg(v string) *CreateWorkitemResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetRequestId(v string) *CreateWorkitemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetSuccess(v bool) *CreateWorkitemResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetWorkitem(v *CreateWorkitemResponseBodyWorkitem) *CreateWorkitemResponseBody {
	s.Workitem = v
	return s
}

func (s *CreateWorkitemResponseBody) Validate() error {
	if s.Workitem != nil {
		if err := s.Workitem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkitemResponseBodyWorkitem struct {
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
	// ABCD-1
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
	// 28
	StatusIdentifier *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	// example:
	//
	// 100005
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

func (s CreateWorkitemResponseBodyWorkitem) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemResponseBodyWorkitem) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponseBodyWorkitem) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *CreateWorkitemResponseBodyWorkitem) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetCreator() *string {
	return s.Creator
}

func (s *CreateWorkitemResponseBodyWorkitem) GetDocument() *string {
	return s.Document
}

func (s *CreateWorkitemResponseBodyWorkitem) GetDocumentFormat() *string {
	return s.DocumentFormat
}

func (s *CreateWorkitemResponseBodyWorkitem) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateWorkitemResponseBodyWorkitem) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CreateWorkitemResponseBodyWorkitem) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *CreateWorkitemResponseBodyWorkitem) GetModifier() *string {
	return s.Modifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateWorkitemResponseBodyWorkitem) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetSpaceName() *string {
	return s.SpaceName
}

func (s *CreateWorkitemResponseBodyWorkitem) GetSpaceType() *string {
	return s.SpaceType
}

func (s *CreateWorkitemResponseBodyWorkitem) GetSprintIdentifier() *string {
	return s.SprintIdentifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetStatus() *string {
	return s.Status
}

func (s *CreateWorkitemResponseBodyWorkitem) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *CreateWorkitemResponseBodyWorkitem) GetSubject() *string {
	return s.Subject
}

func (s *CreateWorkitemResponseBodyWorkitem) GetUpdateStatusAt() *int64 {
	return s.UpdateStatusAt
}

func (s *CreateWorkitemResponseBodyWorkitem) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *CreateWorkitemResponseBodyWorkitem) SetAssignedTo(v string) *CreateWorkitemResponseBodyWorkitem {
	s.AssignedTo = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetCategoryIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.CategoryIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetCreator(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Creator = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetDocument(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Document = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetDocumentFormat(v string) *CreateWorkitemResponseBodyWorkitem {
	s.DocumentFormat = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetGmtCreate(v int64) *CreateWorkitemResponseBodyWorkitem {
	s.GmtCreate = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetGmtModified(v int64) *CreateWorkitemResponseBodyWorkitem {
	s.GmtModified = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetLogicalStatus(v string) *CreateWorkitemResponseBodyWorkitem {
	s.LogicalStatus = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetModifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Modifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetParentIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSerialNumber(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SerialNumber = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSpaceIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSpaceName(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SpaceName = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSpaceType(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SpaceType = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSprintIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SprintIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetStatus(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Status = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetStatusIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.StatusIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetStatusStageIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.StatusStageIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSubject(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Subject = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetUpdateStatusAt(v int64) *CreateWorkitemResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) Validate() error {
	return dara.Validate(s)
}
