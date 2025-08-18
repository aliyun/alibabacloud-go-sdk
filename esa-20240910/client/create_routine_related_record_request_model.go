// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRelatedRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateRoutineRelatedRecordRequest
	GetName() *string
	SetRecordName(v string) *CreateRoutineRelatedRecordRequest
	GetRecordName() *string
	SetSiteId(v int64) *CreateRoutineRelatedRecordRequest
	GetSiteId() *int64
}

type CreateRoutineRelatedRecordRequest struct {
	// The routine name.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateRoutineRelatedRecord
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The record name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-record-1.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The website ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54362329990032
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateRoutineRelatedRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRelatedRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineRelatedRecordRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineRelatedRecordRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateRoutineRelatedRecordRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRoutineRelatedRecordRequest) SetName(v string) *CreateRoutineRelatedRecordRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineRelatedRecordRequest) SetRecordName(v string) *CreateRoutineRelatedRecordRequest {
	s.RecordName = &v
	return s
}

func (s *CreateRoutineRelatedRecordRequest) SetSiteId(v int64) *CreateRoutineRelatedRecordRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRoutineRelatedRecordRequest) Validate() error {
	return dara.Validate(s)
}
