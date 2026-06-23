// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationMaintainConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *DescribeActiveOperationMaintainConfResponseBodyConfig) *DescribeActiveOperationMaintainConfResponseBody
	GetConfig() *DescribeActiveOperationMaintainConfResponseBodyConfig
	SetHasConfig(v int32) *DescribeActiveOperationMaintainConfResponseBody
	GetHasConfig() *int32
	SetRequestId(v string) *DescribeActiveOperationMaintainConfResponseBody
	GetRequestId() *string
}

type DescribeActiveOperationMaintainConfResponseBody struct {
	// The active O\\&M configuration items.
	Config *DescribeActiveOperationMaintainConfResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// Indicates whether a configuration is set.
	//
	// - 1: A configuration is set.
	//
	// - 0: No configuration is set.
	//
	// example:
	//
	// 1
	HasConfig *int32 `json:"HasConfig,omitempty" xml:"HasConfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A23C87D-87DF-4DA0-A50E-CB13F4F7923D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeActiveOperationMaintainConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfResponseBody) GetConfig() *DescribeActiveOperationMaintainConfResponseBodyConfig {
	return s.Config
}

func (s *DescribeActiveOperationMaintainConfResponseBody) GetHasConfig() *int32 {
	return s.HasConfig
}

func (s *DescribeActiveOperationMaintainConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetConfig(v *DescribeActiveOperationMaintainConfResponseBodyConfig) *DescribeActiveOperationMaintainConfResponseBody {
	s.Config = v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetHasConfig(v int32) *DescribeActiveOperationMaintainConfResponseBody {
	s.HasConfig = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetRequestId(v string) *DescribeActiveOperationMaintainConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeActiveOperationMaintainConfResponseBodyConfig struct {
	// The time when the configuration was created.
	//
	// example:
	//
	// 2025-03-18T06:51:11+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The day of the cycle.
	//
	// - If CycleType is Month, this parameter returns one or more numbers from 1 to 28 that indicate the day of the month. Multiple numbers are separated by commas (,).
	//
	// - If CycleType is Week, this parameter returns one or more numbers from 1 to 7 that indicate the day of the week. Multiple numbers are separated by commas (,).
	//
	// example:
	//
	// 1
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The cycle type.
	//
	// - Month
	//
	// - Week
	//
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The time when the configuration was last modified.
	//
	// example:
	//
	// 2025-06-27T07:12:53Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The status of the configuration. A value of 1 indicates that the configuration is enabled. A value of 2 indicates that the configuration is disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeActiveOperationMaintainConfResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetCycleTime() *string {
	return s.CycleTime
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetCycleType() *string {
	return s.CycleType
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetCreatedTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetCycleTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.CycleTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetCycleType(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.CycleType = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetMaintainEndTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetMaintainStartTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetModifiedTime(v string) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetStatus(v int32) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
