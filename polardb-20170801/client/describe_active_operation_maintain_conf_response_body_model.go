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
	SetHasConfig(v int64) *DescribeActiveOperationMaintainConfResponseBody
	GetHasConfig() *int64
	SetRequestId(v string) *DescribeActiveOperationMaintainConfResponseBody
	GetRequestId() *string
}

type DescribeActiveOperationMaintainConfResponseBody struct {
	// The configuration information.
	Config *DescribeActiveOperationMaintainConfResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// Indicates whether the configuration has been set. Valid values: 1 (yes) and 0 (no). On the first visit, HasConfig is 0.
	//
	// example:
	//
	// 0
	HasConfig *int64 `json:"HasConfig,omitempty" xml:"HasConfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14109129-EF13-5C83-AD86-7581D9552603
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

func (s *DescribeActiveOperationMaintainConfResponseBody) GetHasConfig() *int64 {
	return s.HasConfig
}

func (s *DescribeActiveOperationMaintainConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetConfig(v *DescribeActiveOperationMaintainConfResponseBodyConfig) *DescribeActiveOperationMaintainConfResponseBody {
	s.Config = v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBody) SetHasConfig(v int64) *DescribeActiveOperationMaintainConfResponseBody {
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
	// The creation time.
	//
	// The time is in the YYYY-MM-DDTHH:mm:ss format.
	//
	// example:
	//
	// 2023-07-04T19:28:46
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cycle time.
	//
	// - If CycleType is set to Month, a number from 1 to 28 is returned, indicating the day of the month. Multiple values are separated by commas (,).
	//
	// - If CycleType is set to Week, a number from 1 to 7 is returned, indicating the day of the week. Multiple values are separated by commas (,).
	//
	// example:
	//
	// 1
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
	CycleType *int32 `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The end time of the maintenance window.
	//
	// The time is in UTC in the HH:mmZ format.
	//
	// example:
	//
	// 09:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// The time is in UTC. Example: 8:00Z.
	//
	// example:
	//
	// 8:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The modification time.
	//
	// The time is in the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-04-02T02:10:08Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Indicates whether the configuration is effective. Valid values: 1 (yes) and 2 (no).
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetCycleType() *int32 {
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

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) GetStatus() *int64 {
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

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetCycleType(v int32) *DescribeActiveOperationMaintainConfResponseBodyConfig {
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

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) SetStatus(v int64) *DescribeActiveOperationMaintainConfResponseBodyConfig {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
