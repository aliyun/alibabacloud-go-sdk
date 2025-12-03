// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListWorkitemsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListWorkitemsResponseBody
	GetErrorMsg() *string
	SetMaxResults(v int64) *ListWorkitemsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListWorkitemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkitemsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkitemsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListWorkitemsResponseBody
	GetTotalCount() *int64
	SetWorkitems(v []*ListWorkitemsResponseBodyWorkitems) *ListWorkitemsResponseBody
	GetWorkitems() []*ListWorkitemsResponseBodyWorkitems
}

type ListWorkitemsResponseBody struct {
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
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Workitems  []*ListWorkitemsResponseBodyWorkitems `json:"workitems,omitempty" xml:"workitems,omitempty" type:"Repeated"`
}

func (s ListWorkitemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkitemsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkitemsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListWorkitemsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkitemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkitemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkitemsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkitemsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkitemsResponseBody) GetWorkitems() []*ListWorkitemsResponseBodyWorkitems {
	return s.Workitems
}

func (s *ListWorkitemsResponseBody) SetErrorCode(v string) *ListWorkitemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetErrorMsg(v string) *ListWorkitemsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetMaxResults(v int64) *ListWorkitemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetNextToken(v string) *ListWorkitemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetRequestId(v string) *ListWorkitemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetSuccess(v bool) *ListWorkitemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetTotalCount(v int64) *ListWorkitemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetWorkitems(v []*ListWorkitemsResponseBodyWorkitems) *ListWorkitemsResponseBody {
	s.Workitems = v
	return s
}

func (s *ListWorkitemsResponseBody) Validate() error {
	if s.Workitems != nil {
		for _, item := range s.Workitems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkitemsResponseBodyWorkitems struct {
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
	Document   *string `json:"document,omitempty" xml:"document,omitempty"`
	FinishTime *int64  `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
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
	// 100005
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

func (s ListWorkitemsResponseBodyWorkitems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemsResponseBodyWorkitems) GoString() string {
	return s.String()
}

func (s *ListWorkitemsResponseBodyWorkitems) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *ListWorkitemsResponseBodyWorkitems) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkitemsResponseBodyWorkitems) GetDocument() *string {
	return s.Document
}

func (s *ListWorkitemsResponseBodyWorkitems) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListWorkitemsResponseBodyWorkitems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListWorkitemsResponseBodyWorkitems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListWorkitemsResponseBodyWorkitems) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *ListWorkitemsResponseBodyWorkitems) GetModifier() *string {
	return s.Modifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListWorkitemsResponseBodyWorkitems) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetSpaceName() *string {
	return s.SpaceName
}

func (s *ListWorkitemsResponseBodyWorkitems) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListWorkitemsResponseBodyWorkitems) GetSprintIdentifier() *string {
	return s.SprintIdentifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetStatus() *string {
	return s.Status
}

func (s *ListWorkitemsResponseBodyWorkitems) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *ListWorkitemsResponseBodyWorkitems) GetSubject() *string {
	return s.Subject
}

func (s *ListWorkitemsResponseBodyWorkitems) GetUpdateStatusAt() *int64 {
	return s.UpdateStatusAt
}

func (s *ListWorkitemsResponseBodyWorkitems) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *ListWorkitemsResponseBodyWorkitems) SetAssignedTo(v string) *ListWorkitemsResponseBodyWorkitems {
	s.AssignedTo = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetCategoryIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.CategoryIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetCreator(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Creator = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetDocument(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Document = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetFinishTime(v int64) *ListWorkitemsResponseBodyWorkitems {
	s.FinishTime = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetGmtCreate(v int64) *ListWorkitemsResponseBodyWorkitems {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetGmtModified(v int64) *ListWorkitemsResponseBodyWorkitems {
	s.GmtModified = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Identifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetLogicalStatus(v string) *ListWorkitemsResponseBodyWorkitems {
	s.LogicalStatus = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetModifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Modifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetParentIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.ParentIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSerialNumber(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SerialNumber = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSpaceIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSpaceName(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SpaceName = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSpaceType(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SpaceType = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSprintIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SprintIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetStatus(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Status = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetStatusIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.StatusIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetStatusStageIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.StatusStageIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSubject(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Subject = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetUpdateStatusAt(v int64) *ListWorkitemsResponseBodyWorkitems {
	s.UpdateStatusAt = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetWorkitemTypeIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) Validate() error {
	return dara.Validate(s)
}
