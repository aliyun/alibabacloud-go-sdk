// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditLog interface {
	dara.Model
	String() string
	GoString() string
	SetActedAt(v string) *AuditLog
	GetActedAt() *string
	SetActionCategory(v string) *AuditLog
	GetActionCategory() *string
	SetActionType(v string) *AuditLog
	GetActionType() *string
	SetActorId(v string) *AuditLog
	GetActorId() *string
	SetActorName(v string) *AuditLog
	GetActorName() *string
	SetActorType(v string) *AuditLog
	GetActorType() *string
	SetClientDevice(v string) *AuditLog
	GetClientDevice() *string
	SetClientIp(v string) *AuditLog
	GetClientIp() *string
	SetClientType(v string) *AuditLog
	GetClientType() *string
	SetClientVersion(v string) *AuditLog
	GetClientVersion() *string
	SetDetail(v *AuditLogDetail) *AuditLog
	GetDetail() *AuditLogDetail
	SetFilePathType(v string) *AuditLog
	GetFilePathType() *string
	SetLogId(v string) *AuditLog
	GetLogId() *string
	SetObjectId(v string) *AuditLog
	GetObjectId() *string
	SetObjectName(v string) *AuditLog
	GetObjectName() *string
}

type AuditLog struct {
	ActedAt        *string         `json:"acted_at,omitempty" xml:"acted_at,omitempty"`
	ActionCategory *string         `json:"action_category,omitempty" xml:"action_category,omitempty"`
	ActionType     *string         `json:"action_type,omitempty" xml:"action_type,omitempty"`
	ActorId        *string         `json:"actor_id,omitempty" xml:"actor_id,omitempty"`
	ActorName      *string         `json:"actor_name,omitempty" xml:"actor_name,omitempty"`
	ActorType      *string         `json:"actor_type,omitempty" xml:"actor_type,omitempty"`
	ClientDevice   *string         `json:"client_device,omitempty" xml:"client_device,omitempty"`
	ClientIp       *string         `json:"client_ip,omitempty" xml:"client_ip,omitempty"`
	ClientType     *string         `json:"client_type,omitempty" xml:"client_type,omitempty"`
	ClientVersion  *string         `json:"client_version,omitempty" xml:"client_version,omitempty"`
	Detail         *AuditLogDetail `json:"detail,omitempty" xml:"detail,omitempty"`
	FilePathType   *string         `json:"file_path_type,omitempty" xml:"file_path_type,omitempty"`
	LogId          *string         `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ObjectId       *string         `json:"object_id,omitempty" xml:"object_id,omitempty"`
	ObjectName     *string         `json:"object_name,omitempty" xml:"object_name,omitempty"`
}

func (s AuditLog) String() string {
	return dara.Prettify(s)
}

func (s AuditLog) GoString() string {
	return s.String()
}

func (s *AuditLog) GetActedAt() *string {
	return s.ActedAt
}

func (s *AuditLog) GetActionCategory() *string {
	return s.ActionCategory
}

func (s *AuditLog) GetActionType() *string {
	return s.ActionType
}

func (s *AuditLog) GetActorId() *string {
	return s.ActorId
}

func (s *AuditLog) GetActorName() *string {
	return s.ActorName
}

func (s *AuditLog) GetActorType() *string {
	return s.ActorType
}

func (s *AuditLog) GetClientDevice() *string {
	return s.ClientDevice
}

func (s *AuditLog) GetClientIp() *string {
	return s.ClientIp
}

func (s *AuditLog) GetClientType() *string {
	return s.ClientType
}

func (s *AuditLog) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *AuditLog) GetDetail() *AuditLogDetail {
	return s.Detail
}

func (s *AuditLog) GetFilePathType() *string {
	return s.FilePathType
}

func (s *AuditLog) GetLogId() *string {
	return s.LogId
}

func (s *AuditLog) GetObjectId() *string {
	return s.ObjectId
}

func (s *AuditLog) GetObjectName() *string {
	return s.ObjectName
}

func (s *AuditLog) SetActedAt(v string) *AuditLog {
	s.ActedAt = &v
	return s
}

func (s *AuditLog) SetActionCategory(v string) *AuditLog {
	s.ActionCategory = &v
	return s
}

func (s *AuditLog) SetActionType(v string) *AuditLog {
	s.ActionType = &v
	return s
}

func (s *AuditLog) SetActorId(v string) *AuditLog {
	s.ActorId = &v
	return s
}

func (s *AuditLog) SetActorName(v string) *AuditLog {
	s.ActorName = &v
	return s
}

func (s *AuditLog) SetActorType(v string) *AuditLog {
	s.ActorType = &v
	return s
}

func (s *AuditLog) SetClientDevice(v string) *AuditLog {
	s.ClientDevice = &v
	return s
}

func (s *AuditLog) SetClientIp(v string) *AuditLog {
	s.ClientIp = &v
	return s
}

func (s *AuditLog) SetClientType(v string) *AuditLog {
	s.ClientType = &v
	return s
}

func (s *AuditLog) SetClientVersion(v string) *AuditLog {
	s.ClientVersion = &v
	return s
}

func (s *AuditLog) SetDetail(v *AuditLogDetail) *AuditLog {
	s.Detail = v
	return s
}

func (s *AuditLog) SetFilePathType(v string) *AuditLog {
	s.FilePathType = &v
	return s
}

func (s *AuditLog) SetLogId(v string) *AuditLog {
	s.LogId = &v
	return s
}

func (s *AuditLog) SetObjectId(v string) *AuditLog {
	s.ObjectId = &v
	return s
}

func (s *AuditLog) SetObjectName(v string) *AuditLog {
	s.ObjectName = &v
	return s
}

func (s *AuditLog) Validate() error {
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
