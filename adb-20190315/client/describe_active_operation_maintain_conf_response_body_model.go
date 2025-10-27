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
	// The queried configuration information.
	Config *DescribeActiveOperationMaintainConfResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// Indicates whether the O\\&M task is configured. Valid values:
	//
	// 	- 1: yes.
	//
	// 	- 0: no.
	//
	// example:
	//
	// 1
	HasConfig *int32 `json:"HasConfig,omitempty" xml:"HasConfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92B3D2F0-B5E3-5592-9A6F-D0A6C34ACBA9
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
	// The time when the O\\&M task was created. The time follows the ISO 8601 standard in the *yyyy-mm-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-08T02:23:13Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The days on which the O\\&M task is executed every week.
	//
	// example:
	//
	// 0
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The unit of the O\\&M task cycle. Valid values:
	//
	// 	- Month
	//
	// 	- Week
	//
	// example:
	//
	// Month
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 20:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 18:00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 2023-05-12T12:34:40Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Indicates whether the configuration takes effect. Valid values:
	//
	// 	- 1: yes.
	//
	// 	- 2: no.
	//
	// example:
	//
	// 2
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
