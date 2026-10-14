// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventMigrateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventInfos(v []*BatchEventMigrateInstanceRequestEventInfos) *BatchEventMigrateInstanceRequest
	GetEventInfos() []*BatchEventMigrateInstanceRequestEventInfos
}

type BatchEventMigrateInstanceRequest struct {
	// The list of events.
	EventInfos []*BatchEventMigrateInstanceRequestEventInfos `json:"EventInfos,omitempty" xml:"EventInfos,omitempty" type:"Repeated"`
}

func (s BatchEventMigrateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEventMigrateInstanceRequest) GoString() string {
	return s.String()
}

func (s *BatchEventMigrateInstanceRequest) GetEventInfos() []*BatchEventMigrateInstanceRequestEventInfos {
	return s.EventInfos
}

func (s *BatchEventMigrateInstanceRequest) SetEventInfos(v []*BatchEventMigrateInstanceRequestEventInfos) *BatchEventMigrateInstanceRequest {
	s.EventInfos = v
	return s
}

func (s *BatchEventMigrateInstanceRequest) Validate() error {
	if s.EventInfos != nil {
		for _, item := range s.EventInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchEventMigrateInstanceRequestEventInfos struct {
	// The data migration policy. Valid values:
	//
	// - abandon: Does not migrate data. This is the default value.
	//
	// - force_transfer: Performs forcible migration of EPS resource data.
	//
	// - try_transfer: Migrates data on a best-effort basis.
	//
	// example:
	//
	// abandon
	DataPolicy *string `json:"DataPolicy,omitempty" xml:"DataPolicy,omitempty"`
	// The event ID, in the format of e-*.
	//
	// This parameter is required.
	//
	// example:
	//
	// e-d71ff150945b9c02eb6ebc0016328468
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The O&M task type. Valid values:
	//
	// - immediate: immediate execution.
	//
	// - scheduled: scheduled execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// immediate
	OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	// The instance password. This parameter is optional. If you do not specify this parameter, a random password is used.
	//
	// The password must be 8 to 30 characters in length and must contain uppercase letters, lowercase letters, and digits.
	//
	// Note: Password specification is not supported for scheduled execution.
	//
	// example:
	//
	// AAaa1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The scheduled execution time. The value is a timestamp in milliseconds.
	//
	// This parameter is required when OpsType is set to scheduled.
	//
	// example:
	//
	// 1742452232000
	PlanTime    *int64  `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
	PlanUtcTime *string `json:"PlanUtcTime,omitempty" xml:"PlanUtcTime,omitempty"`
	// The resource ID, in the format of i-*.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-55qi8m11rr53c4i964md8a00l
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s BatchEventMigrateInstanceRequestEventInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchEventMigrateInstanceRequestEventInfos) GoString() string {
	return s.String()
}

func (s *BatchEventMigrateInstanceRequestEventInfos) GetDataPolicy() *string {
	return s.DataPolicy
}

func (s *BatchEventMigrateInstanceRequestEventInfos) GetEventId() *string {
	return s.EventId
}

func (s *BatchEventMigrateInstanceRequestEventInfos) GetOpsType() *string {
	return s.OpsType
}

func (s *BatchEventMigrateInstanceRequestEventInfos) GetPassword() *string {
	return s.Password
}

func (s *BatchEventMigrateInstanceRequestEventInfos) GetPlanTime() *int64 {
	return s.PlanTime
}

func (s *BatchEventMigrateInstanceRequestEventInfos) GetPlanUtcTime() *string {
	return s.PlanUtcTime
}

func (s *BatchEventMigrateInstanceRequestEventInfos) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchEventMigrateInstanceRequestEventInfos) SetDataPolicy(v string) *BatchEventMigrateInstanceRequestEventInfos {
	s.DataPolicy = &v
	return s
}

func (s *BatchEventMigrateInstanceRequestEventInfos) SetEventId(v string) *BatchEventMigrateInstanceRequestEventInfos {
	s.EventId = &v
	return s
}

func (s *BatchEventMigrateInstanceRequestEventInfos) SetOpsType(v string) *BatchEventMigrateInstanceRequestEventInfos {
	s.OpsType = &v
	return s
}

func (s *BatchEventMigrateInstanceRequestEventInfos) SetPassword(v string) *BatchEventMigrateInstanceRequestEventInfos {
	s.Password = &v
	return s
}

func (s *BatchEventMigrateInstanceRequestEventInfos) SetPlanTime(v int64) *BatchEventMigrateInstanceRequestEventInfos {
	s.PlanTime = &v
	return s
}

func (s *BatchEventMigrateInstanceRequestEventInfos) SetPlanUtcTime(v string) *BatchEventMigrateInstanceRequestEventInfos {
	s.PlanUtcTime = &v
	return s
}

func (s *BatchEventMigrateInstanceRequestEventInfos) SetResourceId(v string) *BatchEventMigrateInstanceRequestEventInfos {
	s.ResourceId = &v
	return s
}

func (s *BatchEventMigrateInstanceRequestEventInfos) Validate() error {
	return dara.Validate(s)
}
