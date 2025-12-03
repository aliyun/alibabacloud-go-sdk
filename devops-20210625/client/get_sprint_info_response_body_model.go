// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSprintInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetSprintInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSprintInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetSprintInfoResponseBody
	GetRequestId() *string
	SetSprint(v *GetSprintInfoResponseBodySprint) *GetSprintInfoResponseBody
	GetSprint() *GetSprintInfoResponseBodySprint
	SetSuccess(v bool) *GetSprintInfoResponseBody
	GetSuccess() *bool
}

type GetSprintInfoResponseBody struct {
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
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Sprint    *GetSprintInfoResponseBodySprint `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSprintInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSprintInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSprintInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSprintInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSprintInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSprintInfoResponseBody) GetSprint() *GetSprintInfoResponseBodySprint {
	return s.Sprint
}

func (s *GetSprintInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSprintInfoResponseBody) SetErrorCode(v string) *GetSprintInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSprintInfoResponseBody) SetErrorMessage(v string) *GetSprintInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSprintInfoResponseBody) SetRequestId(v string) *GetSprintInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSprintInfoResponseBody) SetSprint(v *GetSprintInfoResponseBodySprint) *GetSprintInfoResponseBody {
	s.Sprint = v
	return s
}

func (s *GetSprintInfoResponseBody) SetSuccess(v bool) *GetSprintInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSprintInfoResponseBody) Validate() error {
	if s.Sprint != nil {
		if err := s.Sprint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSprintInfoResponseBodySprint struct {
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
	Name   *string   `json:"name,omitempty" xml:"name,omitempty"`
	Owners []*string `json:"owners,omitempty" xml:"owners,omitempty" type:"Repeated"`
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
	// Todo
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetSprintInfoResponseBodySprint) String() string {
	return dara.Prettify(s)
}

func (s GetSprintInfoResponseBodySprint) GoString() string {
	return s.String()
}

func (s *GetSprintInfoResponseBodySprint) GetCreator() *string {
	return s.Creator
}

func (s *GetSprintInfoResponseBodySprint) GetDescription() *string {
	return s.Description
}

func (s *GetSprintInfoResponseBodySprint) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSprintInfoResponseBodySprint) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetSprintInfoResponseBodySprint) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetSprintInfoResponseBodySprint) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetSprintInfoResponseBodySprint) GetModifier() *string {
	return s.Modifier
}

func (s *GetSprintInfoResponseBodySprint) GetName() *string {
	return s.Name
}

func (s *GetSprintInfoResponseBodySprint) GetOwners() []*string {
	return s.Owners
}

func (s *GetSprintInfoResponseBodySprint) GetScope() *string {
	return s.Scope
}

func (s *GetSprintInfoResponseBodySprint) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *GetSprintInfoResponseBodySprint) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSprintInfoResponseBodySprint) GetStatus() *string {
	return s.Status
}

func (s *GetSprintInfoResponseBodySprint) SetCreator(v string) *GetSprintInfoResponseBodySprint {
	s.Creator = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetDescription(v string) *GetSprintInfoResponseBodySprint {
	s.Description = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetEndDate(v int64) *GetSprintInfoResponseBodySprint {
	s.EndDate = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetGmtCreate(v int64) *GetSprintInfoResponseBodySprint {
	s.GmtCreate = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetGmtModified(v int64) *GetSprintInfoResponseBodySprint {
	s.GmtModified = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetIdentifier(v string) *GetSprintInfoResponseBodySprint {
	s.Identifier = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetModifier(v string) *GetSprintInfoResponseBodySprint {
	s.Modifier = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetName(v string) *GetSprintInfoResponseBodySprint {
	s.Name = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetOwners(v []*string) *GetSprintInfoResponseBodySprint {
	s.Owners = v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetScope(v string) *GetSprintInfoResponseBodySprint {
	s.Scope = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetSpaceIdentifier(v string) *GetSprintInfoResponseBodySprint {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetStartDate(v int64) *GetSprintInfoResponseBodySprint {
	s.StartDate = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetStatus(v string) *GetSprintInfoResponseBodySprint {
	s.Status = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) Validate() error {
	return dara.Validate(s)
}
