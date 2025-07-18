// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructIdpSyncTask interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentFailed(v int64) *OpenStructIdpSyncTask
	GetDepartmentFailed() *int64
	SetDepartmentTotal(v int64) *OpenStructIdpSyncTask
	GetDepartmentTotal() *int64
	SetEndTimeUnix(v int64) *OpenStructIdpSyncTask
	GetEndTimeUnix() *int64
	SetFailType(v string) *OpenStructIdpSyncTask
	GetFailType() *string
	SetIdpConfigId(v []byte) *OpenStructIdpSyncTask
	GetIdpConfigId() []byte
	SetStartTimeUnix(v int64) *OpenStructIdpSyncTask
	GetStartTimeUnix() *int64
	SetStatus(v string) *OpenStructIdpSyncTask
	GetStatus() *string
	SetSyncTaskId(v string) *OpenStructIdpSyncTask
	GetSyncTaskId() *string
	SetUpdateTimeUnix(v int64) *OpenStructIdpSyncTask
	GetUpdateTimeUnix() *int64
	SetUserFailed(v int64) *OpenStructIdpSyncTask
	GetUserFailed() *int64
	SetUserTotal(v int64) *OpenStructIdpSyncTask
	GetUserTotal() *int64
}

type OpenStructIdpSyncTask struct {
	DepartmentFailed *int64  `json:"DepartmentFailed,omitempty" xml:"DepartmentFailed,omitempty"`
	DepartmentTotal  *int64  `json:"DepartmentTotal,omitempty" xml:"DepartmentTotal,omitempty"`
	EndTimeUnix      *int64  `json:"EndTimeUnix,omitempty" xml:"EndTimeUnix,omitempty"`
	FailType         *string `json:"FailType,omitempty" xml:"FailType,omitempty"`
	IdpConfigId      []byte  `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	StartTimeUnix    *int64  `json:"StartTimeUnix,omitempty" xml:"StartTimeUnix,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SyncTaskId       *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	UpdateTimeUnix   *int64  `json:"UpdateTimeUnix,omitempty" xml:"UpdateTimeUnix,omitempty"`
	UserFailed       *int64  `json:"UserFailed,omitempty" xml:"UserFailed,omitempty"`
	UserTotal        *int64  `json:"UserTotal,omitempty" xml:"UserTotal,omitempty"`
}

func (s OpenStructIdpSyncTask) String() string {
	return dara.Prettify(s)
}

func (s OpenStructIdpSyncTask) GoString() string {
	return s.String()
}

func (s *OpenStructIdpSyncTask) GetDepartmentFailed() *int64 {
	return s.DepartmentFailed
}

func (s *OpenStructIdpSyncTask) GetDepartmentTotal() *int64 {
	return s.DepartmentTotal
}

func (s *OpenStructIdpSyncTask) GetEndTimeUnix() *int64 {
	return s.EndTimeUnix
}

func (s *OpenStructIdpSyncTask) GetFailType() *string {
	return s.FailType
}

func (s *OpenStructIdpSyncTask) GetIdpConfigId() []byte {
	return s.IdpConfigId
}

func (s *OpenStructIdpSyncTask) GetStartTimeUnix() *int64 {
	return s.StartTimeUnix
}

func (s *OpenStructIdpSyncTask) GetStatus() *string {
	return s.Status
}

func (s *OpenStructIdpSyncTask) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *OpenStructIdpSyncTask) GetUpdateTimeUnix() *int64 {
	return s.UpdateTimeUnix
}

func (s *OpenStructIdpSyncTask) GetUserFailed() *int64 {
	return s.UserFailed
}

func (s *OpenStructIdpSyncTask) GetUserTotal() *int64 {
	return s.UserTotal
}

func (s *OpenStructIdpSyncTask) SetDepartmentFailed(v int64) *OpenStructIdpSyncTask {
	s.DepartmentFailed = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetDepartmentTotal(v int64) *OpenStructIdpSyncTask {
	s.DepartmentTotal = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetEndTimeUnix(v int64) *OpenStructIdpSyncTask {
	s.EndTimeUnix = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetFailType(v string) *OpenStructIdpSyncTask {
	s.FailType = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetIdpConfigId(v []byte) *OpenStructIdpSyncTask {
	s.IdpConfigId = v
	return s
}

func (s *OpenStructIdpSyncTask) SetStartTimeUnix(v int64) *OpenStructIdpSyncTask {
	s.StartTimeUnix = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetStatus(v string) *OpenStructIdpSyncTask {
	s.Status = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetSyncTaskId(v string) *OpenStructIdpSyncTask {
	s.SyncTaskId = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetUpdateTimeUnix(v int64) *OpenStructIdpSyncTask {
	s.UpdateTimeUnix = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetUserFailed(v int64) *OpenStructIdpSyncTask {
	s.UserFailed = &v
	return s
}

func (s *OpenStructIdpSyncTask) SetUserTotal(v int64) *OpenStructIdpSyncTask {
	s.UserTotal = &v
	return s
}

func (s *OpenStructIdpSyncTask) Validate() error {
	return dara.Validate(s)
}
