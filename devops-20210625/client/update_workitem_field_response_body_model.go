// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkitemFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateWorkitemFieldResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateWorkitemFieldResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *UpdateWorkitemFieldResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkitemFieldResponseBody
	GetSuccess() *bool
	SetWorkitem(v *UpdateWorkitemFieldResponseBodyWorkitem) *UpdateWorkitemFieldResponseBody
	GetWorkitem() *UpdateWorkitemFieldResponseBodyWorkitem
}

type UpdateWorkitemFieldResponseBody struct {
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
	Success  *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	Workitem *UpdateWorkitemFieldResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
}

func (s UpdateWorkitemFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemFieldResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateWorkitemFieldResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateWorkitemFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkitemFieldResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkitemFieldResponseBody) GetWorkitem() *UpdateWorkitemFieldResponseBodyWorkitem {
	return s.Workitem
}

func (s *UpdateWorkitemFieldResponseBody) SetErrorCode(v string) *UpdateWorkitemFieldResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBody) SetErrorMsg(v string) *UpdateWorkitemFieldResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBody) SetRequestId(v string) *UpdateWorkitemFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBody) SetSuccess(v bool) *UpdateWorkitemFieldResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBody) SetWorkitem(v *UpdateWorkitemFieldResponseBodyWorkitem) *UpdateWorkitemFieldResponseBody {
	s.Workitem = v
	return s
}

func (s *UpdateWorkitemFieldResponseBody) Validate() error {
	if s.Workitem != nil {
		if err := s.Workitem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkitemFieldResponseBodyWorkitem struct {
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
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
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
	// 1
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

func (s UpdateWorkitemFieldResponseBodyWorkitem) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemFieldResponseBodyWorkitem) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetCreator() *string {
	return s.Creator
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetDocument() *string {
	return s.Document
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetModifier() *string {
	return s.Modifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetSpaceName() *string {
	return s.SpaceName
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetSpaceType() *string {
	return s.SpaceType
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetSprintIdentifier() *string {
	return s.SprintIdentifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetStatus() *string {
	return s.Status
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetSubject() *string {
	return s.Subject
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetUpdateStatusAt() *int64 {
	return s.UpdateStatusAt
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetAssignedTo(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.AssignedTo = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetCategoryIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.CategoryIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetCreator(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.Creator = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetDocument(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.Document = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetGmtCreate(v int64) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.GmtCreate = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetGmtModified(v int64) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.GmtModified = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.Identifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetLogicalStatus(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.LogicalStatus = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetModifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.Modifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetParentIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.ParentIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetSerialNumber(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.SerialNumber = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetSpaceIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.SpaceIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetSpaceName(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.SpaceName = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetSpaceType(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.SpaceType = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetSprintIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.SprintIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetStatus(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.Status = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetStatusIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.StatusIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetStatusStageIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.StatusStageIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetSubject(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.Subject = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetUpdateStatusAt(v int64) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *UpdateWorkitemFieldResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldResponseBodyWorkitem) Validate() error {
	return dara.Validate(s)
}
