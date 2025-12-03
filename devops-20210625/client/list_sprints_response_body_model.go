// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSprintsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSprintsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListSprintsResponseBody
	GetErrorMsg() *string
	SetMaxResults(v int64) *ListSprintsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListSprintsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSprintsResponseBody
	GetRequestId() *string
	SetSprints(v []*ListSprintsResponseBodySprints) *ListSprintsResponseBody
	GetSprints() []*ListSprintsResponseBodySprints
	SetSuccess(v bool) *ListSprintsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSprintsResponseBody
	GetTotalCount() *int64
}

type ListSprintsResponseBody struct {
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
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Sprints   []*ListSprintsResponseBodySprints `json:"sprints,omitempty" xml:"sprints,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSprintsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSprintsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSprintsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSprintsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListSprintsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListSprintsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSprintsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSprintsResponseBody) GetSprints() []*ListSprintsResponseBodySprints {
	return s.Sprints
}

func (s *ListSprintsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSprintsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSprintsResponseBody) SetErrorCode(v string) *ListSprintsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSprintsResponseBody) SetErrorMsg(v string) *ListSprintsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListSprintsResponseBody) SetMaxResults(v int64) *ListSprintsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSprintsResponseBody) SetNextToken(v string) *ListSprintsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSprintsResponseBody) SetRequestId(v string) *ListSprintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSprintsResponseBody) SetSprints(v []*ListSprintsResponseBodySprints) *ListSprintsResponseBody {
	s.Sprints = v
	return s
}

func (s *ListSprintsResponseBody) SetSuccess(v bool) *ListSprintsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSprintsResponseBody) SetTotalCount(v int64) *ListSprintsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSprintsResponseBody) Validate() error {
	if s.Sprints != nil {
		for _, item := range s.Sprints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSprintsResponseBodySprints struct {
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// xxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1623916393000
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 1623916393000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1623916393000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// demo示例项目
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// public
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// 5e70xxxxxxcd000xxxxe96
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// 1638403200000
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// TODO
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListSprintsResponseBodySprints) String() string {
	return dara.Prettify(s)
}

func (s ListSprintsResponseBodySprints) GoString() string {
	return s.String()
}

func (s *ListSprintsResponseBodySprints) GetCreator() *string {
	return s.Creator
}

func (s *ListSprintsResponseBodySprints) GetDescription() *string {
	return s.Description
}

func (s *ListSprintsResponseBodySprints) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListSprintsResponseBodySprints) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSprintsResponseBodySprints) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListSprintsResponseBodySprints) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListSprintsResponseBodySprints) GetModifier() *string {
	return s.Modifier
}

func (s *ListSprintsResponseBodySprints) GetName() *string {
	return s.Name
}

func (s *ListSprintsResponseBodySprints) GetScope() *string {
	return s.Scope
}

func (s *ListSprintsResponseBodySprints) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListSprintsResponseBodySprints) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListSprintsResponseBodySprints) GetStatus() *string {
	return s.Status
}

func (s *ListSprintsResponseBodySprints) SetCreator(v string) *ListSprintsResponseBodySprints {
	s.Creator = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetDescription(v string) *ListSprintsResponseBodySprints {
	s.Description = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetEndDate(v int64) *ListSprintsResponseBodySprints {
	s.EndDate = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetGmtCreate(v int64) *ListSprintsResponseBodySprints {
	s.GmtCreate = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetGmtModified(v int64) *ListSprintsResponseBodySprints {
	s.GmtModified = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetIdentifier(v string) *ListSprintsResponseBodySprints {
	s.Identifier = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetModifier(v string) *ListSprintsResponseBodySprints {
	s.Modifier = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetName(v string) *ListSprintsResponseBodySprints {
	s.Name = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetScope(v string) *ListSprintsResponseBodySprints {
	s.Scope = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetSpaceIdentifier(v string) *ListSprintsResponseBodySprints {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetStartDate(v int64) *ListSprintsResponseBodySprints {
	s.StartDate = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetStatus(v string) *ListSprintsResponseBodySprints {
	s.Status = &v
	return s
}

func (s *ListSprintsResponseBodySprints) Validate() error {
	return dara.Validate(s)
}
