// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetWorkItemInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetWorkItemInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetWorkItemInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkItemInfoResponseBody
	GetSuccess() *bool
	SetWorkitem(v *GetWorkItemInfoResponseBodyWorkitem) *GetWorkItemInfoResponseBody
	GetWorkitem() *GetWorkItemInfoResponseBodyWorkitem
}

type GetWorkItemInfoResponseBody struct {
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
	Success  *bool                                `json:"success,omitempty" xml:"success,omitempty"`
	Workitem *GetWorkItemInfoResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
}

func (s GetWorkItemInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkItemInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetWorkItemInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkItemInfoResponseBody) GetWorkitem() *GetWorkItemInfoResponseBodyWorkitem {
	return s.Workitem
}

func (s *GetWorkItemInfoResponseBody) SetErrorCode(v string) *GetWorkItemInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetErrorMessage(v string) *GetWorkItemInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetRequestId(v string) *GetWorkItemInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetSuccess(v bool) *GetWorkItemInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetWorkitem(v *GetWorkItemInfoResponseBodyWorkitem) *GetWorkItemInfoResponseBody {
	s.Workitem = v
	return s
}

func (s *GetWorkItemInfoResponseBody) Validate() error {
	if s.Workitem != nil {
		if err := s.Workitem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkItemInfoResponseBodyWorkitem struct {
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
	Creator      *string                                            `json:"creator,omitempty" xml:"creator,omitempty"`
	CustomFields []*GetWorkItemInfoResponseBodyWorkitemCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// example:
	//
	// html格式
	Document       *string `json:"document,omitempty" xml:"document,omitempty"`
	DocumentFormat *string `json:"documentFormat,omitempty" xml:"documentFormat,omitempty"`
	FinishTime     *int64  `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
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
	ParentIdentifier *string   `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	Participant      []*string `json:"participant,omitempty" xml:"participant,omitempty" type:"Repeated"`
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
	SpaceType *string   `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Sprint    []*string `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Repeated"`
	// example:
	//
	// 待处理
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 例：100005
	StatusIdentifier *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	// example:
	//
	// 1
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// example:
	//
	// 测试工作项
	Subject    *string                                          `json:"subject,omitempty" xml:"subject,omitempty"`
	Tag        []*string                                        `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	TagDetails []*GetWorkItemInfoResponseBodyWorkitemTagDetails `json:"tagDetails,omitempty" xml:"tagDetails,omitempty" type:"Repeated"`
	Tracker    []*string                                        `json:"tracker,omitempty" xml:"tracker,omitempty" type:"Repeated"`
	// example:
	//
	// 1640850328000
	UpdateStatusAt *int64                                         `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	Verifier       []*string                                      `json:"verifier,omitempty" xml:"verifier,omitempty" type:"Repeated"`
	Versions       []*GetWorkItemInfoResponseBodyWorkitemVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
	// example:
	//
	// 9uxxxxxxre573f5xxxxxx0
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitem) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitem) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetCreator() *string {
	return s.Creator
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetCustomFields() []*GetWorkItemInfoResponseBodyWorkitemCustomFields {
	return s.CustomFields
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetDocument() *string {
	return s.Document
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetDocumentFormat() *string {
	return s.DocumentFormat
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetModifier() *string {
	return s.Modifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetParticipant() []*string {
	return s.Participant
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetSpaceName() *string {
	return s.SpaceName
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetSpaceType() *string {
	return s.SpaceType
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetSprint() []*string {
	return s.Sprint
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetStatus() *string {
	return s.Status
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetSubject() *string {
	return s.Subject
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetTag() []*string {
	return s.Tag
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetTagDetails() []*GetWorkItemInfoResponseBodyWorkitemTagDetails {
	return s.TagDetails
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetTracker() []*string {
	return s.Tracker
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetUpdateStatusAt() *int64 {
	return s.UpdateStatusAt
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetVerifier() []*string {
	return s.Verifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetVersions() []*GetWorkItemInfoResponseBodyWorkitemVersions {
	return s.Versions
}

func (s *GetWorkItemInfoResponseBodyWorkitem) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetAssignedTo(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.AssignedTo = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetCategoryIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetCreator(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Creator = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetCustomFields(v []*GetWorkItemInfoResponseBodyWorkitemCustomFields) *GetWorkItemInfoResponseBodyWorkitem {
	s.CustomFields = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetDocument(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Document = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetDocumentFormat(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.DocumentFormat = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetFinishTime(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.FinishTime = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetGmtCreate(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetGmtModified(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.GmtModified = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetLogicalStatus(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.LogicalStatus = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetModifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Modifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetParentIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.ParentIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetParticipant(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Participant = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSerialNumber(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SerialNumber = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSpaceIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSpaceName(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SpaceName = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSpaceType(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SpaceType = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSprint(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Sprint = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetStatus(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Status = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetStatusIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.StatusIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetStatusStageIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.StatusStageIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSubject(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Subject = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetTag(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Tag = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetTagDetails(v []*GetWorkItemInfoResponseBodyWorkitemTagDetails) *GetWorkItemInfoResponseBodyWorkitem {
	s.TagDetails = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetTracker(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Tracker = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetUpdateStatusAt(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetVerifier(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Verifier = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetVersions(v []*GetWorkItemInfoResponseBodyWorkitemVersions) *GetWorkItemInfoResponseBodyWorkitem {
	s.Versions = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) Validate() error {
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagDetails != nil {
		for _, item := range s.TagDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkItemInfoResponseBodyWorkitemCustomFields struct {
	// example:
	//
	// 例如：date
	FieldClassName *string `json:"fieldClassName,omitempty" xml:"fieldClassName,omitempty"`
	// example:
	//
	// 例：input
	FieldFormat *string `json:"fieldFormat,omitempty" xml:"fieldFormat,omitempty"`
	// example:
	//
	// 例：80
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// example:
	//
	// 1
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// null
	ObjectValue *string `json:"objectValue,omitempty" xml:"objectValue,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// 例：2022-01-06 00:00:00
	Value     *string                                                     `json:"value,omitempty" xml:"value,omitempty"`
	ValueList []*GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList `json:"valueList,omitempty" xml:"valueList,omitempty" type:"Repeated"`
	// example:
	//
	// 5daa9a15c7fd55523996......
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFields) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFields) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetFieldClassName() *string {
	return s.FieldClassName
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetFieldFormat() *string {
	return s.FieldFormat
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetLevel() *int64 {
	return s.Level
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetObjectValue() *string {
	return s.ObjectValue
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetPosition() *int64 {
	return s.Position
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetValue() *string {
	return s.Value
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetValueList() []*GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	return s.ValueList
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetFieldClassName(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.FieldClassName = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetFieldFormat(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.FieldFormat = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetFieldIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.FieldIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetLevel(v int64) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.Level = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetObjectValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.ObjectValue = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetPosition(v int64) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.Position = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.Value = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetValueList(v []*GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.ValueList = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetWorkitemIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.WorkitemIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) Validate() error {
	if s.ValueList != nil {
		for _, item := range s.ValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList struct {
	// example:
	//
	// 2022-02-01 00:00:00
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	// example:
	//
	// 2022-02-01 00:00:00
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 1
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// 2022-02-01 00:00:00
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// null
	ValueEn *string `json:"valueEn,omitempty" xml:"valueEn,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) GetDisplayValue() *string {
	return s.DisplayValue
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) GetLevel() *int64 {
	return s.Level
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) GetValue() *string {
	return s.Value
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) GetValueEn() *string {
	return s.ValueEn
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetDisplayValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.DisplayValue = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetLevel(v int64) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.Level = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.Value = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetValueEn(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.ValueEn = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) Validate() error {
	return dara.Validate(s)
}

type GetWorkItemInfoResponseBodyWorkitemTagDetails struct {
	Color      *string `json:"color,omitempty" xml:"color,omitempty"`
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitemTagDetails) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitemTagDetails) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitemTagDetails) GetColor() *string {
	return s.Color
}

func (s *GetWorkItemInfoResponseBodyWorkitemTagDetails) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkItemInfoResponseBodyWorkitemTagDetails) GetName() *string {
	return s.Name
}

func (s *GetWorkItemInfoResponseBodyWorkitemTagDetails) SetColor(v string) *GetWorkItemInfoResponseBodyWorkitemTagDetails {
	s.Color = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemTagDetails) SetIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemTagDetails {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemTagDetails) SetName(v string) *GetWorkItemInfoResponseBodyWorkitemTagDetails {
	s.Name = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemTagDetails) Validate() error {
	return dara.Validate(s)
}

type GetWorkItemInfoResponseBodyWorkitemVersions struct {
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitemVersions) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitemVersions) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitemVersions) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkItemInfoResponseBodyWorkitemVersions) GetName() *string {
	return s.Name
}

func (s *GetWorkItemInfoResponseBodyWorkitemVersions) SetIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemVersions {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemVersions) SetName(v string) *GetWorkItemInfoResponseBodyWorkitemVersions {
	s.Name = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemVersions) Validate() error {
	return dara.Validate(s)
}
