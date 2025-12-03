// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActualTime(v string) *CreateWorkitemRecordRequest
	GetActualTime() *string
	SetDescription(v string) *CreateWorkitemRecordRequest
	GetDescription() *string
	SetGmtEnd(v string) *CreateWorkitemRecordRequest
	GetGmtEnd() *string
	SetGmtStart(v string) *CreateWorkitemRecordRequest
	GetGmtStart() *string
	SetRecordUserIdentifier(v string) *CreateWorkitemRecordRequest
	GetRecordUserIdentifier() *string
	SetType(v string) *CreateWorkitemRecordRequest
	GetType() *string
	SetWorkitemIdentifier(v string) *CreateWorkitemRecordRequest
	GetWorkitemIdentifier() *string
}

type CreateWorkitemRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12
	ActualTime  *string `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1646323200000
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1667205617061
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1316458xxxxx41068
	RecordUserIdentifier *string `json:"recordUserIdentifier,omitempty" xml:"recordUserIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordRequest) GetActualTime() *string {
	return s.ActualTime
}

func (s *CreateWorkitemRecordRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkitemRecordRequest) GetGmtEnd() *string {
	return s.GmtEnd
}

func (s *CreateWorkitemRecordRequest) GetGmtStart() *string {
	return s.GmtStart
}

func (s *CreateWorkitemRecordRequest) GetRecordUserIdentifier() *string {
	return s.RecordUserIdentifier
}

func (s *CreateWorkitemRecordRequest) GetType() *string {
	return s.Type
}

func (s *CreateWorkitemRecordRequest) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *CreateWorkitemRecordRequest) SetActualTime(v string) *CreateWorkitemRecordRequest {
	s.ActualTime = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetDescription(v string) *CreateWorkitemRecordRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetGmtEnd(v string) *CreateWorkitemRecordRequest {
	s.GmtEnd = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetGmtStart(v string) *CreateWorkitemRecordRequest {
	s.GmtStart = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetRecordUserIdentifier(v string) *CreateWorkitemRecordRequest {
	s.RecordUserIdentifier = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetType(v string) *CreateWorkitemRecordRequest {
	s.Type = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetWorkitemIdentifier(v string) *CreateWorkitemRecordRequest {
	s.WorkitemIdentifier = &v
	return s
}

func (s *CreateWorkitemRecordRequest) Validate() error {
	return dara.Validate(s)
}
