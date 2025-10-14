// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckSqlFlashbackTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *PreCheckSqlFlashbackTaskRequest
	GetDbName() *string
	SetEndTime(v string) *PreCheckSqlFlashbackTaskRequest
	GetEndTime() *string
	SetPolardbxInstanceId(v string) *PreCheckSqlFlashbackTaskRequest
	GetPolardbxInstanceId() *string
	SetRegionId(v string) *PreCheckSqlFlashbackTaskRequest
	GetRegionId() *string
	SetStartTime(v string) *PreCheckSqlFlashbackTaskRequest
	GetStartTime() *string
}

type PreCheckSqlFlashbackTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-09-21 14:41:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	PolardbxInstanceId *string `json:"PolardbxInstanceId,omitempty" xml:"PolardbxInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-09-21 14:35:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PreCheckSqlFlashbackTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PreCheckSqlFlashbackTaskRequest) GoString() string {
	return s.String()
}

func (s *PreCheckSqlFlashbackTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *PreCheckSqlFlashbackTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *PreCheckSqlFlashbackTaskRequest) GetPolardbxInstanceId() *string {
	return s.PolardbxInstanceId
}

func (s *PreCheckSqlFlashbackTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PreCheckSqlFlashbackTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PreCheckSqlFlashbackTaskRequest) SetDbName(v string) *PreCheckSqlFlashbackTaskRequest {
	s.DbName = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskRequest) SetEndTime(v string) *PreCheckSqlFlashbackTaskRequest {
	s.EndTime = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskRequest) SetPolardbxInstanceId(v string) *PreCheckSqlFlashbackTaskRequest {
	s.PolardbxInstanceId = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskRequest) SetRegionId(v string) *PreCheckSqlFlashbackTaskRequest {
	s.RegionId = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskRequest) SetStartTime(v string) *PreCheckSqlFlashbackTaskRequest {
	s.StartTime = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskRequest) Validate() error {
	return dara.Validate(s)
}
