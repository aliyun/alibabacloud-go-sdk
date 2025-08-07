// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *GetBackupResponseBody
	GetBackupId() *string
	SetCreateTime(v string) *GetBackupResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetBackupResponseBody
	GetDescription() *string
	SetEndTime(v string) *GetBackupResponseBody
	GetEndTime() *string
	SetModifiedTime(v string) *GetBackupResponseBody
	GetModifiedTime() *string
	SetRequestId(v string) *GetBackupResponseBody
	GetRequestId() *string
	SetServiceInstanceId(v string) *GetBackupResponseBody
	GetServiceInstanceId() *string
	SetStartTime(v string) *GetBackupResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetBackupResponseBody
	GetStatus() *string
	SetStatusDetail(v string) *GetBackupResponseBody
	GetStatusDetail() *string
}

type GetBackupResponseBody struct {
	// The backup ID.
	//
	// example:
	//
	// backup-cad4a85ff5e340388b93
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The creation time of the backup task.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the backup task.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the backup task.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The update time of the backup task.
	//
	// example:
	//
	// 2025-03-10T19:26:20Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The start time of the backup task.
	//
	// example:
	//
	// 2022-01-01T11:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the backup task.
	//
	// 	- Creating
	//
	// 	- Created
	//
	// 	- CreateFailed
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeleteFailed
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment instance status.
	//
	// example:
	//
	// Disk i-xxxx backup failed, error message: error
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
}

func (s GetBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBackupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *GetBackupResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetBackupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetBackupResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetBackupResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBackupResponseBody) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetBackupResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetBackupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetBackupResponseBody) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *GetBackupResponseBody) SetBackupId(v string) *GetBackupResponseBody {
	s.BackupId = &v
	return s
}

func (s *GetBackupResponseBody) SetCreateTime(v string) *GetBackupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetBackupResponseBody) SetDescription(v string) *GetBackupResponseBody {
	s.Description = &v
	return s
}

func (s *GetBackupResponseBody) SetEndTime(v string) *GetBackupResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetBackupResponseBody) SetModifiedTime(v string) *GetBackupResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetBackupResponseBody) SetRequestId(v string) *GetBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupResponseBody) SetServiceInstanceId(v string) *GetBackupResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetBackupResponseBody) SetStartTime(v string) *GetBackupResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetBackupResponseBody) SetStatus(v string) *GetBackupResponseBody {
	s.Status = &v
	return s
}

func (s *GetBackupResponseBody) SetStatusDetail(v string) *GetBackupResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
