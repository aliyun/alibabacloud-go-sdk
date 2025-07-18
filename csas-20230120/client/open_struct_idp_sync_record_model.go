// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructIdpSyncRecord interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *OpenStructIdpSyncRecord
	GetAction() *string
	SetIdpRawJson(v string) *OpenStructIdpSyncRecord
	GetIdpRawJson() *string
	SetIdpResourceId(v string) *OpenStructIdpSyncRecord
	GetIdpResourceId() *string
	SetName(v string) *OpenStructIdpSyncRecord
	GetName() *string
	SetRecordType(v string) *OpenStructIdpSyncRecord
	GetRecordType() *string
	SetSaseRawJson(v string) *OpenStructIdpSyncRecord
	GetSaseRawJson() *string
	SetSaseResourceId(v string) *OpenStructIdpSyncRecord
	GetSaseResourceId() *string
	SetSuccess(v bool) *OpenStructIdpSyncRecord
	GetSuccess() *bool
	SetSyncRecordId(v string) *OpenStructIdpSyncRecord
	GetSyncRecordId() *string
	SetSyncTaskId(v string) *OpenStructIdpSyncRecord
	GetSyncTaskId() *string
	SetTimeUnix(v string) *OpenStructIdpSyncRecord
	GetTimeUnix() *string
}

type OpenStructIdpSyncRecord struct {
	Action         *string `json:"Action,omitempty" xml:"Action,omitempty"`
	IdpRawJson     *string `json:"IdpRawJson,omitempty" xml:"IdpRawJson,omitempty"`
	IdpResourceId  *string `json:"IdpResourceId,omitempty" xml:"IdpResourceId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecordType     *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	SaseRawJson    *string `json:"SaseRawJson,omitempty" xml:"SaseRawJson,omitempty"`
	SaseResourceId *string `json:"SaseResourceId,omitempty" xml:"SaseResourceId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SyncRecordId   *string `json:"SyncRecordId,omitempty" xml:"SyncRecordId,omitempty"`
	SyncTaskId     *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	TimeUnix       *string `json:"TimeUnix,omitempty" xml:"TimeUnix,omitempty"`
}

func (s OpenStructIdpSyncRecord) String() string {
	return dara.Prettify(s)
}

func (s OpenStructIdpSyncRecord) GoString() string {
	return s.String()
}

func (s *OpenStructIdpSyncRecord) GetAction() *string {
	return s.Action
}

func (s *OpenStructIdpSyncRecord) GetIdpRawJson() *string {
	return s.IdpRawJson
}

func (s *OpenStructIdpSyncRecord) GetIdpResourceId() *string {
	return s.IdpResourceId
}

func (s *OpenStructIdpSyncRecord) GetName() *string {
	return s.Name
}

func (s *OpenStructIdpSyncRecord) GetRecordType() *string {
	return s.RecordType
}

func (s *OpenStructIdpSyncRecord) GetSaseRawJson() *string {
	return s.SaseRawJson
}

func (s *OpenStructIdpSyncRecord) GetSaseResourceId() *string {
	return s.SaseResourceId
}

func (s *OpenStructIdpSyncRecord) GetSuccess() *bool {
	return s.Success
}

func (s *OpenStructIdpSyncRecord) GetSyncRecordId() *string {
	return s.SyncRecordId
}

func (s *OpenStructIdpSyncRecord) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *OpenStructIdpSyncRecord) GetTimeUnix() *string {
	return s.TimeUnix
}

func (s *OpenStructIdpSyncRecord) SetAction(v string) *OpenStructIdpSyncRecord {
	s.Action = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetIdpRawJson(v string) *OpenStructIdpSyncRecord {
	s.IdpRawJson = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetIdpResourceId(v string) *OpenStructIdpSyncRecord {
	s.IdpResourceId = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetName(v string) *OpenStructIdpSyncRecord {
	s.Name = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetRecordType(v string) *OpenStructIdpSyncRecord {
	s.RecordType = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetSaseRawJson(v string) *OpenStructIdpSyncRecord {
	s.SaseRawJson = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetSaseResourceId(v string) *OpenStructIdpSyncRecord {
	s.SaseResourceId = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetSuccess(v bool) *OpenStructIdpSyncRecord {
	s.Success = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetSyncRecordId(v string) *OpenStructIdpSyncRecord {
	s.SyncRecordId = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetSyncTaskId(v string) *OpenStructIdpSyncRecord {
	s.SyncTaskId = &v
	return s
}

func (s *OpenStructIdpSyncRecord) SetTimeUnix(v string) *OpenStructIdpSyncRecord {
	s.TimeUnix = &v
	return s
}

func (s *OpenStructIdpSyncRecord) Validate() error {
	return dara.Validate(s)
}
