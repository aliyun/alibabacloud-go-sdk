// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProjectsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListProjectsResponseBody
	GetErrorMsg() *string
	SetMaxResults(v int64) *ListProjectsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListProjectsResponseBody
	GetNextToken() *string
	SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody
	GetProjects() []*ListProjectsResponseBodyProjects
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListProjectsResponseBody
	GetTotalCount() *int64
}

type ListProjectsResponseBody struct {
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
	NextToken *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Projects  []*ListProjectsResponseBodyProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
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
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProjectsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListProjectsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsResponseBody) GetProjects() []*ListProjectsResponseBodyProjects {
	return s.Projects
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProjectsResponseBody) SetErrorCode(v string) *ListProjectsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectsResponseBody) SetErrorMsg(v string) *ListProjectsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectsResponseBody) SetMaxResults(v int64) *ListProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetSuccess(v bool) *ListProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int64) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyProjects struct {
	// example:
	//
	// Project
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// OJAY
	CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	// example:
	//
	// null
	DeleteTime *int64 `json:"deleteTime,omitempty" xml:"deleteTime,omitempty"`
	// example:
	//
	// xxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1640778694000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// https://xxxxxx.png
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// null
	LogicalStatus *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	// example:
	//
	// xxxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// public
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// null
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// example:
	//
	// null
	TypeIdentifier *string `json:"typeIdentifier,omitempty" xml:"typeIdentifier,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *ListProjectsResponseBodyProjects) GetCreator() *string {
	return s.Creator
}

func (s *ListProjectsResponseBodyProjects) GetCustomCode() *string {
	return s.CustomCode
}

func (s *ListProjectsResponseBodyProjects) GetDeleteTime() *int64 {
	return s.DeleteTime
}

func (s *ListProjectsResponseBodyProjects) GetDescription() *string {
	return s.Description
}

func (s *ListProjectsResponseBodyProjects) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListProjectsResponseBodyProjects) GetIcon() *string {
	return s.Icon
}

func (s *ListProjectsResponseBodyProjects) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListProjectsResponseBodyProjects) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *ListProjectsResponseBodyProjects) GetName() *string {
	return s.Name
}

func (s *ListProjectsResponseBodyProjects) GetScope() *string {
	return s.Scope
}

func (s *ListProjectsResponseBodyProjects) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *ListProjectsResponseBodyProjects) GetTypeIdentifier() *string {
	return s.TypeIdentifier
}

func (s *ListProjectsResponseBodyProjects) SetCategoryIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.CategoryIdentifier = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCreator(v string) *ListProjectsResponseBodyProjects {
	s.Creator = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCustomCode(v string) *ListProjectsResponseBodyProjects {
	s.CustomCode = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDeleteTime(v int64) *ListProjectsResponseBodyProjects {
	s.DeleteTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtCreate(v int64) *ListProjectsResponseBodyProjects {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetIcon(v string) *ListProjectsResponseBodyProjects {
	s.Icon = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.Identifier = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetLogicalStatus(v string) *ListProjectsResponseBodyProjects {
	s.LogicalStatus = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetName(v string) *ListProjectsResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetScope(v string) *ListProjectsResponseBodyProjects {
	s.Scope = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetStatusStageIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.StatusStageIdentifier = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetTypeIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.TypeIdentifier = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) Validate() error {
	return dara.Validate(s)
}
