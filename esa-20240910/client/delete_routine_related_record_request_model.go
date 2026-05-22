// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineRelatedRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteRoutineRelatedRecordRequest
	GetName() *string
	SetRecordId(v int64) *DeleteRoutineRelatedRecordRequest
	GetRecordId() *int64
	SetRecordName(v string) *DeleteRoutineRelatedRecordRequest
	GetRecordName() *string
	SetSiteId(v int64) *DeleteRoutineRelatedRecordRequest
	GetSiteId() *int64
}

type DeleteRoutineRelatedRecordRequest struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// This parameter is required.
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteRoutineRelatedRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineRelatedRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRelatedRecordRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineRelatedRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DeleteRoutineRelatedRecordRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *DeleteRoutineRelatedRecordRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteRoutineRelatedRecordRequest) SetName(v string) *DeleteRoutineRelatedRecordRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineRelatedRecordRequest) SetRecordId(v int64) *DeleteRoutineRelatedRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DeleteRoutineRelatedRecordRequest) SetRecordName(v string) *DeleteRoutineRelatedRecordRequest {
	s.RecordName = &v
	return s
}

func (s *DeleteRoutineRelatedRecordRequest) SetSiteId(v int64) *DeleteRoutineRelatedRecordRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteRoutineRelatedRecordRequest) Validate() error {
	return dara.Validate(s)
}
