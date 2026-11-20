// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledBackupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetScheduledBackupConfigRequest
	GetRegionId() *string
	SetInstanceId(v string) *GetScheduledBackupConfigRequest
	GetInstanceId() *string
	SetScheduleType(v string) *GetScheduledBackupConfigRequest
	GetScheduleType() *string
}

type GetScheduledBackupConfigRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgprecn-cn-9lb3bjg1n003
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The backup type. By default, periodic and manual backup configurations are returned. To return the remote backup configuration, set this parameter to remote.
	//
	// example:
	//
	// remote
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
}

func (s GetScheduledBackupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledBackupConfigRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledBackupConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScheduledBackupConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetScheduledBackupConfigRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetScheduledBackupConfigRequest) SetRegionId(v string) *GetScheduledBackupConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetScheduledBackupConfigRequest) SetInstanceId(v string) *GetScheduledBackupConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScheduledBackupConfigRequest) SetScheduleType(v string) *GetScheduledBackupConfigRequest {
	s.ScheduleType = &v
	return s
}

func (s *GetScheduledBackupConfigRequest) Validate() error {
	return dara.Validate(s)
}
