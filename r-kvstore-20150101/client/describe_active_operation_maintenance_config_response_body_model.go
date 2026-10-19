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
	// The configuration details of the O&M task.
	Config *DescribeActiveOperationMaintenanceConfigResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// Indicates whether the configuration exists. Valid values:
	//
	// - 1: Configured.
	//
	// - 0: Not configured.
	//
	// example:
	//
	// 0
	HasConfig *int32 `json:"HasConfig,omitempty" xml:"HasConfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 794120D1-E0CF-4713-BAE4-EBAEA04506AF
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
	// The time when the O&M task configuration was created. The time is in the <i>yyyy-mm-dd</i>T<i>hh:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-05-30T14:30:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cycle time.
	//
	// - If CycleType is set to Month, a number from 1 to 28 is returned, indicating the day of the month. Multiple values are separated by commas (,). The value is 0 if the configuration is inactive.
	//
	// - If CycleType is set to Week, a number from 1 to 7 is returned, indicating the day of the week. Multiple values are separated by commas (,). The value is 0 if the configuration is inactive.
	//
	// example:
	//
	// 1,2,3
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The cycle type. Valid values:
	//
	// - Month: monthly
	//
	// - Week: weekly
	//
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The end time of the O&M time window. The time is in the <i>hh:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 20:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the O&M time window. The time is in the <i>hh:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 18:00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The time when the O&M task configuration was last modified. The time is in the <i>yyyy-mm-dd</i>T<i>hh:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2024-05-30T14:30:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Indicates whether the configuration is active. Valid values:
	//
	// - 1: Active.
	//
	// - 2: Inactive.
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
