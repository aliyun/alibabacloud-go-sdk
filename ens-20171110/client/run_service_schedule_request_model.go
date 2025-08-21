// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunServiceScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunServiceScheduleRequest
	GetAppId() *string
	SetClientIp(v string) *RunServiceScheduleRequest
	GetClientIp() *string
	SetDirectorys(v string) *RunServiceScheduleRequest
	GetDirectorys() *string
	SetPodConfigName(v string) *RunServiceScheduleRequest
	GetPodConfigName() *string
	SetPreLockedTimeout(v int32) *RunServiceScheduleRequest
	GetPreLockedTimeout() *int32
	SetScheduleStrategy(v string) *RunServiceScheduleRequest
	GetScheduleStrategy() *string
	SetServiceAction(v string) *RunServiceScheduleRequest
	GetServiceAction() *string
	SetServiceCommands(v string) *RunServiceScheduleRequest
	GetServiceCommands() *string
	SetUuid(v string) *RunServiceScheduleRequest
	GetUuid() *string
}

type RunServiceScheduleRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 474bdef0-d149-4695-abfb-52912d9143f0
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The IP address of the client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 180.166.45.146
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The directory to which the data file is mounted. The value must be a full path and cannot be \\"/../\\". Example: ["/data/app01", "/data/user"]. Specify the relative path of the virtual device for this parameter. For example, specify /data for this parameter when the actual path of the virtual device is /data/{input path}.
	//
	// example:
	//
	// [\\"/data/app01\\", \\"/data/user\\"]
	Directorys *string `json:"Directorys,omitempty" xml:"Directorys,omitempty"`
	// The parameter does not take effect.
	//
	// example:
	//
	// android
	PodConfigName *string `json:"PodConfigName,omitempty" xml:"PodConfigName,omitempty"`
	// The maximum duration for locking an idle device. Unit: seconds. This parameter takes effect only if you set ServiceAction to PreSchedule. Default value: 300.
	//
	// example:
	//
	// 300
	PreLockedTimeout *int32 `json:"PreLockedTimeout,omitempty" xml:"PreLockedTimeout,omitempty"`
	// The scheduling policy of the device. The value must be a JSON string.
	//
	// example:
	//
	// {\\"selectLevel\\": \\"RegionId\\", \\"values\\": [\\"cn-chengdu-telecom-2\\"]  }
	ScheduleStrategy *string `json:"ScheduleStrategy,omitempty" xml:"ScheduleStrategy,omitempty"`
	// The scheduling operation. The value must be of the enumeration type. Valid values:
	//
	// Container scenario:
	//
	// 	- Start: selects and activates an idle cloud device.
	//
	// 	- Stop: stops and releases the cloud device.
	//
	// 	- Console: performs the scheduling operation when the device is in the scheduling state.
	//
	// Bare metal instance or virtual machine scenario:
	//
	// 	- PreSchedule: locks a virtual machine instance for scheduling.
	//
	// 	- Confirm: confirms the scheduling operation.
	//
	// 	- Cancel: cancels the scheduling operation.
	//
	// 	- Console: performs the scheduling operation when the device is in the scheduling state.
	//
	// This parameter is required.
	//
	// example:
	//
	// Start
	ServiceAction *string `json:"ServiceAction,omitempty" xml:"ServiceAction,omitempty"`
	// The service commands. The value must be a JSON string.
	//
	// example:
	//
	// [{\\"containerName\\": \\"android\\",       \\"commands\\":[\\"cat /tmp/token.json\\"]    },    {      \\"commands\\":[\\"ls -l /data/data\\"]    }]
	ServiceCommands *string `json:"ServiceCommands,omitempty" xml:"ServiceCommands,omitempty"`
	// The UUID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// hdm_d4f1059a8e1afc0956bd71b497faa433
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RunServiceScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s RunServiceScheduleRequest) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunServiceScheduleRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *RunServiceScheduleRequest) GetDirectorys() *string {
	return s.Directorys
}

func (s *RunServiceScheduleRequest) GetPodConfigName() *string {
	return s.PodConfigName
}

func (s *RunServiceScheduleRequest) GetPreLockedTimeout() *int32 {
	return s.PreLockedTimeout
}

func (s *RunServiceScheduleRequest) GetScheduleStrategy() *string {
	return s.ScheduleStrategy
}

func (s *RunServiceScheduleRequest) GetServiceAction() *string {
	return s.ServiceAction
}

func (s *RunServiceScheduleRequest) GetServiceCommands() *string {
	return s.ServiceCommands
}

func (s *RunServiceScheduleRequest) GetUuid() *string {
	return s.Uuid
}

func (s *RunServiceScheduleRequest) SetAppId(v string) *RunServiceScheduleRequest {
	s.AppId = &v
	return s
}

func (s *RunServiceScheduleRequest) SetClientIp(v string) *RunServiceScheduleRequest {
	s.ClientIp = &v
	return s
}

func (s *RunServiceScheduleRequest) SetDirectorys(v string) *RunServiceScheduleRequest {
	s.Directorys = &v
	return s
}

func (s *RunServiceScheduleRequest) SetPodConfigName(v string) *RunServiceScheduleRequest {
	s.PodConfigName = &v
	return s
}

func (s *RunServiceScheduleRequest) SetPreLockedTimeout(v int32) *RunServiceScheduleRequest {
	s.PreLockedTimeout = &v
	return s
}

func (s *RunServiceScheduleRequest) SetScheduleStrategy(v string) *RunServiceScheduleRequest {
	s.ScheduleStrategy = &v
	return s
}

func (s *RunServiceScheduleRequest) SetServiceAction(v string) *RunServiceScheduleRequest {
	s.ServiceAction = &v
	return s
}

func (s *RunServiceScheduleRequest) SetServiceCommands(v string) *RunServiceScheduleRequest {
	s.ServiceCommands = &v
	return s
}

func (s *RunServiceScheduleRequest) SetUuid(v string) *RunServiceScheduleRequest {
	s.Uuid = &v
	return s
}

func (s *RunServiceScheduleRequest) Validate() error {
	return dara.Validate(s)
}
