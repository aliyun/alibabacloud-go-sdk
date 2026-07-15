// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationMaintenanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) *DescribeActiveOperationMaintenanceConfigResponseBody
	GetConfig() *DescribeActiveOperationMaintenanceConfigResponseBodyConfig
	SetHasConfig(v int32) *DescribeActiveOperationMaintenanceConfigResponseBody
	GetHasConfig() *int32
	SetRequestId(v string) *DescribeActiveOperationMaintenanceConfigResponseBody
	GetRequestId() *string
}

type DescribeActiveOperationMaintenanceConfigResponseBody struct {
	// The details of the configuration.
	Config *DescribeActiveOperationMaintenanceConfigResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// Indicates whether a configuration is set.
	//
	// - 1: A configuration is set.
	//
	// - 0: No configuration is set.
	//
	// example:
	//
	// 0
	HasConfig *int32 `json:"HasConfig,omitempty" xml:"HasConfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 72651AF9-7897-41A7-8B67-6C15C7F26A0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeActiveOperationMaintenanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintenanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBody) GetConfig() *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	return s.Config
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBody) GetHasConfig() *int32 {
	return s.HasConfig
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBody) SetConfig(v *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) *DescribeActiveOperationMaintenanceConfigResponseBody {
	s.Config = v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBody) SetHasConfig(v int32) *DescribeActiveOperationMaintenanceConfigResponseBody {
	s.HasConfig = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBody) SetRequestId(v string) *DescribeActiveOperationMaintenanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeActiveOperationMaintenanceConfigResponseBodyConfig struct {
	// The time when the O\\&M task configuration was created. The time is in the *yyyy-mm-dd*T*hh:mm:ss*Z format and is displayed in UTC.
	//
	// example:
	//
	// 2018-05-30T14:30:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The days of the cycle.
	//
	// - If CycleType is set to Month, this parameter returns the days of the month. Valid values: 1 to 28. Multiple values are separated by commas (,).
	//
	// - If CycleType is set to Week, this parameter returns the days of the week. Valid values: 1 to 7. Multiple values are separated by commas (,).
	//
	// example:
	//
	// 1,2,3
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The cycle type.
	//
	// - Month
	//
	// - Week
	//
	// example:
	//
	// Month
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The end time of the O\\&M time window. The time is in the *hh:mm:ss*Z format and is displayed in UTC.
	//
	// example:
	//
	// 20:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the O\\&M time window. The time is in the *hh:mm:ss*Z format and is displayed in UTC.
	//
	// example:
	//
	// 18:00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The time when the operations and maintenance (O\\&M) task configuration was modified. The time is in the *yyyy-mm-dd*T*hh:mm:ss*Z format and is displayed in UTC.
	//
	// example:
	//
	// 2024-05-30T14:30:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Indicates whether the configuration is enabled.
	//
	// - 1: Enabled
	//
	// - 2: Disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeActiveOperationMaintenanceConfigResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GetCycleTime() *string {
	return s.CycleTime
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GetCycleType() *string {
	return s.CycleType
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) SetCreatedTime(v string) *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) SetCycleTime(v string) *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	s.CycleTime = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) SetCycleType(v string) *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	s.CycleType = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) SetMaintainEndTime(v string) *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) SetMaintainStartTime(v string) *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) SetModifiedTime(v string) *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) SetStatus(v int32) *DescribeActiveOperationMaintenanceConfigResponseBodyConfig {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
