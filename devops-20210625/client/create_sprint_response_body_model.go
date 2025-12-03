// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSprintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateSprintResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateSprintResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateSprintResponseBody
	GetRequestId() *string
	SetSprint(v *CreateSprintResponseBodySprint) *CreateSprintResponseBody
	GetSprint() *CreateSprintResponseBodySprint
	SetSuccess(v bool) *CreateSprintResponseBody
	GetSuccess() *bool
}

type CreateSprintResponseBody struct {
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
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Sprint    *CreateSprintResponseBodySprint `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSprintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSprintResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSprintResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateSprintResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateSprintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSprintResponseBody) GetSprint() *CreateSprintResponseBodySprint {
	return s.Sprint
}

func (s *CreateSprintResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSprintResponseBody) SetErrorCode(v string) *CreateSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSprintResponseBody) SetErrorMsg(v string) *CreateSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateSprintResponseBody) SetRequestId(v string) *CreateSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSprintResponseBody) SetSprint(v *CreateSprintResponseBodySprint) *CreateSprintResponseBody {
	s.Sprint = v
	return s
}

func (s *CreateSprintResponseBody) SetSuccess(v bool) *CreateSprintResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSprintResponseBody) Validate() error {
	if s.Sprint != nil {
		if err := s.Sprint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSprintResponseBodySprint struct {
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

func (s CreateSprintResponseBodySprint) String() string {
	return dara.Prettify(s)
}

func (s CreateSprintResponseBodySprint) GoString() string {
	return s.String()
}

func (s *CreateSprintResponseBodySprint) GetCreator() *string {
	return s.Creator
}

func (s *CreateSprintResponseBodySprint) GetDescription() *string {
	return s.Description
}

func (s *CreateSprintResponseBodySprint) GetEndDate() *int64 {
	return s.EndDate
}

func (s *CreateSprintResponseBodySprint) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateSprintResponseBodySprint) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CreateSprintResponseBodySprint) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateSprintResponseBodySprint) GetModifier() *string {
	return s.Modifier
}

func (s *CreateSprintResponseBodySprint) GetName() *string {
	return s.Name
}

func (s *CreateSprintResponseBodySprint) GetScope() *string {
	return s.Scope
}

func (s *CreateSprintResponseBodySprint) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *CreateSprintResponseBodySprint) GetStartDate() *int64 {
	return s.StartDate
}

func (s *CreateSprintResponseBodySprint) GetStatus() *string {
	return s.Status
}

func (s *CreateSprintResponseBodySprint) SetCreator(v string) *CreateSprintResponseBodySprint {
	s.Creator = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetDescription(v string) *CreateSprintResponseBodySprint {
	s.Description = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetEndDate(v int64) *CreateSprintResponseBodySprint {
	s.EndDate = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetGmtCreate(v int64) *CreateSprintResponseBodySprint {
	s.GmtCreate = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetGmtModified(v int64) *CreateSprintResponseBodySprint {
	s.GmtModified = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetIdentifier(v string) *CreateSprintResponseBodySprint {
	s.Identifier = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetModifier(v string) *CreateSprintResponseBodySprint {
	s.Modifier = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetName(v string) *CreateSprintResponseBodySprint {
	s.Name = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetScope(v string) *CreateSprintResponseBodySprint {
	s.Scope = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetSpaceIdentifier(v string) *CreateSprintResponseBodySprint {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetStartDate(v int64) *CreateSprintResponseBodySprint {
	s.StartDate = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetStatus(v string) *CreateSprintResponseBodySprint {
	s.Status = &v
	return s
}

func (s *CreateSprintResponseBodySprint) Validate() error {
	return dara.Validate(s)
}
